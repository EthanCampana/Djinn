package evaluator

import (
	"djinn/ast"
	"djinn/object"
)

// STOPS the AST and starts defining the macros
func DefineMacros(program *ast.Program, env *object.Enviroment) {
	definitions := []int{}
	for i, statement := range program.Statements {
		if isMacroDefinition(statement) {
			addMacro(statement, env)
			definitions = append(definitions, i)
		}
	}
	for i := len(definitions) - 1; i >= 0; i = i - 1 {
		definitionIndex := definitions[i]
		program.Statements = append(
			program.Statements[:definitionIndex],
			program.Statements[definitionIndex+1:]...,
		)
	}

}

//Maps a variable that maps to a macro and makes it a macro i.e a = macro(a,b).. b = a  b is a macro now
func addMacro(stmt ast.Statement, env *object.Enviroment) {
	CreateStatement, _ := stmt.(*ast.CreateStatement)
	macroLiteral, _ := CreateStatement.Value.(*ast.MacroLiteral)
	macro := &object.Macro{
		Parameters: macroLiteral.Parameters,
		Env:        env,
		Body:       macroLiteral.Body,
	}
	env.Set(CreateStatement.Name.Value, macro)
}

//Defines the macro from create statement cr a = macro(a,b) yada yada
func isMacroDefinition(node ast.Statement) bool {
	CreateStatement, ok := node.(*ast.CreateStatement)
	if !ok {
		return false
	}
	_, ok = CreateStatement.Value.(*ast.MacroLiteral)

	if !ok {
		return false
	}
	return true
}

func ExpandMacros(program ast.Node, env *object.Enviroment) ast.Node {
	return ast.Modify(program, func(node ast.Node) ast.Node {
		callExpression, ok := node.(*ast.CallExpression)
		if !ok {
			return node
		}
		macro, ok := isMacroCall(callExpression, env)
		if !ok {
			return node
		}
		args := quoteArgs(callExpression)
		evalEnv := extendMacroEnv(macro, args)
		evaluated := Eval(macro.Body, evalEnv)
		quote, ok := evaluated.(*object.Quote)
		if !ok {
			panic("we only support returning AST-nodes from macros")
		}
		return quote.Node
	})
}

func isMacroCall(
	exp *ast.CallExpression,
	env *object.Enviroment,
) (*object.Macro, bool) {
	identifier, ok := exp.Function.(*ast.Identifier)
	if !ok {
		return nil, false
	}
	obj, ok := env.Get(identifier.Value)
	if !ok {
		return nil, false
	}
	macro, ok := obj.(*object.Macro)
	if !ok {
		return nil, false
	}
	return macro, true
}
func quoteArgs(exp *ast.CallExpression) []*object.Quote {
	args := []*object.Quote{}
	for _, a := range exp.Arguments {
		args = append(args, &object.Quote{Node: a})
	}
	return args
}

func extendMacroEnv(
	macro *object.Macro,
	args []*object.Quote,
) *object.Enviroment {
	extended := object.NewEnclosedEnviroment(macro.Env)
	for paramIdx, param := range macro.Parameters {
		extended.Set(param.Value, args[paramIdx])
	}
	return extended
}
