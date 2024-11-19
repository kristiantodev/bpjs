package in

type SkillRequest struct {
	Id       int64    `json:"-"`
	UserId   int64    `json:"user_id"`
	Skill    string   `json:"skill"`
	Level    string   `json:"level"`
}
