package sap_api_input_reader

type EC_MC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Document      struct {
		DocumentNo     string `json:"document_no"`
		DeliverTo      string `json:"deliver_to"`
		Quantity       string `json:"quantity"`
		PickedQuantity string `json:"picked_quantity"`
		Price          string `json:"price"`
		Batch          string `json:"batch"`
	} `json:"document"`
	ProductionOrder struct {
		DocumentNo           string `json:"document_no"`
		Status               string `json:"status"`
		DeliverTo            string `json:"deliver_to"`
		Quantity             string `json:"quantity"`
		CompletedQuantity    string `json:"completed_quantity"`
		PlannedStartDate     string `json:"planned_start_date"`
		PlannedValidatedDate string `json:"planned_validated_date"`
		ActualStartDate      string `json:"actual_start_date"`
		ActualValidatedDate  string `json:"actual_validated_date"`
		Batch                string `json:"batch"`
		Work                 struct {
			WorkNo                   string `json:"work_no"`
			Quantity                 string `json:"quantity"`
			CompletedQuantity        string `json:"completed_quantity"`
			ErroredQuantity          string `json:"errored_quantity"`
			Component                string `json:"component"`
			PlannedComponentQuantity string `json:"planned_component_quantity"`
			PlannedStartDate         string `json:"planned_start_date"`
			PlannedStartTime         string `json:"planned_start_time"`
			PlannedValidatedDate     string `json:"planned_validated_date"`
			PlannedValidatedTime     string `json:"planned_validated_time"`
			ActualStartDate          string `json:"actual_start_date"`
			ActualStartTime          string `json:"actual_start_time"`
			ActualValidatedDate      string `json:"actual_validated_date"`
			ActualValidatedTime      string `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"production_order"`
	APISchema      string `json:"api_schema"`
	MaterialCode   string `json:"material_code"`
	Plant_Supplier string `json:"plant/supplier"`
	Stock          string `json:"stock"`
	DocumentType   string `json:"document_type"`
	DocumentNo     string `json:"document_no"`
	PlannedDate    string `json:"planned_date"`
	ValidatedDate  string `json:"validated_date"`
	Deleted        bool   `json:"deleted"`
}

type SDC struct {
	ConnectionKey  string `json:"connection_key"`
	Result         bool   `json:"result"`
	RedisKey       string `json:"redis_key"`
	Filepath       string `json:"filepath"`
	SurveyResponse struct {
		ObjectID            string `json:"ObjectID"`
		ID                  string `json:"ID"`
		EntityLastChangedOn string `json:"EntityLastChangedOn"`
		ETag                string `json:"ETag"`
		SurveyCreationDate  string `json:"SurveyCreationDate"`
		SurveyValuation     struct {
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
			RepresentationBinaryObjectResponseIndicator string `json:"RepresentationBinaryObjectResponseIndicator"`
			DesignTimeVersionUUID                       string `json:"DesignTimeVersionUUID"`
			LifeCycleStatusCode                         string `json:"LifeCycleStatusCode"`
			LifeCycleStatusCodeText                     string `json:"LifeCycleStatusCodeText"`
			EntityLastChangedOn                         string `json:"EntityLastChangedOn"`
			ValuationCollectionLastChangedBy            string `json:"ValuationCollectionLastChangedBy"`
			ETag                                        string `json:"ETag"`
			SignatureUUID                               string `json:"SignatureUUID"`
			TransactionDocumentUUID                     string `json:"TransactionDocumentUUID"`
			SurveyValuationItem                         struct {
				ObjectID                 string `json:"ObjectID"`
				ParentObjectID           string `json:"ParentObjectID"`
				ProductID                string `json:"ProductID"`
				ProductUUID              string `json:"ProductUUID"`
				ProductCategoryID        string `json:"ProductCategoryID"`
				ProductCategoryUUID      string `json:"ProductCategoryUUID"`
				PartyID                  string `json:"PartyID"`
				PartyUUID                string `json:"PartyUUID"`
				FinishedIndicator        string `json:"FinishedIndicator"`
				ETag                     string `json:"ETag"`
				CompletionRatePerProduct string `json:"CompletionRatePerProduct"`
				SurveyQuestionAnswers    struct {
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
					NumericValue               string `json:"NumericValue"`
					Quantity                   string `json:"Quantity"`
					UnitOfMeasure              string `json:"UnitOfMeasure"`
					UnitOfMeasureText          string `json:"UnitOfMeasureText"`
					AttributeAssigneeID        string `json:"AttributeAssigneeID"`
					AttributeSetID             string `json:"AttributeSetID"`
					AttributeID                string `json:"AttributeID"`
					Text                       string `json:"Text"`
					ETag                       string `json:"ETag"`
				} `json:"SurveyQuestionAnswers"`
			} `json:"SurveyValuationItem"`
		} `json:"SurveyValuation"`
	} `json:"SurveyResponse"`
	APISchema    string   `json:"api_schema"`
	Accepter     []string `json:"accepter"`
	CampaignCode string   `json:"campaign_code"`
	Deleted      bool     `json:"deleted"`
}
