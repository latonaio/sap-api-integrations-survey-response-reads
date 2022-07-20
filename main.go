package main

import (
	sap_api_caller "sap-api-integrations-survey-response-reads/SAP_API_Caller"
	sap_api_input_reader "sap-api-integrations-survey-response-reads/SAP_API_Input_Reader"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
)

func main() {
	l := logger.NewLogger()
	fr := sap_api_input_reader.NewFileReader()
	inoutSDC := fr.ReadSDC("./Inputs/SDC_Survey_Response_Survey_Response_sample.json")
	caller := sap_api_caller.NewSAPAPICaller(
		"https://sandbox.api.sap.com/sap/c4c/odata/v1/", l,
	)

	accepter := inoutSDC.Accepter
	if len(accepter) == 0 || accepter[0] == "All" {
		accepter = []string{
			"SurveyResponse",
		}
	}

	caller.AsyncGetSurveyResponse(
		inoutSDC.SurveyResponse.ID,
		accepter,
	)
}
