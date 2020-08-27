package main

type TimeoutException Struct{
    exception string
}

type InterruptedException Struct{
    exception string
}

type CancellationException Struct{
    exception string
}

type ExecutionException Struct{
    exception string
}

func(e *TimeoutException) Exception() String{
    return e.exception
}

func(e *InterruptedException) Exception() String{
    return e.exception
}

func(e *CancellationException) Exception() String{
    return e.exception
}

func(e *ExecutionException) Exception() String{
    return e.exception
}
