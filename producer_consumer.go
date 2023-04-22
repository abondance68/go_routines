 



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






