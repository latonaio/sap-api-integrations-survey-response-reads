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

type SurveyValuation struct {
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
}

type SurveyValuationItem struct {
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
}

type SurveyQuestionAnswers struct {
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
}
