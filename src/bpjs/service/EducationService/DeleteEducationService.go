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

func DeleteEducation(response http.ResponseWriter, request *http.Request) (err error) {
	now := time.Now()
	fmt.Println("HIT -> DeleteEducationService.go On ", now.Format("2006-01-02 15:04:05"))

	id, err := utils.ReadParam(request)
	if err != nil{
		out.ResponseOut(response, nil, false, constanta.CodeInternalServerErrorResponse, err.Error())
		return
	}
	db := confiq.Connect()

	detail, err := dao.EducationDAO.GetDetailEducation(db, id)
	if err != nil {
		out.ResponseOut(response, nil, false, constanta.CodeInternalServerErrorResponse, constanta.ErrorInternalDB)
		return
	}

	if detail.ID.Int64 == 0{
		out.ResponseOut(response, nil, false, constanta.CodeBadRequestResponse, "Id Skill tidak diketahui")
		return
	}

	err = dao.EducationDAO.DeleteEducation(db, id)
	if err != nil {
		out.ResponseOut(response, nil, false, constanta.CodeInternalServerErrorResponse, constanta.ErrorInternalDB)
		return
	}

	db.Close()
	out.ResponseOut(response, nil, true, constanta.CodeSuccessResponse, constanta.SuccessDeleteData)
	return nil
}
