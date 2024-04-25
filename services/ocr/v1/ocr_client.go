package v1

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ocr/v1/model"
)

type OcrClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewOcrClient(hcClient *httpclient.HcHttpClient) *OcrClient {
	return &OcrClient{HcClient: hcClient}
}

func OcrClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// RecognizeAcceptanceBill 承兑汇票识别
//
// 识别承兑汇票中的关键信息, 并以json格式返回结构化结果。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OcrClient) RecognizeAcceptanceBill(request *model.RecognizeAcceptanceBillRequest) (*model.RecognizeAcceptanceBillResponse, error) {
	requestDef := GenReqDefForRecognizeAcceptanceBill()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizeAcceptanceBillResponse), nil
	}
}

// RecognizeAcceptanceBillInvoker 承兑汇票识别
func (c *OcrClient) RecognizeAcceptanceBillInvoker(request *model.RecognizeAcceptanceBillRequest) *RecognizeAcceptanceBillInvoker {
	requestDef := GenReqDefForRecognizeAcceptanceBill()
	return &RecognizeAcceptanceBillInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RecognizeAutoClassification 智能分类识别
//
// 检测定位图片上指定要识别的票证（票据、证件或其他文字载体），并对其进行结构化识别。接口以列表形式返回图片上要识别票证的位置坐标、结构化识别的内容以及对应的类别。该接口的使用限制请参见[约束与限制](https://support.huaweicloud.com/productdesc-ocr/ocr_01_0006.html#section3)，详细使用指导请参见[OCR服务使用简介](https://support.huaweicloud.com/qs-ocr/ocr_05_0001.html)章节。
// 计费次数说明：
// 只对识别成功的票证进行计费，识别失败的票证不计费。例如图片中包含三张票证，有两张识别成功，一张识别失败，此时接口计费两次。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OcrClient) RecognizeAutoClassification(request *model.RecognizeAutoClassificationRequest) (*model.RecognizeAutoClassificationResponse, error) {
	requestDef := GenReqDefForRecognizeAutoClassification()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizeAutoClassificationResponse), nil
	}
}

// RecognizeAutoClassificationInvoker 智能分类识别
func (c *OcrClient) RecognizeAutoClassificationInvoker(request *model.RecognizeAutoClassificationRequest) *RecognizeAutoClassificationInvoker {
	requestDef := GenReqDefForRecognizeAutoClassification()
	return &RecognizeAutoClassificationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RecognizeBankReceipt 银行回单识别
//
// 支持对银行回单版式进行文字识别及键值对提取，实现高效的自动化结构化返回。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OcrClient) RecognizeBankReceipt(request *model.RecognizeBankReceiptRequest) (*model.RecognizeBankReceiptResponse, error) {
	requestDef := GenReqDefForRecognizeBankReceipt()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizeBankReceiptResponse), nil
	}
}

// RecognizeBankReceiptInvoker 银行回单识别
func (c *OcrClient) RecognizeBankReceiptInvoker(request *model.RecognizeBankReceiptRequest) *RecognizeBankReceiptInvoker {
	requestDef := GenReqDefForRecognizeBankReceipt()
	return &RecognizeBankReceiptInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RecognizeBankcard 银行卡识别
//
// 识别银行卡上的关键文字信息，并返回识别的结构化结果。该接口的使用限制请参见[约束与限制](https://support.huaweicloud.com/productdesc-ocr/ocr_01_0006.html#section9)，详细使用指导请参见[OCR服务使用简介](https://support.huaweicloud.com/qs-ocr/ocr_05_0001.html)章节。
// 说明：
// 如果图片中包含多张卡证票据，请调用[智能分类识别](https://apiexplorer.developer.huaweicloud.com/apiexplorer/doc?product&#x3D;OCR&amp;api&#x3D;AutoClassification)服务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OcrClient) RecognizeBankcard(request *model.RecognizeBankcardRequest) (*model.RecognizeBankcardResponse, error) {
	requestDef := GenReqDefForRecognizeBankcard()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizeBankcardResponse), nil
	}
}

// RecognizeBankcardInvoker 银行卡识别
func (c *OcrClient) RecognizeBankcardInvoker(request *model.RecognizeBankcardRequest) *RecognizeBankcardInvoker {
	requestDef := GenReqDefForRecognizeBankcard()
	return &RecognizeBankcardInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RecognizeBusinessCard 名片识别
//
// 识别名片图片上的文字信息，并返回识别的结构化结果。支持对多种不同版式名片进行结构化信息提取。该接口的使用限制请参见[约束与限制](https://support.huaweicloud.com/productdesc-ocr/ocr_01_0006.html#section13)，详细使用指导请参见[OCR服务使用简介](https://support.huaweicloud.com/qs-ocr/ocr_05_0001.html)章节。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OcrClient) RecognizeBusinessCard(request *model.RecognizeBusinessCardRequest) (*model.RecognizeBusinessCardResponse, error) {
	requestDef := GenReqDefForRecognizeBusinessCard()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizeBusinessCardResponse), nil
	}
}

// RecognizeBusinessCardInvoker 名片识别
func (c *OcrClient) RecognizeBusinessCardInvoker(request *model.RecognizeBusinessCardRequest) *RecognizeBusinessCardInvoker {
	requestDef := GenReqDefForRecognizeBusinessCard()
	return &RecognizeBusinessCardInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RecognizeBusinessLicense 营业执照识别
//
// 识别营业执照首页图片中的文字信息，并返回识别的结构化结果。该接口的使用限制请参见[约束与限制](https://support.huaweicloud.com/productdesc-ocr/ocr_01_0006.html#section10)，详细使用指导请参见[OCR服务使用简介](https://support.huaweicloud.com/qs-ocr/ocr_05_0001.html)章节。
//
// 说明：
//
// 如果图片中包含多张卡证票据，请调用[智能分类识别](https://apiexplorer.developer.huaweicloud.com/apiexplorer/doc?product&#x3D;OCR&amp;api&#x3D;AutoClassification)服务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OcrClient) RecognizeBusinessLicense(request *model.RecognizeBusinessLicenseRequest) (*model.RecognizeBusinessLicenseResponse, error) {
	requestDef := GenReqDefForRecognizeBusinessLicense()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizeBusinessLicenseResponse), nil
	}
}

// RecognizeBusinessLicenseInvoker 营业执照识别
func (c *OcrClient) RecognizeBusinessLicenseInvoker(request *model.RecognizeBusinessLicenseRequest) *RecognizeBusinessLicenseInvoker {
	requestDef := GenReqDefForRecognizeBusinessLicense()
	return &RecognizeBusinessLicenseInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RecognizeCambodianIdCard 柬文身份证识别
//
// 识别柬文身份证图片中的文字内容，并将识别的结构化结果返回给用户。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OcrClient) RecognizeCambodianIdCard(request *model.RecognizeCambodianIdCardRequest) (*model.RecognizeCambodianIdCardResponse, error) {
	requestDef := GenReqDefForRecognizeCambodianIdCard()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizeCambodianIdCardResponse), nil
	}
}

// RecognizeCambodianIdCardInvoker 柬文身份证识别
func (c *OcrClient) RecognizeCambodianIdCardInvoker(request *model.RecognizeCambodianIdCardRequest) *RecognizeCambodianIdCardInvoker {
	requestDef := GenReqDefForRecognizeCambodianIdCard()
	return &RecognizeCambodianIdCardInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RecognizeChileIdCard 智利身份证识别
//
// 识别智利身份证图片中的文字内容，并返回识别的结构化结果。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OcrClient) RecognizeChileIdCard(request *model.RecognizeChileIdCardRequest) (*model.RecognizeChileIdCardResponse, error) {
	requestDef := GenReqDefForRecognizeChileIdCard()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizeChileIdCardResponse), nil
	}
}

// RecognizeChileIdCardInvoker 智利身份证识别
func (c *OcrClient) RecognizeChileIdCardInvoker(request *model.RecognizeChileIdCardRequest) *RecognizeChileIdCardInvoker {
	requestDef := GenReqDefForRecognizeChileIdCard()
	return &RecognizeChileIdCardInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RecognizeColombiaIdCard 哥伦比亚身份证识别
//
// 识别哥伦比亚身份证中的文字信息，并将识别的结构化结果返回给用户。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OcrClient) RecognizeColombiaIdCard(request *model.RecognizeColombiaIdCardRequest) (*model.RecognizeColombiaIdCardResponse, error) {
	requestDef := GenReqDefForRecognizeColombiaIdCard()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizeColombiaIdCardResponse), nil
	}
}

// RecognizeColombiaIdCardInvoker 哥伦比亚身份证识别
func (c *OcrClient) RecognizeColombiaIdCardInvoker(request *model.RecognizeColombiaIdCardRequest) *RecognizeColombiaIdCardInvoker {
	requestDef := GenReqDefForRecognizeColombiaIdCard()
	return &RecognizeColombiaIdCardInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RecognizeDriverLicense 驾驶证识别
//
// 识别用户上传的驾驶证图片（或者用户提供的华为云上OBS的驾驶证图片文件的URL）中主页与副页的文字内容，并将识别的结果返回给用户。该接口的使用限制请参见[约束与限制](https://support.huaweicloud.com/productdesc-ocr/ocr_01_0006.html#section6)，详细使用指导请参见[OCR服务使用简介](https://support.huaweicloud.com/qs-ocr/ocr_05_0001.html)章节。
//
// 说明：
//
// 如果图片中包含多张卡证票据，请调用[智能分类识别](https://apiexplorer.developer.huaweicloud.com/apiexplorer/doc?product&#x3D;OCR&amp;api&#x3D;AutoClassification)服务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OcrClient) RecognizeDriverLicense(request *model.RecognizeDriverLicenseRequest) (*model.RecognizeDriverLicenseResponse, error) {
	requestDef := GenReqDefForRecognizeDriverLicense()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizeDriverLicenseResponse), nil
	}
}

// RecognizeDriverLicenseInvoker 驾驶证识别
func (c *OcrClient) RecognizeDriverLicenseInvoker(request *model.RecognizeDriverLicenseRequest) *RecognizeDriverLicenseInvoker {
	requestDef := GenReqDefForRecognizeDriverLicense()
	return &RecognizeDriverLicenseInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RecognizeExitEntryPermit 往来港澳台通行证识别
//
// 识别往来港澳台证件图片中的文字内容，并将识别的结构化结果返回给用户。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OcrClient) RecognizeExitEntryPermit(request *model.RecognizeExitEntryPermitRequest) (*model.RecognizeExitEntryPermitResponse, error) {
	requestDef := GenReqDefForRecognizeExitEntryPermit()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizeExitEntryPermitResponse), nil
	}
}

// RecognizeExitEntryPermitInvoker 往来港澳台通行证识别
func (c *OcrClient) RecognizeExitEntryPermitInvoker(request *model.RecognizeExitEntryPermitRequest) *RecognizeExitEntryPermitInvoker {
	requestDef := GenReqDefForRecognizeExitEntryPermit()
	return &RecognizeExitEntryPermitInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RecognizeFinancialStatement 财务报表识别
//
// 识别用户上传的表格图片中的文字内容，并将识别的结果返回给用户。该接口的使用限制请参见[约束与限制](https://support.huaweicloud.com/productdesc-ocr/ocr_01_0006.html#section24)，详细使用指导请参见[OCR服务使用简介](https://support.huaweicloud.com/qs-ocr/ocr_05_0001.html)章节。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OcrClient) RecognizeFinancialStatement(request *model.RecognizeFinancialStatementRequest) (*model.RecognizeFinancialStatementResponse, error) {
	requestDef := GenReqDefForRecognizeFinancialStatement()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizeFinancialStatementResponse), nil
	}
}

// RecognizeFinancialStatementInvoker 财务报表识别
func (c *OcrClient) RecognizeFinancialStatementInvoker(request *model.RecognizeFinancialStatementRequest) *RecognizeFinancialStatementInvoker {
	requestDef := GenReqDefForRecognizeFinancialStatement()
	return &RecognizeFinancialStatementInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RecognizeFlightItinerary 飞机行程单识别
//
// 识别飞机行程单中的文字信息，并返回识别的结构化结果。该接口的使用限制请参见[约束与限制](https://support.huaweicloud.com/productdesc-ocr/ocr_01_0006.html#section20)，详细使用指导请参见[OCR服务使用简介](https://support.huaweicloud.com/qs-ocr/ocr_05_0001.html)章节。
// 说明：
// 如果图片中包含多张卡证票据，请调用[智能分类识别](https://apiexplorer.developer.huaweicloud.com/apiexplorer/doc?product&#x3D;OCR&amp;api&#x3D;AutoClassification)服务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OcrClient) RecognizeFlightItinerary(request *model.RecognizeFlightItineraryRequest) (*model.RecognizeFlightItineraryResponse, error) {
	requestDef := GenReqDefForRecognizeFlightItinerary()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizeFlightItineraryResponse), nil
	}
}

// RecognizeFlightItineraryInvoker 飞机行程单识别
func (c *OcrClient) RecognizeFlightItineraryInvoker(request *model.RecognizeFlightItineraryRequest) *RecognizeFlightItineraryInvoker {
	requestDef := GenReqDefForRecognizeFlightItinerary()
	return &RecognizeFlightItineraryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RecognizeGeneralTable 通用表格识别
//
// 用于识别用户上传的通用表格图片（或者用户提供的华为云上OBS的通用表格图片文件的URL）中的文字内容，并将识别的结果返回给用户。该接口的使用限制请参见[约束与限制](https://support.huaweicloud.com/productdesc-ocr/ocr_01_0006.html#section0)，详细使用指导请参见[OCR服务使用简介](https://support.huaweicloud.com/qs-ocr/ocr_05_0001.html)章节。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OcrClient) RecognizeGeneralTable(request *model.RecognizeGeneralTableRequest) (*model.RecognizeGeneralTableResponse, error) {
	requestDef := GenReqDefForRecognizeGeneralTable()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizeGeneralTableResponse), nil
	}
}

// RecognizeGeneralTableInvoker 通用表格识别
func (c *OcrClient) RecognizeGeneralTableInvoker(request *model.RecognizeGeneralTableRequest) *RecognizeGeneralTableInvoker {
	requestDef := GenReqDefForRecognizeGeneralTable()
	return &RecognizeGeneralTableInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RecognizeGeneralText 通用文字识别
//
// 识别图片上的文字信息，返回识别的文字和坐标。支持扫描文件、电子文档、书籍、票据和表单等多种场景的文字识别。该接口的使用限制请参见[约束与限制](https://support.huaweicloud.com/productdesc-ocr/ocr_01_0006.html#section1)，详细使用指导请参见[OCR服务使用简介](https://support.huaweicloud.com/qs-ocr/ocr_05_0001.html)章节。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OcrClient) RecognizeGeneralText(request *model.RecognizeGeneralTextRequest) (*model.RecognizeGeneralTextResponse, error) {
	requestDef := GenReqDefForRecognizeGeneralText()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizeGeneralTextResponse), nil
	}
}

// RecognizeGeneralTextInvoker 通用文字识别
func (c *OcrClient) RecognizeGeneralTextInvoker(request *model.RecognizeGeneralTextRequest) *RecognizeGeneralTextInvoker {
	requestDef := GenReqDefForRecognizeGeneralText()
	return &RecognizeGeneralTextInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RecognizeHandwriting 手写文字识别
//
// 识别文档中的手写文字信息，并将识别的结构化结果返回给用户。该接口的使用限制请参见[约束与限制](https://support.huaweicloud.com/productdesc-ocr/ocr_01_0006.html#section4)，详细使用指导请参见[OCR服务使用简介](https://support.huaweicloud.com/qs-ocr/ocr_05_0001.html)章节。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OcrClient) RecognizeHandwriting(request *model.RecognizeHandwritingRequest) (*model.RecognizeHandwritingResponse, error) {
	requestDef := GenReqDefForRecognizeHandwriting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizeHandwritingResponse), nil
	}
}

// RecognizeHandwritingInvoker 手写文字识别
func (c *OcrClient) RecognizeHandwritingInvoker(request *model.RecognizeHandwritingRequest) *RecognizeHandwritingInvoker {
	requestDef := GenReqDefForRecognizeHandwriting()
	return &RecognizeHandwritingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RecognizeHealthCode 防疫健康码识别
//
// 支持对全国各地区不同版式的防疫健康码、核酸检测记录、行程卡中的14个关键字段进行结构化识别；支持识别4种健康码颜色，包括绿码、黄码、红码、灰码；支持返回各个关键字段的置信度，以便提高人工校验效率。该接口的使用限制请参见[约束与限制](https://support.huaweicloud.com/productdesc-ocr/ocr_01_0006.html#section26)，详细使用指导请参见[OCR服务使用简介](https://support.huaweicloud.com/qs-ocr/ocr_05_0001.html)章节。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OcrClient) RecognizeHealthCode(request *model.RecognizeHealthCodeRequest) (*model.RecognizeHealthCodeResponse, error) {
	requestDef := GenReqDefForRecognizeHealthCode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizeHealthCodeResponse), nil
	}
}

// RecognizeHealthCodeInvoker 防疫健康码识别
func (c *OcrClient) RecognizeHealthCodeInvoker(request *model.RecognizeHealthCodeRequest) *RecognizeHealthCodeInvoker {
	requestDef := GenReqDefForRecognizeHealthCode()
	return &RecognizeHealthCodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RecognizeHkIdCard 香港身份证识别
//
// 识别香港身份证中的文字内容，并将识别的结果返回给用户。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OcrClient) RecognizeHkIdCard(request *model.RecognizeHkIdCardRequest) (*model.RecognizeHkIdCardResponse, error) {
	requestDef := GenReqDefForRecognizeHkIdCard()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizeHkIdCardResponse), nil
	}
}

// RecognizeHkIdCardInvoker 香港身份证识别
func (c *OcrClient) RecognizeHkIdCardInvoker(request *model.RecognizeHkIdCardRequest) *RecognizeHkIdCardInvoker {
	requestDef := GenReqDefForRecognizeHkIdCard()
	return &RecognizeHkIdCardInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RecognizeHouseholdRegister 户口本识别
//
// 识别户口本中的文字信息，并返回识别的结构化结果。该接口的使用限制请参见[约束与限制](https://support.huaweicloud.com/productdesc-ocr/ocr_01_0006.html#section11)，详细使用指导请参见[OCR服务使用简介](https://support.huaweicloud.com/qs-ocr/ocr_05_0001.html)章节。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OcrClient) RecognizeHouseholdRegister(request *model.RecognizeHouseholdRegisterRequest) (*model.RecognizeHouseholdRegisterResponse, error) {
	requestDef := GenReqDefForRecognizeHouseholdRegister()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizeHouseholdRegisterResponse), nil
	}
}

// RecognizeHouseholdRegisterInvoker 户口本识别
func (c *OcrClient) RecognizeHouseholdRegisterInvoker(request *model.RecognizeHouseholdRegisterRequest) *RecognizeHouseholdRegisterInvoker {
	requestDef := GenReqDefForRecognizeHouseholdRegister()
	return &RecognizeHouseholdRegisterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RecognizeIdCard 身份证识别
//
// 识别身份证图片中的文字内容，并将识别的结果返回给用户。该接口的使用限制请参见[约束与限制](https://support.huaweicloud.com/productdesc-ocr/ocr_01_0006.html#section5)，详细使用指导请参见[OCR服务使用简介](https://support.huaweicloud.com/qs-ocr/ocr_05_0001.html)章节。
//
// 说明：
//
// 身份证识别支持中华人民共和国居民身份证识别。
//
// 如果图片中包含多张卡证票据，请调用[智能分类识别](https://apiexplorer.developer.huaweicloud.com/apiexplorer/doc?product&#x3D;OCR&amp;api&#x3D;AutoClassification)服务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OcrClient) RecognizeIdCard(request *model.RecognizeIdCardRequest) (*model.RecognizeIdCardResponse, error) {
	requestDef := GenReqDefForRecognizeIdCard()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizeIdCardResponse), nil
	}
}

// RecognizeIdCardInvoker 身份证识别
func (c *OcrClient) RecognizeIdCardInvoker(request *model.RecognizeIdCardRequest) *RecognizeIdCardInvoker {
	requestDef := GenReqDefForRecognizeIdCard()
	return &RecognizeIdCardInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RecognizeIdDocument 通用证件识别
//
// 识别身份证件图像，并将识别的结构化结果返回给用户。支持多个国家/地区的身份证、驾驶证和护照，具体国家/地区和证件列表详见表1国家/地区和证件列表。
//
// **表1国家/地区和证件列表**
//
// | 国家/地区  | 英文名称    | 国家/地区代码 country_region | 支持证件类型 id_type      |
// | ---------- | ----------- | ---------------------------- | ------------------------- |
// | 越南       | Vietnam     | VNM                          | PP、DL、ID                |
// | 印度       | India       | IND                          | PP                        |
// | 菲律宾     | Philippines | PHL                          | PP、DL、ID（仅支持UUMID） |
// | 阿尔巴尼亚 | Albania     | ALB                          | PP、DL、ID                |
// | 巴西       | BRAZIL      | BRA                          | PP                        |
// | 印度尼西亚 | INDONESIA   | IDN                          | PP                        |
// | 马来西亚   | MALAYSIA    | MYS                          | PP                        |
// | 尼日利亚   | NIGERIA     | NGA                          | PP                        |
// | 巴基斯坦   | PAKISTAN    | PAK                          | PP                        |
// | 俄罗斯     | RUSSIA      | RUS                          | PP（仅支持国际标准版本）  |
// | 中国台湾   | TAIWAN      | TWN                          | PP                        |
// | 乌克兰     | UKRAINE     | UKR                          | PP                        |
// | 泰国       | THAILAND    | THA                          | ID、PP                    |
// | 智利       | CHILE       | CHL                          | ID、PP                    |
// | 中国香港   | HONGKONG    | HKG                          | ID                        |
//
// - PP: passport,国际护照
// - DL: driving license,驾驶证
// - ID: identification card,各国颁发的身份证类型证件，比如身份证、选民证、社保卡等。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OcrClient) RecognizeIdDocument(request *model.RecognizeIdDocumentRequest) (*model.RecognizeIdDocumentResponse, error) {
	requestDef := GenReqDefForRecognizeIdDocument()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizeIdDocumentResponse), nil
	}
}

// RecognizeIdDocumentInvoker 通用证件识别
func (c *OcrClient) RecognizeIdDocumentInvoker(request *model.RecognizeIdDocumentRequest) *RecognizeIdDocumentInvoker {
	requestDef := GenReqDefForRecognizeIdDocument()
	return &RecognizeIdDocumentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RecognizeInsurancePolicy 保险单识别
//
// 识别保险单图片上的文字信息，并将识别的结构化结果返回给用户。支持对多种版式保险单的扫描图片及手机照片进行结构化信息提取。该接口的使用限制请参见[约束与限制](https://support.huaweicloud.com/productdesc-ocr/ocr_01_0006.html#section23)，详细使用指导请参见[OCR服务使用简介](https://support.huaweicloud.com/qs-ocr/ocr_05_0001.html)章节。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OcrClient) RecognizeInsurancePolicy(request *model.RecognizeInsurancePolicyRequest) (*model.RecognizeInsurancePolicyResponse, error) {
	requestDef := GenReqDefForRecognizeInsurancePolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizeInsurancePolicyResponse), nil
	}
}

// RecognizeInsurancePolicyInvoker 保险单识别
func (c *OcrClient) RecognizeInsurancePolicyInvoker(request *model.RecognizeInsurancePolicyRequest) *RecognizeInsurancePolicyInvoker {
	requestDef := GenReqDefForRecognizeInsurancePolicy()
	return &RecognizeInsurancePolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RecognizeInvoiceVerification 发票验真
//
// 发票验真服务支持10种增值税发票的信息核验，包括增值税专用发票、增值税普通发票、增值税普通发票（卷式）、增值税电子专用发票、增值税电子普通发票、增值税电子普通发票（通行费）、二手车销售统一发票、机动车销售统一发票、区块链电子发票、全电发票，支持返回票面的全部信息。该接口的使用限制请参见[约束与限制](https://support.huaweicloud.com/productdesc-ocr/ocr_01_0006.html#section16)，详细使用指导请参见[OCR服务使用简介](https://support.huaweicloud.com/qs-ocr/ocr_05_0001.html)章节。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OcrClient) RecognizeInvoiceVerification(request *model.RecognizeInvoiceVerificationRequest) (*model.RecognizeInvoiceVerificationResponse, error) {
	requestDef := GenReqDefForRecognizeInvoiceVerification()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizeInvoiceVerificationResponse), nil
	}
}

// RecognizeInvoiceVerificationInvoker 发票验真
func (c *OcrClient) RecognizeInvoiceVerificationInvoker(request *model.RecognizeInvoiceVerificationRequest) *RecognizeInvoiceVerificationInvoker {
	requestDef := GenReqDefForRecognizeInvoiceVerification()
	return &RecognizeInvoiceVerificationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RecognizeLicensePlate 车牌识别
//
// 识别输入图片中的车牌信息，并返回其坐标和内容。该接口的使用限制请参见[约束与限制](https://support.huaweicloud.com/productdesc-ocr/ocr_01_0006.html#section12)，详细使用指导请参见[OCR服务使用简介](https://support.huaweicloud.com/qs-ocr/ocr_05_0001.html)章节。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OcrClient) RecognizeLicensePlate(request *model.RecognizeLicensePlateRequest) (*model.RecognizeLicensePlateResponse, error) {
	requestDef := GenReqDefForRecognizeLicensePlate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizeLicensePlateResponse), nil
	}
}

// RecognizeLicensePlateInvoker 车牌识别
func (c *OcrClient) RecognizeLicensePlateInvoker(request *model.RecognizeLicensePlateRequest) *RecognizeLicensePlateInvoker {
	requestDef := GenReqDefForRecognizeLicensePlate()
	return &RecognizeLicensePlateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RecognizeMacaoIdCard 澳门身份证识别
//
// 识别澳门身份证图片中的文字内容，并将识别的结果返回给用户。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OcrClient) RecognizeMacaoIdCard(request *model.RecognizeMacaoIdCardRequest) (*model.RecognizeMacaoIdCardResponse, error) {
	requestDef := GenReqDefForRecognizeMacaoIdCard()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizeMacaoIdCardResponse), nil
	}
}

// RecognizeMacaoIdCardInvoker 澳门身份证识别
func (c *OcrClient) RecognizeMacaoIdCardInvoker(request *model.RecognizeMacaoIdCardRequest) *RecognizeMacaoIdCardInvoker {
	requestDef := GenReqDefForRecognizeMacaoIdCard()
	return &RecognizeMacaoIdCardInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RecognizeMainlandTravelPermit 港澳台居民来往内地通行证识别
//
// 识别港澳居民来往内地通行证上的文字内容，并将识别的结构化结果返回给用户。支持港澳居民来往内地通行证和台湾居民来往内地通行证两种卡证。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OcrClient) RecognizeMainlandTravelPermit(request *model.RecognizeMainlandTravelPermitRequest) (*model.RecognizeMainlandTravelPermitResponse, error) {
	requestDef := GenReqDefForRecognizeMainlandTravelPermit()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizeMainlandTravelPermitResponse), nil
	}
}

// RecognizeMainlandTravelPermitInvoker 港澳台居民来往内地通行证识别
func (c *OcrClient) RecognizeMainlandTravelPermitInvoker(request *model.RecognizeMainlandTravelPermitRequest) *RecognizeMainlandTravelPermitInvoker {
	requestDef := GenReqDefForRecognizeMainlandTravelPermit()
	return &RecognizeMainlandTravelPermitInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RecognizeMvsInvoice 机动车销售发票识别
//
// 识别机动车销售发票、二手车销售发票图片（服务能自动分辨两种类型，返回对应的字段）中的文字内容，并将识别的结果以JSON格式返回给用户。该接口的使用限制请参见[约束与限制](https://support.huaweicloud.com/productdesc-ocr/ocr_01_0006.html#section17)，详细使用指导请参见[OCR服务使用简介](https://support.huaweicloud.com/qs-ocr/ocr_05_0001.html)章节。
// 说明：
// 该增值税发票仅限于中华人民共和国境内使用的增值税发票。
// 如果图片中包含多张卡证票据，请调用[智能分类识别](https://apiexplorer.developer.huaweicloud.com/apiexplorer/doc?product&#x3D;OCR&amp;api&#x3D;AutoClassification)服务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OcrClient) RecognizeMvsInvoice(request *model.RecognizeMvsInvoiceRequest) (*model.RecognizeMvsInvoiceResponse, error) {
	requestDef := GenReqDefForRecognizeMvsInvoice()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizeMvsInvoiceResponse), nil
	}
}

// RecognizeMvsInvoiceInvoker 机动车销售发票识别
func (c *OcrClient) RecognizeMvsInvoiceInvoker(request *model.RecognizeMvsInvoiceRequest) *RecognizeMvsInvoiceInvoker {
	requestDef := GenReqDefForRecognizeMvsInvoice()
	return &RecognizeMvsInvoiceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RecognizeMyanmarDriverLicense 缅文驾驶证识别
//
// 识别缅甸驾驶证中的文字信息，并返回识别的结构化结果。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OcrClient) RecognizeMyanmarDriverLicense(request *model.RecognizeMyanmarDriverLicenseRequest) (*model.RecognizeMyanmarDriverLicenseResponse, error) {
	requestDef := GenReqDefForRecognizeMyanmarDriverLicense()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizeMyanmarDriverLicenseResponse), nil
	}
}

// RecognizeMyanmarDriverLicenseInvoker 缅文驾驶证识别
func (c *OcrClient) RecognizeMyanmarDriverLicenseInvoker(request *model.RecognizeMyanmarDriverLicenseRequest) *RecognizeMyanmarDriverLicenseInvoker {
	requestDef := GenReqDefForRecognizeMyanmarDriverLicense()
	return &RecognizeMyanmarDriverLicenseInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RecognizeMyanmarIdcard 缅文身份证识别
//
// 识别缅文身份证中的文字信息，并返回识别的结构化结果。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OcrClient) RecognizeMyanmarIdcard(request *model.RecognizeMyanmarIdcardRequest) (*model.RecognizeMyanmarIdcardResponse, error) {
	requestDef := GenReqDefForRecognizeMyanmarIdcard()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizeMyanmarIdcardResponse), nil
	}
}

// RecognizeMyanmarIdcardInvoker 缅文身份证识别
func (c *OcrClient) RecognizeMyanmarIdcardInvoker(request *model.RecognizeMyanmarIdcardRequest) *RecognizeMyanmarIdcardInvoker {
	requestDef := GenReqDefForRecognizeMyanmarIdcard()
	return &RecognizeMyanmarIdcardInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RecognizePassport 护照识别
//
// 识别用户上传的护照首页图片中的文字信息，并返回识别的结构化结果。当前版本支持中国护照的全字段识别。外国护照支持护照下方两行国际标准化的机读码识别，并可从中提取6-7个关键字段信息。该接口的使用限制请参见[约束与限制](https://support.huaweicloud.com/productdesc-ocr/ocr_01_0006.html#section8)，详细使用指导请参见[OCR服务使用简介](https://support.huaweicloud.com/qs-ocr/ocr_05_0001.html)章节。
// 说明：
// 如果图片中包含多张卡证票据，请调用[智能分类识别](https://apiexplorer.developer.huaweicloud.com/apiexplorer/doc?product&#x3D;OCR&amp;api&#x3D;AutoClassification)服务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OcrClient) RecognizePassport(request *model.RecognizePassportRequest) (*model.RecognizePassportResponse, error) {
	requestDef := GenReqDefForRecognizePassport()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizePassportResponse), nil
	}
}

// RecognizePassportInvoker 护照识别
func (c *OcrClient) RecognizePassportInvoker(request *model.RecognizePassportRequest) *RecognizePassportInvoker {
	requestDef := GenReqDefForRecognizePassport()
	return &RecognizePassportInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RecognizePcrTestRecord 核酸检测记录识别
//
// 识别核酸检测记录中的文字信息，并将识别的结构化结果返回给用户。PCR，全称Polymerase chain reaction,即聚合酶链式反应。PCR-test也为大众所认知为新型冠状病毒核酸检测测试。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OcrClient) RecognizePcrTestRecord(request *model.RecognizePcrTestRecordRequest) (*model.RecognizePcrTestRecordResponse, error) {
	requestDef := GenReqDefForRecognizePcrTestRecord()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizePcrTestRecordResponse), nil
	}
}

// RecognizePcrTestRecordInvoker 核酸检测记录识别
func (c *OcrClient) RecognizePcrTestRecordInvoker(request *model.RecognizePcrTestRecordRequest) *RecognizePcrTestRecordInvoker {
	requestDef := GenReqDefForRecognizePcrTestRecord()
	return &RecognizePcrTestRecordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RecognizePeruIdCard 秘鲁身份证识别
//
// 识别秘鲁身份证图片中的文字内容，并将识别的结构化结果返回给用户。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OcrClient) RecognizePeruIdCard(request *model.RecognizePeruIdCardRequest) (*model.RecognizePeruIdCardResponse, error) {
	requestDef := GenReqDefForRecognizePeruIdCard()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizePeruIdCardResponse), nil
	}
}

// RecognizePeruIdCardInvoker 秘鲁身份证识别
func (c *OcrClient) RecognizePeruIdCardInvoker(request *model.RecognizePeruIdCardRequest) *RecognizePeruIdCardInvoker {
	requestDef := GenReqDefForRecognizePeruIdCard()
	return &RecognizePeruIdCardInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RecognizeQualificationCertificate 道路运输从业资格证识别
//
// 识别道路运输从业资格证上的关键文字信息，并返回识别的结构化结果。该接口的使用限制请参见[约束与限制](https://support.huaweicloud.com/productdesc-ocr/ocr_01_0006.html#section25)，详细使用指导请参见[OCR服务使用简介](https://support.huaweicloud.com/qs-ocr/ocr_05_0001.html)章节。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OcrClient) RecognizeQualificationCertificate(request *model.RecognizeQualificationCertificateRequest) (*model.RecognizeQualificationCertificateResponse, error) {
	requestDef := GenReqDefForRecognizeQualificationCertificate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizeQualificationCertificateResponse), nil
	}
}

// RecognizeQualificationCertificateInvoker 道路运输从业资格证识别
func (c *OcrClient) RecognizeQualificationCertificateInvoker(request *model.RecognizeQualificationCertificateRequest) *RecognizeQualificationCertificateInvoker {
	requestDef := GenReqDefForRecognizeQualificationCertificate()
	return &RecognizeQualificationCertificateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RecognizeQuotaInvoice 定额发票识别
//
// 识别定额发票中的文字信息，并返回识别的结构化结果。该接口的使用限制请参见[约束与限制](https://support.huaweicloud.com/productdesc-ocr/ocr_01_0006.html#section21)，详细使用指导请参见[OCR服务使用简介](https://support.huaweicloud.com/qs-ocr/ocr_05_0001.html)章节。
//
// 说明：
//
// 如果图片中包含多张卡证票据，请调用[智能分类识别](https://apiexplorer.developer.huaweicloud.com/apiexplorer/doc?product&#x3D;OCR&amp;api&#x3D;AutoClassification)服务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OcrClient) RecognizeQuotaInvoice(request *model.RecognizeQuotaInvoiceRequest) (*model.RecognizeQuotaInvoiceResponse, error) {
	requestDef := GenReqDefForRecognizeQuotaInvoice()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizeQuotaInvoiceResponse), nil
	}
}

// RecognizeQuotaInvoiceInvoker 定额发票识别
func (c *OcrClient) RecognizeQuotaInvoiceInvoker(request *model.RecognizeQuotaInvoiceRequest) *RecognizeQuotaInvoiceInvoker {
	requestDef := GenReqDefForRecognizeQuotaInvoice()
	return &RecognizeQuotaInvoiceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RecognizeRealEstateCertificate 不动产证识别
//
// 识别不动产证中的文字信息，并返回识别的结构化结果。该接口的使用限制请参见[约束与限制](https://support.huaweicloud.com/productdesc-ocr/ocr_01_0006.html#section11)，详细使用指导请参见[OCR服务使用简介](https://support.huaweicloud.com/qs-ocr/ocr_05_0001.html)章节。
// 说明： 如果图片中包含多张卡证票据，请调用[智能分类识别](https://apiexplorer.developer.huaweicloud.com/apiexplorer/doc?product&#x3D;OCR&amp;api&#x3D;AutoClassification)服务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OcrClient) RecognizeRealEstateCertificate(request *model.RecognizeRealEstateCertificateRequest) (*model.RecognizeRealEstateCertificateResponse, error) {
	requestDef := GenReqDefForRecognizeRealEstateCertificate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizeRealEstateCertificateResponse), nil
	}
}

// RecognizeRealEstateCertificateInvoker 不动产证识别
func (c *OcrClient) RecognizeRealEstateCertificateInvoker(request *model.RecognizeRealEstateCertificateRequest) *RecognizeRealEstateCertificateInvoker {
	requestDef := GenReqDefForRecognizeRealEstateCertificate()
	return &RecognizeRealEstateCertificateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RecognizeSeal 印章识别
//
// 检测和识别合同文件或常用票据中的印章，并可擦除和提取图片中的印章，通过JSON格式返回印章检测、识别、擦除和提取的结果。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OcrClient) RecognizeSeal(request *model.RecognizeSealRequest) (*model.RecognizeSealResponse, error) {
	requestDef := GenReqDefForRecognizeSeal()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizeSealResponse), nil
	}
}

// RecognizeSealInvoker 印章识别
func (c *OcrClient) RecognizeSealInvoker(request *model.RecognizeSealRequest) *RecognizeSealInvoker {
	requestDef := GenReqDefForRecognizeSeal()
	return &RecognizeSealInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RecognizeSmartDocumentRecognizer 智能文档解析
//
// 对证件、票据、表单等任意版式文档进行键值对提取、文字识别、以及表格识别等任务，实现进阶高效的自动化结构化返回。该接口的使用限制请参见[约束与限制](https://support.huaweicloud.com/productdesc-ocr/ocr_01_0006.html#section11)，详细使用指导请参见[OCR服务使用简介](https://support.huaweicloud.com/qs-ocr/ocr_05_0001.html)章节。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OcrClient) RecognizeSmartDocumentRecognizer(request *model.RecognizeSmartDocumentRecognizerRequest) (*model.RecognizeSmartDocumentRecognizerResponse, error) {
	requestDef := GenReqDefForRecognizeSmartDocumentRecognizer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizeSmartDocumentRecognizerResponse), nil
	}
}

// RecognizeSmartDocumentRecognizerInvoker 智能文档解析
func (c *OcrClient) RecognizeSmartDocumentRecognizerInvoker(request *model.RecognizeSmartDocumentRecognizerRequest) *RecognizeSmartDocumentRecognizerInvoker {
	requestDef := GenReqDefForRecognizeSmartDocumentRecognizer()
	return &RecognizeSmartDocumentRecognizerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RecognizeTaxiInvoice 出租车发票识别
//
// 识别出租车发票中的文字信息，并返回识别的结构化结果。该接口的使用限制请参见[约束与限制](https://support.huaweicloud.com/productdesc-ocr/ocr_01_0006.html#section18)，详细使用指导请参见[OCR服务使用简介](https://support.huaweicloud.com/qs-ocr/ocr_05_0001.html)章节。
//
// 说明：
// 如果图片中包含多张卡证票据，请调用[智能分类识别](https://apiexplorer.developer.huaweicloud.com/apiexplorer/doc?product&#x3D;OCR&amp;api&#x3D;AutoClassification)服务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OcrClient) RecognizeTaxiInvoice(request *model.RecognizeTaxiInvoiceRequest) (*model.RecognizeTaxiInvoiceResponse, error) {
	requestDef := GenReqDefForRecognizeTaxiInvoice()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizeTaxiInvoiceResponse), nil
	}
}

// RecognizeTaxiInvoiceInvoker 出租车发票识别
func (c *OcrClient) RecognizeTaxiInvoiceInvoker(request *model.RecognizeTaxiInvoiceRequest) *RecognizeTaxiInvoiceInvoker {
	requestDef := GenReqDefForRecognizeTaxiInvoice()
	return &RecognizeTaxiInvoiceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RecognizeThailandIdcard 泰文身份证识别
//
// 识别泰国身份证中的文字信息，并返回识别的结构化结果。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OcrClient) RecognizeThailandIdcard(request *model.RecognizeThailandIdcardRequest) (*model.RecognizeThailandIdcardResponse, error) {
	requestDef := GenReqDefForRecognizeThailandIdcard()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizeThailandIdcardResponse), nil
	}
}

// RecognizeThailandIdcardInvoker 泰文身份证识别
func (c *OcrClient) RecognizeThailandIdcardInvoker(request *model.RecognizeThailandIdcardRequest) *RecognizeThailandIdcardInvoker {
	requestDef := GenReqDefForRecognizeThailandIdcard()
	return &RecognizeThailandIdcardInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RecognizeThailandLicensePlate 泰国车牌识别
//
// 识别泰国车牌图片中的车牌信息，并返回识别的结构化结果。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OcrClient) RecognizeThailandLicensePlate(request *model.RecognizeThailandLicensePlateRequest) (*model.RecognizeThailandLicensePlateResponse, error) {
	requestDef := GenReqDefForRecognizeThailandLicensePlate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizeThailandLicensePlateResponse), nil
	}
}

// RecognizeThailandLicensePlateInvoker 泰国车牌识别
func (c *OcrClient) RecognizeThailandLicensePlateInvoker(request *model.RecognizeThailandLicensePlateRequest) *RecognizeThailandLicensePlateInvoker {
	requestDef := GenReqDefForRecognizeThailandLicensePlate()
	return &RecognizeThailandLicensePlateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RecognizeTollInvoice 车辆通行费发票识别
//
// 识别车辆通行费发票中的文字信息，并返回识别的结构化结果。该接口的使用限制请参见[约束与限制](https://support.huaweicloud.com/productdesc-ocr/ocr_01_0006.html#section19)，详细使用指导请参见[OCR服务使用简介](https://support.huaweicloud.com/qs-ocr/ocr_05_0001.html)章节。
// 说明：
// 如果图片中包含多张卡证票据，请调用[智能分类识别](https://apiexplorer.developer.huaweicloud.com/apiexplorer/doc?product&#x3D;OCR&amp;api&#x3D;AutoClassification)服务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OcrClient) RecognizeTollInvoice(request *model.RecognizeTollInvoiceRequest) (*model.RecognizeTollInvoiceResponse, error) {
	requestDef := GenReqDefForRecognizeTollInvoice()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizeTollInvoiceResponse), nil
	}
}

// RecognizeTollInvoiceInvoker 车辆通行费发票识别
func (c *OcrClient) RecognizeTollInvoiceInvoker(request *model.RecognizeTollInvoiceRequest) *RecognizeTollInvoiceInvoker {
	requestDef := GenReqDefForRecognizeTollInvoice()
	return &RecognizeTollInvoiceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RecognizeTrainTicket 火车票识别
//
// 识别火车票中的文字信息，并返回识别的结构化结果。该接口的使用限制请参见[约束与限制](https://support.huaweicloud.com/productdesc-ocr/ocr_01_0006.html#section22)，详细使用指导请参见[OCR服务使用简介](https://support.huaweicloud.com/qs-ocr/ocr_05_0001.html)章节。
// 说明：
// 如果图片中包含多张卡证票据，请调用[智能分类识别](https://apiexplorer.developer.huaweicloud.com/apiexplorer/doc?product&#x3D;OCR&amp;api&#x3D;AutoClassification)服务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OcrClient) RecognizeTrainTicket(request *model.RecognizeTrainTicketRequest) (*model.RecognizeTrainTicketResponse, error) {
	requestDef := GenReqDefForRecognizeTrainTicket()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizeTrainTicketResponse), nil
	}
}

// RecognizeTrainTicketInvoker 火车票识别
func (c *OcrClient) RecognizeTrainTicketInvoker(request *model.RecognizeTrainTicketRequest) *RecognizeTrainTicketInvoker {
	requestDef := GenReqDefForRecognizeTrainTicket()
	return &RecognizeTrainTicketInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RecognizeTransportationLicense 道路运输证识别
//
// 识别道路运输证首页中的文字信息，并返回识别的结构化结果。该接口的使用限制请参见[约束与限制](https://support.huaweicloud.com/productdesc-ocr/ocr_01_0006.html#section11)，详细使用指导请参见[OCR服务使用简介](https://support.huaweicloud.com/qs-ocr/ocr_05_0001.html)章节。
// 说明： 如果图片中包含多张卡证票据，请调用[智能分类识别](https://apiexplorer.developer.huaweicloud.com/apiexplorer/doc?product&#x3D;OCR&amp;api&#x3D;AutoClassification)服务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OcrClient) RecognizeTransportationLicense(request *model.RecognizeTransportationLicenseRequest) (*model.RecognizeTransportationLicenseResponse, error) {
	requestDef := GenReqDefForRecognizeTransportationLicense()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizeTransportationLicenseResponse), nil
	}
}

// RecognizeTransportationLicenseInvoker 道路运输证识别
func (c *OcrClient) RecognizeTransportationLicenseInvoker(request *model.RecognizeTransportationLicenseRequest) *RecognizeTransportationLicenseInvoker {
	requestDef := GenReqDefForRecognizeTransportationLicense()
	return &RecognizeTransportationLicenseInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RecognizeVatInvoice 增值税发票识别
//
// 识别增值税发票的类别，以及图片中的文字内容，并以json格式返回识别的结构化结果，不支持真伪验证。该接口的使用限制请参见[约束与限制](https://support.huaweicloud.com/productdesc-ocr/ocr_01_0006.html#section15)，详细使用指导请参见[OCR服务使用简介](https://support.huaweicloud.com/qs-ocr/ocr_05_0001.html)章节。
// 说明：
// 该增值税发票仅限于中华人民共和国境内使用的增值税发票。
// 支持的增值税发票包括：增值税专用发票、增值税普通发票、增值税电子普通发票、增值税电子专用发票、增值税电子普通发票（通行费）、增值税普通发票（卷票）。
// 如果图片中包含多张卡证票据，请调用智能分类识别服务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OcrClient) RecognizeVatInvoice(request *model.RecognizeVatInvoiceRequest) (*model.RecognizeVatInvoiceResponse, error) {
	requestDef := GenReqDefForRecognizeVatInvoice()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizeVatInvoiceResponse), nil
	}
}

// RecognizeVatInvoiceInvoker 增值税发票识别
func (c *OcrClient) RecognizeVatInvoiceInvoker(request *model.RecognizeVatInvoiceRequest) *RecognizeVatInvoiceInvoker {
	requestDef := GenReqDefForRecognizeVatInvoice()
	return &RecognizeVatInvoiceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RecognizeVehicleCertificate 车辆合格证识别
//
// 识别车辆合格证中的文字信息，并返回识别的结构化结果。该接口的使用限制请参见[约束与限制](https://support.huaweicloud.com/productdesc-ocr/ocr_01_0006.html#section11)，详细使用指导请参见[OCR服务使用简介](https://support.huaweicloud.com/qs-ocr/ocr_05_0001.html)章节。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OcrClient) RecognizeVehicleCertificate(request *model.RecognizeVehicleCertificateRequest) (*model.RecognizeVehicleCertificateResponse, error) {
	requestDef := GenReqDefForRecognizeVehicleCertificate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizeVehicleCertificateResponse), nil
	}
}

// RecognizeVehicleCertificateInvoker 车辆合格证识别
func (c *OcrClient) RecognizeVehicleCertificateInvoker(request *model.RecognizeVehicleCertificateRequest) *RecognizeVehicleCertificateInvoker {
	requestDef := GenReqDefForRecognizeVehicleCertificate()
	return &RecognizeVehicleCertificateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RecognizeVehicleLicense 行驶证识别
//
// 识别用户上传的行驶证图片（或者用户提供的华为云上OBS的行驶证图片文件的URL）中主页和副页的文字内容，并将识别的结果返回给用户。该接口的使用限制请参见[约束与限制](https://support.huaweicloud.com/productdesc-ocr/ocr_01_0006.html#section7)，详细使用指导请参见[OCR服务使用简介](https://support.huaweicloud.com/qs-ocr/ocr_05_0001.html)章节。
// 说明：
// 如果图片中包含多张卡证票据，请调用[智能分类识别](https://apiexplorer.developer.huaweicloud.com/apiexplorer/doc?product&#x3D;OCR&amp;api&#x3D;AutoClassification)服务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OcrClient) RecognizeVehicleLicense(request *model.RecognizeVehicleLicenseRequest) (*model.RecognizeVehicleLicenseResponse, error) {
	requestDef := GenReqDefForRecognizeVehicleLicense()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizeVehicleLicenseResponse), nil
	}
}

// RecognizeVehicleLicenseInvoker 行驶证识别
func (c *OcrClient) RecognizeVehicleLicenseInvoker(request *model.RecognizeVehicleLicenseRequest) *RecognizeVehicleLicenseInvoker {
	requestDef := GenReqDefForRecognizeVehicleLicense()
	return &RecognizeVehicleLicenseInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RecognizeVietnamIdCard 越南身份证识别
//
// 识别越南身份证中的文字信息，并将识别的结构化结果返回给用户。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OcrClient) RecognizeVietnamIdCard(request *model.RecognizeVietnamIdCardRequest) (*model.RecognizeVietnamIdCardResponse, error) {
	requestDef := GenReqDefForRecognizeVietnamIdCard()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizeVietnamIdCardResponse), nil
	}
}

// RecognizeVietnamIdCardInvoker 越南身份证识别
func (c *OcrClient) RecognizeVietnamIdCardInvoker(request *model.RecognizeVietnamIdCardRequest) *RecognizeVietnamIdCardInvoker {
	requestDef := GenReqDefForRecognizeVietnamIdCard()
	return &RecognizeVietnamIdCardInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RecognizeWaybillElectronic 电子面单识别
//
// 识别用户上传的电子面单图片中的文字内容，并将识别的结果以json格式返回给用户。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OcrClient) RecognizeWaybillElectronic(request *model.RecognizeWaybillElectronicRequest) (*model.RecognizeWaybillElectronicResponse, error) {
	requestDef := GenReqDefForRecognizeWaybillElectronic()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizeWaybillElectronicResponse), nil
	}
}

// RecognizeWaybillElectronicInvoker 电子面单识别
func (c *OcrClient) RecognizeWaybillElectronicInvoker(request *model.RecognizeWaybillElectronicRequest) *RecognizeWaybillElectronicInvoker {
	requestDef := GenReqDefForRecognizeWaybillElectronic()
	return &RecognizeWaybillElectronicInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RecognizeWebImage 网络图片识别
//
// 识别网络图片中的文字内容，并返回识别的结构化结果。该接口的使用限制请参见[约束与限制](https://support.huaweicloud.com/productdesc-ocr/ocr_01_0006.html#section2)，详细使用指导请参见[OCR服务使用简介](https://support.huaweicloud.com/qs-ocr/ocr_05_0001.html)章节。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OcrClient) RecognizeWebImage(request *model.RecognizeWebImageRequest) (*model.RecognizeWebImageResponse, error) {
	requestDef := GenReqDefForRecognizeWebImage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizeWebImageResponse), nil
	}
}

// RecognizeWebImageInvoker 网络图片识别
func (c *OcrClient) RecognizeWebImageInvoker(request *model.RecognizeWebImageRequest) *RecognizeWebImageInvoker {
	requestDef := GenReqDefForRecognizeWebImage()
	return &RecognizeWebImageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RecognizeCustomTemplate 自定义模板OCR
//
// 自定义模板OCR，支持用户自定义模板，对于版式固定的各种票据和卡证，通过可视化界面操作，指定需要识别的关键字段，实现用户特定格式图片的自动识别和结构化提取。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OcrClient) RecognizeCustomTemplate(request *model.RecognizeCustomTemplateRequest) (*model.RecognizeCustomTemplateResponse, error) {
	requestDef := GenReqDefForRecognizeCustomTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizeCustomTemplateResponse), nil
	}
}

// RecognizeCustomTemplateInvoker 自定义模板OCR
func (c *OcrClient) RecognizeCustomTemplateInvoker(request *model.RecognizeCustomTemplateRequest) *RecognizeCustomTemplateInvoker {
	requestDef := GenReqDefForRecognizeCustomTemplate()
	return &RecognizeCustomTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RecognizeVin VIN码识别
//
// 识别图片中的车架号信息，并将识别结果返回给用户。该接口的使用限制请参见[约束与限制](https://support.huaweicloud.com/productdesc-ocr/ocr_01_0006.html#section14)，详细使用指导请参见[OCR服务使用简介](https://support.huaweicloud.com/qs-ocr/ocr_05_0001.html)章节。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OcrClient) RecognizeVin(request *model.RecognizeVinRequest) (*model.RecognizeVinResponse, error) {
	requestDef := GenReqDefForRecognizeVin()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizeVinResponse), nil
	}
}

// RecognizeVinInvoker VIN码识别
func (c *OcrClient) RecognizeVinInvoker(request *model.RecognizeVinRequest) *RecognizeVinInvoker {
	requestDef := GenReqDefForRecognizeVin()
	return &RecognizeVinInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
