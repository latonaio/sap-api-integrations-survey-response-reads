# sap-api-integrations-survey-response-reads  
sap-api-integrations-survey-response-reads は、外部システム(特にエッジコンピューティング環境)をSAPと統合することを目的に、SAP API 調査回答データを取得するマイクロサービスです。  
sap-api-integrations-survey-response-reads には、サンプルのAPI Json フォーマットが含まれています。  
sap-api-integrations-survey-response-reads は、オンプレミス版である（＝クラウド版ではない）SAPC4HANA API の利用を前提としています。クラウド版APIを利用する場合は、ご注意ください。  
https://api.sap.com/api/campaign/overview  

## 動作環境
sap-api-integrations-survey-response-reads は、主にエッジコンピューティング環境における動作にフォーカスしています。   
使用する際は、事前に下記の通り エッジコンピューティングの動作環境（推奨/必須）を用意してください。   
・ エッジ Kubernetes （推奨）    
・ AION のリソース （推奨)    
・ OS: LinuxOS （必須）    
・ CPU: ARM/AMD/Intel（いずれか必須） 

## クラウド環境での利用  
sap-api-integrations-survey-response-reads は、外部システムがクラウド環境である場合にSAPと統合するときにおいても、利用可能なように設計されています。  

## 本レポジトリ が 対応する API サービス
sap-api-integrations-survey-response-reads が対応する APIサービス は、次のものです。

* APIサービス概要説明 URL: https://api.sap.com/api/campaign/overview 
* APIサービス名(=baseURL): c4codataapi

## 本レポジトリ に 含まれる API名
sap-api-integrations-survey-response-reads には、次の API をコールするためのリソースが含まれています。  

* SurveyResponse（調査回答 - 調査回答）
* SurveyValuation（調査回答 - 調査評価）
* SurveyValuationItem（調査回答 - 調査評価項目）
* SurveyQuestionAnswers（調査回答 - 調査質問回答）

## API への 値入力条件 の 初期値
sap-api-integrations-survey-response-reads において、API への値入力条件の初期値は、入力ファイルレイアウトの種別毎に、次の通りとなっています。  

### SDC レイアウト

* inoutSDC.SurveyResponse.ID（ID）  
* inoutSDC.SurveyValuation.Version（バージョン）
* inoutSDC.SurveyValuationItem.ProductID（製品ID）
* inoutSDC.SurveyQuestionAnswers.QuestionUUID（質問UUID）

## SAP API Bussiness Hub の API の選択的コール

Latona および AION の SAP 関連リソースでは、Inputs フォルダ下の sample.json の accepter に取得したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAll(もしくは空白)の値を入力することで、全データ（＝全APIの種別）をまとめて取得することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、"SurveyResponse" が指定されています。    
  
```
	"api_schema": "SurveyResponseSurveyResponse",
	"accepter": ["SurveyResponse"],
	"survey_response_code": "4",
	"deleted": false
```
  
* 全データを取得する際のsample.jsonの記載例(2)  

全データを取得する場合、sample.json は以下のように記載します。  

```
	"api_schema": "SurveyResponseSurveyResponse",
	"accepter": ["All"],
	"survey_response_code": "4",
	"deleted": false
```

## 指定されたデータ種別のコール

accepter における データ種別 の指定に基づいて SAP_API_Caller 内の caller.go で API がコールされます。  
caller.go の func() 毎 の 以下の箇所が、指定された API をコールするソースコードです。  

```
func (c *SAPAPICaller) AsyncGetSurveyResponse(iD, version, productID, questionUUID string, accepter []string) {
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
				c.SurveyValuation(version)
				wg.Done()
			}()
		case "SurveyValuationItem":
			func() {
				c.SurveyValuationItem(productID)
				wg.Done()
			}()
		case "SurveyQuestionAnswers":
			func() {
				c.SurveyQuestionAnswers(questionUUID)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}
```

## Output  
本マイクロサービスでは、[golang-logging-library-for-sap](https://github.com/latonaio/golang-logging-library-for-sap) により、以下のようなデータがJSON形式で出力されます。  
以下の sample.json の例は、SAP 調査回答  の 調査回答データ が取得された結果の JSON の例です。  
以下の項目のうち、"ObjectID" ～ "CampaignInboundBusinessTransactionDocumentReference" は、/SAP_API_Output_Formatter/type.go 内 の Type SurveyResponse {} による出力結果です。"cursor" ～ "time"は、golang-logging-library-for-sap による 定型フォーマットの出力結果です。  

```
{
	"cursor": "/Users/latona5/bitbucket/sap-api-integrations-survey-response-reads/SAP_API_Caller/caller.go#L53",
	"function": "sap-api-integrations-survey-response-reads/SAP_API_Caller.(*SAPAPICaller).SurveyResponse",
	"level": "INFO",
	"message": [
		{
			"ObjectID": "00163E04DCEB1ED2BFF75DD91A32A7FF",
			"ID": "4",
			"EntityLastChangedOn": "2019-04-24T17:53:13+09:00",
			"ETag": "2019-04-24T17:53:13+09:00",
			"SurveyCreationDate": "2013-08-08T08:42:14+09:00",
			"SurveyResponse": "https://sandbox.api.sap.com/sap/c4c/odata/v1/c4codataapi/SurveyResponseRootCollection('00163E04DCEB1ED2BFF75DD91A32A7FF')/SurveyResponse"
		}
	],
	"time": "2022-07-20T13:40:17+09:00"
}
```