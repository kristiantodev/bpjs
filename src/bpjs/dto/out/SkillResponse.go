package out

type SkillResponse struct {
	Id          int64   `json:"id"`
	Skill       string  `json:"skill"`
	Level       string   `json:"level"`
	CreatedAt   string   `json:"created_at"`
	UpdatedAt   string   `json:"updated_at"`
}
