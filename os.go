package tomato

import "time"

type GenericOS struct{}

func (this *GenericOS) Sleep(duration time.Duration) {
	time.Sleep(duration)
}
