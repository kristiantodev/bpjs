package EducationService

import (
	"bpjs/confiq"
	"bpjs/constanta"
	"bpjs/dao"
	"bpjs/dto/out"
	"bpjs/utils"
	"fmt"
	"net/http"
	"time"
)

func UpdateSkill(response http.ResponseWriter, request *http.Request) (err error) {
	now := time.Now()
	fmt.Println("HIT -> UpdateSkillService.go On ", now.Format("2006-01-02 15:04:05"))

	tokenString := request.Header.Get("Authorization")
	claims, err := confiq.DecodeToken(tokenString)
	if err != nil {
		out.ResponseOut(response, nil, false, constanta.CodeInternalServerErrorResponse, err.Error())
		return
	}

	id, err := utils.ReadParam(request)
	if err != nil{
		out.ResponseOut(response, nil, false, constanta.CodeInternalServerErrorResponse, err.Error())
		return
	}

	body := utils.EducationBody(request)
	body.UserId = claims.Id
	body.Id = id
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

	detail, err := dao.EducationDAO.GetDetailEducation(db, id)
	if err != nil {
		out.ResponseOut(response, nil, false, constanta.CodeInternalServerErrorResponse, constanta.ErrorInternalDB)
		return
	}

	if detail.ID.Int64 == 0{
		out.ResponseOut(response, nil, false, constanta.CodeBadRequestResponse, "Id Education tidak diketahui")
		return
	}

	err = dao.EducationDAO.UpdateEducation(db, bodyRepo)
	if err != nil{
		out.ResponseOut(response, nil, false, constanta.CodeInternalServerErrorResponse, constanta.ErrorInternalDB)
		return
	}
	db.Close()
	out.ResponseOut(response, nil, true, constanta.CodeSuccessResponse, constanta.SuccessEditData)
	return nil
}
