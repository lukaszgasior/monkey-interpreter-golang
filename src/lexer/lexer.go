package lexer

type Lexer struct {
	input       string
	position    int
	readPostion int
	ch          byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	return l
}
