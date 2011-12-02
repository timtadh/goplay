package main

import "os"
import "strings"

import "stack"
import "fmt"

import "reflect"

import "go/parser"
import "go/ast"
import "go/token"
// import "go/printer"

type IntStack stack.Stack
type StrStack stack.Stack

func (s *IntStack) Peek() int {
	self := (*stack.Stack)(s)
	i, _ := self.Peek().(int)
	return i
}

func (s *IntStack) Push(i int) {
	self := (*stack.Stack)(s)
	self.Push(i)
}

func (s *IntStack) Pop() int {
	self := (*stack.Stack)(s)
	i, _ := self.Pop().(int)
	return i
}

func (s *StrStack) Peek() string {
	self := (*stack.Stack)(s)
	i, _ := self.Peek().(string)
	return i
}

func (s *StrStack) Push(i string) {
	self := (*stack.Stack)(s)
	self.Push(i)
}

func (s *StrStack) Pop() string {
	self := (*stack.Stack)(s)
	i, _ := self.Pop().(string)
	return i
}

func main() {
	{
		s := (*IntStack)(stack.NewStack())
		s.Push(1)
		fmt.Println(s.Peek())
		s.Push(2)
		fmt.Println(s.Pop())
		fmt.Println(s.Pop())
	}
	{
		s := (*StrStack)(stack.NewStack())
		s.Push("Hello")
		fmt.Println(s.Peek())
		s.Push("World")
		fmt.Println(s.Pop())
		fmt.Println(s.Pop())
	}
	fmt.Println()
	fmt.Println()
	if asts, err := parser.ParseDir("stack",
		func(finfo *os.FileInfo) (bool) {
			return strings.HasSuffix(finfo.Name, ".go")
		}, 0); err != nil {
		fmt.Println(err)
	} else {
		indent := func(i int) string {
			s := ""
			for j := 0; j < i; j++ {
				s += " "
			}
			return s
		}
		var decl func(d ast.Decl)
		var gendecl func(d *ast.GenDecl)
		var funcdecl func(d *ast.FuncDecl)
		var blockstmt func(d *ast.BlockStmt)
		decl = func (d ast.Decl) {
			switch d.(type) {
				case *ast.GenDecl:
					gendecl(d.(*ast.GenDecl))
				case *ast.FuncDecl:
					funcdecl(d.(*ast.FuncDecl))
				default:
					fmt.Println(indent(8), "unknown decl", reflect.Typeof(d))
			}
		}
		gendecl = func (d *ast.GenDecl) {
			fmt.Println(indent(8), "GenDecl", d.Tok)
			switch d.Tok {
				case token.IMPORT:
					fmt.Println(indent(12), "import", d.Tok)
					for i := range d.Specs {
						imp := d.Specs[i].(*ast.ImportSpec)
						l := len(imp.Path.Value)
						fmt.Println(indent(16), string(imp.Path.Value[1:l-1]))
					}
				case token.TYPE:
					fmt.Println(indent(12), "type", d.Tok)
				default:
					fmt.Println(indent(12), "unknown tok", d.Tok)
			}
		}
		funcdecl = func (d *ast.FuncDecl) {
			fmt.Println(indent(8), "FuncDecl", d.Type.Results)
			blockstmt(d.Body)
		}
		blockstmt = func (d *ast.BlockStmt) {
			fmt.Println(indent(12), "BlockStmt", d)
			for i := range d.List {
				fmt.Println(indent(16), "Stmt", d.List[i])
			}
		}
		for name,pack := range asts {
			fmt.Println(name)
			for name,file := range pack.Files {
				fmt.Println(indent(4), name)
				for i := range file.Decls {
					decl(file.Decls[i])
				}
			}
		}
	}
}
