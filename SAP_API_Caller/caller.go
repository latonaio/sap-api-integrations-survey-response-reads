package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	sap_api_output_formatter "sap-api-integrations-survey-response-reads/SAP_API_Output_Formatter"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

type SAPAPICaller struct {
	baseURL string
	apiKey  string
	log     *logger.Logger
}

func NewSAPAPICaller(baseUrl string, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL: baseUrl,
		apiKey:  GetApiKey(),
		log:     l,
	}
}

func (c *SAPAPICaller) AsyncGetSurveyResponse(objectID, iD, version, productID, questionUUID string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "SurveyResponse":
			func() {
				c.SurveyResponse(iD)
				wg.Done()
			}()
		case "SurveyValuation":
			func() {
				c.SurveyValuation(iD, version)
				wg.Done()
			}()
		case "SurveyValuationItem":
			func() {
				c.SurveyValuationItem(objectID, productID)
				wg.Done()
			}()
		case "SurveyQuestionAnswers":
			func() {
				c.SurveyQuestionAnswers(objectID, questionUUID)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) SurveyResponse(iD string) {
	surveyResponseData, err := c.callSurveyResponseSrvAPIRequirementSurveyResponse("SurveyResponseRootCollection", iD)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(surveyResponseData)

}

func (c *SAPAPICaller) callSurveyResponseSrvAPIRequirementSurveyResponse(api, iD string) ([]sap_api_output_formatter.SurveyResponse, error) {
	url := strings.Join([]string{c.baseURL, "c4codataapi", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithSurveyResponse(req, iD)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToSurveyResponse(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) SurveyValuation(iD, version string) {
	surveyValuationData, err := c.callSurveyResponseSrvAPIRequirementSurveyValuation("SurveyResponseCollection", iD, version)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(surveyValuationData)

}

func (c *SAPAPICaller) callSurveyResponseSrvAPIRequirementSurveyValuation(api, iD, version string) ([]sap_api_output_formatter.SurveyValuation, error) {
	url := strings.Join([]string{c.baseURL, "c4codataapi", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithSurveyValuation(req, iD, version)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToSurveyValuation(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) SurveyValuationItem(objectID, productID string) {
	surveyValuationItemData, err := c.callSurveyResponseSrvAPIRequirementSurveyValuationItem("SurveyResponseItemCollection", objectID, productID)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(surveyValuationItemData)

}

func (c *SAPAPICaller) callSurveyResponseSrvAPIRequirementSurveyValuationItem(api, objectID, productID string) ([]sap_api_output_formatter.SurveyValuationItem, error) {
	url := strings.Join([]string{c.baseURL, "c4codataapi", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithSurveyValuationItem(req, objectID, productID)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToSurveyValuationItem(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) SurveyQuestionAnswers(objectID, questionUUID string) {
	surveyQuestionAnswersData, err := c.callSurveyResponseSrvAPIRequirementSurveyQuestionAnswers("SurveyQuestionAnswersCollection", objectID, questionUUID)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(surveyQuestionAnswersData)

}

func (c *SAPAPICaller) callSurveyResponseSrvAPIRequirementSurveyQuestionAnswers(api, objectID, questionUUID string) ([]sap_api_output_formatter.SurveyQuestionAnswers, error) {
	url := strings.Join([]string{c.baseURL, "c4codataapi", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithSurveyQuestionAnswers(req, objectID, questionUUID)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToSurveyQuestionAnswers(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) setHeaderAPIKeyAccept(req *http.Request) {
	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")
}

func (c *SAPAPICaller) getQueryWithSurveyResponse(req *http.Request, iD string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("ID eq '%s'", iD))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithSurveyValuation(req *http.Request, iD, version string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("ID eq '%s' and Version eq '%s'", iD, version))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithSurveyValuationItem(req *http.Request, objectID, productID string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("ObjectID eq '%s' and ProductID eq '%s'", objectID, productID))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithSurveyQuestionAnswers(req *http.Request, objectID, questionUUID string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("ObjectID eq '%s' and QuestionUUID eq '%s'", objectID, questionUUID))
	req.URL.RawQuery = params.Encode()
}
