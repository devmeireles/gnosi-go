package validations

import (
	"fmt"

	"github.com/devmeireles/gnosi-api/app/models"
	"github.com/thedevsaddam/govalidator"
)

// ValidateCatalogue validates the catalogue
func ValidateCatalogue(catalogue *models.Catalogue) map[string]interface{} {
	rules := govalidator.MapData{
		"title":       []string{"required"},
		"description": []string{"required"},
		"price":       []string{"required"},
		"owner_id":    []string{"required"},
	}

	opts := govalidator.Options{
		Data:  catalogue,
		Rules: rules,
	}

	v := govalidator.New(opts)
	e := v.ValidateStruct()
	if len(e) > 0 {
		fmt.Println(e)
		return map[string]interface{}{"error": e}
	}

	return nil
}
