package in

type UserRequest struct {
	Id          int64    `json:"-"`
	Username    string   `json:"username"`
	Password    string   `json:"password"`
	FirstName   string   `json:"first_name"`
	LastName    string   `json:"last_name"`
	Gender      string   `json:"gender"`
	Telephone   string   `json:"telephone"`
	Email       string   `json:"email"`
	Address     string   `json:"address"`
}
