package main

import {
     "fmt"
     "time"
}

//function to return a reference to a future object
func ReturnFuture(task func() (Result)) *FutureTask{

	channel := make(chan Result)

	futureObject := FutureTask{
		success :          false,
		done    :          false,
		exception   :          nil,
		result  : 		   Result{},
		interfaceChannel : channel,
	}

	go func(){
		fmt.Println("go routine start")
		defer close(channel)
		resultObject := task()
		channel <- resultObject
		fmt.Println("go routine end")
	}()
	return &futureObject
}

//Main method
func main(){

	testCase1:= ReturnFuture(func() (Result){
			var res interface{}
			res=1+2
			time.Sleep(4*time.Second)
			return Result{resultValue:res}
	})
	testCase2:= ReturnFuture(func() (Result){
			var res interface{}
			res="4"
			time.Sleep(1*time.Second)
			return Result{resultValue:res}
	})

	exception1 := testCase1.setException(&ExecutionException{exception:"Mark the task as complete"})
	fmt.Println(exception1)

	fmt.Println(testCase2.getWithTimeout(4*time.Second))
	fmt.Println(testCase1.get())

	exception2 := testCase1.setException(&ExecutionException{exception:"Mark the task as complete"})
	fmt.Println(exception2)
}
