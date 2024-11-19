package endpoint

import (
	"bpjs/service/EducationService"
	"net/http"
)

func EducationEntpointWithId(response http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "PUT":
		EducationService.UpdateSkill(response, request)
		break
	case "GET":
		EducationService.GetDetailEducation(response, request)
		break
	case "DELETE":
		EducationService.DeleteEducation(response, request)
		break
	}
}

func EducationEntpointWithoutId(response http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "POST":
		EducationService.EducationInsert(response, request)
		break
	case "GET":
		EducationService.GetEducationist(response, request)
		break
	}
}
