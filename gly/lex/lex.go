package lex

import ("os"; "fmt");

type Token int;
type Chr uint8;

func (chr Chr) String() string
{
    return string(chr);
}

func Lex(input <-chan Chr, tokens chan<- Token, lex_error chan<- os.Error, quit chan<- bool)
{
    for cur := <-input; cur != 255; cur = <-input
    {
        fmt.Print(cur)
    }
    quit <- true;
}

