package interpreter

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/Sameetpatro/odlang/ast"
)

var stdinScanner = bufio.NewScanner(os.Stdin)

type Environment struct {
	store map[string]interface{}
	types map[string]string
	outer *Environment
}

func NewEnvironment() *Environment {
	env := &Environment{
		store: make(map[string]interface{}),
		types: make(map[string]string),
	}
	registerBuiltins(env)
	return env
}

func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

func (env *Environment) Get(name string) (interface{}, bool) {
	val, ok := env.store[name]
	if ok {
		return val, true
	}
	if env.outer != nil {
		return env.outer.Get(name)
	}
	return nil, false
}

func (env *Environment) GetType(name string) string {
	if t, ok := env.types[name]; ok {
		return t
	}
	if env.outer != nil {
		return env.outer.GetType(name)
	}
	return ""
}

func (env *Environment) Set(name string, val interface{}) {
	env.store[name] = val
}

func (env *Environment) SetType(name string, typeName string) {
	env.types[name] = typeName
}

type breakSignal struct{}

type continueSignal struct{}

type returnSignal struct {
	values []interface{}
}

type FunctionValue struct {
	Parameters  []ast.Parameter
	ReturnTypes []string
	Body        []ast.Statement
	Closure     *Environment
}

type ClassValue struct {
	Name    string
	Fields  map[string]string
	Methods map[string]*FunctionValue
}

type InstanceValue struct {
	Class  *ClassValue
	Fields map[string]interface{}
}

func Eval(program *ast.Program) {
	env := NewEnvironment()
	for _, stmt := range program.Statements {
		evalStatement(stmt, env)
	}

	fnVal, ok := env.Get("aarambha")
	if !ok {
		fmt.Println("[OdLang Error] karya aarambha() not found")
		os.Exit(1)
	}

	fn, ok := fnVal.(*FunctionValue)
	if !ok {
		fmt.Println("[OdLang Error] karya aarambha() not found")
		os.Exit(1)
	}

	evalCallFunction(fn, nil)
}

func EvalProgram(program *ast.Program) interface{} {
	env := NewEnvironment()
	var last interface{}
	for _, stmt := range program.Statements {
		evalStatement(stmt, env)
		if es, ok := stmt.(*ast.ExpressionStatement); ok {
			last = evalExpression(es.Expression, env)
		}
	}
	return last
}

func evalStatement(stmt ast.Statement, env *Environment) interface{} {
	switch s := stmt.(type) {
	case *ast.VarStatement:
		var val interface{}
		if s.Value != nil {
			// krama nums(5, 0) → []interface{}{0,0,0,0,0}
			if s.TypeName == "krama" {
				if call, ok := s.Value.(*ast.CallExpression); ok {
					size := int(toInt64(evalExpression(call.Arguments[0], env)))
					var fill interface{}
					if len(call.Arguments) > 1 {
						fill = evalExpression(call.Arguments[1], env)
					}
					arr := make([]interface{}, size)
					for i := range arr {
						arr[i] = fill
					}
					val = arr
				}
			}
			if val == nil {
				val = evalExpression(s.Value, env)
			}
		} else {
			val = defaultValueForType(s.TypeName)
		}
		env.Set(s.Name, val)
		typeName := s.TypeName
		if inst, ok := val.(*InstanceValue); ok {
			typeName = inst.Class.Name
		}
		env.SetType(s.Name, typeName)
	case *ast.AssignStatement:
		val := evalExpression(s.Value, env)
		if s.LeftHand != nil {
			assignIndex(s.LeftHand, val, env)
		} else if values, ok := val.([]interface{}); ok && len(s.Targets) > 1 {
			for i, name := range s.Targets {
				if i < len(values) {
					env.Set(name, values[i])
				}
			}
		} else {
			for _, name := range s.Targets {
				env.Set(name, val)
			}
		}
	case *ast.LekhaStatement:
		parts := make([]string, 0, len(s.Arguments))
		for _, arg := range s.Arguments {
			parts = append(parts, formatValue(evalExpression(arg, env)))
		}
		fmt.Println(strings.Join(parts, " "))
	case *ast.DiaStatement:
		var fields []string
		if stdinScanner.Scan() {
			fields = strings.Fields(stdinScanner.Text())
		}
		for i, name := range s.Targets {
			typeName := env.GetType(name)
			var val interface{} = ""
			if i < len(fields) {
				val = fields[i]
			}
			switch typeName {
			case "sankhya":
				if s, ok := val.(string); ok {
					if parsed, err := strconv.ParseInt(s, 10, 64); err == nil {
						val = parsed
					}
				}
			case "dasmik":
				if s, ok := val.(string); ok {
					if parsed, err := strconv.ParseFloat(s, 64); err == nil {
						val = parsed
					}
				}
			case "akshara":
				if s, ok := val.(string); ok && len(s) > 0 {
					val = rune(s[0])
				}
			}
			env.Set(name, val)
		}
	case *ast.DeideStatement:
		values := make([]interface{}, 0, len(s.Values))
		for _, v := range s.Values {
			values = append(values, evalExpression(v, env))
		}
		return returnSignal{values: values}
	case *ast.IfStatement:
		if isTruthy(evalExpression(s.Condition, env)) {
			return evalBlock(s.Consequence, env)
		}
		for _, clause := range s.ElseIfs {
			if isTruthy(evalExpression(clause.Condition, env)) {
				return evalBlock(clause.Body, env)
			}
		}
		if s.Alternative != nil {
			return evalBlock(s.Alternative, env)
		}
	case *ast.GhuraStatement:
		start := toInt64(evalExpression(s.Start, env))
		end := toInt64(evalExpression(s.End, env))
		i := start
		// Step is stored as varName+"++" or varName+"--"; suffix check is intentional.
		increment := strings.HasSuffix(s.Step, "++")
		for {
			if increment {
				if i > end {
					break
				}
			} else if i < end {
				break
			}
			loopEnv := NewEnclosedEnvironment(env)
			loopEnv.Set(s.VarName, i)
			loopEnv.SetType(s.VarName, s.TypeName)
			sig := evalBlock(s.Body, loopEnv)
			if _, ok := sig.(breakSignal); ok {
				break
			}
			if _, ok := sig.(continueSignal); ok {
				if increment {
					i++
				} else {
					i--
				}
				continue
			}
			if increment {
				i++
			} else {
				i--
			}
		}
	case *ast.JetebeleJainStatement:
		for isTruthy(evalExpression(s.Condition, env)) {
			sig := evalBlock(s.Body, env)
			if _, ok := sig.(breakSignal); ok {
				break
			}
			if _, ok := sig.(continueSignal); ok {
				continue
			}
		}
	case *ast.BaharipadeStatement:
		return breakSignal{}
	case *ast.ChadideStatement:
		return continueSignal{}
	case *ast.KaryaStatement:
		env.Set(s.Name, &FunctionValue{
			Parameters:  s.Parameters,
			ReturnTypes: s.ReturnTypes,
			Body:        s.Body,
			Closure:     env,
		})
	case *ast.SreniStatement:
		fieldTypes := make(map[string]string)
		for _, field := range s.Fields {
			fieldTypes[field.Name] = field.TypeName
		}
		methods := make(map[string]*FunctionValue)
		for _, method := range s.Methods {
			methods[method.Name] = &FunctionValue{
				Parameters:  method.Parameters,
				ReturnTypes: method.ReturnTypes,
				Body:        method.Body,
				Closure:     env,
			}
		}
		env.Set(s.Name, &ClassValue{
			Name:    s.Name,
			Fields:  fieldTypes,
			Methods: methods,
		})
	case *ast.ChestaStatement:
		panicked := false
		func() {
			defer func() {
				if recover() != nil {
					panicked = true
				}
			}()
			evalBlock(s.TryBody, env)
		}()
		if panicked {
			evalBlock(s.CatchBody, env)
		}
	case *ast.AnaaStatement:
		// module imports not implemented yet
	case *ast.ExpressionStatement:
		evalExpression(s.Expression, env)
	}
	return nil
}

func evalBlock(stmts []ast.Statement, env *Environment) interface{} {
	for _, stmt := range stmts {
		result := evalStatement(stmt, env)
		switch result.(type) {
		case breakSignal, continueSignal, returnSignal:
			return result
		}
	}
	return nil
}

func evalExpression(expr ast.Expression, env *Environment) interface{} {
	if expr == nil {
		return nil
	}
	switch e := expr.(type) {
	case *ast.IntegerLiteral:
		return e.Value
	case *ast.FloatLiteral:
		return e.Value
	case *ast.StringLiteral:
		return e.Value
	case *ast.CharLiteral:
		return e.Value
	case *ast.BooleanLiteral:
		return e.Value
	case *ast.NullLiteral:
		return nil
	case *ast.Identifier:
		val, ok := env.Get(e.Name)
		if !ok {
			fmt.Printf("[OdLang Error] undefined variable: %s\n", e.Name)
			return nil
		}
		return val
	case *ast.PrefixExpression:
		right := evalExpression(e.Right, env)
		switch e.Operator {
		case "!":
			if b, ok := right.(bool); ok {
				return !b
			}
		case "-":
			if i, ok := right.(int64); ok {
				return -i
			}
			if f, ok := right.(float64); ok {
				return -f
			}
		}
	case *ast.InfixExpression:
		left := evalExpression(e.Left, env)
		right := evalExpression(e.Right, env)
		return evalInfix(e.Operator, left, right)
	case *ast.CallExpression:
		return evalCall(e, env)
	case *ast.MemberExpression:
		return evalMember(e, env)
	case *ast.IndexExpression:
		left := evalExpression(e.Left, env)
		index := evalExpression(e.Index, env)
		switch container := left.(type) {
		case []interface{}:
			i := toInt64(index)
			if i < 0 || int(i) >= len(container) {
				return nil
			}
			return container[i]
		case map[string]interface{}:
			key := fmt.Sprintf("%v", index)
			return container[key]
		}
		return nil
	case *ast.TypeCastExpression:
		return evalTypeCast(e, env)
	}
	return nil
}

func evalInfix(operator string, left, right interface{}) interface{} {
	if operator == "==" {
		return reflect.DeepEqual(left, right)
	}
	if operator == "!=" {
		return !reflect.DeepEqual(left, right)
	}

	if ls, lok := left.(string); lok {
		if rs, rok := right.(string); rok && operator == "+" {
			return ls + rs
		}
	}

	if lb, lok := left.(bool); lok {
		if rb, rok := right.(bool); rok {
			switch operator {
			case "aau":
				return lb || rb
			case "sahita":
				return lb && rb
			}
		}
	}

	lf, leftIsFloat := toFloat64(left)
	rf, rightIsFloat := toFloat64(right)
	if leftIsFloat || rightIsFloat {
		return evalFloatInfix(operator, lf, rf)
	}

	li, lok := left.(int64)
	ri, rok := right.(int64)
	if lok && rok {
		return evalIntInfix(operator, li, ri)
	}

	return nil
}

func evalIntInfix(operator string, left, right int64) interface{} {
	switch operator {
	case "+":
		return left + right
	case "-":
		return left - right
	case "*":
		return left * right
	case "/":
		if right == 0 {
			panic("division by zero")
		}
		return left / right
	case "**":
		return int64(math.Pow(float64(left), float64(right)))
	case "<":
		return left < right
	case ">":
		return left > right
	case "<=":
		return left <= right
	case ">=":
		return left >= right
	}
	return nil
}

func evalFloatInfix(operator string, left, right float64) interface{} {
	switch operator {
	case "+":
		return left + right
	case "-":
		return left - right
	case "*":
		return left * right
	case "/":
		if right == 0.0 {
			return math.NaN()
		}
		return left / right
	case "**":
		return math.Pow(left, right)
	case "<":
		return left < right
	case ">":
		return left > right
	case "<=":
		return left <= right
	case ">=":
		return left >= right
	}
	return nil
}

func evalCall(expr *ast.CallExpression, env *Environment) interface{} {
	if expr.Receiver != nil {
		return evalMethodCall(expr, env)
	}

	val, ok := env.Get(expr.Function)
	if !ok {
		fmt.Printf("[OdLang Error] unknown function: %s\n", expr.Function)
		return nil
	}

	if class, ok := val.(*ClassValue); ok {
		return newInstance(class)
	}

	fn, ok := val.(*FunctionValue)
	if !ok {
		fmt.Printf("[OdLang Error] unknown function: %s\n", expr.Function)
		return nil
	}

	args := make([]interface{}, 0, len(expr.Arguments))
	for _, arg := range expr.Arguments {
		args = append(args, evalExpression(arg, env))
	}
	return evalCallFunction(fn, args)
}

func evalMethodCall(expr *ast.CallExpression, env *Environment) interface{} {
	receiver := evalExpression(expr.Receiver, env)
	inst, ok := receiver.(*InstanceValue)
	if !ok {
		fmt.Printf("[OdLang Error] %s is not an object\n", expr.Function)
		return nil
	}

	method, ok := inst.Class.Methods[expr.Function]
	if !ok {
		fmt.Printf("[OdLang Error] unknown method: %s\n", expr.Function)
		return nil
	}

	args := make([]interface{}, 0, len(expr.Arguments))
	for _, arg := range expr.Arguments {
		args = append(args, evalExpression(arg, env))
	}
	return evalMethodFunction(inst, method, args)
}

func evalMember(expr *ast.MemberExpression, env *Environment) interface{} {
	object := evalExpression(expr.Object, env)
	inst, ok := object.(*InstanceValue)
	if !ok {
		fmt.Printf("[OdLang Error] cannot read field %s on non-object\n", expr.Member)
		return nil
	}
	val, ok := inst.Fields[expr.Member]
	if !ok {
		fmt.Printf("[OdLang Error] unknown field: %s\n", expr.Member)
		return nil
	}
	return val
}

func newInstance(class *ClassValue) *InstanceValue {
	inst := &InstanceValue{
		Class:  class,
		Fields: make(map[string]interface{}),
	}
	for name, typeName := range class.Fields {
		inst.Fields[name] = defaultValueForType(typeName)
	}
	return inst
}

func evalMethodFunction(inst *InstanceValue, fn *FunctionValue, args []interface{}) interface{} {
	callEnv := NewEnclosedEnvironment(fn.Closure)
	for name, typeName := range inst.Class.Fields {
		callEnv.Set(name, inst.Fields[name])
		callEnv.SetType(name, typeName)
	}
	for i, param := range fn.Parameters {
		var arg interface{}
		if i < len(args) {
			arg = args[i]
		}
		callEnv.Set(param.Name, arg)
		callEnv.SetType(param.Name, param.TypeName)
	}

	var returnValues []interface{}
	for _, stmt := range fn.Body {
		result := evalStatement(stmt, callEnv)
		if sig, ok := result.(returnSignal); ok {
			returnValues = sig.values
			break
		}
	}

	for name := range inst.Class.Fields {
		if val, ok := callEnv.Get(name); ok {
			inst.Fields[name] = val
		}
	}

	if len(fn.ReturnTypes) == 0 {
		return nil
	}
	if len(fn.ReturnTypes) == 1 {
		if len(returnValues) > 0 {
			return returnValues[0]
		}
		return nil
	}
	return returnValues
}

func evalCallFunction(fn *FunctionValue, args []interface{}) interface{} {
	callEnv := NewEnclosedEnvironment(fn.Closure)
	for i, param := range fn.Parameters {
		var arg interface{}
		if i < len(args) {
			arg = args[i]
		}
		callEnv.Set(param.Name, arg)
		callEnv.SetType(param.Name, param.TypeName)
	}

	var returnValues []interface{}
	for _, stmt := range fn.Body {
		result := evalStatement(stmt, callEnv)
		if sig, ok := result.(returnSignal); ok {
			returnValues = sig.values
			break
		}
	}

	if len(fn.ReturnTypes) == 0 {
		return nil
	}
	if len(fn.ReturnTypes) == 1 {
		if len(returnValues) > 0 {
			return returnValues[0]
		}
		return nil
	}
	return returnValues
}

func evalTypeCast(expr *ast.TypeCastExpression, env *Environment) interface{} {
	value := evalExpression(expr.Value, env)
	switch expr.TargetType {
	case "sabda":
		return fmt.Sprintf("%v", value)
	case "sankhya":
		switch v := value.(type) {
		case string:
			parsed, err := strconv.ParseInt(v, 10, 64)
			if err == nil {
				return parsed
			}
		case float64:
			return int64(v)
		case int64:
			return v
		}
	case "dasmik":
		switch v := value.(type) {
		case string:
			parsed, err := strconv.ParseFloat(v, 64)
			if err == nil {
				return parsed
			}
		case int64:
			return float64(v)
		case float64:
			return v
		}
	case "satya":
		return isTruthy(value)
	case "akshara":
		switch v := value.(type) {
		case string:
			if len(v) > 0 {
				return rune(v[0])
			}
			return rune(0)
		case int64:
			return rune(v)
		case rune:
			return v
		}
	}
	return value
}

func assignIndex(left ast.Expression, value interface{}, env *Environment) {
	idx, ok := left.(*ast.IndexExpression)
	if !ok {
		return
	}
	container := evalExpression(idx.Left, env)
	indexVal := evalExpression(idx.Index, env)
	switch c := container.(type) {
	case []interface{}:
		i := toInt64(indexVal)
		if i < 0 || int(i) >= len(c) {
			return
		}
		c[i] = value
		if ident, ok := idx.Left.(*ast.Identifier); ok {
			env.Set(ident.Name, c)
		}
	case map[string]interface{}:
		key := fmt.Sprintf("%v", indexVal)
		c[key] = value
		if ident, ok := idx.Left.(*ast.Identifier); ok {
			env.Set(ident.Name, c)
		}
	}
}

func isTruthy(value interface{}) bool {
	if value == nil {
		return false
	}
	switch v := value.(type) {
	case bool:
		return v
	case int64:
		return v != 0
	case float64:
		return v != 0.0
	case string:
		return len(v) > 0
	default:
		return true
	}
}

func defaultValueForType(typeName string) interface{} {
	switch typeName {
	case "sankhya":
		return int64(0)
	case "dasmik":
		return float64(0)
	case "sabda":
		return ""
	case "akshara":
		return rune(0)
	case "satya":
		return false
	case "krama":
		return []interface{}{}
	case "mana":
		return map[string]interface{}{}
	default:
		return nil
	}
}

func toInt64(value interface{}) int64 {
	switch v := value.(type) {
	case int64:
		return v
	case float64:
		return int64(v)
	case int:
		return int64(v)
	}
	return 0
}

func toFloat64(value interface{}) (float64, bool) {
	switch v := value.(type) {
	case float64:
		return v, true
	case int64:
		return float64(v), true
	case int:
		return float64(v), true
	}
	return 0, false
}

func formatValue(value interface{}) string {
	if value == nil {
		return "khali"
	}
	switch v := value.(type) {
	case int64:
		return fmt.Sprintf("%d", v)
	case float64:
		return fmt.Sprintf("%g", v)
	case string:
		return v
	case rune:
		return string(v)
	case bool:
		if v {
			return "han"
		}
		return "na"
	case []interface{}:
		parts := make([]string, len(v))
		for i, item := range v {
			parts[i] = formatValue(item)
		}
		return "[" + strings.Join(parts, ", ") + "]"
	case map[string]interface{}:
		parts := []string{}
		for k, item := range v {
			parts = append(parts, k+": "+formatValue(item))
		}
		return "{" + strings.Join(parts, ", ") + "}"
	default:
		return fmt.Sprintf("%v", v)
	}
}
