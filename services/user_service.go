package services

import (
	"edot-test/models"
	"edot-test/utils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"time"
)

// UserProcessor interface for user process
type UserProcessor interface {
	FindByMdn(mdn string) (models.User, error)
	Insert(user models.User) (models.User, error)
	UpdateColumns(user models.User, cols ...string) error
}

// UserService struct
type UserService struct {
	Identifier    int64
	userProcessor UserProcessor
}

// NewUserService is func for initialize UserService
func NewUserService(up UserProcessor, i int64) UserService {
	return UserService{
		userProcessor: up,
		Identifier:    i,
	}
}

// Register is func for register new user
func (svc UserService) Register(b []byte) models.Response {
	request := models.RegisterUserRequest{}
	res := ConvertRequest(b, &request, svc.Identifier)
	if res.Code != utils.ErrNone {
		logs.Warn("[%d] Failed to convert request: %+v", svc.Identifier, res)
		return models.ResponseError(res.ErrorMessage, utils.ErrReqInvalid)
	}
	logs.Info("registerUser request: %+v", request)

	user, err := svc.userProcessor.FindByMdn(request.Mdn)
	if err != nil && err != orm.ErrNoRows {
		logs.Warn("[%d] Failed to find user by mdn: %s", svc.Identifier, err.Error())
		return models.ResponseError(utils.MsgErrDefault, utils.ErrDefault)
	}

	if user.PubID != "" {
		logs.Warn("[%d] User already exist: %s", svc.Identifier, user.Mdn)
		return models.ResponseError(utils.MsgUserAlreadyExist, utils.ErrUserAlreadyExist)
	}

	if request.Password != request.ConfirmPassword {
		logs.Warn("[%d] Password not match: %s", svc.Identifier, user)
		return models.ResponseError(utils.MsgPasswordNotMatch, utils.ErrPasswordNotMatch)
	}

	newUserReq := models.NewUser(request)
	user, err = svc.userProcessor.Insert(newUserReq)
	if err != nil {
		logs.Warn("[%d] Failed to register user: %s", svc.Identifier, err.Error())
		return models.ResponseError(utils.MsgErrDefault, utils.ErrDefault)
	}

	return models.ResponseSuccess(user)
}

// Login is function for user login
func (svc UserService) Login(b []byte) models.Response {
	request := models.LoginRequest{}
	res := ConvertRequest(b, &request, svc.Identifier)
	if res.Code != utils.ErrNone {
		logs.Warn("[%d] Failed to convert request: %+v", svc.Identifier, res)
		return models.ResponseError(res.ErrorMessage, utils.ErrReqInvalid)
	}
	logs.Info("login request: %+v", request)

	user, err := svc.userProcessor.FindByMdn(request.Mdn)
	if err != nil {
		logs.Warn("[%d] Failed to find user by mdn: %s", svc.Identifier, err.Error())
		return models.ResponseError(utils.MsgUserNotFound, utils.ErrUserNotFound)
	}

	isValidPassword := utils.CheckPasswordHash(request.Password, user.Password)
	if !isValidPassword {
		counterPassword := beego.AppConfig.DefaultInt("counterPassword", 3)
		if user.CounterPassword >= counterPassword {
			logs.Warn("[%d] User Max Counter Password: %d", svc.Identifier, user.CounterPassword)
			return models.ResponseError(utils.MsgMaxCounterPassword, utils.ErrMaxCounterPassword)
		}

		user.CounterPassword = user.CounterPassword + 1
		user.UpdateBy = "SYSTEM"
		user.UpdateTs = time.Now()

		if user.CounterPassword == counterPassword {
			user.CounterPasswordExpTs = time.Now().Add(time.Minute * 10)
		}

		err = svc.userProcessor.UpdateColumns(user, "CounterPassword", "UpdateBy", "UpdateTs", "CounterPasswordExpTs")
		if err != nil {
			logs.Warn("[%d] Failed to update counter password: %s", svc.Identifier, err.Error())
			return models.ResponseError(utils.MsgErrDefault, utils.ErrDefault)
		}

		logs.Warn("[%d] password not match: %s", svc.Identifier, user.PubID)
		return models.ResponseError(utils.MsgIncorrectPassword, utils.ErrIncorrectPassword)
	}

	token := user.Token
	exp := user.ExpireTokenTs.Unix()
	var columns []string
	if user.ExpireTokenTs.Before(time.Now()) {
		var secretKey = []byte(beego.AppConfig.DefaultString("secretKey", "d3aLls#!23"))
		token, exp = utils.CreateToken(request.Mdn, secretKey)

		user.Token = token
		user.ExpireTokenTs = time.Unix(exp, 10)
		columns = append(columns, "Token", "ExpireTokenTs")
	}

	user.UpdateBy = "SYSTEM"
	user.UpdateTs = time.Now()
	user.CounterPassword = 0
	columns = append(columns, "UpdateBy", "UpdateTs", "CounterPassword")

	err = svc.userProcessor.UpdateColumns(user, columns...)
	if err != nil {
		logs.Warn("[%d] Failed to update user token: %s", svc.Identifier, err.Error())
		return models.ResponseError(utils.MsgErrDefault, utils.ErrDefault)
	}

	resp := models.NewLoginResponse(user.PubID, token, exp)
	return models.ResponseSuccess(resp)
}
