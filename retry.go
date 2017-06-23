package retry

import (
	"fmt"
	//"mytabler-api/logging"
	"time"
)

//var log = logging.GetLogger("retry")

func Retry(attempts int, sleep time.Duration, callback func() error) (err error) {
	for i := 0; ; i++ {
		err = callback()
		if err == nil {
			return
		}

		if i >= (attempts - 1) {
			break
		}

		time.Sleep(sleep)

		//log.Error("retrying after error:", "reason", err)
	}
	return fmt.Errorf("Retry(): after %d attempts, last error: %s", attempts, err)
}

func RetryDuring(duration time.Duration, sleep time.Duration, callback func() error) (err error) {
	t0 := time.Now()
	i := 0
	for {
		i++

		err = callback()
		if err == nil {
			return
		}

		delta := time.Now().Sub(t0)
		if delta > duration {
			return fmt.Errorf("Retry.During(): after %d attempts (during %s), last error: %s", i, delta, err)
		}

		time.Sleep(sleep)

		//log.Error("retrying after error:", "reason", err)
	}
}

