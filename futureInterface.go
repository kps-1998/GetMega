package main

import (
  "time"
)

type Future Interface{
     get() Result
     getWithTimeout(timeout time.Duration) Result 
     isDone() bool
     isCancelled() bool
     cancel(mayInterruptIfRunning bool)
     setException(exception exception)
}
