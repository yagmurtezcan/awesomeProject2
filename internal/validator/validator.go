package validator

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var (
	trans    ut.Translator
	Validate *validator.Validate
)

func Init() {
	en := en.New()
	uni := ut.New(en, en)

	trans, _ = uni.GetTranslator("en")

	Validate = validator.New()
	err := en_translations.RegisterDefaultTranslations(Validate, trans)
	if err != nil {
		return
	}
}

func ValidateRequest(request interface{}) []string {
	var arr []string

	err := Validate.Struct(request)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			validationErr := e.Translate(trans)
			arr = append(arr, validationErr)
		}
		return arr
	}
	return nil
}
