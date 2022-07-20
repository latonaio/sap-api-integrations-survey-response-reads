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
			ObjectID:            data.ObjectID,
			ID:                  data.ID,
			EntityLastChangedOn: data.EntityLastChangedOn,
			ETag:                data.ETag,
			SurveyCreationDate:  data.SurveyCreationDate,
		})
	}

	return surveyResponse, nil
}

func ConvertToSurveyValuation(raw []byte, l *logger.Logger) ([]SurveyValuation, error) {
	pm := &responses.SurveyValuation{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to SurveyValuation. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	surveyValuation := make([]SurveyValuation, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		surveyValuation = append(surveyValuation, SurveyValuation{
			ObjectID:                                data.ObjectID,
			ParentObjectID:                          data.ParentObjectID,
			ID:                                      data.ID,
			Version:                                 data.Version,
			BusinessTransactionDocumentID:           data.BusinessTransactionDocumentID,
			BusinessTransactionDocumentUUID:         data.BusinessTransactionDocumentUUID,
			BusinessTransactionDocumentTypeCode:     data.BusinessTransactionDocumentTypeCode,
			BusinessTransactionDocumentTypeCodeText: data.BusinessTransactionDocumentTypeCodeText,
			ProcessorID:                             data.ProcessorID,
			ProcessorUUID:                           data.ProcessorUUID,
			RepresentationBinaryObject:              data.RepresentationBinaryObject,
			RepresentationBinaryObjectResponseIndicator: data.RepresentationBinaryObjectResponseIndicator,
			DesignTimeVersionUUID:                       data.DesignTimeVersionUUID,
			LifeCycleStatusCode:                         data.LifeCycleStatusCode,
			LifeCycleStatusCodeText:                     data.LifeCycleStatusCodeText,
			EntityLastChangedOn:                         data.EntityLastChangedOn,
			ValuationCollectionLastChangedBy:            data.ValuationCollectionLastChangedBy,
			ETag:                                        data.ETag,
			SignatureUUID:                               data.SignatureUUID,
			TransactionDocumentUUID:                     data.TransactionDocumentUUID,
		})
	}

	return surveyValuation, nil
}

func ConvertToSurveyValuationItem(raw []byte, l *logger.Logger) ([]SurveyValuationItem, error) {
	pm := &responses.SurveyValuationItem{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to SurveyValuationItem. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	surveyValuationItem := make([]SurveyValuationItem, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		surveyValuationItem = append(surveyValuationItem, SurveyValuationItem{
			ObjectID:                 data.ObjectID,
			ParentObjectID:           data.ParentObjectID,
			ProductID:                data.ProductID,
			ProductUUID:              data.ProductUUID,
			ProductCategoryID:        data.ProductCategoryID,
			ProductCategoryUUID:      data.ProductCategoryUUID,
			PartyID:                  data.PartyID,
			PartyUUID:                data.PartyUUID,
			FinishedIndicator:        data.FinishedIndicator,
			ETag:                     data.ETag,
			CompletionRatePerProduct: data.CompletionRatePerProduct,
		})
	}

	return surveyValuationItem, nil
}

func ConvertToSurveyQuestionAnswers(raw []byte, l *logger.Logger) ([]SurveyQuestionAnswers, error) {
	pm := &responses.SurveyQuestionAnswers{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to SurveyQuestionAnswers. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	surveyQuestionAnswers := make([]SurveyQuestionAnswers, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		surveyQuestionAnswers = append(surveyQuestionAnswers, SurveyQuestionAnswers{
			ObjectID:                   data.ObjectID,
			ParentObjectID:             data.ParentObjectID,
			QuestionUUID:               data.QuestionUUID,
			QuestionAnswerOptionUUID:   data.QuestionAnswerOptionUUID,
			QuestionRatingUUID:         data.QuestionRatingUUID,
			QuestionResponseOptionUUID: data.QuestionResponseOptionUUID,
			QuestionTypeCode:           data.QuestionTypeCode,
			QuestionTypeCodeText:       data.QuestionTypeCodeText,
			Amount:                     data.Amount,
			CurrencyCode:               data.CurrencyCode,
			CurrencyCodeText:           data.CurrencyCodeText,
			Date:                       data.Date,
			NumericValue:               data.NumericValue,
			Quantity:                   data.Quantity,
			UnitOfMeasure:              data.UnitOfMeasure,
			UnitOfMeasureText:          data.UnitOfMeasureText,
			AttributeAssigneeID:        data.AttributeAssigneeID,
			AttributeSetID:             data.AttributeSetID,
			AttributeID:                data.AttributeID,
			Text:                       data.Text,
			ETag:                       data.ETag,
		})
	}

	return surveyQuestionAnswers, nil
}
