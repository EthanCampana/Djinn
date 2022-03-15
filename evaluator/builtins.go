package evaluator


import "djinn/object"


var builtins = map[string]*object.Builtin{
	"len": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1{
				return NewError("wrong number of arguments. got=%d, want=1",len(args))
			}
			switch arg := args[0].(type){
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			default:
				return NewError("argument to `len` not supported, got %s",args[0].Type())
			}
			
		},
	},
}