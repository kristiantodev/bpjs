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

func LoginService(response http.ResponseWriter, request *http.Request) (err error) {
	now := time.Now()
	fmt.Println("HIT -> LoginService.go On ", now.Format("2006-01-02 15:04:05"))

	userBody := utils.GetUserBody(request)
	userRepo := userRepository(userBody)
	db := confiq.Connect()

	if userRepo.Username.String == "" || userRepo.Password.String == ""{
		out.ResponseOut(response, nil, false, constanta.CodeBadRequestResponse, "Username/Password tidak boleh kosong")
		return
	}

	user, err := dao.UserDAO.LoginCheck(db, userRepo)
	if err != nil{
		out.ResponseOut(response, nil, false, constanta.CodeInternalServerErrorResponse, constanta.ErrorInternalDB)
		return
	}

	if user.ID.Int64 == 0{
		out.ResponseOut(response, nil, false, constanta.CodeBadRequestResponse, "Username/Password salah")
		return
	}

	token, err := confiq.GenerateToken(user)
	if err != nil{
		out.ResponseOut(response, nil, false, constanta.CodeInternalServerErrorResponse, err.Error())
		return
	}

	db.Close()
	out.ResponseOut(response, token, true, constanta.CodeSuccessResponse, "Login berhasil")
	return nil
}
