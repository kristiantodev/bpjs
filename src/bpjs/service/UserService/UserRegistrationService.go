package UserService

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

func UserRegistration(response http.ResponseWriter, request *http.Request) (err error) {
	now := time.Now()
	fmt.Println("HIT -> UserRegistrationService.go On ", now.Format("2006-01-02 15:04:05"))

	userBody := utils.GetUserBody(request)
	userRepo := userRepository(userBody)
	db := confiq.Connect()

	if userRepo.Username.String == "" || userRepo.Password.String == ""{
		out.ResponseOut(response, nil, false, constanta.CodeBadRequestResponse, "Username/Password tidak boleh kosong")
		return
	}

	if userRepo.Gender.String != "L" && userRepo.Gender.String != "P"{
		out.ResponseOut(response, nil, false, constanta.CodeBadRequestResponse, "Gender hanya bisa diisi dengan L/P")
		return
	}

	idOnDB := dao.UserDAO.CheckUsername(db, userRepo.Username.String)
	if idOnDB >= 1{
		out.ResponseOut(response, nil, false, constanta.CodeBadRequestResponse, "Username sudah digunakan user lain")
		return
	}

    err = dao.UserDAO.InsertUser(db, userRepo)
    if err != nil{
		out.ResponseOut(response, nil, false, constanta.CodeInternalServerErrorResponse, constanta.ErrorInternalDB)
		return
	}
	db.Close()
	out.ResponseOut(response, nil, true, constanta.CodeSuccessResponse, constanta.SuccessRegistrationData)
	return nil
}

func userRepository(userBody in.UserRequest) model.UserModel  {

	return model.UserModel{
		ID     :  sql.NullInt64{Int64: userBody.Id},
		Username: sql.NullString{String: userBody.Username},
		Password: sql.NullString{String: userBody.Password},
		FirstName: sql.NullString{String: userBody.FirstName},
		LastName: sql.NullString{String: userBody.LastName},
		Gender: sql.NullString{String: userBody.Gender},
		Telephone: sql.NullString{String: userBody.Telephone},
		Email: sql.NullString{String: userBody.Email},
		Address: sql.NullString{String: userBody.Address},
	}

}