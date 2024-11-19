package out

type EducationResponse struct {
	Id          int64   `json:"id"`
	School    string   `json:"school"`
	Level    string   `json:"level"`
	Degree    string   `json:"degree"`
	YearIn    int64   `json:"year_in"`
	YearOut   int64   `json:"year_out"`
	CreatedAt   string   `json:"created_at"`
	UpdatedAt   string   `json:"updated_at"`
}
