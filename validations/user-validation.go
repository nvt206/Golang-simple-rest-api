package validations

type UserValidate struct {
	Email string `binding:"email,required"`
	Password string `binding:"required,min=10"`
	Address string `binding:"required"`
}
