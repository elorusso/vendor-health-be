package v1

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/elorusso/vendor-health-be/pkg/utils"
)

const (
	amazonURL = "https://www.amazon.com"
)

func AmazonStatusRequestHandler(writer http.ResponseWriter, request *http.Request) {

	//check amazon status
	status, err := utils.NewStatusChecker().CheckStatus(amazonURL)
	if err != nil {
		fmt.Println(err)
		writeErrorResponse(writer)
	}

	//marshal response
	responseBody, err := json.Marshal(status)
	if err != nil {
		fmt.Println(err)
		writeErrorResponse(writer)
	}

	//write success response
	_, err = writer.Write(responseBody)
	if err != nil {
		fmt.Println(err)
		writeErrorResponse(writer)
	}
}
