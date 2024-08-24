package main

import(
	"fmt"
	"runtime"
	"sync"
	"time"
)
// goroutine it

func main() {
	fmt.Println("CPU's:", runtime.NumCPU())
	fmt.Println("------------------------")
	fmt.Println("Goroutine it man:", runtime.NumGoroutine())

    counter := 0    //starting off at 0

    const gs = 15
    var wg sync.WaitGroup
    wg.Add(gs)

    for i := 0; i < gs; i++ {
    	    go func() {
    	    	v := counter
    	    	time.Sleep(time.Second)
    	    	v++
    	    	counter = v
    	    	wg.Done()
			}()
    	    fmt.Println("Goroutine", runtime.NumGoroutine())
        }
        wg.Wait()
    fmt.Println("Goroutine", runtime.NumGoroutine())
    fmt.Println("Count number is:", counter)
}
