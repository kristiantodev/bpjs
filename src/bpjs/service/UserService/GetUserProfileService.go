package UserService

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

func GetUserProfile(response http.ResponseWriter, request *http.Request) (err error) {
	now := time.Now()
	fmt.Println("HIT -> GetUserProfileService.go On ", now.Format("2006-01-02 15:04:05"))
	tokenString := request.Header.Get("Authorization")
	claims, err := confiq.DecodeToken(tokenString)
	if err != nil {
		out.ResponseOut(response, nil, false, constanta.CodeInternalServerErrorResponse, err.Error())
		return
	}

	db := confiq.Connect()

	user, err := dao.UserDAO.GetUserProfile(db, claims.Id)
	if err != nil{
		out.ResponseOut(response, nil, false, constanta.CodeInternalServerErrorResponse, constanta.ErrorInternalDB)
		return
	}

	outUser := convertRepoToDTOID(user)
	db.Close()
	out.ResponseOut(response, outUser, true, constanta.CodeSuccessResponse, constanta.SuccessGetData)
	return nil
}

func convertRepoToDTOID(datas model.UserModel) out.UserRequest {
	return out.UserRequest{
		FirstName:    datas.FirstName.String,
		LastName:     datas.LastName.String,
		Email:        datas.Email.String,
		Address:      datas.Address.String,
		Gender:       datas.Gender.String,
		Telephone:    datas.Telephone.String,
		CreatedAt:     datas.CreatedAt.String,
		UpdatedAt:     datas.UpdatedAt.String,
	}
}
