package EducationService

import (
	"bpjs/confiq"
	"bpjs/constanta"
	"bpjs/dao"
	"bpjs/dto/in"
	"bpjs/dto/out"
	"bpjs/model"
	"bpjs/utils"
	"database/sql"
	"fmt"
	"net/http"
	"time"
)

func EducationInsert(response http.ResponseWriter, request *http.Request) (err error) {
	now := time.Now()
	fmt.Println("HIT -> InsertEducationService.go On ", now.Format("2006-01-02 15:04:05"))

	tokenString := request.Header.Get("Authorization")
	claims, err := confiq.DecodeToken(tokenString)
	if err != nil {
		out.ResponseOut(response, nil, false, constanta.CodeInternalServerErrorResponse, err.Error())
		return
	}

	body := utils.EducationBody(request)
	body.UserId = claims.Id
	bodyRepo := educationRepository(body)
	db := confiq.Connect()

	if bodyRepo.School.String == "" || bodyRepo.Level.String == ""{
		out.ResponseOut(response, nil, false, constanta.CodeBadRequestResponse, "School/Level tidak boleh kosong")
		return
	}

	if bodyRepo.YearIn.Int64 > bodyRepo.YearOut.Int64{
		out.ResponseOut(response, nil, false, constanta.CodeBadRequestResponse, "YearOut harus lebih besar dari yearIn")
		return
	}

	err = dao.EducationDAO.InsertEducation(db, bodyRepo)
	if err != nil{
		out.ResponseOut(response, nil, false, constanta.CodeInternalServerErrorResponse, constanta.ErrorInternalDB)
		return
	}
	db.Close()
	out.ResponseOut(response, nil, true, constanta.CodeSuccessResponse, constanta.SuccessAddData)
	return nil
}

func educationRepository(body in.EducationRequest) model.EducationModel  {

	return model.EducationModel{
		ID     :  sql.NullInt64{Int64: body.Id},
		UserId     :  sql.NullInt64{Int64: body.UserId},
		School: sql.NullString{String: body.School},
		Degree: sql.NullString{String: body.Degree},
		Level: sql.NullString{String: body.Level},
		YearIn     :  sql.NullInt64{Int64: body.YearIn},
		YearOut     :  sql.NullInt64{Int64: body.YearOut},
	}

}
