package model

import "database/sql"

type EducationModel struct {
	ID          sql.NullInt64
	UserId      sql.NullInt64
	School      sql.NullString
	Level       sql.NullString
	Degree      sql.NullString
	YearIn      sql.NullInt64
	YearOut      sql.NullInt64
	CreatedAt   sql.NullString
	UpdatedAt   sql.NullString
}
