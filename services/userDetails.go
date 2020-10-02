package services

import (
	"affordmed/models"
)

func VerifyCredentails(credentials models.Login) bool {

	for _, users := range models.RegisteredUsers {
		if credentials.Username == users.Username && credentials.Password == users.Password {
			return true
		}
	}
	return true
}

func EditUserDetials(details models.UserDetials) bool {
	for i := range models.RegisteredUsers {
		if details.Username == models.RegisteredUsers[i].Username {
			details.Password = models.RegisteredUsers[i].Password
			models.RegisteredUsers[i] = details
			return true
		}
	}
	return false
}
