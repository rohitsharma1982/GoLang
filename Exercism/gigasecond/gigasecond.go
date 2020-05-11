package gigasecond

import (
	"time"
	"math"
)

func AddGigasecond(t time.Time) time.Time {

	return t.Add(time.Second * time.Duration(math.Pow10(9)))
}
