package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashPassword(t *testing.T) {
	cases := []struct {
		testName string
	}{
		{
			testName: "1. Positive test",
		},
	}

	for _, c := range cases {
		t.Logf(CurrentlyTesting, c.testName)
		password, _ := HashPassword("123")
		assert.NotNil(t, password)
	}
}

func TestCheckPasswordHash(t *testing.T) {
	cases := []struct {
		testName string
	}{
		{
			testName: "1. Positive test",
		},
	}

	for _, c := range cases {
		t.Logf(CurrentlyTesting, c.testName)
		isValid := CheckPasswordHash("123abc", "$2a$14$7x0URT8OS7iSJug7az08IeVwRPHsK6imOw.E2aqJ6u.C2Y291Z0fW")
		assert.NotNil(t, isValid)
	}
}

func TestCreateToken(t *testing.T) {
	cases := []struct {
		testName string
	}{
		{
			testName: "1. Positive test",
		},
	}

	for _, c := range cases {
		t.Logf(CurrentlyTesting, c.testName)
		token, _ := CreateToken("test", []byte(`123`))
		assert.NotNil(t, token)
	}
}
