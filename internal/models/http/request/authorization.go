package request

type SingUpRequest struct {
	FullName string `json:"fullName" required:"true"`
	Email    string `json:"email" required:"true"`
	Password string `json:"password" required:"true"`
	Phone    string `json:"phone" required:"true"`
}

type SingInRequest struct {
	Email    string `json:"email" required:"true"`
	Password string `json:"password" required:"true"`
}
