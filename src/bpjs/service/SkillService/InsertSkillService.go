package SkillService

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

func SkillInsert(response http.ResponseWriter, request *http.Request) (err error) {
	now := time.Now()
	fmt.Println("HIT -> InsertSkillService.go On ", now.Format("2006-01-02 15:04:05"))

	tokenString := request.Header.Get("Authorization")
	claims, err := confiq.DecodeToken(tokenString)
	if err != nil {
		out.ResponseOut(response, nil, false, constanta.CodeInternalServerErrorResponse, err.Error())
		return
	}

	body := utils.SkillBody(request)
	body.UserId = claims.Id
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

	err = dao.SkillDAO.InsertSkill(db, bodyRepo)
	if err != nil{
		out.ResponseOut(response, nil, false, constanta.CodeInternalServerErrorResponse, constanta.ErrorInternalDB)
		return
	}
	db.Close()
	out.ResponseOut(response, nil, true, constanta.CodeSuccessResponse, constanta.SuccessAddData)
	return nil
}

func skillRepository(body in.SkillRequest) model.SkillModel  {

	return model.SkillModel{
		ID     :  sql.NullInt64{Int64: body.Id},
		UserId     :  sql.NullInt64{Int64: body.UserId},
		Skill: sql.NullString{String: body.Skill},
		Level: sql.NullString{String: body.Level},
	}

}
