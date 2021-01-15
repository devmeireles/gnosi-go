package validations

import (
	"github.com/devmeireles/gnosi-api/app/models"
	"github.com/thedevsaddam/govalidator"
)

// ValidateSeason validates the season
func ValidateSeason(season *models.Season) map[string]interface{} {
	rules := govalidator.MapData{
		"title":        []string{"required"},
		"catalogue_id": []string{"required"},
	}

	opts := govalidator.Options{
		Data:  season,
		Rules: rules,
	}

	v := govalidator.New(opts)
	e := v.ValidateStruct()
	if len(e) > 0 {
		return map[string]interface{}{"error": e}
	}

	return nil
}
