package EducationService

import (
	"bpjs/confiq"
	"bpjs/constanta"
	"bpjs/dao"
	"bpjs/dto/out"
	"bpjs/model"
	"bpjs/utils"
	"fmt"
	"net/http"
	"time"
)

func GetDetailEducation(response http.ResponseWriter, request *http.Request) (err error) {
	now := time.Now()
	fmt.Println("HIT -> GetDetailEducationervice.go On ", now.Format("2006-01-02 15:04:05"))

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
		out.ResponseOut(response, nil, false, constanta.CodeBadRequestResponse, "Id Education tidak diketahui")
		return
	}

	db.Close()
	out.ResponseOut(response, convertRepoToDTOID(detail), true, constanta.CodeSuccessResponse, constanta.SuccessGetData)
	return nil
}

func convertRepoToDTOID(datas model.EducationModel) out.EducationResponse {
	return out.EducationResponse{
		Id:            datas.ID.Int64,
		School:         datas.School.String,
		Level:         datas.Level.String,
		Degree:        datas.Degree.String,
		YearOut:       datas.YearOut.Int64,
		YearIn :       datas.YearIn.Int64,
		CreatedAt:     datas.CreatedAt.String,
		UpdatedAt:     datas.UpdatedAt.String,
	}
}
