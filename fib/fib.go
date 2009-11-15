package main

import ("fmt"; "syscall");

func fib(i int) (r int)
{
    if (i == 0) { r = 0; }
    else if (i == 1) { r = 1; }
    else { r = fib(i-2) + fib(i-1)}
    return
}

func video()
{
    fmt.Println(syscall.Getpid());
    //fmt.Println(VideoShutdown());
}

func main()
{
    fmt.Println(fib(5));
}
