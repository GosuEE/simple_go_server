package service

import (
	"databases"
	"responseDTO"
)

func GetUser() []responseDTO.UserResponseDTO{
	users := databases.GetAllUsers()

	ret := make([]responseDTO.UserResponseDTO,0)
	for i := 0; i<len(users); i++{
		u := responseDTO.UserResponseDTO{}

		u.Id = users[i].Id
		u.Age = users[i].Age
		u.Job = users[i].Job
		u.Name = users[i].Name

		ret = append(ret, u)
	}
	
	return ret;
}