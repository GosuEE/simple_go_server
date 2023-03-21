package requestDTO

type UserRequestDTO struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Age int `json:"age"`
	Job string `json:"job"`
}