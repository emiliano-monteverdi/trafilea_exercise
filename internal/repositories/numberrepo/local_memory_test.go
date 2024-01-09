package numberrepo_test

import (
	"github.com/Trafilea/internal/core/domain"
	"github.com/Trafilea/internal/repositories/numberrepo"
	"github.com/Trafilea/mocks"
	"github.com/Trafilea/pkg/errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

type dependencies struct {
	storageClient map[int]string
}

func makeDependencies() dependencies {
	return dependencies{
		storageClient: make(map[int]string),
	}
}

func TestSave(t *testing.T) {

	//
	// Setup
	//

	number := mocks.Number()

	//
	// Tests Cases
	//

	type args struct {
		number domain.Number
	}

	type want struct{}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "Should create a number successfully",
			args: args{
				number: number,
			},
			want: want{},
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

			args := tt.args

			d := makeDependencies()

			repository := numberrepo.New(d.storageClient)

			//
			// Execute
			//

			repository.Save(args.number)

			//
			// Verify
			//

			result, _ := repository.Find(args.number.Value)

			assert.Equal(t, args.number.Type, result)
		})
	}

}

func TestFind(t *testing.T) {

	//
	// Setup
	//

	number := mocks.Number()

	//
	// Errors
	//

	notFoundError := errors.New(http.StatusNotFound, "the number couldn't be found in the local storage")

	//
	// Tests Cases
	//

	type args struct {
		number int
	}

	type want struct {
		result string
		err    errors.Error
	}

	tests := []struct {
		name string
		mock func(d dependencies, args args, want want)
		args args
		want want
	}{
		{
			name: "Should return a number successfully",
			mock: func(d dependencies, args args, want want) {
				d.storageClient[args.number] = want.result
			},
			args: args{
				number: number.Value,
			},
			want: want{
				result: number.Type,
			},
		},
		{
			name: "Should fail due to not found error",
			mock: func(d dependencies, args args, want want) {
			},
			args: args{
				number: number.Value,
			},
			want: want{
				result: "",
				err:    notFoundError,
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

			d := makeDependencies()

			tt.mock(d, args, want)

			repository := numberrepo.New(d.storageClient)

			//
			// Execute
			//

			result, err := repository.Find(args.number)

			//
			// Verify
			//

			if err != nil || want.err != nil {
				assert.Equal(t, want.err.StatusCode(), err.StatusCode())
				assert.Equal(t, want.err.Code(), err.Code())
				assert.Equal(t, want.err.Message(), err.Message())

			} else {
				assert.Equal(t, want.err, err)
				assert.Equal(t, want.result, result)
			}
		})
	}
}

func TestFindAll(t *testing.T) {

	//
	// Setup
	//

	numbersArray := mocks.NumbersArray()

	//
	// Tests Cases
	//

	type args struct{}

	type want struct {
		result []domain.Number
	}

	tests := []struct {
		name string
		mock func(d dependencies, args args, want want)
		args args
		want want
	}{
		{
			name: "Should return a numbers array successfully",
			mock: func(d dependencies, args args, want want) {
				for _, number := range numbersArray {
					d.storageClient[number.Value] = number.Type
				}
			},
			args: args{},
			want: want{
				result: numbersArray,
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

			d := makeDependencies()

			tt.mock(d, args, want)

			repository := numberrepo.New(d.storageClient)

			//
			// Execute
			//

			result := repository.FindAll()

			//
			// Verify
			//

			counter := 0
			for _, wantedResult := range want.result {
				for _, actualResult := range result {
					if actualResult.Value == wantedResult.Value && actualResult.Type == wantedResult.Type {
						counter++
					}
				}
			}

			assert.Equal(t, 15, counter)
		})
	}
}
