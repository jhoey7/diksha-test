package services

import (
	"testing"
)

type allValidationStruct struct {
	Required     string `validate:"required"`
	Email        string `validate:"email"`
	GreaterThan5 int    `validate:"gte=5"`
	LessThan5    int    `validate:"lte=5"`
	Len5         string `validate:"len=5"`
	Numeric      string `validate:"numeric"`
}

const ExpectedRespCode = "Expected resp code to be %d but it was %d"

var (
	positiveValidation = []byte(`{
		"Required":     "asdf",
		"Email":        "email@gmail.com",
		"GreaterThan5": 6,
		"LessThan5":    4,
		"Len5":         "12345",
		"Numeric":      "12345"
	}`)

	negativeReqUnmarshal = []byte(`{`)

	negativeAllValidation = []byte(`{
		"Required":     "",
		"Email":        "email",
		"GreaterThan5": 4,
		"LessThan5":    6,
		"Len5":         "123",
		"Numeric":      "123abc"
	}`)
)

func TestConvertRequest(t *testing.T) {
	cases := []struct {
		testName         string
		req              []byte
		expectedRespCode int
	}{
		{
			testName:         "1. Positive Test Case",
			req:              positiveValidation,
			expectedRespCode: 200,
		},
		{
			testName:         "2. Negative Test Case: failed unmarshal",
			req:              negativeReqUnmarshal,
			expectedRespCode: 400,
		},
		{
			testName:         "3. Negative Test Case: failed all validation",
			req:              negativeAllValidation,
			expectedRespCode: 400,
		},
	}

	var request allValidationStruct
	for _, c := range cases {
		t.Run(c.testName, func(t *testing.T) {
			resp := ConvertRequest(c.req, &request, 123)
			if resp.Code != c.expectedRespCode {
				t.Errorf(ExpectedRespCode, c.expectedRespCode, resp.Code)
			}
		})
	}
}
