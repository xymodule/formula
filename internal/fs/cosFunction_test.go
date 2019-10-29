package fs

import (
	"fmt"
	"github.com/xymodule/formula/internal/exp"
	"github.com/xymodule/formula/opt"
	"math"
	"strconv"
	"testing"
)

func TestCosFunction_Evaluate(t *testing.T) {
	tests := []struct {
		args []string
	}{
		{[]string{"1", "2", "3", "30", "0.5"}},
	}

	f := NewCosFunction()
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.args), func(t *testing.T) {
			for i := 0; i < len(tt.args); i++ {
				var logicalExpression = *exp.NewFloatExpression(tt.args[i])

				result, err := f.Evaluate(nil, []*opt.LogicalExpression{&logicalExpression}...)
				if err != nil {
					t.Fatal(err)
				}

				v, err := result.Float64()
				if err != nil {
					t.Fatal(err)
				}

				if v1, _ := strconv.ParseFloat(tt.args[i], 64); math.Cos(v1) != v {
					t.Fatal()
				}
			}
		})
	}
}
