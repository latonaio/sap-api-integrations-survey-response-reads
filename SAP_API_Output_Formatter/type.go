package sap_api_output_formatter

type Campaign struct {
	ConnectionKey      string `json:"connection_key"`
	Result             bool   `json:"result"`
	RedisKey           string `json:"redis_key"`
	Filepath           string `json:"filepath"`
	APISchema          string `json:"api_schema"`
	SurveyResponseCode string `json:"survey_response_code"`
	Deleted            bool   `json:"deleted"`
}

type SurveyResponse struct {
	ObjectID            string `json:"ObjectID"`
	ID                  string `json:"ID"`
	EntityLastChangedOn string `json:"EntityLastChangedOn"`
	ETag                string `json:"ETag"`
	SurveyCreationDate  string `json:"SurveyCreationDate"`
	SurveyResponse      string `json:"SurveyResponse"`
}