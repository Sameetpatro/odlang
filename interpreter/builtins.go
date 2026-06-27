package interpreter

// registerBuiltins wires standard library into every environment.
// OdLang programs do not use import syntax — I/O, type casts, and
// collection constructors are handled directly by the interpreter.
func registerBuiltins(env *Environment) {
	_ = env
}
