package numberhdl_test

import (
	"encoding/json"
	"github.com/Trafilea/cmd/dependencies"
	"github.com/Trafilea/cmd/httpserver"
	"github.com/Trafilea/internal/core/ports"
	"github.com/Trafilea/internal/core/service/numbersrv"
	"github.com/Trafilea/internal/handlers/numberhdl"
	"github.com/Trafilea/internal/repositories/numberrepo"
	"github.com/Trafilea/mocks"
	"github.com/Trafilea/pkg/errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type numberDependencies struct {
	numberRepository ports.NumberRepository
	numberService    ports.NumberService
	numberHandler    *numberhdl.Handler
}

func makeNumberDependencies() numberDependencies {
	localStorage := make(map[int]string)

	numberRepository := numberrepo.New(localStorage)

	numberService := numbersrv.New(numberRepository)

	numberHandler := numberhdl.New(numberService)

	return numberDependencies{
		numberRepository: numberRepository,
		numberService:    numberService,
		numberHandler:    numberHandler,
	}
}

func TestCreate(t *testing.T) {

	//
	// Setup
	//

	number := numberhdl.CreateRequest{Number: 1}
	numberBody, _ := json.Marshal(number)

	incorrectNumberBody := []byte("")

	resp := mocks.ResponseCreate()

	//
	// Errors
	//

	unmarshalBodyError := errors.New(http.StatusBadRequest, "unmarshal error")

	//
	// Tests Cases
	//

	type args struct {
		url  string
		body []byte
	}

	type want struct {
		status int
		result interface{}
		err    error
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "Should create the number successfully",
			args: args{
				url:  "/numbers/",
				body: numberBody,
			},
			want: want{
				status: 200,
				result: resp,
			},
		},
		{
			name: "Should fail due to unmarshal body error",
			args: args{
				url:  "/numbers/",
				body: incorrectNumberBody,
			},
			want: want{
				status: 400,
				err:    unmarshalBodyError,
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

			m := makeNumberDependencies()

			//
			// Execute
			//

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", args.url, strings.NewReader(string(args.body)))

			//
			// Dependency Injection
			//

			d := dependencies.Definition{
				NumberRepository: m.numberRepository,
				NumberService:    m.numberService,
				NumberHandler:    m.numberHandler,
			}

			httpserver.SetupRouter(d).ServeHTTP(w, req)

			//
			// Verify
			//

			assert.Equal(t, want.status, w.Code)
			if w.Code >= 200 && w.Code <= 299 {
				want, _ := json.Marshal(want.result)
				assert.Equal(t, string(want), removeBreakLine(w.Body.String()))

			} else {
				want, _ := json.Marshal(want.err)
				assert.Equal(t, string(want), removeBreakLine(w.Body.String()))
			}
		})
	}
}

func TestGet(t *testing.T) {

	//
	// Setup
	//

	number := mocks.Number()

	resp := mocks.ResponseGet()

	//
	// Errors
	//

	invalidNumberError := errors.New(http.StatusBadRequest, "invalid number error")

	numberNotFoundErrorCause := errors.New(http.StatusNotFound, "the number couldn't be found in the local storage")
	numberNotFoundError := errors.New(http.StatusNotFound, "getting the number value from the NumberRepository has failed", errors.NewCauseError(numberNotFoundErrorCause))

	//
	// Tests Cases
	//

	type args struct {
		url string
	}

	type want struct {
		status int
		result interface{}
		err    error
	}

	tests := []struct {
		name string
		mock func(d numberDependencies, args args, want want)
		args args
		want want
	}{
		{
			name: "Should get the number value successfully",
			mock: func(d numberDependencies, args args, want want) {
				d.numberRepository.Save(number)
			},
			args: args{
				url: "/numbers/1",
			},
			want: want{
				status: 200,
				result: resp,
			},
		},
		{
			name: "Should fail due to invalid number error",
			mock: func(d numberDependencies, args args, want want) {},
			args: args{
				url: "/numbers/qwerty",
			},
			want: want{
				status: 400,
				err:    invalidNumberError,
			},
		},
		{
			name: "Should fail due to card not found error",
			mock: func(d numberDependencies, args args, want want) {},
			args: args{
				url: "/numbers/1",
			},
			want: want{
				status: 404,
				err:    numberNotFoundError,
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

			m := makeNumberDependencies()

			tt.mock(m, args, want)

			//
			// Execute
			//

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", args.url, nil)

			//
			// Dependency Injection
			//

			d := dependencies.Definition{
				NumberRepository: m.numberRepository,
				NumberService:    m.numberService,
				NumberHandler:    m.numberHandler,
			}

			httpserver.SetupRouter(d).ServeHTTP(w, req)

			//
			// Verify
			//

			assert.Equal(t, want.status, w.Code)
			if w.Code >= 200 && w.Code <= 299 {
				want, _ := json.Marshal(want.result)
				assert.Equal(t, string(want), removeBreakLine(w.Body.String()))

			} else {
				want, _ := json.Marshal(want.err)
				assert.Equal(t, string(want), removeBreakLine(w.Body.String()))
			}
		})
	}
}

func TestBulkGetValue(t *testing.T) {

	//
	// Setup
	//

	numbersArray := mocks.NumbersArray()

	resp := mocks.ResponseBulkGetValue()

	//
	// Errors
	//

	numberNotFoundError := errors.New(http.StatusNotFound, "there's nothing stored in the NumberRepository")

	//
	// Tests Cases
	//

	type args struct {
		url string
	}

	type want struct {
		status int
		result interface{}
		err    error
	}

	tests := []struct {
		name string
		mock func(d numberDependencies, args args, want want)
		args args
		want want
	}{
		{
			name: "Should get the number array successfully",
			mock: func(d numberDependencies, args args, want want) {
				for _, number := range numbersArray {
					d.numberRepository.Save(number)
				}
			},
			args: args{
				url: "/numbers/bulk/value",
			},
			want: want{
				status: 200,
				result: resp,
			},
		},
		{
			name: "Should fail due to empty storage error",
			mock: func(d numberDependencies, args args, want want) {},
			args: args{
				url: "/numbers/bulk/value",
			},
			want: want{
				status: 404,
				err:    numberNotFoundError,
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

			m := makeNumberDependencies()

			tt.mock(m, args, want)

			//
			// Execute
			//

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", args.url, nil)

			//
			// Dependency Injection
			//

			d := dependencies.Definition{
				NumberRepository: m.numberRepository,
				NumberService:    m.numberService,
				NumberHandler:    m.numberHandler,
			}

			httpserver.SetupRouter(d).ServeHTTP(w, req)

			//
			// Verify
			//

			assert.Equal(t, want.status, w.Code)
			if w.Code >= 200 && w.Code <= 299 {
				want, _ := json.Marshal(want.result)

				wantedResponse := numberhdl.BulkGetValueResponse{}
				actualResponse := numberhdl.BulkGetValueResponse{}

				_ = json.Unmarshal(want, &wantedResponse)
				_ = json.Unmarshal(w.Body.Bytes(), &actualResponse)

				counter := 0
				for _, wantedValue := range wantedResponse.Value {
					for _, actualValue := range actualResponse.Value {
						if actualValue == wantedValue {
							counter++
						}
					}
				}

				assert.Equal(t, 15, counter)

			} else {
				want, _ := json.Marshal(want.err)
				assert.Equal(t, string(want), removeBreakLine(w.Body.String()))
			}
		})
	}
}

func TestBulkGetType(t *testing.T) {

	//
	// Setup
	//

	numbersArray := mocks.NumbersArray()

	resp := mocks.ResponseBulkGetType()

	//
	// Errors
	//

	numberNotFoundError := errors.New(http.StatusNotFound, "there's nothing stored in the NumberRepository")

	//
	// Tests Cases
	//

	type args struct {
		url string
	}

	type want struct {
		status int
		result interface{}
		err    error
	}

	tests := []struct {
		name string
		mock func(d numberDependencies, args args, want want)
		args args
		want want
	}{
		{
			name: "Should get the number array successfully",
			mock: func(d numberDependencies, args args, want want) {
				for _, number := range numbersArray {
					d.numberRepository.Save(number)
				}
			},
			args: args{
				url: "/numbers/bulk/type",
			},
			want: want{
				status: 200,
				result: resp,
			},
		},
		{
			name: "Should fail due to empty storage error",
			mock: func(d numberDependencies, args args, want want) {},
			args: args{
				url: "/numbers/bulk/type",
			},
			want: want{
				status: 404,
				err:    numberNotFoundError,
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

			m := makeNumberDependencies()

			tt.mock(m, args, want)

			//
			// Execute
			//

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", args.url, nil)

			//
			// Dependency Injection
			//

			d := dependencies.Definition{
				NumberRepository: m.numberRepository,
				NumberService:    m.numberService,
				NumberHandler:    m.numberHandler,
			}

			httpserver.SetupRouter(d).ServeHTTP(w, req)

			//
			// Verify
			//

			assert.Equal(t, want.status, w.Code)
			if w.Code >= 200 && w.Code <= 299 {
				want, _ := json.Marshal(want.result)

				wantedResponse := numberhdl.BulkGetTypeResponse{}
				actualResponse := numberhdl.BulkGetTypeResponse{}

				_ = json.Unmarshal(want, &wantedResponse)
				_ = json.Unmarshal(w.Body.Bytes(), &actualResponse)

				wantedMap := make(map[string]int)
				for _, value := range wantedResponse.Type {
					wantedMap[value]++
				}

				actualMap := make(map[string]int)
				for _, value := range actualResponse.Type {
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

			} else {
				want, _ := json.Marshal(want.err)
				assert.Equal(t, string(want), removeBreakLine(w.Body.String()))
			}
		})
	}
}

// I am unaware of the reason why the Chi framework adds a newline at the end of each response. I've been
// investigating, but I couldn't find the cause. I would like to analyze this in more detail later, but for now,
// I added this code that simply removes the newline at the end. At MercadoLibre, we use Chi within a larger library,
// so I don't have direct contact with it. I'm confident that by analyzing it, I could discover the problem. Since I
// have limited time for this task, I decided to take this approach right now.
func removeBreakLine(text string) string {
	return strings.TrimSuffix(text, "\n")
}
