//implementation of future interface

package main

import{
  "fmt"
  "time"
  "golang.org/x/net/context"
}

type Result struct{
     resultValue interface{}
     exception Exception
}


type FutureTask struct{
     success bool
     done bool
     exception Exception
     result Result
}

func (futureTask *FutureTask) get() Result{
  if(futureTask.done){
    return futureTask.result
  }
  ctx := context.Background()
  return futureTask.getWithContext(ctx)
}

func (futureTask *FutureTask) getWithTimeout(timeout time.Duration) Result{
  if(futureTask.done){
    return futureTask.result
  }
  ctx, cancel := context.WithTimeout(context.Background(), timeout)
  defer cancel()
  return futureTask.getWithContext(ctx)
}

func (futureTask *FutureTask) getWithContext(ctx context.Context) Result{
  select {
  case <-ctx.Done():
    futureTask.done = true
    futureTask.success = false
    futureTask.exception = &TimeoutException{exception:"Request Timeout!"}
    futureTask.result = Result{resultValue:nil,exception:futureTask.exception}
    return futureTask.result
  case futureTask.result = <-futureTask.interfaceChannel:
    if(futureTask.result.exception!=nil){
      futureTask.done = true
      futureTask.success = false
      futureTask.exception = futureTask.result.exception
    }else{
      futureTask.success = true
      futureTask.done = true
      futureTask.exception = nil
    }
    return futureTask.result
  }
}

func (futureTask *FutureTask) isDone() bool{
     if(futureTask.done){
        return true
     }
     return false
}

func (futureTask *FutureTask) isCancelled() bool{
     if(futureTask.done){
       if(futureTask.exception=="Manually Cancelled"){
          return true
       }
     }
     return false
}

func (futureTask *FutureTask) cancel(){
     if(futureTask.isDone()||futureTask.isCancelled()){
       return
     }
     futureTask.done = true
     futureTask.success = false
     cancelledException = &CancellationException{exception:"Manually Cancelled"}
     futureTask.exception = cancelledException
     futureTask.Result = Result{resultValue:nil, exception:cancelledException}
}

func (futureTask *FutureTask) setException(exception Exception){
     futureTask.exception = exception
     return
 }
