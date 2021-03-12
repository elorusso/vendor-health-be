package v1

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/elorusso/vendor-health-be/pkg/utils"
)

func AllStatusRequestHandler(writer http.ResponseWriter, request *http.Request) {

	//check amazon status
	amazonStatus, err := utils.NewStatusChecker().CheckStatus(amazonURL)
	if err != nil {
		fmt.Println(err)
		writeErrorResponse(writer)
	}

	//check google status
	googleStatus, err := utils.NewStatusChecker().CheckStatus(googleURL)
	if err != nil {
		fmt.Println(err)
		writeErrorResponse(writer)
	}

	body := []*utils.StatusCheckResponse{
		amazonStatus,
		googleStatus,
	}

	//marshal response
	responseBody, err := json.Marshal(body)
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
