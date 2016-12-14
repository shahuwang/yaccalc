package main

import (
	"fmt"
	"os"
	"strconv"
	"unicode"
	"unicode/utf8"
)

type LexerStatus int

const (
	INITIAL_STATUS LexerStatus = iota
	IN_INT_PART_STATUS
	DOT_STATUS
	IN_FRAC_PART_STATUS
)

type GoCalcLex struct {
	Pos    int
	Status LexerStatus
	Input  []byte
}

func (l *GoCalcLex) Lex(lval *yySymType) int {
	var Str []rune
	l.Status = INITIAL_STATUS
	for l.Pos < len(l.Input) {
		r, n := utf8.DecodeRune(l.Input[l.Pos:])
		if (l.Status == IN_INT_PART_STATUS || l.Status == IN_FRAC_PART_STATUS) &&
			!unicode.IsDigit(r) && r != '.' {
			lval.float_value, _ = strconv.ParseFloat(string(Str), 64)
			return DOUBLE_LITERAL
		}
		if unicode.IsSpace(r) {
			l.Pos = l.Pos + n
			if r == '\n' || r == '\r' {
				return CR
			}
			continue
		}
		Str = append(Str, r)
		l.Pos = l.Pos + n
		switch {
		case r == '+':
			return ADD
		case r == '-':
			return SUB
		case r == '*':
			return MUL
		case r == '/':
			return DIV
		case r == '(':
			return LP
		case r == ')':
			return RP
		case r == '.':
			if l.Status == IN_INT_PART_STATUS {
				l.Status = DOT_STATUS
			} else {
				fmt.Println("syntax error in dot")
				os.Exit(1)
			}
		case unicode.IsDigit(r):
			if l.Status == INITIAL_STATUS {
				l.Status = IN_INT_PART_STATUS
			} else if l.Status == DOT_STATUS {
				l.Status = IN_FRAC_PART_STATUS
			}
		default:
			fmt.Println("unknow input type %s", string(r))
			os.Exit(1)
		}
	}
	return 0
}

func (l *GoCalcLex) Error(s string) {
	fmt.Printf("syntax error: %s\n", s)
}
