# GoRoutines #
## Below are basic examples of using goroutines and channels to illustrate the concurrency model of  the Go language . ##

-------------------------------------------------------------------------------------------------------------------------------------------------

The file `producer_consumer.go` illustrates a simple producer/consumer pattern. \
It implements 2 goroutines and  2 channels. \
The `producer` function send 5 `true` values to the `comm` channel. \
The `consumer` function receives values from the `comm` channel in an infinite loop . \
and prints a message for each received value. \
When the `producer` has finished sending all 5 values, it sends a `true` value to the `done` channel to signal that it is done. \
As an the entry point to the program ,  in the `main` function , the `producer` and `consumer` functions are launched \
as separate goroutines. \
The `main` function then waits for a value from the `done` channel and before printing the final  output 'All Done!".


------------------------------------------------------------------------------------------------------------------------------------------------

The file  `multiple.go`  creates a variable number of goroutines.
As a Go program does not wait for  its go routines to end before exiting, 
we need to delay it manually, which is the purpose od the `time.Sleep()` call.
We will correct that by using Waitgroups to wait for all goroutines to finsih before exiting.
