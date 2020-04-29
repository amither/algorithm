package main

import (
	"fmt"
	"math"
	"time"
)

var first_above_time time.Time //Time when we'll declare we're above target(0 if below), drop one packet and cal drop_next when now >= first_above_time + interval
var drop_next time.Time        //Time for drop next packet, drop_next = control_law(drop_next)
var count uint32               //packets dropped since going into drop state
var dropping bool              //drop state
var queue chan *packet_t

const (
	target    = time.Millisecond * 5
	interval  = time.Millisecond * 100
	maxpacket = 512
)

type packet_t struct {
	timestamp time.Time
	no        int
}

func (p *packet_t) String() string {
	return fmt.Sprintf("timestamp[%s], no[%d]", p.timestamp, p.no)
}

type dodeque_result struct {
	p           *packet_t
	ok_to_drop  bool
	actual_drop bool
}

func enque(p *packet_t) {
	queue <- p
}

func drop(r *dodeque_result) {
	r.actual_drop = true
	log(r)
}
func log(r *dodeque_result) {
	fmt.Printf("time: %v, packet: %s , ok_to_drop:%v, actual_drop:%v, first_above_time: %v, drop_next: %v, count: %v, dropping: %v \n", time.Now(), r.p, r.ok_to_drop, r.actual_drop, first_above_time, drop_next, count, dropping)
}

func dodeque(now time.Time) *dodeque_result {
	packet := <-queue
	r := &dodeque_result{packet, false, false}
	sojourn_time := now.Sub(r.p.timestamp)
	if sojourn_time < target {
		first_above_time = time.Time{}
	} else {
		if first_above_time.Equal(time.Time{}) {
			first_above_time = now.Add(interval)
		} else if !now.Before(first_above_time) {
			r.ok_to_drop = true
		}
	}

	return r
}

func deque() {
	now := time.Now()
	r := dodeque(now)
	if dropping { // in drop state
		if !r.ok_to_drop {
			dropping = false
		} else if !now.Before(drop_next) {
			for dropping && !now.Before(drop_next) {
				drop(r)
				count++
				r = dodeque(now)
				if !r.ok_to_drop {
					dropping = false
				} else {
					drop_next = control_law(drop_next)
				}
			}
		}
		// not in drop state
		//If we get here, then we're not in dropping state. If the sojourn time has been
		//above target for interval, then we decide whether it's time to enter dropping state.
		//We do so if we've been either in dropping state recently or above target for a relatively long time.
		//The "recently" check helps ensure that when we're successfully controlling the queue we react quickly (in
		//one interval) and start with the drop rate that controlled the queue last time rather than relearn the
		//correct rate from scratch. If we haven't been dropping recently, the
		//"long time above" check adds some hysteresis to the state entry so we don't drop on a slightly
		//bigger-than-normal traffic pulse into an otherwise quiet queue.
	} else if r.ok_to_drop && (now.Sub(drop_next) < interval || now.Sub(first_above_time) >= interval) {
		dropping = true
		drop(r)
		r = dodeque(now)

		if now.Sub(drop_next) < interval {
			if count > 2 {
				count = count - 2
			} else {
				count = 1
			}
		} else {
			count = 1
		}

		drop_next = control_law(now)
	}

	log(r)
}

func control_law(t time.Time) time.Time {
	return t.Add(interval / time.Duration(math.Sqrt(float64(count))))
}

func init() {
	queue = make(chan *packet_t, 10000)
}

func main() {
	go func() {
		for {
			deque()
			time.Sleep(20 * time.Millisecond)
		}
	}()

	for i := 0; i < 20; i++ {
		p := &packet_t{time.Now(), i}
		enque(p)
	}
	time.Sleep(1000 * time.Second)
}
