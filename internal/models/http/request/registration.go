package request

type RegistrationRequest struct {
	FullName string `json:"fullName" required:"true"`
	Email    string `json:"email" required:"true"`
	Password string `json:"password" required:"true"`
	Phone    string `json:"phone" required:"true"`
}
