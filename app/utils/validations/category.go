package validations

import (
	"fmt"

	"github.com/devmeireles/gnosi-api/app/models"
	"github.com/thedevsaddam/govalidator"
)

// ValidateCategory validates the category
func ValidateCategory(category *models.Category) map[string]interface{} {
	rules := govalidator.MapData{
		"title":       []string{"required"},
		"description": []string{"required"},
	}

	opts := govalidator.Options{
		Data:  category,
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
