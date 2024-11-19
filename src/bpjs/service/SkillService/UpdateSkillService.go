package SkillService

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

	body := utils.SkillBody(request)
	body.UserId = claims.Id
	body.Id = id
	bodyRepo := skillRepository(body)
	db := confiq.Connect()

	if bodyRepo.Skill.String == "" || bodyRepo.Level.String == ""{
		out.ResponseOut(response, nil, false, constanta.CodeBadRequestResponse, "Skill/Level tidak boleh kosong")
		return
	}

	if bodyRepo.Level.String != "Advanced" && bodyRepo.Level.String != "Intermediate" && bodyRepo.Level.String != "Basic"{
		out.ResponseOut(response, nil, false, constanta.CodeBadRequestResponse, "Level hanya boleh diisi dengan Advanced/Intermediate/Basic")
		return
	}

	detail, err := dao.SkillDAO.GetDetailSkill(db, id)
	if err != nil {
		out.ResponseOut(response, nil, false, constanta.CodeInternalServerErrorResponse, constanta.ErrorInternalDB)
		return
	}

	if detail.ID.Int64 == 0{
		out.ResponseOut(response, nil, false, constanta.CodeBadRequestResponse, "Id Skill tidak diketahui")
		return
	}

	err = dao.SkillDAO.UpdateSkill(db, bodyRepo)
	if err != nil{
		out.ResponseOut(response, nil, false, constanta.CodeInternalServerErrorResponse, constanta.ErrorInternalDB)
		return
	}
	db.Close()
	out.ResponseOut(response, nil, true, constanta.CodeSuccessResponse, constanta.SuccessEditData)
	return nil
}
