package lex

import ("os"; "fmt"; re "regexp");

type Type int;

const
(
    Int Type = iota;
    Str;
    Flt;
    None;
);

const
(
    L_BRACE = iota;
    R_BRACE;
    WORD;
    SEMI;
    INT;
    DOT;
    WHITE;
);

type Chr uint8;
type Token struct
{
    name string;
    token uint16;
    attr_type Type;
    int_attr int64;
    str_attr string;
    flt_attr float64;
}

type match_handler func (str string) Token;

var regexs = map[string] *re.Regexp
{
    "[^{} \t\n]+":re.MustCompile("^[a-zA-Z0-9_]+$"),
    "{":re.MustCompile("^{$"),
    "}":re.MustCompile("^}$"),
    "^( +|\t+|\n)$":re.MustCompile("^( +|\t+|\n)$"),
    "^;$":re.MustCompile("^;$")
};

func l_brace(str string) Token
{
    return Token{"L_BRACE", L_BRACE, None, 0, "", 0.0};
}

func r_brace(str string) Token
{
    return Token{"R_BRACE", R_BRACE, None, 0, "", 0.0};
}

func word(str string) Token
{
    return Token{"WORD", WORD, Str, 0, str, 0.0};
}

func white(str string) Token
{
    return Token{"WHITE", WHITE, None, 0, "", 0.0};
}

func (chr Chr) String() string
{
    return string(chr);
}

var handlers = map[string] match_handler {"[^{} \t\n]+":word, "{":l_brace, "}":r_brace, " |\t|\n":white};

func Lex(input <-chan Chr, tokens chan<- Token, lex_error chan<- os.Error, quit chan<- bool)
{
    buf := make([]byte, 50);
    i := 0;
    skip := false;
    var cur Chr;
    for true
    {
        if !skip
        {
            cur = <-input;
            if cur == 255 { break; }
            buf[i] = uint8(cur); i++;
        }
        else
        {
            skip = false
        }
        for s, r := range regexs
        {
            fmt.Printf("match='%s' to buf='%s'", s, string(buf[0:i]));
            fmt.Println(r.Match(buf[0:i]));
            if r.Match(buf[0:i])
            {
                fmt.Println(s);
                for r.Match(buf[0:i])
                {
                    fmt.Println(string(buf[0:i]));
                    cur = <-input;
                    if cur == 255 { quit <- true; return; }
                    buf[i] = uint8(cur);
                    i++;
                }
                match := string(buf[0:i-1]);
                fmt.Println(s, "---->", match, string(cur));
                fmt.Printf("MATCH: '%s'\n", match);
                i = 1;
                buf[0] = uint8(cur);
                skip = true;
                break;
            }
        }
        fmt.Println(i, string(buf[0:i]));
    }
    quit <- true;
}

