package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strconv"
	"strings"
)

func main() {

	expr := getExpr()

	res, _ := parser.ParseExpr(expr)
	//fs := token.NewFileSet()
	fmt.Printf("%d\n", Eval(res))

}

func getExpr() string {

	// get the math expression
	fmt.Println(" Please enter a math expression \n ")
	var expr string
	_, err := fmt.Scan(&expr)

	if err != nil {
		panic("crap!!!")
	}

	//validExpr(expr)

	return expr
}

func validExpr(exprs string) {
	fmt.Print(len(exprs))
	if len(exprs) < 5 {
		panic("use the right format of expr")
	}

	parms := strings.Fields(exprs)
	if len(parms) < 3 {
		panic("not enough pieces in expression")
	}
}

func Eval(exp ast.Expr) int {
	switch exp := exp.(type) {
	case *ast.BinaryExpr:
		return EvalBinaryExpr(exp)
	case *ast.BasicLit:
		switch exp.Kind {
		case token.INT:
			i, _ := strconv.Atoi(exp.Value)
			return i
		}
	}

	return 0
}

func EvalBinaryExpr(exp *ast.BinaryExpr) int {
	left := Eval(exp.X)
	right := Eval(exp.Y)

	switch exp.Op {
	case token.ADD:
		return left + right
	case token.SUB:
		return left - right
	case token.MUL:
		return left * right
	case token.QUO:
		return left / right
	}

	return 0
}
