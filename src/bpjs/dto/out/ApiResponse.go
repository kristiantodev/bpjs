package out

import (
	"bpjs/utils"
	"encoding/json"
	"net/http"
)

type APIResponse struct {
	Bpjs BpjsMessage `json:"bpjs"`
}

type BpjsMessage struct {
	Success bool     `json:"success"`
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Content interface{} `json:"content"`
}

func (ar APIResponse) String() string {
	return utils.StructToJSON(ar)
}

func ResponseOut(response http.ResponseWriter, data interface{}, success bool, code int, message string){
	response.Header().Set("Content-type", "application/json")
	var apiResponse APIResponse
	apiResponse.Bpjs.Success = success
	apiResponse.Bpjs.Code = code
	apiResponse.Bpjs.Message = message
	apiResponse.Bpjs.Content = data
	response.WriteHeader(code)
	json.NewEncoder(response).Encode(apiResponse)
}