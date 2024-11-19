package UserService

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

func UserProfileUpdate(response http.ResponseWriter, request *http.Request) (err error) {
	now := time.Now()
	fmt.Println("HIT -> UserProfileUpdateService.go On ", now.Format("2006-01-02 15:04:05"))
	tokenString := request.Header.Get("Authorization")
	claims, err := confiq.DecodeToken(tokenString)
	if err != nil {
		out.ResponseOut(response, nil, false, constanta.CodeInternalServerErrorResponse, err.Error())
		return
	}

	userBody := utils.GetUserBody(request)
	userBody.Id = claims.Id
	userRepo := userRepository(userBody)
	db := confiq.Connect()

	if userRepo.FirstName.String == "" {
		out.ResponseOut(response, nil, false, constanta.CodeBadRequestResponse, "First Name tidak boleh kosong")
		return
	}

	if userRepo.Gender.String != "L" && userRepo.Gender.String != "P"{
		out.ResponseOut(response, nil, false, constanta.CodeBadRequestResponse, "Gender hanya bisa diisi dengan L/P")
		return
	}


	err = dao.UserDAO.UpdateUser(db, userRepo)
	if err != nil{
		out.ResponseOut(response, nil, false, constanta.CodeInternalServerErrorResponse, constanta.ErrorInternalDB)
		return
	}
	db.Close()
	out.ResponseOut(response, nil, true, constanta.CodeSuccessResponse, constanta.SuccessEditData)
	return nil
}