package responses

type SurveyValuation struct {
	D struct {
		Results []struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
			ObjectID                                    string `json:"ObjectID"`
			ParentObjectID                              string `json:"ParentObjectID"`
			ID                                          string `json:"ID"`
			Version                                     string `json:"Version"`
			BusinessTransactionDocumentID               string `json:"BusinessTransactionDocumentID"`
			BusinessTransactionDocumentUUID             string `json:"BusinessTransactionDocumentUUID"`
			BusinessTransactionDocumentTypeCode         string `json:"BusinessTransactionDocumentTypeCode"`
			BusinessTransactionDocumentTypeCodeText     string `json:"BusinessTransactionDocumentTypeCodeText"`
			ProcessorID                                 string `json:"ProcessorID"`
			ProcessorUUID                               string `json:"ProcessorUUID"`
			RepresentationBinaryObject                  string `json:"RepresentationBinaryObject"`
			RepresentationBinaryObjectResponseIndicator bool   `json:"RepresentationBinaryObjectResponseIndicator"`
			DesignTimeVersionUUID                       string `json:"DesignTimeVersionUUID"`
			LifeCycleStatusCode                         string `json:"LifeCycleStatusCode"`
			LifeCycleStatusCodeText                     string `json:"LifeCycleStatusCodeText"`
			EntityLastChangedOn                         string `json:"EntityLastChangedOn"`
			ValuationCollectionLastChangedBy            string `json:"ValuationCollectionLastChangedBy"`
			ETag                                        string `json:"ETag"`
			SignatureUUID                               string `json:"SignatureUUID"`
			TransactionDocumentUUID                     string `json:"TransactionDocumentUUID"`
		} `json:"results"`
	} `json:"d"`
}
