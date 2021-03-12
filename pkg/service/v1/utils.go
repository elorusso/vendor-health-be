package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func writeErrorResponse(writer http.ResponseWriter) {
	body := &struct {
		Message string `json:"message"`
	}{
		Message: "Internal Server Error",
	}

	bodyData, err := json.Marshal(body)
	if err != nil {
		fmt.Println(err)
		//just send an empty body back
	}

	writer.WriteHeader(http.StatusInternalServerError)
	_, err = writer.Write(bodyData)
	if err != nil {
		fmt.Println(err)
	}
}
