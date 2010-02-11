package main;
import "exp/draw/x11";
import "fmt";
//import "time";
//import "exp/draw";
//import "os";

func init(){
    return;
}

func main() {
    //var x draw.Context;
    //var y os.Error;
    x, y := x11.NewWindow();
    if (y != nil) { fmt.Println("error"); }
    if (x == nil) { fmt.Println("error"); } 
    fmt.Println("hello");
    for i := 0; i < 1000000000 ; i++ {
        fmt.Print("x");
    } 
}


