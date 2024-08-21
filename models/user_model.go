package models

import (
	"edot-test/utils"
	"time"
)

// User struct for user
type User struct {
	PubID                string    `orm:"pk;column(pubid)" json:"pubId"`
	Email                string    `orm:"column(email);null" json:"email"`
	Mdn                  string    `orm:"column(mdn);null" json:"mdn"`
	UserName             string    `orm:"column(username);null" json:"userName"`
	FullName             string    `orm:"column(full_name);null" json:"name"`
	Address              string    `orm:"column(address);null" json:"address"`
	DateOfBirth          time.Time `orm:"column(date_of_birth);null" json:"dateOfBirth"`
	Gender               string    `orm:"column(gender);null" json:"gender"`
	IsUnlimitedSwipe     bool      `orm:"column(is_unlimited_swipe);null" json:"isUnlimitedSwipe"`
	IsVerified           bool      `orm:"column(is_verified);null" json:"isVerified"`
	Password             string    `orm:"column(password);null" json:"-"`
	CounterPassword      int       `orm:"column(counter_password);null" json:"counterPassword"`
	CounterPasswordExpTs time.Time `orm:"column(counter_password_exp_ts);null" json:"counterPasswordExpTs"`
	Token                string    `orm:"column(token);null" json:"-"`
	ExpireTokenTs        time.Time `orm:"column(expire_token_ts);null" json:"-"`
	CreateTs             time.Time `orm:"column(create_ts);null" json:"createTs"`
	CreateBy             string    `orm:"column(create_by);null" json:"createBy"`
	UpdateTs             time.Time `orm:"column(update_ts);null" json:"updateTs"`
	UpdateBy             string    `orm:"column(update_by);null" json:"updateBy"`
}

// TableName for users
func (u *User) TableName() string {

	return "users"
}

// NewUser is func for initialize User
func NewUser(req RegisterUserRequest) User {
	pubID := utils.NewV4().String()
	dob, _ := time.Parse("2006-01-02", req.DateOfBirth)
	password, _ := utils.HashPassword(req.Password)
	return User{
		PubID:       pubID,
		Email:       req.Email,
		Mdn:         req.Mdn,
		UserName:    req.UserName,
		FullName:    req.FullName,
		Address:     req.Address,
		DateOfBirth: dob,
		Gender:      req.Gender,
		Password:    password,
		CreateTs:    time.Now(),
		CreateBy:    "SYSTEM",
		UpdateTs:    time.Now(),
	}
}

// RegisterUserRequest struct for register new user
type RegisterUserRequest struct {
	Email           string `json:"email"`
	Mdn             string `json:"mdn" valid:"Required"`
	UserName        string `json:"userName"`
	FullName        string `json:"name"`
	Address         string `json:"address"`
	DateOfBirth     string `json:"dateOfBirth"`
	Gender          string `json:"gender"`
	Password        string `json:"password" valid:"Required"`
	ConfirmPassword string `json:"confirmPassword" valid:"Required"`
}

// LoginRequest struct for login request
type LoginRequest struct {
	Mdn      string `json:"mdn" valid:"Required"`
	Password string `json:"password" valid:"Required"`
}

// LoginResponse struct for response login request
type LoginResponse struct {
	UserPubID   string `json:"userPubId"`
	Token       string `json:"token"`
	ExpireToken int64  `json:"expireToken"`
}

// NewLoginResponse is function for initialize LoginResponse
func NewLoginResponse(userPubID, token string, expToken int64) LoginResponse {
	return LoginResponse{
		UserPubID:   userPubID,
		Token:       token,
		ExpireToken: expToken,
	}
}
