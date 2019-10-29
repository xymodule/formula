package fs

import (
	"formula/opt"
	"reflect"
)

type DivideFunction struct {
}

func (*DivideFunction) Name() string {
	return "/"
}

func (f *DivideFunction) Evaluate(context *opt.FormulaContext, args ...*opt.LogicalExpression) (*opt.Argument, error) {
	err := opt.MatchTwoArgument(f.Name(), args...)
	if err != nil {
		return nil, err
	}

	arg0, err := (*args[0]).Evaluate(context)
	if err != nil {
		return nil, err
	}

	arg1, err := (*args[1]).Evaluate(context)
	if err != nil {
		return nil, err
	}

	v1, err := arg1.Float64()
	if err != nil {
		return nil, err
	}

	if v1 == 0 {
		//fmt.Println("divide by zero")
		return opt.NewArgumentWithType(float64(0), reflect.Float64), nil
		//return nil, fmt.Errorf("divide by zero")
	}

	v0, err := arg0.Float64()
	if err != nil {
		return nil, err
	}
	return opt.NewArgumentWithType(v0/v1, reflect.Float64), nil
}

func NewDivideFunction() *DivideFunction {
	return &DivideFunction{}
}
