package validations

import (
	"github.com/devmeireles/gnosi-api/app/models"
	"github.com/thedevsaddam/govalidator"
)

// ValidateEpisode validates the episode
func ValidateEpisode(episode *models.Episode) map[string]interface{} {
	rules := govalidator.MapData{
		"title":     []string{"required"},
		"season_id": []string{"required"},
	}

	opts := govalidator.Options{
		Data:  episode,
		Rules: rules,
	}

	v := govalidator.New(opts)
	e := v.ValidateStruct()
	if len(e) > 0 {
		return map[string]interface{}{"error": e}
	}

	return nil
}
