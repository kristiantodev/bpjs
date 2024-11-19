package model

import "database/sql"

type SkillModel struct {
	ID          sql.NullInt64
	UserId      sql.NullInt64
	Skill       sql.NullString
	Level       sql.NullString
	CreatedAt   sql.NullString
	UpdatedAt   sql.NullString
}
