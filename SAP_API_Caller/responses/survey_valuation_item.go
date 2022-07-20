package responses

type SurveyValuationItem struct {
	D struct {
		Results []struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
			ObjectID                 string `json:"ObjectID"`
			ParentObjectID           string `json:"ParentObjectID"`
			ProductID                string `json:"ProductID"`
			ProductUUID              string `json:"ProductUUID"`
			ProductCategoryID        string `json:"ProductCategoryID"`
			ProductCategoryUUID      string `json:"ProductCategoryUUID"`
			PartyID                  string `json:"PartyID"`
			PartyUUID                string `json:"PartyUUID"`
			FinishedIndicator        bool   `json:"FinishedIndicator"`
			ETag                     string `json:"ETag"`
			CompletionRatePerProduct string `json:"CompletionRatePerProduct"`
		} `json:"results"`
	} `json:"d"`
}
