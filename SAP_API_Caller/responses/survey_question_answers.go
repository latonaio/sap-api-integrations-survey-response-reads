package responses

type SurveyQuestionAnswers struct {
	D struct {
		Results []struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
			ObjectID                   string `json:"ObjectID"`
			ParentObjectID             string `json:"ParentObjectID"`
			QuestionUUID               string `json:"QuestionUUID"`
			QuestionAnswerOptionUUID   string `json:"QuestionAnswerOptionUUID"`
			QuestionRatingUUID         string `json:"QuestionRatingUUID"`
			QuestionResponseOptionUUID string `json:"QuestionResponseOptionUUID"`
			QuestionTypeCode           string `json:"QuestionTypeCode"`
			QuestionTypeCodeText       string `json:"QuestionTypeCodeText"`
			Amount                     string `json:"Amount"`
			CurrencyCode               string `json:"CurrencyCode"`
			CurrencyCodeText           string `json:"CurrencyCodeText"`
			Date                       string `json:"Date"`
			NumericValue               int    `json:"NumericValue"`
			Quantity                   string `json:"Quantity"`
			UnitOfMeasure              string `json:"UnitOfMeasure"`
			UnitOfMeasureText          string `json:"UnitOfMeasureText"`
			AttributeAssigneeID        string `json:"AttributeAssigneeID"`
			AttributeSetID             string `json:"AttributeSetID"`
			AttributeID                string `json:"AttributeID"`
			Text                       string `json:"Text"`
			ETag                       string `json:"ETag"`
		} `json:"results"`
	} `json:"d"`
}
