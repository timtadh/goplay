package main

import "fmt"

func fib(i int) (r int)
{
    if (i == 0) { r = 0; }
    else if (i == 1) { r = 1; }
    else { r = fib(i-2) + fib(i-1)}
    fmt.Println(r);
    return
}

func main()
{
    fmt.Println(fib(5));
}
