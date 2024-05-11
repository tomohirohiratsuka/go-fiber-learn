package dto

type UserCreateServiceOutput struct {
	BaseUser
	Email string `json:"email"`
}

type UserFindServiceOutput struct {
	BaseUser
	Email string `json:"email"`
}
