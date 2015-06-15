package parser

import (
	"container/list"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/iNamik/go_lexer"
	"github.com/tiago4orion/scherzo/lang"
)

type parserState struct {
	IsLParen                    bool
	IsRParen                    bool
	IsDoubleQuotedString        bool
	IsSingleQuotedString        bool
	IsEscapedDoubleQuotedString bool
	IsEscapedSingleQuotedString bool
	listOp                      *list.List
}

// We define our lexer tokens starting from the pre-defined EOF token
const (
	TokenEOF = lexer.TokenTypeEOF
	TokenNil = lexer.TokenTypeEOF + iota
	TokenLParen
	TokenRParen
	TokenSpace
	TokenNewline
	TokenDoubleQuotedString
	TokenSingleQuotedString
	TokenEscapedDoubleQuotedString
	TokenEscapedSingleQuotedString
	TokenSemiColon
	TokenWord
	TokenNumbers
	TokenUsing
)

var (
	lparen                         = []byte{'('}
	rparen                         = []byte{')'}
	bytesNonWord                   = []byte{' ', '\t', '\f', '\v', '\n', '\r', ';', '"', '\'', '\\', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '(', ')'}
	bytesIntegers                  = []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	bytesSpace                     = []byte{' ', '\t', '\f', '\v'}
	bytesDoubleQuotedStrings       = []byte{'"'}
	bytesSingleQuotedStrings       = []byte{'\''}
	bytesEscapedDoubleQuotedString = []byte{'\\', '"'}
	bytesEscapedSingleQuotedString = []byte{'\\', '\''}
)

const (
	charNewLine   = '\n'
	charReturn    = '\r'
	charSemicolon = ';'
)

func parseCons(conString string) (lang.SExprs, int, error) {
	var (
		pState        = parserState{}
		ignoreCurrent int
		ignoreLength  int
		ignoreRunes   bool
	)

	fmt.Println("Parsing: ", conString)

	pState.listOp = list.New()

	// Create our lexer
	// NewSize(startState, reader, readerBufLen, channelCap)
	lex := lexer.NewFromString(lexFunc, conString, 1)
	//	var lastTokenType = TokenNil

	// Process lexer-emitted tokens
Scanner:
	for t := lex.NextToken(); lexer.TokenTypeEOF != t.Type(); t = lex.NextToken() {
		if ignoreRunes && ignoreCurrent < (lex.Column()+ignoreLength) {
			fmt.Println("Ignoring: ", string([]byte(conString)[lex.Column()-1:]))
			ignoreCurrent++
			continue
		} else {
			ignoreRunes = false
			ignoreCurrent = 0
			ignoreLength = 0
		}

		switch t.Type() {
		case TokenWord:
			word := string(t.Bytes())

			fmt.Println("TOKENWORD: ", word)

			if pState.IsDoubleQuotedString || pState.IsSingleQuotedString {
				fmt.Println("TOKENWORD inside quoted")
			} else {
				pState.listOp.PushBack(lang.NewAtom(strings.Trim(word, " ")))
			}
		case TokenLParen:
			fmt.Println("LPAREN: ", string(t.Bytes()))
			fmt.Println("Column: ", lex.Column())

			column := lex.Column()
			if column == 1 {
				pState.IsLParen = true
			} else {
				subExprSrc := string([]byte(conString)[lex.Column()-1:])
				subExpr, column, err := parseCons(subExprSrc)
				if err != nil {
					return nil, column, err
				}

				pState.listOp.PushBack(lang.NewAtom(subExpr))

				currentColumn := lex.Column()
				fmt.Println("Advancing from ", currentColumn, " to ", currentColumn+column)

				ignoreRunes = true
				ignoreLength = column
				ignoreCurrent = 0
			}
		case TokenRParen:
			fmt.Println("RPAREN: ", string(t.Bytes()))

			if pState.IsDoubleQuotedString || pState.IsSingleQuotedString {
				fmt.Println("RParen inside string")
			} else if pState.IsLParen {
				break Scanner
			}
		case TokenDoubleQuotedString:
			fmt.Println("TOKENDOUBLEQUOTE: ", string(t.Bytes()))

			pState.IsDoubleQuotedString = !pState.IsDoubleQuotedString
		case TokenSingleQuotedString:
			fmt.Println("TOKENSINGLEQUOTE: ", string(t.Bytes()))

			pState.IsSingleQuotedString = !pState.IsSingleQuotedString
		case TokenEscapedDoubleQuotedString:
			if pState.IsSingleQuotedString {
				panic("Escaped double quoted string inside single quoted string...")
			} else if pState.IsDoubleQuotedString {
				fmt.Println("TOKENESCAPEDDOUBLEQUOTE: ", string(t.Bytes()))
			}

			pState.IsEscapedDoubleQuotedString = !pState.IsEscapedDoubleQuotedString
		case TokenEscapedSingleQuotedString:
			if pState.IsDoubleQuotedString {
				return nil, lex.Column(), errors.New("Escaped single quoted string inside double quoted string...")
			} else if pState.IsSingleQuotedString {
				fmt.Println("TOKENESCAPEDSINGLEQUOTE: ", string(t.Bytes()))
			}

			pState.IsEscapedSingleQuotedString = !pState.IsEscapedSingleQuotedString
		case TokenSpace:
			// Spaces only makes difference inside quotes
			if pState.IsDoubleQuotedString || pState.IsSingleQuotedString {
				fmt.Println("SPACE: ", string(t.Bytes()))
			}
		case TokenNewline:
			// New lines only makes difference inside quotes
			if pState.IsDoubleQuotedString || pState.IsSingleQuotedString {
				fmt.Println("NEWLINE: ", string(t.Bytes()))
			}
		case TokenSemiColon:
			if pState.IsSingleQuotedString || pState.IsDoubleQuotedString {
				fmt.Println("SEMICOLON INSIDE STRING: ", string(t.Bytes()))
			} else {
				fmt.Println("SEMICOLON: ", string(t.Bytes()))
			}
		case TokenNumbers:
			if pState.IsSingleQuotedString || pState.IsDoubleQuotedString {
				fmt.Println("NUMBER INSIDE STRING: ", string(t.Bytes()))
			} else {
				numberStr := string(t.Bytes())
				fmt.Println("NUMBER: ", string(t.Bytes()))

				number, err := strconv.Atoi(numberStr)

				if err != nil {
					return nil, lex.Column(), err
				}

				pState.listOp.PushBack(lang.NewAtom(number))
			}
		default:
			return nil, lex.Column(), errors.New("Failed to parse line at '" + string(t.Bytes()) + "'")
		}

		//		lastTokenType = t.Type()
	}

	var finalExpr lang.SExprs

	last := lang.Nil
	finalExpr = lang.Nil

	for expr := pState.listOp.Back(); expr != nil; expr = pState.listOp.Back() {
		fmt.Println("last: ", last(1))
		var cdr lang.SExprs

		if last(1) == nil {
			cdr = last
		} else {
			cdr = lang.NewAtom(last)
		}

		finalExpr = lang.Cons(expr.Value.(lang.SExprs), cdr)
		last = finalExpr
		pState.listOp.Remove(expr)
	}

	fmt.Println("Finish expression")
	return finalExpr, lex.Column(), nil
}

// FromReader parse the file
func FromString(src string) (lang.SExprs, error) {
	sexpr, _, err := parseCons(src)
	return sexpr, err
}

func lexFunc(l lexer.Lexer) lexer.StateFn {
	// EOF
	if l.MatchEOF() {
		l.EmitEOF()
		return nil // We're done here
	}

	if l.MatchMinMaxBytes(bytesEscapedDoubleQuotedString, 2, 2) {
		l.EmitTokenWithBytes(TokenEscapedDoubleQuotedString)

	} else if l.MatchMinMaxBytes(bytesEscapedSingleQuotedString, 2, 2) {
		l.EmitTokenWithBytes(TokenEscapedSingleQuotedString)

	} else if l.MatchOneOrMoreBytes(bytesDoubleQuotedStrings) {
		l.EmitTokenWithBytes(TokenDoubleQuotedString)

	} else if l.MatchOneOrMoreBytes(bytesSingleQuotedStrings) {
		l.EmitTokenWithBytes(TokenSingleQuotedString)

	} else if l.NonMatchOneOrMoreBytes(bytesNonWord) {
		l.EmitTokenWithBytes(TokenWord)

	} else if l.MatchOneOrMoreBytes(bytesIntegers) {
		l.EmitTokenWithBytes(TokenNumbers)

		// Space run
	} else if l.MatchOneOrMoreBytes(bytesSpace) {
		l.EmitTokenWithBytes(TokenSpace)

		// (
	} else if l.MatchOneBytes(lparen) {
		l.EmitTokenWithBytes(TokenLParen)

		// )
	} else if l.MatchOneBytes(rparen) {
		l.EmitTokenWithBytes(TokenRParen)

		// Line Feed
	} else if charNewLine == l.PeekRune(0) {
		l.NextRune()
		l.EmitTokenWithBytes(TokenNewline)
		l.NewLine()

		// Carriage-Return with optional line-feed immediately following
	} else if charReturn == l.PeekRune(0) {
		l.NextRune()
		if charNewLine == l.PeekRune(0) {
			l.NextRune()
		}
		l.EmitTokenWithBytes(TokenNewline)
		l.NewLine()
	} else if charSemicolon == l.PeekRune(0) {
		l.NextRune()
		l.EmitTokenWithBytes(TokenSemiColon)
	} else {
		panic("Failed to parse line at '" + string(l.PeekRune(0)))
	}

	return lexFunc
}
