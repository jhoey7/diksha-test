package utils

import (
	"database/sql/driver"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	CurrentlyTesting = "Currently testing case: %s"
	ExpectedResponse = "Expected resp to be %v but it was %v"
)

var (
	uuid                  = "6ba7b810-9dad-11d1-80b4-00c04fd430c8"
	uuidVariantNCS, _     = FromString("34dc23469000.0d.00.00.7c.5f.00.00.00")
	uuidVariantRFC4122, _ = FromString("36626137-6238-3130-2d39-6461642d3131")
	uuidNil, _            = FromString("00000000-0000-0000-0000-000000000000")
	byteUUID              = []byte(uuid)
	shortUUID             = "6ba7b810-9dad-11"
	byteShortUUID         = []byte(shortUUID)
	nullUUID              = NullUUID{
		UUID:  NamespaceDNS,
		Valid: true,
	}
	invalidNullUUID = NullUUID{
		UUID:  NamespaceDNS,
		Valid: false,
	}
)

func TestInitClockSequence(t *testing.T) {
	cases := []struct {
		testName string
	}{
		{
			testName: "1. Positive test",
		},
	}

	for _, c := range cases {
		t.Logf(CurrentlyTesting, c.testName)
		initClockSequence()
	}
}

func TestInitHardwareAddr(t *testing.T) {
	cases := []struct {
		testName string
	}{
		{
			testName: "1. Positive test",
		},
	}

	for _, c := range cases {
		t.Logf(CurrentlyTesting, c.testName)
		initHardwareAddr()
	}
}

func TestInitStorage(t *testing.T) {
	cases := []struct {
		testName string
	}{
		{
			testName: "1. Positive test",
		},
	}

	for _, c := range cases {
		t.Logf(CurrentlyTesting, c.testName)
		initStorage()
	}
}

func TestUnixTimeFunc(t *testing.T) {
	cases := []struct {
		testName string
	}{
		{
			testName: "1. Positive test",
		},
	}

	for _, c := range cases {
		t.Logf(CurrentlyTesting, c.testName)
		resp := unixTimeFunc()
		assert.NotNil(t, resp)
	}
}

func TestAnd(t *testing.T) {
	cases := []struct {
		testName string
		u1       UUID
		u2       UUID
	}{
		{
			testName: "1. Positive test",
			u1:       NamespaceDNS,
			u2:       NamespaceDNS,
		},
	}

	for _, c := range cases {
		t.Logf(CurrentlyTesting, c.testName)
		uuid := And(c.u1, c.u2)
		assert.NotNil(t, uuid)
	}
}

func TestOr(t *testing.T) {
	cases := []struct {
		testName string
		u1       UUID
		u2       UUID
	}{
		{
			testName: "1. Positive test",
			u1:       NamespaceDNS,
			u2:       NamespaceDNS,
		},
	}

	for _, c := range cases {
		t.Logf(CurrentlyTesting, c.testName)
		uuid := Or(c.u1, c.u2)
		assert.NotNil(t, uuid)
	}
}

func TestEqual(t *testing.T) {
	cases := []struct {
		testName string
		u1       UUID
		u2       UUID
	}{
		{
			testName: "1. Positive test",
			u1:       NamespaceDNS,
			u2:       NamespaceDNS,
		},
	}

	for _, c := range cases {
		t.Logf(CurrentlyTesting, c.testName)
		uuid := Equal(c.u1, c.u2)
		assert.Equal(t, true, uuid)
	}
}

func TestUUID_Version(t *testing.T) {
	cases := []struct {
		testName string
		u        UUID
	}{
		{
			testName: "1. Positive test",
			u:        NamespaceDNS,
		},
	}

	for _, c := range cases {
		t.Logf(CurrentlyTesting, c.testName)
		version := c.u.Version()
		assert.NotNil(t, version)
	}
}

func TestUUID_Variant(t *testing.T) {
	cases := []struct {
		testName        string
		u               UUID
		expectedVariant uint
	}{
		{
			testName:        "1. Positive test : Variant NCS",
			u:               uuidVariantNCS,
			expectedVariant: VariantNCS,
		},
		{
			testName:        "2. Positive test : Variant RFC4122",
			u:               NamespaceDNS,
			expectedVariant: VariantRFC4122,
		},
	}

	for _, c := range cases {
		t.Logf(CurrentlyTesting, c.testName)
		variant := c.u.Variant()
		if variant != c.expectedVariant {
			t.Errorf(ExpectedResponse, c.expectedVariant, variant)
		}
	}
}

func TestUUID_Bytes(t *testing.T) {
	cases := []struct {
		testName string
		u        UUID
	}{
		{
			testName: "1. Positive test",
			u:        NamespaceDNS,
		},
	}

	for _, c := range cases {
		t.Logf(CurrentlyTesting, c.testName)
		bytes := c.u.Bytes()
		assert.NotNil(t, bytes)
	}
}

func TestUUID_String(t *testing.T) {
	cases := []struct {
		testName string
		u        UUID
	}{
		{
			testName: "1. Positive test",
			u:        NamespaceDNS,
		},
	}

	for _, c := range cases {
		t.Logf(CurrentlyTesting, c.testName)
		canonicalString := c.u.String()
		assert.NotNil(t, canonicalString)
	}
}

func TestUUID_MarshalText(t *testing.T) {
	cases := []struct {
		testName      string
		u             UUID
		expectedError error
	}{
		{
			testName:      "1. Positive test",
			u:             NamespaceDNS,
			expectedError: nil,
		},
	}

	for _, c := range cases {
		t.Logf(CurrentlyTesting, c.testName)
		_, err := c.u.MarshalText()
		if err != nil {
			t.Errorf(ExpectedResponse, c.expectedError, err)
		}
	}
}

func TestUUID_MarshalBinary(t *testing.T) {
	cases := []struct {
		testName string
		u        UUID
	}{
		{
			testName: "1. Positive test",
			u:        NamespaceDNS,
		},
	}

	for _, c := range cases {
		t.Logf(CurrentlyTesting, c.testName)
		_, err := c.u.MarshalBinary()
		if err != nil {
			t.Errorf(ExpectedResponse, nil, err)
		}
	}
}

func TestUUID_UnmarshalBinary(t *testing.T) {
	cases := []struct {
		testName      string
		u             UUID
		data          []byte
		expectedError error
	}{
		{
			testName:      "1. Positive test",
			u:             NamespaceDNS,
			data:          byteShortUUID,
			expectedError: nil,
		},
	}

	for _, c := range cases {
		t.Logf(CurrentlyTesting, c.testName)
		err := c.u.UnmarshalBinary(c.data)
		if err != c.expectedError {
			t.Errorf(ExpectedResponse, c.expectedError, err)
		}
	}
}

func TestUUID_Value(t *testing.T) {
	cases := []struct {
		testName      string
		u             UUID
		expectedValue driver.Value
	}{
		{
			testName:      "1. Positive test",
			u:             NamespaceDNS,
			expectedValue: uuid,
		},
	}

	for _, c := range cases {
		t.Logf(CurrentlyTesting, c.testName)
		value, _ := c.u.Value()
		if value != c.expectedValue {
			t.Errorf(ExpectedResponse, c.expectedValue, value)
		}
	}
}

func TestUUID_Scan(t *testing.T) {
	cases := []struct {
		testName      string
		u             UUID
		src           interface{}
		expectedError error
	}{
		{
			testName:      "1. Positive test : Unmarshal Binary",
			u:             NamespaceDNS,
			src:           byteShortUUID,
			expectedError: nil,
		},
		{
			testName:      "2. Positive test : Unmarshal Text (Case []byte)",
			u:             NamespaceDNS,
			src:           byteUUID,
			expectedError: nil,
		},
		{
			testName:      "3. Positive test : Unmarshal Text (Case string)",
			u:             NamespaceDNS,
			src:           uuid,
			expectedError: nil,
		},
	}

	for _, c := range cases {
		t.Logf(CurrentlyTesting, c.testName)
		err := c.u.Scan(c.src)
		if err != c.expectedError {
			t.Errorf(ExpectedResponse, c.expectedError, err)
		}
	}
}

func TestNullUUID_Value(t *testing.T) {
	cases := []struct {
		testName      string
		u             NullUUID
		expectedValue driver.Value
	}{
		{
			testName:      "1. Positive test",
			u:             nullUUID,
			expectedValue: nullUUID.UUID.String(),
		},
		{
			testName:      "2. Negative test",
			u:             invalidNullUUID,
			expectedValue: nil,
		},
	}

	for _, c := range cases {
		t.Logf(CurrentlyTesting, c.testName)
		value, _ := c.u.Value()
		if value != c.expectedValue {
			t.Errorf(ExpectedResponse, c.expectedValue, value)
		}
	}
}

func TestNullUUID_Scan(t *testing.T) {
	cases := []struct {
		testName      string
		u             NullUUID
		src           interface{}
		expectedError error
	}{
		{
			testName:      "1. Positive test",
			u:             nullUUID,
			src:           byteUUID,
			expectedError: nil,
		},
		{
			testName:      "2. Negative test",
			u:             nullUUID,
			src:           nil,
			expectedError: nil,
		},
	}

	for _, c := range cases {
		t.Logf(CurrentlyTesting, c.testName)
		err := c.u.Scan(c.src)
		if err != c.expectedError {
			t.Errorf(ExpectedResponse, c.expectedError, err)
		}
	}
}

func TestFromBytes(t *testing.T) {
	cases := []struct {
		testName      string
		input         []byte
		expectedError error
	}{
		{
			testName:      "1. Positive test",
			input:         byteShortUUID,
			expectedError: nil,
		},
	}

	for _, c := range cases {
		t.Logf(CurrentlyTesting, c.testName)
		_, err := FromBytes(c.input)
		if err != c.expectedError {
			t.Errorf(ExpectedResponse, c.expectedError, err)
		}
	}
}

func TestFromBytesOrNil(t *testing.T) {
	cases := []struct {
		testName     string
		input        []byte
		expectedResp UUID
	}{
		{
			testName:     "1. Positive test",
			input:        byteShortUUID,
			expectedResp: uuidVariantRFC4122,
		},
		{
			testName:     "2. Negative test",
			input:        byteUUID,
			expectedResp: uuidNil,
		},
	}

	for _, c := range cases {
		t.Logf(CurrentlyTesting, c.testName)
		resp := FromBytesOrNil(c.input)
		if resp != c.expectedResp {
			t.Errorf(ExpectedResponse, c.expectedResp, resp)
		}
	}
}

func TestFromStringOrNil(t *testing.T) {
	cases := []struct {
		testName     string
		input        string
		expectedResp UUID
	}{
		{
			testName:     "1. Positive test",
			input:        "36626137-6238-3130-2d39-6461642d3131",
			expectedResp: uuidVariantRFC4122,
		},
		{
			testName:     "2. Negative test",
			input:        shortUUID,
			expectedResp: uuidNil,
		},
	}

	for _, c := range cases {
		t.Logf(CurrentlyTesting, c.testName)
		resp := FromStringOrNil(c.input)
		if resp != c.expectedResp {
			t.Errorf(ExpectedResponse, c.expectedResp, resp)
		}
	}
}

func TestGetStorage(t *testing.T) {
	cases := []struct {
		testName string
	}{
		{
			testName: "1. Positive test",
		},
	}

	for _, c := range cases {
		t.Logf(CurrentlyTesting, c.testName)
		getStorage()
	}
}

func TestNewV1(t *testing.T) {
	cases := []struct {
		testName string
	}{
		{
			testName: "1. Positive test",
		},
	}

	for _, c := range cases {
		t.Logf(CurrentlyTesting, c.testName)
		uuid := NewV1()
		assert.NotNil(t, uuid)
	}
}

func TestNewV2(t *testing.T) {
	cases := []struct {
		testName string
		domain   byte
	}{
		{
			testName: "1. Positive test : Domain Person",
			domain:   DomainPerson,
		},
		{
			testName: "2. Positive test : Domain Group",
			domain:   DomainGroup,
		},
	}

	for _, c := range cases {
		t.Logf(CurrentlyTesting, c.testName)
		uuid := NewV2(c.domain)
		assert.NotNil(t, uuid)
	}
}

func TestNewV3(t *testing.T) {
	cases := []struct {
		testName string
		ns       UUID
		name     string
	}{
		{
			testName: "1. Positive test",
			ns:       NamespaceDNS,
			name:     "UUID",
		},
	}

	for _, c := range cases {
		t.Logf(CurrentlyTesting, c.testName)
		uuid := NewV3(c.ns, c.name)
		assert.NotNil(t, uuid)
	}
}

func TestNewV4(t *testing.T) {
	cases := []struct {
		testName string
	}{
		{
			testName: "1. Positive test",
		},
	}

	for _, c := range cases {
		t.Logf(CurrentlyTesting, c.testName)
		uuid := NewV4()
		assert.NotNil(t, uuid)
	}
}

func TestNewV5(t *testing.T) {
	cases := []struct {
		testName string
		ns       UUID
		name     string
	}{
		{
			testName: "1. Positive test",
			ns:       NamespaceDNS,
			name:     "UUID",
		},
	}

	for _, c := range cases {
		t.Logf(CurrentlyTesting, c.testName)
		uuid := NewV5(c.ns, c.name)
		assert.NotNil(t, uuid)
	}
}
