package v1

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ocr/v1/model"
)

type OcrClient struct {
	HcClient *http_client.HcHttpClient
}

func NewOcrClient(hcClient *http_client.HcHttpClient) *OcrClient {
	return &OcrClient{HcClient: hcClient}
}

func OcrClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// 智能分类识别
//
// 检测定位图片上指定要识别的票证（票据、证件或其他文字载体），并对其进行结构化识别。接口以列表形式返回图片上要识别票证的位置坐标、结构化识别的内容以及对应的类别。
//
// 计费次数说明：
//
// 只对识别成功的票证进行计费，识别失败的票证不计费。例如图片中包含三张票证，有两张识别成功，一张识别失败，此时接口计费两次。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *OcrClient) RecognizeAutoClassification(request *model.RecognizeAutoClassificationRequest) (*model.RecognizeAutoClassificationResponse, error) {
	requestDef := GenReqDefForRecognizeAutoClassification()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizeAutoClassificationResponse), nil
	}
}

// 银行卡识别
//
// 识别银行卡上的关键文字信息，并返回识别的结构化结果。
//
// 说明：
//
// 如果图片中包含多张卡证票据，请调用[智能分类识别](https://apiexplorer.developer.huaweicloud.com/apiexplorer/doc?product&#x3D;OCR&amp;api&#x3D;AutoClassification)服务。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *OcrClient) RecognizeBankcard(request *model.RecognizeBankcardRequest) (*model.RecognizeBankcardResponse, error) {
	requestDef := GenReqDefForRecognizeBankcard()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizeBankcardResponse), nil
	}
}

// 名片识别
//
// 识别名片图片上的文字信息，并返回识别的结构化结果。支持对多种不同版式名片进行结构化信息提取。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *OcrClient) RecognizeBusinessCard(request *model.RecognizeBusinessCardRequest) (*model.RecognizeBusinessCardResponse, error) {
	requestDef := GenReqDefForRecognizeBusinessCard()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizeBusinessCardResponse), nil
	}
}

// 营业执照识别
//
// 识别营业执照首页图片中的文字信息，并返回识别的结构化结果。
//
// 说明：
//
// 如果图片中包含多张卡证票据，请调用[智能分类识别](https://apiexplorer.developer.huaweicloud.com/apiexplorer/doc?product&#x3D;OCR&amp;api&#x3D;AutoClassification)服务。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *OcrClient) RecognizeBusinessLicense(request *model.RecognizeBusinessLicenseRequest) (*model.RecognizeBusinessLicenseResponse, error) {
	requestDef := GenReqDefForRecognizeBusinessLicense()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizeBusinessLicenseResponse), nil
	}
}

// 智利身份证识别
//
// 识别智利身份证图片中的文字内容，并返回识别的结构化结果。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *OcrClient) RecognizeChileIdCard(request *model.RecognizeChileIdCardRequest) (*model.RecognizeChileIdCardResponse, error) {
	requestDef := GenReqDefForRecognizeChileIdCard()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizeChileIdCardResponse), nil
	}
}

// 驾驶证识别
//
// 识别用户上传的驾驶证图片（或者用户提供的华为云上OBS的驾驶证图片文件的URL）中主页与副页的文字内容，并将识别的结果返回给用户。
//
// 说明：
//
// 如果图片中包含多张卡证票据，请调用智能分类识别服务。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *OcrClient) RecognizeDriverLicense(request *model.RecognizeDriverLicenseRequest) (*model.RecognizeDriverLicenseResponse, error) {
	requestDef := GenReqDefForRecognizeDriverLicense()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizeDriverLicenseResponse), nil
	}
}

// 财务报表识别
//
// 识别用户上传的表格图片中的文字内容，并将识别的结果返回给用户。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *OcrClient) RecognizeFinancialStatement(request *model.RecognizeFinancialStatementRequest) (*model.RecognizeFinancialStatementResponse, error) {
	requestDef := GenReqDefForRecognizeFinancialStatement()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizeFinancialStatementResponse), nil
	}
}

// 飞机行程单识别
//
// 识别飞机行程单中的文字信息，并返回识别的结构化结果。
//
// 说明：
//
// 如果图片中包含多张卡证票据，请调用智能分类识别服务。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *OcrClient) RecognizeFlightItinerary(request *model.RecognizeFlightItineraryRequest) (*model.RecognizeFlightItineraryResponse, error) {
	requestDef := GenReqDefForRecognizeFlightItinerary()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizeFlightItineraryResponse), nil
	}
}

// 通用表格识别
//
// 识别用于识别用户上传的通用表格图片（或者用户提供的华为云上OBS的通用表格图片文件的URL）中的文字内容，并将识别的结果返回给用户。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *OcrClient) RecognizeGeneralTable(request *model.RecognizeGeneralTableRequest) (*model.RecognizeGeneralTableResponse, error) {
	requestDef := GenReqDefForRecognizeGeneralTable()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizeGeneralTableResponse), nil
	}
}

// 通用文字识别
//
// 识别图片上的文字信息，返回识别的文字和坐标。支持扫描文件、电子文档、书籍、票据和表单等多种场景的文字识别。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *OcrClient) RecognizeGeneralText(request *model.RecognizeGeneralTextRequest) (*model.RecognizeGeneralTextResponse, error) {
	requestDef := GenReqDefForRecognizeGeneralText()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizeGeneralTextResponse), nil
	}
}

// 手写文字识别
//
// 识别文档中的手写文字信息，并将识别的结构化结果返回给用户。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *OcrClient) RecognizeHandwriting(request *model.RecognizeHandwritingRequest) (*model.RecognizeHandwritingResponse, error) {
	requestDef := GenReqDefForRecognizeHandwriting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizeHandwritingResponse), nil
	}
}

// 身份证识别
//
// 识别身份证图片中的文字内容，并将识别的结果返回给用户。
//
// 说明：
//
// 身份证识别支持中华人民共和国居民身份证识别。
//
// 如果图片中包含多张卡证票据，请调用[智能分类识别](https://apiexplorer.developer.huaweicloud.com/apiexplorer/doc?product&#x3D;OCR&amp;api&#x3D;AutoClassification)服务。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *OcrClient) RecognizeIdCard(request *model.RecognizeIdCardRequest) (*model.RecognizeIdCardResponse, error) {
	requestDef := GenReqDefForRecognizeIdCard()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizeIdCardResponse), nil
	}
}

// 保险单识别
//
// 识别保险单图片上的文字信息，并将识别的结构化结果返回给用户。支持对多板式保险单的扫描图片及手机照片进行结构化信息提取。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *OcrClient) RecognizeInsurancePolicy(request *model.RecognizeInsurancePolicyRequest) (*model.RecognizeInsurancePolicyResponse, error) {
	requestDef := GenReqDefForRecognizeInsurancePolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizeInsurancePolicyResponse), nil
	}
}

// 发票验真
//
// 发票验真服务支持9种增值税发票的信息核验，包括增值税专用发票、增值税普通发票、增值税普通发票（卷式）、增值税电子专用发票、增值税电子普通发票、增值税电子普通发票（通行费）、二手车销售统一发票、机动车销售统一发票、区块链电子发票，支持返回票面的全部信息。该接口的使用限制请参见[约束与限制](https://support.huaweicloud.com/productdesc-ocr/ocr_01_0006.html)，详细使用指导请参见[OCR服务使用简介](https://support.huaweicloud.com/qs-ocr/ocr_05_0001.html)章节。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *OcrClient) RecognizeInvoiceVerification(request *model.RecognizeInvoiceVerificationRequest) (*model.RecognizeInvoiceVerificationResponse, error) {
	requestDef := GenReqDefForRecognizeInvoiceVerification()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizeInvoiceVerificationResponse), nil
	}
}

// 车牌识别
//
// 识别输入图片中的车牌信息，并返回其坐标和内容。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *OcrClient) RecognizeLicensePlate(request *model.RecognizeLicensePlateRequest) (*model.RecognizeLicensePlateResponse, error) {
	requestDef := GenReqDefForRecognizeLicensePlate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizeLicensePlateResponse), nil
	}
}

// 机动车销售发票识别
//
// 识别机动车销售发票图片中的文字内容，并将识别的结果返回给用户。
//
// 说明：
//
// 该增值税发票仅限于中华人民共和国境内使用的增值税发票。
//
// 如果图片中包含多张卡证票据，请调用[智能分类识别](https://apiexplorer.developer.huaweicloud.com/apiexplorer/doc?product&#x3D;OCR&amp;api&#x3D;AutoClassification)服务。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *OcrClient) RecognizeMvsInvoice(request *model.RecognizeMvsInvoiceRequest) (*model.RecognizeMvsInvoiceResponse, error) {
	requestDef := GenReqDefForRecognizeMvsInvoice()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizeMvsInvoiceResponse), nil
	}
}

// 缅文驾驶证识别
//
// 识别缅甸驾驶证中的文字信息，并返回识别的结构化结果。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *OcrClient) RecognizeMyanmarDriverLicense(request *model.RecognizeMyanmarDriverLicenseRequest) (*model.RecognizeMyanmarDriverLicenseResponse, error) {
	requestDef := GenReqDefForRecognizeMyanmarDriverLicense()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizeMyanmarDriverLicenseResponse), nil
	}
}

// 缅文身份证识别
//
// 识别缅文身份证中的文字信息，并返回识别的结构化结果。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *OcrClient) RecognizeMyanmarIdcard(request *model.RecognizeMyanmarIdcardRequest) (*model.RecognizeMyanmarIdcardResponse, error) {
	requestDef := GenReqDefForRecognizeMyanmarIdcard()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizeMyanmarIdcardResponse), nil
	}
}

// 护照识别
//
// 识别用户上传的护照首页图片中的文字信息，并返回识别的结构化结果。当前版本支持中国护照的全字段识别。外国护照支持护照下方两行国际标准化的机读码识别，并可从中提取6-7个关键字段信息。
//
// 说明：
//
// 如果图片中包含多张卡证票据，请调用[智能分类识别](https://apiexplorer.developer.huaweicloud.com/apiexplorer/doc?product&#x3D;OCR&amp;api&#x3D;AutoClassification)服务。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *OcrClient) RecognizePassport(request *model.RecognizePassportRequest) (*model.RecognizePassportResponse, error) {
	requestDef := GenReqDefForRecognizePassport()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizePassportResponse), nil
	}
}

// 从业资格证识别
//
// 识别道路运输从业资格证上的关键文字信息，并返回识别的结构化结果。
//
// 说明：
//
// 如果图片中包含多张卡证票据，请调用[智能分类识别](https://apiexplorer.developer.huaweicloud.com/apiexplorer/doc?product&#x3D;OCR&amp;api&#x3D;AutoClassification)服务。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *OcrClient) RecognizeQualificationCertificate(request *model.RecognizeQualificationCertificateRequest) (*model.RecognizeQualificationCertificateResponse, error) {
	requestDef := GenReqDefForRecognizeQualificationCertificate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizeQualificationCertificateResponse), nil
	}
}

// 定额发票识别
//
// 识别定额发票中的文字信息，并返回识别的结构化结果。
//
// 说明：
//
// 如果图片中包含多张卡证票据，请调用[智能分类识别](https://apiexplorer.developer.huaweicloud.com/apiexplorer/doc?product&#x3D;OCR&amp;api&#x3D;AutoClassification)服务。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *OcrClient) RecognizeQuotaInvoice(request *model.RecognizeQuotaInvoiceRequest) (*model.RecognizeQuotaInvoiceResponse, error) {
	requestDef := GenReqDefForRecognizeQuotaInvoice()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizeQuotaInvoiceResponse), nil
	}
}

// 出租车发票识别
//
// 识别出租车发票中的文字信息，并返回识别的结构化结果。
//
// 说明：
//
// 如果图片中包含多张卡证票据，请调用智能分类识别服务。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *OcrClient) RecognizeTaxiInvoice(request *model.RecognizeTaxiInvoiceRequest) (*model.RecognizeTaxiInvoiceResponse, error) {
	requestDef := GenReqDefForRecognizeTaxiInvoice()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizeTaxiInvoiceResponse), nil
	}
}

// 泰文身份证识别
//
// 识别泰国身份证中的文字信息，并返回识别的结构化结果。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *OcrClient) RecognizeThailandIdcard(request *model.RecognizeThailandIdcardRequest) (*model.RecognizeThailandIdcardResponse, error) {
	requestDef := GenReqDefForRecognizeThailandIdcard()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizeThailandIdcardResponse), nil
	}
}

// 泰国车牌识别
//
// 识别泰国车牌图片中的车牌信息，并返回识别的结构化结果。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *OcrClient) RecognizeThailandLicensePlate(request *model.RecognizeThailandLicensePlateRequest) (*model.RecognizeThailandLicensePlateResponse, error) {
	requestDef := GenReqDefForRecognizeThailandLicensePlate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizeThailandLicensePlateResponse), nil
	}
}

// 车辆通行费发票识别
//
// 识别车辆通行费发票中的文字信息，并返回识别的结构化结果。
//
// 说明：
//
// 如果图片中包含多张卡证票据，请调用[智能分类识别](https://apiexplorer.developer.huaweicloud.com/apiexplorer/doc?product&#x3D;OCR&amp;api&#x3D;AutoClassification)服务。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *OcrClient) RecognizeTollInvoice(request *model.RecognizeTollInvoiceRequest) (*model.RecognizeTollInvoiceResponse, error) {
	requestDef := GenReqDefForRecognizeTollInvoice()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizeTollInvoiceResponse), nil
	}
}

// 火车票识别
//
// 识别火车票中的文字信息，并返回识别的结构化结果。
//
// 说明：
//
// 如果图片中包含多张卡证票据，请调用[智能分类识别](https://apiexplorer.developer.huaweicloud.com/apiexplorer/doc?product&#x3D;OCR&amp;api&#x3D;AutoClassification)服务。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *OcrClient) RecognizeTrainTicket(request *model.RecognizeTrainTicketRequest) (*model.RecognizeTrainTicketResponse, error) {
	requestDef := GenReqDefForRecognizeTrainTicket()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizeTrainTicketResponse), nil
	}
}

// 道路运输证识别
//
// 识别道路运输证首页中的文字信息，并返回识别的结构化结果。
//
// 说明： 如果图片中包含多张卡证票据，请调用智能分类识别服务。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *OcrClient) RecognizeTransportationLicense(request *model.RecognizeTransportationLicenseRequest) (*model.RecognizeTransportationLicenseResponse, error) {
	requestDef := GenReqDefForRecognizeTransportationLicense()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizeTransportationLicenseResponse), nil
	}
}

// 增值税发票识别
//
// 识别用户上传的增值税发票图片（或者用户提供的华为云上OBS的增值税发票图片文件的URL）中的文字内容，并将识别的结果返回给用户。
//
// 说明：
//
// 该增值税发票仅限于中华人民共和国境内使用的增值税发票。
//
// 如果图片中包含多张卡证票据，请调用智能分类识别服务。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *OcrClient) RecognizeVatInvoice(request *model.RecognizeVatInvoiceRequest) (*model.RecognizeVatInvoiceResponse, error) {
	requestDef := GenReqDefForRecognizeVatInvoice()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizeVatInvoiceResponse), nil
	}
}

// 行驶证识别
//
// 识别用户上传的行驶证图片（或者用户提供的华为云上OBS的行驶证图片文件的URL）中主页和副页的文字内容，并将识别的结果返回给用户。
//
// 说明：
//
// 如果图片中包含多张卡证票据，请调用智能分类识别服务。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *OcrClient) RecognizeVehicleLicense(request *model.RecognizeVehicleLicenseRequest) (*model.RecognizeVehicleLicenseResponse, error) {
	requestDef := GenReqDefForRecognizeVehicleLicense()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizeVehicleLicenseResponse), nil
	}
}

// 电子面单识别
//
// 识别用户上传的韵达电子面单图片中的文字内容，并将识别的结果以json格式返回给用户。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *OcrClient) RecognizeWaybillElectronic(request *model.RecognizeWaybillElectronicRequest) (*model.RecognizeWaybillElectronicResponse, error) {
	requestDef := GenReqDefForRecognizeWaybillElectronic()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizeWaybillElectronicResponse), nil
	}
}

// 网络图片识别
//
// 识别网络图片中的文字内容，并返回识别的结构化结果。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *OcrClient) RecognizeWebImage(request *model.RecognizeWebImageRequest) (*model.RecognizeWebImageResponse, error) {
	requestDef := GenReqDefForRecognizeWebImage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizeWebImageResponse), nil
	}
}

// VIN码识别
//
// 识别图片中的车架号信息，并将识别结果返回给用户。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *OcrClient) RecognizeVin(request *model.RecognizeVinRequest) (*model.RecognizeVinResponse, error) {
	requestDef := GenReqDefForRecognizeVin()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizeVinResponse), nil
	}
}
