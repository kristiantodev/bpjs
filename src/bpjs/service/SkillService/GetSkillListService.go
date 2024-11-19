package SkillService

import (
	"bpjs/confiq"
	"bpjs/constanta"
	"bpjs/dao"
	"bpjs/dto/out"
	"bpjs/model"
	"fmt"
	"net/http"
	"time"
)

func GetSkillList(response http.ResponseWriter, request *http.Request) (err error) {
	now := time.Now()
	fmt.Println("HIT -> GetSkillListService.go On ", now.Format("2006-01-02 15:04:05"))
	db := confiq.Connect()

	tokenString := request.Header.Get("Authorization")
	claims, err := confiq.DecodeToken(tokenString)
	params := request.URL.Query()
	if err != nil {
		out.ResponseOut(response, nil, false, constanta.CodeInternalServerErrorResponse, err.Error())
		return
	}

	datas, err := dao.SkillDAO.GetSkillList(db, dao.CustomQueryModel{
		Page:  params.Get("page"),
		Limit: params.Get("limit"),
		Keyword: params.Get("keyword"),
		Id : fmt.Sprintf("%d", claims.Id),
	})

	if err != nil {
		out.ResponseOut(response, nil, false, constanta.CodeInternalServerErrorResponse, constanta.ErrorInternalDB)
		return
	}


	db.Close()
	out.ResponseOut(response, convertRepoToDTO(datas), true, constanta.CodeSuccessResponse, constanta.SuccessGetData)
	return nil
}

func convertRepoToDTO(datas []model.SkillModel) (output []out.SkillResponse) {
	for i:=0;i<len(datas);i++ {
		output = append(output, out.SkillResponse{
			Id:           datas[i].ID.Int64,
			Skill:         datas[i].Skill.String,
			Level:         datas[i].Level.String,
			CreatedAt:     datas[i].CreatedAt.String,
			UpdatedAt:     datas[i].UpdatedAt.String,
		})
	}
	return output
}
