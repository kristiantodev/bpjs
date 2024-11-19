package dao

import (
	"bpjs/model"
	"database/sql"
	"fmt"
	"strconv"
)

type skillDAO struct {
	AbstractDAO
}

var SkillDAO = skillDAO{}.New()

func (input skillDAO) New() (output skillDAO) {
	output.TableName = "skills"
	output.FileName = "SkillDAO.go"
	return
}

func (input skillDAO) InsertSkill(db *sql.DB, inputStruct model.SkillModel) (err error) {
	query := "INSERT INTO " + input.TableName + " (" +
		" user_id, skill, level)  " +
		"VALUES (?, ?, ?) "
	params := []interface{}{
		inputStruct.UserId.Int64, inputStruct.Skill.String, inputStruct.Level.String,
	}

	_, err = db.Exec(query, params...)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return
}

func (input skillDAO) UpdateSkill(db *sql.DB, inputStruct model.SkillModel) (err error) {
	query := "UPDATE " + input.TableName + " SET " +
		"skill = ?, level = ? " +
		"WHERE id = ?"

	params := []interface{}{
		inputStruct.Skill.String, inputStruct.Level.String,
		inputStruct.ID.Int64,
	}

	_, err = db.Exec(query, params...)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (input skillDAO) GetDetailSkill(db *sql.DB, id int64) (model.SkillModel, error) {
	query := "SELECT id, skill, level, created_at, updated_at " +
		"FROM " + input.TableName + " WHERE id = ?"

	row := db.QueryRow(query, id)
	var skill model.SkillModel

	err := row.Scan(&skill.ID, &skill.Skill, &skill.Level,
		&skill.CreatedAt, &skill.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return skill, nil
		}
		return skill, err
	}
	return skill, nil
}

func (input skillDAO) GetSkillList(db *sql.DB, param CustomQueryModel) (results []model.SkillModel, err error) {
	query := "SELECT id, skill, level, created_at, updated_at " +
		" FROM " + input.TableName +
		" WHERE user_id = ? "

	var params []interface{}

	params = append(params, param.Id)

	if param.Keyword != "" {
		query += " WHERE skill LIKE ?"
		params = append(params, "%"+param.Keyword+"%")
	}

	if param.Page != "" && param.Limit != "" {
		page, err := strconv.Atoi(param.Page)
		if err != nil {
			return nil, err
		}
		limit, err := strconv.Atoi(param.Limit)
		if err != nil {
			return nil, err
		}
		offset := limit * (page - 1)

		query += " LIMIT ? OFFSET ?"
		params = append(params, limit, offset)
	}

	rows, err := db.Query(query, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var data model.SkillModel
		if err := rows.Scan(&data.ID,
			&data.Skill,
			&data.Level,
			&data.CreatedAt,
			&data.UpdatedAt); err != nil {
			return nil, err
		}
		results = append(results, data)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func (input skillDAO) DeleteSkill(db *sql.DB, id int64) (err error) {
	query := "DELETE FROM " + input.TableName + " " +
		"WHERE id = ? "
	params := []interface{}{
		id,
	}

	_, err = db.Exec(query, params...)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return
}