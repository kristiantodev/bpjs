package dao

import (
	"bpjs/model"
	"database/sql"
	"fmt"
	"strconv"
)

type educationDAO struct {
	AbstractDAO
}

var EducationDAO = educationDAO{}.New()

func (input educationDAO) New() (output educationDAO) {
	output.TableName = "education"
	output.FileName = "educationDAO.go"
	return
}

func (input educationDAO) InsertEducation(db *sql.DB, inputStruct model.EducationModel) (err error) {
	query := "INSERT INTO " + input.TableName + " (" +
		" user_id, school, level, degree, year_in, year_out)  " +
		"VALUES (?, ?, ?, ?, ?, ?) "
	params := []interface{}{
		inputStruct.UserId.Int64, inputStruct.School.String, inputStruct.Level.String,
		inputStruct.Degree.String, inputStruct.YearIn.Int64, inputStruct.YearOut.Int64,
	}

	_, err = db.Exec(query, params...)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return
}

func (input educationDAO) UpdateEducation(db *sql.DB, inputStruct model.EducationModel) (err error) {
	query := "UPDATE " + input.TableName + " SET " +
		"school = ?, level = ? , degree = ?, year_in = ?, year_out = ?" +
		" WHERE id = ?"

	params := []interface{}{
		inputStruct.School.String, inputStruct.Level.String,
		inputStruct.Degree.String, inputStruct.YearIn.Int64,
		inputStruct.YearOut.Int64,inputStruct.ID.Int64,
	}

	_, err = db.Exec(query, params...)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (input educationDAO) GetDetailEducation(db *sql.DB, id int64) (model.EducationModel, error) {
	query := "SELECT id, school, level, degree, year_in, year_out, created_at, updated_at " +
		"FROM " + input.TableName + " WHERE id = ?"

	row := db.QueryRow(query, id)
	var data model.EducationModel

	err := row.Scan(&data.ID,
		&data.School,
		&data.Level,
		&data.Degree,
		&data.YearIn,
		&data.YearOut,
		&data.CreatedAt,
		&data.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return data, nil
		}
		return data, err
	}
	return data, nil
}

func (input educationDAO) GetEducationList(db *sql.DB, param CustomQueryModel) (results []model.EducationModel, err error) {
	query := "SELECT id, school, level, degree, year_in, year_out, created_at, updated_at " +
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
		var data model.EducationModel
		if err := rows.Scan(&data.ID,
			&data.School,
			&data.Level,
			&data.Degree,
			&data.YearIn,
			&data.YearOut,
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

func (input educationDAO) DeleteEducation(db *sql.DB, id int64) (err error) {
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