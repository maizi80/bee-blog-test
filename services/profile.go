package services

import (
	"bee-blog/models"
)

type profileService struct{}

func GetProfileByMap(p interface{}, profile models.Profile) models.Profile {
	for i, v := range p.(map[string]interface{}) {
		switch i {
		case "alias":
			profile.Alias = v.(string)
		case "name":
			profile.Name = v.(string)
		case "type":
			in := v.(int)
			profile.Type = uint(in)
		}
	}
	return profile
}
