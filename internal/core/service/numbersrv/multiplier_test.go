package numbersrv

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetTypeIfIsAnMultiplier(t *testing.T) {

	//
	// Tests Cases
	//

	type args struct {
		number int
	}

	type want struct {
		result string
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "Should get a type 1 response",
			args: args{
				number: multiplierThree,
			},
			want: want{
				result: "Type 1",
			},
		},
		{
			name: "Should get a type 2 response",
			args: args{
				number: multiplierFive,
			},
			want: want{
				result: "Type 2",
			},
		},
		{
			name: "Should get a type 3 response",
			args: args{
				number: multiplierThree * multiplierFive,
			},
			want: want{
				result: "Type 3",
			},
		},
		{
			name: "Should get a number response",
			args: args{
				number: 1,
			},
			want: want{
				result: "1",
			},
		},
	}

	//
	// Tests Runner
	//

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			//
			// Setup
			//

			args, want := tt.args, tt.want

			//
			// Execute
			//

			result := getTypeIfIsAnMultiplier(args.number)

			//
			// Verify
			//

			assert.Equal(t, want.result, result)
		})
	}
}
