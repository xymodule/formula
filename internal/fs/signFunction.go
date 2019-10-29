package fs

import (
	"github.com/xymodule/formula/opt"
	"math"
	"reflect"
)

type SignFunction struct {
}

func (*SignFunction) Name() string {
	return "sign"
}

func (f *SignFunction) Evaluate(context *opt.FormulaContext, args ...*opt.LogicalExpression) (*opt.Argument, error) {
	v, err := ParseFloat(f.Name(), context, args...)
	if err != nil {
		return nil, err
	}
	
	return opt.NewArgumentWithType(math.Signbit(v), reflect.Bool), nil
}

func NewSignFunction() *SignFunction {
	return &SignFunction{}
}
