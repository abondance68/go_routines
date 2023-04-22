/*
The main function starts two goroutines using the go keyword: one for the producer function and 
another for the consumer function. 
The producer function writes to the done channel after it finishes sending five messages to the comm channel.
The consumer function reads from the comm channel indefinitely
until it reads a false value, at which point it breaks out of the loop.
After starting the two goroutines, the main function blocks on the <-done statement, waiting for a value to be written to the done channel.
Once it receives a value, it proceeds to print "All Done!".








*/



package main

import ("fmt") 


var comm = make(chan bool)
var done = make(chan bool) 

func producer () { 
for  i:=0; i<5; i++ { 
comm<- true
}

comm <-false
done <-true

}

func consumer() { 
for { 
communication :=<-comm 
if !communication { 
break 
}
fmt.Println("Communication from producer received!", communication) 
}
}

func main() { 
go producer() 
go consumer() 
<-done
fmt.Println("All Done!") 
}






