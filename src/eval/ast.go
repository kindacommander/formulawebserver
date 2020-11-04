package eval

// Expr defines arithmetic expression.
type Expr interface {
	// Eval returns current Expr value in env environment.
	Eval(env Env) float64
	// Check reports errors in this Expr and adds its Vars to the set.
	Check(vars map[Var]bool) error
}

// Var defines variable (x for example).
type Var string

// literal defines a numeric constant.
type literal float64

// unary defines an expression with the unary operator, such as -x.
type unary struct {
	op rune // '+' or '-'
	x  Expr
}

// binary defines an expression with a binary operator, such as x + y.
type binary struct {
	op   rune // '+', '-', '*' or '/'
	x, y Expr
}

// call defines a function call expression, such as sin (x).
type call struct {
	fn   string // pow, sin or sqrt
	args []Expr
}
