%{
package main
import(
    "fmt"
    "bufio"
    "os"
)
const(
    Debug = 4
    ErrorVerbose = true
)
%}

%union {
    int_value int
    float_value float64
}

%token <float_value> DOUBLE_LITERAL
%token ADD SUB MUL DIV CR LP RP

%type <float_value> expression term primary_expression

%%
line_list
    : line
    | line_list line
    ;
line
    : expression CR
    {
        fmt.Printf(">>%1f\n", $1);
    }
    ;
expression
    : term
    | expression ADD term
    {
        $$ = $1 + $3;
    }
    | expression SUB term
    {
        $$ = $1 - $3;
    }
    ;
term
    : primary_expression
    | term MUL primary_expression
    {
        $$ = $1 * $3;
    }
    | term DIV primary_expression
    {
        $$ = $1 / $3;
    }
    ;
primary_expression
    : DOUBLE_LITERAL
    | LP expression RP
    {
        $$ = $2;
    }
    | SUB primary_expression
    {
        $$ = -$2;
    }
    ;
%%

func main(){
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan(){
        text := scanner.Text()
        text = fmt.Sprintf("%s\n", text)
        yyParse(&GoCalcLex{Input: []byte(text)})
    }
}

