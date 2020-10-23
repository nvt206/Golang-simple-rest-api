package validations

type LoginValidation struct {
	Email string `binding:"email,required"`
	Password string `binding:"required,min=10"`
}
