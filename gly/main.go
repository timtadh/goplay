package main

import ("./lex/lex"; "os"; "fmt"; "io"; "bufio")

func main()
{
    fmt.Println("Welcome!!");
    input_stream := make(chan lex.Chr);
    token_stream := make(chan lex.Token);
    lex_errors := make(chan os.Error);
    lex_done := make(chan bool);
    go lex.Lex(input_stream, token_stream, lex_errors, lex_done);

    f, err := os.Open("test.java", os.O_RDONLY, 0666);
    if f == nil { fmt.Println(err); input_stream<-0; os.Exit(1); }

    reader := bufio.NewReader(io.Reader(f));
    b, err := reader.ReadByte();
    for ; err == nil;  b, err = reader.ReadByte()
    {
        input_stream<-lex.Chr(b);
    }
    input_stream<-255;
    f.Close();
    <-lex_done;
}

