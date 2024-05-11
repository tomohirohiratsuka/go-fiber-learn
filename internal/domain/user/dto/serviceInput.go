package dto

type UserCreateServiceInput struct {
	Email    string
	Password string
}

func (r *UserCreateServiceInput) FromRequest(body UserCreateRequest) {
	r.Email = body.Email
	r.Password = body.Password
}
