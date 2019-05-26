package bruteforce

import (
	"testing"
)

func TestListTravel(t *testing.T) {
	l := list{}
	l.insert(1)
	l.insert(2)
	l.insert(3)
	l.travel()
	l.reverse()
	l.travel()
}
