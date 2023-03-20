package responseDTO

type UserResponseDTO struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Age int `json:"age"`
	Job string `json:"job"`
}