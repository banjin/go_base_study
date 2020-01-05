package racer

import "time"
import "net/http"



func Racer(slow, fast string) (winner string){
	startA := time.Now()
	http.Get(slow)
	aDuration := time.Since(startA)

	startB := time.Now()
    http.Get(fast)
	bDuration := time.Since(startB)
	
	if aDuration < bDuration {
        return slow
	}
	
	return fast
}