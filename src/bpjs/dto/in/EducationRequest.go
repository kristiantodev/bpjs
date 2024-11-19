package in

type EducationRequest struct {
	Id       int64    `json:"-"`
	UserId   int64    `json:"user_id"`
	School    string   `json:"school"`
	Level    string   `json:"level"`
	Degree    string   `json:"degree"`
	YearIn    int64   `json:"year_in"`
	YearOut   int64   `json:"year_out"`
}

