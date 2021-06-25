package models

type User struct {
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
	Email string `json:"email"`
	Address string `json:"address"`
}

func (user User) ValidateEmail() bool {
	if len(user.Email) == 0  {
		return false
	}
	if !(user.Email[0] >= 'A' && user.Email[0] <= 'Z') && !(user.Email[0] >= 'a' && user.Email[0] <= 'z') {
		return false
	}
	var index = 0
	for ;index < len(user.Email);index++ {
		if user.Email[index] == '@' {
			break
		}
	}
	index++
	for ;index < len(user.Email);index++ {
		if user.Email[index] == '.' {
			break
		}
	}
	index++
	if index == len(user.Email) {
		return false
	}
	return true
}