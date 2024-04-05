package services

import (
	"diskha-test/models"
	"diskha-test/utils"
	"errors"
	"testing"
	"time"
)

const expectedResponseCode = "Expected resp code to be %v but it was %v"

type fakeUserProcessor struct {
	respFindByMdn  models.User
	errFindByMdn   error
	respInsertUser models.User
	errInsertUser  error
	errUpdateToken error
}

func (f fakeUserProcessor) FindByMdn(mdn string) (models.User, error) {
	return f.respFindByMdn, f.errFindByMdn
}
func (f fakeUserProcessor) Insert(user models.User) (models.User, error) {
	return f.respInsertUser, f.errInsertUser
}

func (f fakeUserProcessor) UpdateColumns(user models.User, cols ...string) error {
	return f.errUpdateToken
}

var (
	positiveRegisterUser = fakeUserProcessor{
		errFindByMdn:  nil,
		errInsertUser: nil,
	}

	negativeFindUserByMdn = fakeUserProcessor{
		errFindByMdn: errors.New("DB Error"),
	}

	existFindUserByMdn = fakeUserProcessor{
		errFindByMdn:  nil,
		respFindByMdn: models.User{PubID: "123"},
	}

	negativeInsertUser = fakeUserProcessor{
		errInsertUser: errors.New("DB Error"),
	}

	positiveRequest         = []byte(`{"mdn":"6281987876654", "password":"123", "confirmPassword": "123"}`)
	negativeRequest         = []byte(`I AM JSON`)
	passwordNotMatchRequest = []byte(`{"mdn":"6281987876654", "password":"123", "confirmPassword": "abc"}`)
)

func TestUserService_Register(t *testing.T) {
	cases := []struct {
		testName         string
		request          []byte
		userProcessor    UserProcessor
		expectedResponse int
	}{
		{
			testName:         "Positive: Success Flow",
			expectedResponse: utils.ErrNone,
			userProcessor:    positiveRegisterUser,
			request:          positiveRequest,
		},
		{
			testName:         "Negative Test: Invalid Request",
			request:          negativeRequest,
			expectedResponse: utils.ErrReqInvalid,
		},
		{
			testName:         "Negative Test: Failed Find Use By MDN",
			request:          positiveRequest,
			userProcessor:    negativeFindUserByMdn,
			expectedResponse: utils.ErrDefault,
		},
		{
			testName:         "Negative Test: User Already Exist",
			request:          positiveRequest,
			userProcessor:    existFindUserByMdn,
			expectedResponse: utils.ErrUserAlreadyExist,
		},
		{
			testName:         "Negative Test: Password Not Match",
			request:          passwordNotMatchRequest,
			userProcessor:    positiveRegisterUser,
			expectedResponse: utils.ErrPasswordNotMatch,
		},
		{
			testName:         "Negative Test: Failed Register User",
			request:          positiveRequest,
			userProcessor:    negativeInsertUser,
			expectedResponse: utils.ErrDefault,
		},
	}

	for _, c := range cases {
		svc := NewUserService(c.userProcessor, 123)
		resp := svc.Register(c.request)
		if resp.Code != c.expectedResponse {
			t.Errorf(expectedResponseCode, c.expectedResponse, resp.Code)
		}
	}
}

var (
	expTokenDateSting     = "2070-08-15 02:30:45"
	expTokenDate, _       = time.Parse("2006-01-02 03:04:05", expTokenDateSting)
	positiveFindUserLogin = fakeUserProcessor{
		respFindByMdn: models.User{
			Password:      "$2a$14$7x0URT8OS7iSJug7az08IeVwRPHsK6imOw.E2aqJ6u.C2Y291Z0fW",
			ExpireTokenTs: expTokenDate,
		},
	}

	negativePasswordNotMatch = fakeUserProcessor{
		respFindByMdn: models.User{
			Password:        "$2a$14$7x0URT8OS7iSJug7az08IeVwRPHsK6imOw.E2aqJ6u",
			CounterPassword: 2,
		},
	}

	negativePasswordNotMatch3Times = fakeUserProcessor{
		respFindByMdn: models.User{
			Password:        "$2a$14$7x0URT8OS7iSJug7az08IeVwRPHsK6imOw.E2aqJ6u",
			CounterPassword: 3,
		},
	}

	negativeUpdateCounterPasswordNotMatch = fakeUserProcessor{
		respFindByMdn: models.User{
			Password:        "$2a$14$7x0URT8OS7iSJug7az08IeVwRPHsK6imOw.E2aqJ6u",
			CounterPassword: 2,
		},
		errUpdateToken: errors.New("DB Error"),
	}

	expTokenDateString1          = "2020-08-15 02:30:45"
	expTokenDate1, _             = time.Parse("2006-01-02 03:04:05", expTokenDateString1)
	positiveFindUserWithExpToken = fakeUserProcessor{
		respFindByMdn: models.User{
			Password:      "$2a$14$7x0URT8OS7iSJug7az08IeVwRPHsK6imOw.E2aqJ6u.C2Y291Z0fW",
			ExpireTokenTs: expTokenDate1,
		},
	}

	negativeUpdateUserToken = fakeUserProcessor{
		respFindByMdn: models.User{
			Password: "$2a$14$7x0URT8OS7iSJug7az08IeVwRPHsK6imOw.E2aqJ6u.C2Y291Z0fW",
		},
		errUpdateToken: errors.New("Error DB"),
	}

	positiveLoginReq = []byte(`{"username": "6287716253625", "password": "123abc"}`)
)

func TestUserService_Login(t *testing.T) {
	cases := []struct {
		testName         string
		request          []byte
		userProcessor    UserProcessor
		expectedResponse int
	}{
		{
			testName:         "Positive: Success Flow",
			expectedResponse: utils.ErrNone,
			userProcessor:    positiveFindUserLogin,
			request:          positiveLoginReq,
		},
		{
			testName:         "Positive Test: Success flow with exp token",
			request:          positiveLoginReq,
			userProcessor:    positiveFindUserWithExpToken,
			expectedResponse: utils.ErrNone,
		},
		{
			testName:         "Negative Test: Invalid Request",
			request:          negativeRequest,
			expectedResponse: utils.ErrReqInvalid,
		},
		{
			testName:         "Negative Test: Failed Find Use By MDN",
			request:          positiveRequest,
			userProcessor:    negativeFindUserByMdn,
			expectedResponse: utils.ErrUserNotFound,
		},
		{
			testName:         "Negative Test: Incorrect Password",
			request:          positiveRequest,
			userProcessor:    negativePasswordNotMatch,
			expectedResponse: utils.ErrIncorrectPassword,
		},
		{
			testName:         "Negative Test: Incorrect Password 3 Times",
			request:          positiveRequest,
			userProcessor:    negativePasswordNotMatch3Times,
			expectedResponse: utils.ErrMaxCounterPassword,
		},
		{
			testName:         "Negative Test: Incorrect Password Failed Update Counter",
			request:          positiveRequest,
			userProcessor:    negativeUpdateCounterPasswordNotMatch,
			expectedResponse: utils.ErrDefault,
		},
		{
			testName:         "Negative Test: Failed Update Token",
			request:          positiveLoginReq,
			userProcessor:    negativeUpdateUserToken,
			expectedResponse: utils.ErrDefault,
		},
	}

	for _, c := range cases {
		svc := NewUserService(c.userProcessor, 123)
		resp := svc.Login(c.request)
		if resp.Code != c.expectedResponse {
			t.Errorf(expectedResponseCode, c.expectedResponse, resp.Code)
		}
	}
}
