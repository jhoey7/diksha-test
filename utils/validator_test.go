package utils

import (
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRequiredIfFieldEqual(t *testing.T) {
	type SomeRequest struct {
		Instrument         string
		OutletExternalName string `validate:"required_if_field_equal=Instrument:PQR"`
	}
	var nonPointerCases = []struct {
		testName           string
		someRequest        SomeRequest
		somePointerRequest *SomeRequest
		result             []string
	}{
		{
			testName: "1. Positive Test: Example Instrument non PQR, Name is not empty",
			someRequest: SomeRequest{
				Instrument:         "Beacon",
				OutletExternalName: "Name is not empty",
			},
			somePointerRequest: &SomeRequest{
				Instrument:         "Beacon",
				OutletExternalName: "Name is not empty",
			},
			result: []string(nil),
		},
		{
			testName: "2. Positive Test: Example Instrument non PQR, Name is empty",
			someRequest: SomeRequest{
				Instrument:         "Beacon",
				OutletExternalName: "",
			},
			somePointerRequest: &SomeRequest{
				Instrument:         "Beacon",
				OutletExternalName: "",
			},
			result: []string(nil),
		},
		{
			testName: "3. Positive Test: Example Instrument PQR, Name is not empty",
			someRequest: SomeRequest{
				Instrument:         "PQR",
				OutletExternalName: "Name is not empty",
			},
			somePointerRequest: &SomeRequest{
				Instrument:         "PQR",
				OutletExternalName: "Name is not empty",
			},
			result: []string(nil),
		},
		{
			testName: "4. Positive Test: Example Instrument is empty, Name is empty",
			someRequest: SomeRequest{
				Instrument:         "",
				OutletExternalName: "",
			},
			somePointerRequest: &SomeRequest{
				Instrument:         "",
				OutletExternalName: "",
			},
			result: []string(nil),
		},
		{
			testName: "5. Negative Test: Example Instrument PQR, Name is empty",
			someRequest: SomeRequest{
				Instrument:         "PQR",
				OutletExternalName: "",
			},
			somePointerRequest: &SomeRequest{
				Instrument:         "PQR",
				OutletExternalName: "",
			},
			result: []string{"OutletExternalName"},
		},
	}

	validate := validator.New()
	validate.RegisterValidation("required_if_field_equal", RequiredIfFieldEqual)

	for _, c := range nonPointerCases {
		t.Logf(CurrentlyTesting, c.testName)
		errors := validate.Struct(c.someRequest)
		var fields []string
		if errors != nil {
			for _, err := range errors.(validator.ValidationErrors) {
				fields = append(fields, err.Field())
			}
		}
		assert.Equal(t, c.result, fields)

		errorPointerReqs := validate.Struct(c.somePointerRequest)
		var fieldPointers []string
		if errorPointerReqs != nil {
			for _, err := range errors.(validator.ValidationErrors) {
				fieldPointers = append(fieldPointers, err.Field())
			}
		}
		assert.Equal(t, c.result, fieldPointers)
	}
}
