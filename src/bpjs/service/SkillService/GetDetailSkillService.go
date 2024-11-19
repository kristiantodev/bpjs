package SkillService

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

func GetDetailSkill(response http.ResponseWriter, request *http.Request) (err error) {
	now := time.Now()
	fmt.Println("HIT -> GetDetailSkillService.go On ", now.Format("2006-01-02 15:04:05"))

	id, err := utils.ReadParam(request)
	if err != nil{
		out.ResponseOut(response, nil, false, constanta.CodeInternalServerErrorResponse, err.Error())
		return
	}
	db := confiq.Connect()

	detail, err := dao.SkillDAO.GetDetailSkill(db, id)
	if err != nil {
		out.ResponseOut(response, nil, false, constanta.CodeInternalServerErrorResponse, constanta.ErrorInternalDB)
		return
	}

	if detail.ID.Int64 == 0{
		out.ResponseOut(response, nil, false, constanta.CodeBadRequestResponse, "Id Skill tidak diketahui")
		return
	}

	db.Close()
	out.ResponseOut(response, convertRepoToDTOID(detail), true, constanta.CodeSuccessResponse, constanta.SuccessGetData)
	return nil
}

func convertRepoToDTOID(datas model.SkillModel) out.SkillResponse {
	return out.SkillResponse{
		Id:           datas.ID.Int64,
		Skill:         datas.Skill.String,
		Level:         datas.Level.String,
		CreatedAt:     datas.CreatedAt.String,
		UpdatedAt:     datas.UpdatedAt.String,
	}
}