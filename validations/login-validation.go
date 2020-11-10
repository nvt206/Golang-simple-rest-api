package validations

type LoginValidation struct {
	Email string `binding:"email,required" example:"abc@gmail.com"`
	Password string `binding:"required,min=10" example:"123123@X"`
}
