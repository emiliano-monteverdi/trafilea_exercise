package numbersrv_test

import (
	"github.com/Trafilea/internal/core/ports"
	"github.com/Trafilea/internal/core/service/numbersrv"
	"github.com/Trafilea/internal/repositories/numberrepo"
	"github.com/Trafilea/mocks"
	"github.com/Trafilea/pkg/errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

type dependencies struct {
	numberRepository ports.NumberRepository
}

func makeDependencies() dependencies {
	numberRepository := numberrepo.New(make(map[int]string))

	return dependencies{
		numberRepository: numberRepository,
	}
}

func TestCreate(t *testing.T) {

	//
	// Setup
	//

	number := mocks.Number()

	//
	// Tests Cases
	//

	type args struct {
		number int
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
				number: number.Value,
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

			service := numbersrv.New(d.numberRepository)

			//
			// Execute
			//

			service.Create(args.number)

			//
			// Verify
			//

			result, _ := service.Get(args.number)

			assert.Equal(t, number.Type, result)
		})
	}
}

func TestGet(t *testing.T) {

	//
	// Setup
	//

	number := mocks.Number()

	//
	// Errors
	//

	getNumberRepositoryCauseError := errors.New(http.StatusInternalServerError, "")
	getNumberRepositoryError := errors.New(http.StatusNotFound, "getting the number value from the NumberRepository has failed", errors.NewCauseError(getNumberRepositoryCauseError))

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
			name: "Should get a number successfully",
			mock: func(d dependencies, args args, want want) {
				d.numberRepository.Save(number)
			},
			args: args{
				number: number.Value,
			},
			want: want{
				result: number.Type,
			},
		},
		{
			name: "Should fail when tries to get the number from the NumberRepository",
			mock: func(d dependencies, args args, want want) {},
			args: args{
				number: number.Value,
			},
			want: want{
				err: getNumberRepositoryError,
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

			service := numbersrv.New(d.numberRepository)

			//
			// Execute
			//

			result, err := service.Get(args.number)

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

func TestBulkGet(t *testing.T) {

	//
	// Setup
	//

	numbersArray := mocks.NumbersArray()
	numbersArrayAsString := mocks.NumbersArrayAsString("int")

	//
	// Errors
	//

	numberNotFoundError := errors.New(http.StatusNotFound, "there's nothing stored in the NumberRepository")

	//
	// Tests Cases
	//

	type args struct{}

	type want struct {
		result []string
		err    errors.Error
	}

	tests := []struct {
		name string
		mock func(d dependencies, args args, want want)
		args args
		want want
	}{
		{
			name: "Should get a numbers array successfully",
			mock: func(d dependencies, args args, want want) {
				for _, number := range numbersArray {
					d.numberRepository.Save(number)
				}
			},
			args: args{},
			want: want{
				result: numbersArrayAsString,
			},
		},
		{
			name: "Should fail due to empty storage error",
			mock: func(d dependencies, args args, want want) {},
			args: args{},
			want: want{
				err: numberNotFoundError,
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

			service := numbersrv.New(d.numberRepository)

			//
			// Execute
			//

			result, err := service.BulkGet()

			//
			// Verify
			//

			if err != nil || want.err != nil {
				assert.Equal(t, want.err.StatusCode(), err.StatusCode())
				assert.Equal(t, want.err.Code(), err.Code())
				assert.Equal(t, want.err.Message(), err.Message())

			} else {
				assert.Equal(t, want.err, err)

				counter := 0
				for _, wantedResult := range want.result {
					for _, actualResult := range result {
						if actualResult == wantedResult {
							counter++
						}
					}
				}

				assert.Equal(t, 15, counter)
			}

		})
	}
}

func TestBulkGetTypes(t *testing.T) {

	//
	// Setup
	//

	numbersArray := mocks.NumbersArray()
	numbersArrayAsString := mocks.NumbersArrayAsString("type")

	//
	// Errors
	//

	numberNotFoundError := errors.New(http.StatusNotFound, "there's nothing stored in the NumberRepository")

	//
	// Tests Cases
	//

	type args struct{}

	type want struct {
		result []string
		err    errors.Error
	}

	tests := []struct {
		name string
		mock func(d dependencies, args args, want want)
		args args
		want want
	}{
		{
			name: "Should get a numbers array successfully",
			mock: func(d dependencies, args args, want want) {
				for _, number := range numbersArray {
					d.numberRepository.Save(number)
				}
			},
			args: args{},
			want: want{
				result: numbersArrayAsString,
			},
		},
		{
			name: "Should fail due to empty storage error",
			mock: func(d dependencies, args args, want want) {},
			args: args{},
			want: want{
				err: numberNotFoundError,
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

			service := numbersrv.New(d.numberRepository)

			//
			// Execute
			//

			result, err := service.BulkGetTypes()

			//
			// Verify
			//

			if err != nil || want.err != nil {
				assert.Equal(t, want.err.StatusCode(), err.StatusCode())
				assert.Equal(t, want.err.Code(), err.Code())
				assert.Equal(t, want.err.Message(), err.Message())

			} else {
				assert.Equal(t, want.err, err)

				wantedMap := make(map[string]int)
				for _, value := range numbersArrayAsString {
					wantedMap[value]++
				}

				actualMap := make(map[string]int)
				for _, value := range result {
					actualMap[value]++
				}

				counter := 0
				for key1, value1 := range wantedMap {
					for key2, value2 := range actualMap {
						if key1 == key2 && value1 == value2 {
							counter++
						}
					}
				}

				assert.Equal(t, 11, counter)
			}
		})
	}
}
