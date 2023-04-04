# GoRoutines #
## Below are basic examples of using go_routines and channels to illustrate the concurrency model of  the Go language . ##

The file  `multiple.go`  creates a variable number of goroutines.
As a Go program does not wait for  its go routines to end before exiting, 
we need to delay it manually, which is the purpose od the `time.Sleep()` call.
We will correct that by using Waitgroups to wait for all goroutines to finsih before exiting.
