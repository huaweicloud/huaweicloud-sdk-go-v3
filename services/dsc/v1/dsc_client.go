package v1

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dsc/v1/model"
)

type DscClient struct {
	HcClient *http_client.HcHttpClient
}

func NewDscClient(hcClient *http_client.HcHttpClient) *DscClient {
	return &DscClient{HcClient: hcClient}
}

func DscClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// BatchAddDataMask 对数据进行脱敏
//
// 对数据进行脱敏
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DscClient) BatchAddDataMask(request *model.BatchAddDataMaskRequest) (*model.BatchAddDataMaskResponse, error) {
	requestDef := GenReqDefForBatchAddDataMask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchAddDataMaskResponse), nil
	}
}

// BatchAddDataMaskInvoker 对数据进行脱敏
func (c *DscClient) BatchAddDataMaskInvoker(request *model.BatchAddDataMaskRequest) *BatchAddDataMaskInvoker {
	requestDef := GenReqDefForBatchAddDataMask()
	return &BatchAddDataMaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDatabaseWaterMark 嵌入数据水印
//
// 对json体中数据动态添加水印
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DscClient) CreateDatabaseWaterMark(request *model.CreateDatabaseWaterMarkRequest) (*model.CreateDatabaseWaterMarkResponse, error) {
	requestDef := GenReqDefForCreateDatabaseWaterMark()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDatabaseWaterMarkResponse), nil
	}
}

// CreateDatabaseWaterMarkInvoker 嵌入数据水印
func (c *DscClient) CreateDatabaseWaterMarkInvoker(request *model.CreateDatabaseWaterMarkRequest) *CreateDatabaseWaterMarkInvoker {
	requestDef := GenReqDefForCreateDatabaseWaterMark()
	return &CreateDatabaseWaterMarkInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDocWatermark 文档嵌入水印
//
// 对WORD(.docx)，PPT(.pptx)，EXCEL(.xlsx)，PDF(.pdf) 类型的文件嵌入文字暗水印、文字明水印或者图片明水印，用户以formData的格式传入待加水印的文件和水印相关信息，DSC服务给文件加完水印后返回给用户已嵌入水印的文件的二进制流。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DscClient) CreateDocWatermark(request *model.CreateDocWatermarkRequest) (*model.CreateDocWatermarkResponse, error) {
	requestDef := GenReqDefForCreateDocWatermark()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDocWatermarkResponse), nil
	}
}

// CreateDocWatermarkInvoker 文档嵌入水印
func (c *DscClient) CreateDocWatermarkInvoker(request *model.CreateDocWatermarkRequest) *CreateDocWatermarkInvoker {
	requestDef := GenReqDefForCreateDocWatermark()
	return &CreateDocWatermarkInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDocWatermarkByAddress 文档嵌入水印（文件地址版本）
//
// 对WORD(.docx)，PPT(.pptx)，EXCEL(.xlsx)，PDF(.pdf)*类型的文档嵌入文字暗水印、文字明水印或者图片明水印，用户传入待加水印的文档地址（目前支持OBS)和水印相关信息，DSC服务对文档加完水印后返回给用户已嵌入水印的文档的存放地址。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DscClient) CreateDocWatermarkByAddress(request *model.CreateDocWatermarkByAddressRequest) (*model.CreateDocWatermarkByAddressResponse, error) {
	requestDef := GenReqDefForCreateDocWatermarkByAddress()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDocWatermarkByAddressResponse), nil
	}
}

// CreateDocWatermarkByAddressInvoker 文档嵌入水印（文件地址版本）
func (c *DscClient) CreateDocWatermarkByAddressInvoker(request *model.CreateDocWatermarkByAddressRequest) *CreateDocWatermarkByAddressInvoker {
	requestDef := GenReqDefForCreateDocWatermarkByAddress()
	return &CreateDocWatermarkByAddressInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateImageWatermark 图片嵌入暗水印
//
// 对图片嵌入文字暗水印或者图片暗水印，用户以formData的格式传入待加水印图片和水印相关信息，DSC服务对图片加完水印后返回给用户已嵌入水印的图片二进制流，目前支持的图片格式为：*.jpg, *.jpeg, *.jpe, *.png, *.bmp, *.dib, *.rle, *.tiff, *.tif, *.ppm, *.webp, *.tga, *.tpic, *.gif。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DscClient) CreateImageWatermark(request *model.CreateImageWatermarkRequest) (*model.CreateImageWatermarkResponse, error) {
	requestDef := GenReqDefForCreateImageWatermark()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateImageWatermarkResponse), nil
	}
}

// CreateImageWatermarkInvoker 图片嵌入暗水印
func (c *DscClient) CreateImageWatermarkInvoker(request *model.CreateImageWatermarkRequest) *CreateImageWatermarkInvoker {
	requestDef := GenReqDefForCreateImageWatermark()
	return &CreateImageWatermarkInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateImageWatermarkByAddress 图片嵌入暗水印（文件地址版本）
//
// 对指定存储地址信息（目前支持华为云OBS）的图片嵌入文字暗水印或者图片暗水印，已嵌入的水印的图片将存放在用户指定的位置（目前支持华为云OBS），支持的图片格式为：*.jpg, *.jpeg, *.jpe, *.png, *.bmp, *.dib, *.rle, *.tiff, *.tif, *.ppm, *.webp, *.tga, *.tpic, *.gif。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DscClient) CreateImageWatermarkByAddress(request *model.CreateImageWatermarkByAddressRequest) (*model.CreateImageWatermarkByAddressResponse, error) {
	requestDef := GenReqDefForCreateImageWatermarkByAddress()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateImageWatermarkByAddressResponse), nil
	}
}

// CreateImageWatermarkByAddressInvoker 图片嵌入暗水印（文件地址版本）
func (c *DscClient) CreateImageWatermarkByAddressInvoker(request *model.CreateImageWatermarkByAddressRequest) *CreateImageWatermarkByAddressInvoker {
	requestDef := GenReqDefForCreateImageWatermarkByAddress()
	return &CreateImageWatermarkByAddressInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDatabaseWaterMark 提取数据水印
//
// 提取请求数据中水印内容
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DscClient) ShowDatabaseWaterMark(request *model.ShowDatabaseWaterMarkRequest) (*model.ShowDatabaseWaterMarkResponse, error) {
	requestDef := GenReqDefForShowDatabaseWaterMark()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDatabaseWaterMarkResponse), nil
	}
}

// ShowDatabaseWaterMarkInvoker 提取数据水印
func (c *DscClient) ShowDatabaseWaterMarkInvoker(request *model.ShowDatabaseWaterMarkRequest) *ShowDatabaseWaterMarkInvoker {
	requestDef := GenReqDefForShowDatabaseWaterMark()
	return &ShowDatabaseWaterMarkInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDocWatermark 文档提取暗水印
//
// 对已嵌入文字暗水印的WORD(.docx)，PPT(.pptx)，EXCEL(.xlsx)，PDF(.pdf)类型的文档进行文字暗水印提取，用户以formData的格式传入待提取水印的文件，DSC服务以JSON的格式返回从文档里提取的出的文字暗水印内容。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DscClient) ShowDocWatermark(request *model.ShowDocWatermarkRequest) (*model.ShowDocWatermarkResponse, error) {
	requestDef := GenReqDefForShowDocWatermark()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDocWatermarkResponse), nil
	}
}

// ShowDocWatermarkInvoker 文档提取暗水印
func (c *DscClient) ShowDocWatermarkInvoker(request *model.ShowDocWatermarkRequest) *ShowDocWatermarkInvoker {
	requestDef := GenReqDefForShowDocWatermark()
	return &ShowDocWatermarkInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDocWatermarkByAddress 文档提取暗水印（文档地址版本）
//
// 支持对已嵌入文字暗水印的WORD(.docx)，PPT(.pptx)，EXCEL(.xlsx)，PDF(.pdf)类型的文档进行水印提取，用户传入待提取水印的文档地址（目前支持OBS），DSC服务以JSON的格式返回从文档里提取的出的文字暗水印内容。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DscClient) ShowDocWatermarkByAddress(request *model.ShowDocWatermarkByAddressRequest) (*model.ShowDocWatermarkByAddressResponse, error) {
	requestDef := GenReqDefForShowDocWatermarkByAddress()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDocWatermarkByAddressResponse), nil
	}
}

// ShowDocWatermarkByAddressInvoker 文档提取暗水印（文档地址版本）
func (c *DscClient) ShowDocWatermarkByAddressInvoker(request *model.ShowDocWatermarkByAddressRequest) *ShowDocWatermarkByAddressInvoker {
	requestDef := GenReqDefForShowDocWatermarkByAddress()
	return &ShowDocWatermarkByAddressInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowImageWatermark 提取图片中的文字暗水印
//
// 对已嵌入文字暗水印的图片进行水印提取，用户以formData的格式传入待提取水印的图片，DSC服务以JSON的格式返回从图片里提取的出的文字暗水印。目前支持的图片格式为：*.jpg, *.jpeg, *.jpe, *.png, *.bmp, *.dib, *.rle, *.tiff, *.tif, *.ppm, *.webp, *.tga, *.tpic, *.gif。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DscClient) ShowImageWatermark(request *model.ShowImageWatermarkRequest) (*model.ShowImageWatermarkResponse, error) {
	requestDef := GenReqDefForShowImageWatermark()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowImageWatermarkResponse), nil
	}
}

// ShowImageWatermarkInvoker 提取图片中的文字暗水印
func (c *DscClient) ShowImageWatermarkInvoker(request *model.ShowImageWatermarkRequest) *ShowImageWatermarkInvoker {
	requestDef := GenReqDefForShowImageWatermark()
	return &ShowImageWatermarkInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowImageWatermarkByAddress 提取图片中的文字暗水印（文件地址版本）
//
// 对指定存储地址信息（目前支持华为云OBS）的已嵌入文字暗水印的图片提取文字暗水印，支持的图片格式为：*.jpg, *.jpeg, *.jpe, *.png, *.bmp, *.dib, *.rle, *.tiff, *.tif, *.ppm, *.webp, *.tga, *.tpic, *.gif。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DscClient) ShowImageWatermarkByAddress(request *model.ShowImageWatermarkByAddressRequest) (*model.ShowImageWatermarkByAddressResponse, error) {
	requestDef := GenReqDefForShowImageWatermarkByAddress()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowImageWatermarkByAddressResponse), nil
	}
}

// ShowImageWatermarkByAddressInvoker 提取图片中的文字暗水印（文件地址版本）
func (c *DscClient) ShowImageWatermarkByAddressInvoker(request *model.ShowImageWatermarkByAddressRequest) *ShowImageWatermarkByAddressInvoker {
	requestDef := GenReqDefForShowImageWatermarkByAddress()
	return &ShowImageWatermarkByAddressInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowImageWatermarkWithImage 提取图片中的图片暗水印
//
// 对已嵌入图片暗水印的图片进行水印提取，用户以formData的格式传入待提取水印的图片，DSC服务以图片二进制流的格式返回从图片里提取的出的图片暗水印。目前支持的图片格式为：*.jpg, *.jpeg, *.jpe, *.png, *.bmp, *.dib, *.rle, *.tiff, *.tif, *.ppm, *.webp, *.tga, *.tpic, *.gif。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DscClient) ShowImageWatermarkWithImage(request *model.ShowImageWatermarkWithImageRequest) (*model.ShowImageWatermarkWithImageResponse, error) {
	requestDef := GenReqDefForShowImageWatermarkWithImage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowImageWatermarkWithImageResponse), nil
	}
}

// ShowImageWatermarkWithImageInvoker 提取图片中的图片暗水印
func (c *DscClient) ShowImageWatermarkWithImageInvoker(request *model.ShowImageWatermarkWithImageRequest) *ShowImageWatermarkWithImageInvoker {
	requestDef := GenReqDefForShowImageWatermarkWithImage()
	return &ShowImageWatermarkWithImageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowImageWatermarkWithImageByAddress 提取图片中的图片暗水印（文件地址版本）
//
// 对指定存储地址信息（目前支持华为云OBS）的已嵌入图片暗水印的图片提取图片暗水印，提取出的水印图片将存放在用户指定的位置（目前支持华为云OBS），支持的图片格式为：*.jpg, *.jpeg, *.jpe, *.png, *.bmp, *.dib, *.rle, *.tiff, *.tif, *.ppm, *.webp, *.tga, *.tpic, *.gif。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DscClient) ShowImageWatermarkWithImageByAddress(request *model.ShowImageWatermarkWithImageByAddressRequest) (*model.ShowImageWatermarkWithImageByAddressResponse, error) {
	requestDef := GenReqDefForShowImageWatermarkWithImageByAddress()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowImageWatermarkWithImageByAddressResponse), nil
	}
}

// ShowImageWatermarkWithImageByAddressInvoker 提取图片中的图片暗水印（文件地址版本）
func (c *DscClient) ShowImageWatermarkWithImageByAddressInvoker(request *model.ShowImageWatermarkWithImageByAddressRequest) *ShowImageWatermarkWithImageByAddressInvoker {
	requestDef := GenReqDefForShowImageWatermarkWithImageByAddress()
	return &ShowImageWatermarkWithImageByAddressInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowScanJobResults 查询指定任务扫描结果
//
// 查询指定任务扫描结果
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DscClient) ShowScanJobResults(request *model.ShowScanJobResultsRequest) (*model.ShowScanJobResultsResponse, error) {
	requestDef := GenReqDefForShowScanJobResults()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowScanJobResultsResponse), nil
	}
}

// ShowScanJobResultsInvoker 查询指定任务扫描结果
func (c *DscClient) ShowScanJobResultsInvoker(request *model.ShowScanJobResultsRequest) *ShowScanJobResultsInvoker {
	requestDef := GenReqDefForShowScanJobResults()
	return &ShowScanJobResultsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowScanJobs 查询扫描任务列表
//
// 查询扫描任务列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DscClient) ShowScanJobs(request *model.ShowScanJobsRequest) (*model.ShowScanJobsResponse, error) {
	requestDef := GenReqDefForShowScanJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowScanJobsResponse), nil
	}
}

// ShowScanJobsInvoker 查询扫描任务列表
func (c *DscClient) ShowScanJobsInvoker(request *model.ShowScanJobsRequest) *ShowScanJobsInvoker {
	requestDef := GenReqDefForShowScanJobs()
	return &ShowScanJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowOpenApiCalledRecords 查询OpenApi调用记录
//
// 查询OpenApi调用记录
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DscClient) ShowOpenApiCalledRecords(request *model.ShowOpenApiCalledRecordsRequest) (*model.ShowOpenApiCalledRecordsResponse, error) {
	requestDef := GenReqDefForShowOpenApiCalledRecords()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowOpenApiCalledRecordsResponse), nil
	}
}

// ShowOpenApiCalledRecordsInvoker 查询OpenApi调用记录
func (c *DscClient) ShowOpenApiCalledRecordsInvoker(request *model.ShowOpenApiCalledRecordsRequest) *ShowOpenApiCalledRecordsInvoker {
	requestDef := GenReqDefForShowOpenApiCalledRecords()
	return &ShowOpenApiCalledRecordsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
