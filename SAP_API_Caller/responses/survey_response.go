package responses

type SurveyResponse struct {
	D struct {
		Results []struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
			ObjectID            string `json:"ObjectID"`
			ID                  string `json:"ID"`
			EntityLastChangedOn string `json:"EntityLastChangedOn"`
			ETag                string `json:"ETag"`
			SurveyCreationDate  string `json:"SurveyCreationDate"`
		} `json:"results"`
	} `json:"d"`
}
