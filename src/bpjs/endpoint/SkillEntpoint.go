package endpoint

import (
	"bpjs/service/SkillService"
	"net/http"
)

func SkillEntpointWithId(response http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "PUT":
		SkillService.UpdateSkill(response, request)
		break
	case "GET":
		SkillService.GetDetailSkill(response, request)
		break
	case "DELETE":
		SkillService.DeleteSkill(response, request)
		break
	}
}

func SkillEntpointWithoutId(response http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "POST":
		SkillService.SkillInsert(response, request)
		break
	case "GET":
		SkillService.GetSkillList(response, request)
		break
	}
}