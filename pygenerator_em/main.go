package main

import "fmt"

func main() {
    // yield and results are the same channel.
    results := make(chan int)

    // ack acknowledge reciept and processing of last result
    ack := make(chan bool)

    // Caculating function
    go func(yield chan<- int, ack <-chan bool) {
        
        for i := 0; i < 10; i++ {
            yield<-i
            <-ack
        }
        close(yield)                            // causes the processing loop to quit
    }(results, ack)

    // Processing Loop
    for result := range results {
        fmt.Println(result);
        ack<-true;                              // ack<-true must be the last line of the loop.
    }
}
