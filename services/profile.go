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

func GetMapByProfile(profiles []models.Profile) map[string]string {
	p := make(map[string]string)
	for _, profile := range profiles {
		switch profile.Alias {
		case "motto":
			p["motto"] = profile.Content
		case "motto_e":
			p["motto_e"] = profile.Content
		case "introduction":
			p["introduction"] = profile.Content
		case "qq":
			p["qq"] = profile.Content
		case "email":
			p["email"] = profile.Content
		case "github":
			p["github"] = profile.Content
		case "nickname":
			p["nickname"] = profile.Content
		case "avatar":
			p["avatar"] = profile.Content
		}
	}
	return p
}
