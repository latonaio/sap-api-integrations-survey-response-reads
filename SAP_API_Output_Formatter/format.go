package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-survey-response-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToSurveyResponse(raw []byte, l *logger.Logger) ([]SurveyResponse, error) {
	pm := &responses.SurveyResponse{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to SurveyResponse. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	surveyResponse := make([]SurveyResponse, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		surveyResponse = append(surveyResponse, SurveyResponse{
			ObjectID:             data.ObjectID,                         
			ID:                   data.ID,                         
			EntityLastChangedOn:  data.EntityLastChangedOn,                         
			ETag:                 data.ETag,                         
			SurveyCreationDate:   data.SurveyCreationDate,                         
			SurveyResponse:       data.SurveyResponse.Deferred.URI,
		})
	}

	return surveyResponse, nil
}