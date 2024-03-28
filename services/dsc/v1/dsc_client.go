package v1

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dsc/v1/model"
)

type DscClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewDscClient(hcClient *httpclient.HcHttpClient) *DscClient {
	return &DscClient{HcClient: hcClient}
}

func DscClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// AddBuckets 添加资产授权
//
// 添加数据资产扫描授权
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DscClient) AddBuckets(request *model.AddBucketsRequest) (*model.AddBucketsResponse, error) {
	requestDef := GenReqDefForAddBuckets()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddBucketsResponse), nil
	}
}

// AddBucketsInvoker 添加资产授权
func (c *DscClient) AddBucketsInvoker(request *model.AddBucketsRequest) *AddBucketsInvoker {
	requestDef := GenReqDefForAddBuckets()
	return &AddBucketsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddRule 创建扫描规则
//
// 根据指定的规则名称、规则类型、风险等级、最小匹配次数等参数创建自定义的敏感数据识别规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DscClient) AddRule(request *model.AddRuleRequest) (*model.AddRuleResponse, error) {
	requestDef := GenReqDefForAddRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddRuleResponse), nil
	}
}

// AddRuleInvoker 创建扫描规则
func (c *DscClient) AddRuleInvoker(request *model.AddRuleRequest) *AddRuleInvoker {
	requestDef := GenReqDefForAddRule()
	return &AddRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddRuleGroup 创建扫描规则组
//
// 根据指定的规则组名称和扫描规则列表创建敏感数据扫描规则组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DscClient) AddRuleGroup(request *model.AddRuleGroupRequest) (*model.AddRuleGroupResponse, error) {
	requestDef := GenReqDefForAddRuleGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddRuleGroupResponse), nil
	}
}

// AddRuleGroupInvoker 创建扫描规则组
func (c *DscClient) AddRuleGroupInvoker(request *model.AddRuleGroupRequest) *AddRuleGroupInvoker {
	requestDef := GenReqDefForAddRuleGroup()
	return &AddRuleGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddScanJob 创建扫描任务
//
// 根据指定的任务名称、扫描方式、扫描周期、扫描规则组、扫描时间创建扫描任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DscClient) AddScanJob(request *model.AddScanJobRequest) (*model.AddScanJobResponse, error) {
	requestDef := GenReqDefForAddScanJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddScanJobResponse), nil
	}
}

// AddScanJobInvoker 创建扫描任务
func (c *DscClient) AddScanJobInvoker(request *model.AddScanJobRequest) *AddScanJobInvoker {
	requestDef := GenReqDefForAddScanJob()
	return &AddScanJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchAddDataMask 对数据进行脱敏
//
// 对数据进行脱敏
//
// Please refer to HUAWEI cloud API Explorer for details.
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

// ChangeDbTemplateOperation 开启/停止脱敏任务
//
// 开启/停止脱敏任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DscClient) ChangeDbTemplateOperation(request *model.ChangeDbTemplateOperationRequest) (*model.ChangeDbTemplateOperationResponse, error) {
	requestDef := GenReqDefForChangeDbTemplateOperation()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeDbTemplateOperationResponse), nil
	}
}

// ChangeDbTemplateOperationInvoker 开启/停止脱敏任务
func (c *DscClient) ChangeDbTemplateOperationInvoker(request *model.ChangeDbTemplateOperationRequest) *ChangeDbTemplateOperationInvoker {
	requestDef := GenReqDefForChangeDbTemplateOperation()
	return &ChangeDbTemplateOperationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeRule 修改扫描规则
//
// 修改自定义的敏感数据识别规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DscClient) ChangeRule(request *model.ChangeRuleRequest) (*model.ChangeRuleResponse, error) {
	requestDef := GenReqDefForChangeRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeRuleResponse), nil
	}
}

// ChangeRuleInvoker 修改扫描规则
func (c *DscClient) ChangeRuleInvoker(request *model.ChangeRuleRequest) *ChangeRuleInvoker {
	requestDef := GenReqDefForChangeRule()
	return &ChangeRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDatabaseWaterMark 嵌入数据水印
//
// 对json体中数据动态添加水印
//
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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

// CreateProductOrder 实例下单
//
// 根据计费方式、计费周期等信息进行实例下单
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DscClient) CreateProductOrder(request *model.CreateProductOrderRequest) (*model.CreateProductOrderResponse, error) {
	requestDef := GenReqDefForCreateProductOrder()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateProductOrderResponse), nil
	}
}

// CreateProductOrderInvoker 实例下单
func (c *DscClient) CreateProductOrderInvoker(request *model.CreateProductOrderRequest) *CreateProductOrderInvoker {
	requestDef := GenReqDefForCreateProductOrder()
	return &CreateProductOrderInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteBucket 删除资产授权
//
// 删除数据资产扫描授权
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DscClient) DeleteBucket(request *model.DeleteBucketRequest) (*model.DeleteBucketResponse, error) {
	requestDef := GenReqDefForDeleteBucket()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteBucketResponse), nil
	}
}

// DeleteBucketInvoker 删除资产授权
func (c *DscClient) DeleteBucketInvoker(request *model.DeleteBucketRequest) *DeleteBucketInvoker {
	requestDef := GenReqDefForDeleteBucket()
	return &DeleteBucketInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteRule 删除扫描规则
//
// 删除指定的敏感数据识别规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DscClient) DeleteRule(request *model.DeleteRuleRequest) (*model.DeleteRuleResponse, error) {
	requestDef := GenReqDefForDeleteRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteRuleResponse), nil
	}
}

// DeleteRuleInvoker 删除扫描规则
func (c *DscClient) DeleteRuleInvoker(request *model.DeleteRuleRequest) *DeleteRuleInvoker {
	requestDef := GenReqDefForDeleteRule()
	return &DeleteRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteRuleGroup 删除扫描规则组
//
// 根据扫描规则组ID删除指定的扫描规则组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DscClient) DeleteRuleGroup(request *model.DeleteRuleGroupRequest) (*model.DeleteRuleGroupResponse, error) {
	requestDef := GenReqDefForDeleteRuleGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteRuleGroupResponse), nil
	}
}

// DeleteRuleGroupInvoker 删除扫描规则组
func (c *DscClient) DeleteRuleGroupInvoker(request *model.DeleteRuleGroupRequest) *DeleteRuleGroupInvoker {
	requestDef := GenReqDefForDeleteRuleGroup()
	return &DeleteRuleGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteScanJob 删除扫描任务
//
// 删除扫描任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DscClient) DeleteScanJob(request *model.DeleteScanJobRequest) (*model.DeleteScanJobResponse, error) {
	requestDef := GenReqDefForDeleteScanJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteScanJobResponse), nil
	}
}

// DeleteScanJobInvoker 删除扫描任务
func (c *DscClient) DeleteScanJobInvoker(request *model.DeleteScanJobRequest) *DeleteScanJobInvoker {
	requestDef := GenReqDefForDeleteScanJob()
	return &DeleteScanJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListBuckets 查看资产列表
//
// 查询数据资产扫描授权列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DscClient) ListBuckets(request *model.ListBucketsRequest) (*model.ListBucketsResponse, error) {
	requestDef := GenReqDefForListBuckets()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBucketsResponse), nil
	}
}

// ListBucketsInvoker 查看资产列表
func (c *DscClient) ListBucketsInvoker(request *model.ListBucketsRequest) *ListBucketsInvoker {
	requestDef := GenReqDefForListBuckets()
	return &ListBucketsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDbMaskTask 查询脱敏任务执行列表
//
// 查询脱敏任务执行列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DscClient) ListDbMaskTask(request *model.ListDbMaskTaskRequest) (*model.ListDbMaskTaskResponse, error) {
	requestDef := GenReqDefForListDbMaskTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDbMaskTaskResponse), nil
	}
}

// ListDbMaskTaskInvoker 查询脱敏任务执行列表
func (c *DscClient) ListDbMaskTaskInvoker(request *model.ListDbMaskTaskRequest) *ListDbMaskTaskInvoker {
	requestDef := GenReqDefForListDbMaskTask()
	return &ListDbMaskTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRuleGroups 查询扫描规则组列表
//
// 根据指定的项目ID查询扫描规则组列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DscClient) ListRuleGroups(request *model.ListRuleGroupsRequest) (*model.ListRuleGroupsResponse, error) {
	requestDef := GenReqDefForListRuleGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRuleGroupsResponse), nil
	}
}

// ListRuleGroupsInvoker 查询扫描规则组列表
func (c *DscClient) ListRuleGroupsInvoker(request *model.ListRuleGroupsRequest) *ListRuleGroupsInvoker {
	requestDef := GenReqDefForListRuleGroups()
	return &ListRuleGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDatabaseWaterMark 提取数据水印
//
// 提取请求数据中水印内容
//
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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

// ShowRules 查看规则列表
//
// 查询扫描规则列表，返回扫描规则总数和扫描规则列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DscClient) ShowRules(request *model.ShowRulesRequest) (*model.ShowRulesResponse, error) {
	requestDef := GenReqDefForShowRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRulesResponse), nil
	}
}

// ShowRulesInvoker 查看规则列表
func (c *DscClient) ShowRulesInvoker(request *model.ShowRulesRequest) *ShowRulesInvoker {
	requestDef := GenReqDefForShowRules()
	return &ShowRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowScanJobResults 查询指定任务扫描结果
//
// 查询指定任务扫描结果
//
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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

// ShowSpecification 查询资源开通信息
//
// 查询资源开通信息，根据项目ID查询订单详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DscClient) ShowSpecification(request *model.ShowSpecificationRequest) (*model.ShowSpecificationResponse, error) {
	requestDef := GenReqDefForShowSpecification()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSpecificationResponse), nil
	}
}

// ShowSpecificationInvoker 查询资源开通信息
func (c *DscClient) ShowSpecificationInvoker(request *model.ShowSpecificationRequest) *ShowSpecificationInvoker {
	requestDef := GenReqDefForShowSpecification()
	return &ShowSpecificationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTopics 查询告警通知主题
//
// 查询告警通知主题，返回默认主题、已确认主题数量及列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DscClient) ShowTopics(request *model.ShowTopicsRequest) (*model.ShowTopicsResponse, error) {
	requestDef := GenReqDefForShowTopics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTopicsResponse), nil
	}
}

// ShowTopicsInvoker 查询告警通知主题
func (c *DscClient) ShowTopicsInvoker(request *model.ShowTopicsRequest) *ShowTopicsInvoker {
	requestDef := GenReqDefForShowTopics()
	return &ShowTopicsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAssetName 编辑资产名称
//
// 编辑数据资产名称
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DscClient) UpdateAssetName(request *model.UpdateAssetNameRequest) (*model.UpdateAssetNameResponse, error) {
	requestDef := GenReqDefForUpdateAssetName()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAssetNameResponse), nil
	}
}

// UpdateAssetNameInvoker 编辑资产名称
func (c *DscClient) UpdateAssetNameInvoker(request *model.UpdateAssetNameRequest) *UpdateAssetNameInvoker {
	requestDef := GenReqDefForUpdateAssetName()
	return &UpdateAssetNameInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateDefaultTopic 修改告警通知主题
//
// 修改告警通知的关联项目ID、通知主题、通知状态(0为关闭通知，1为开启通知)等通用配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DscClient) UpdateDefaultTopic(request *model.UpdateDefaultTopicRequest) (*model.UpdateDefaultTopicResponse, error) {
	requestDef := GenReqDefForUpdateDefaultTopic()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDefaultTopicResponse), nil
	}
}

// UpdateDefaultTopicInvoker 修改告警通知主题
func (c *DscClient) UpdateDefaultTopicInvoker(request *model.UpdateDefaultTopicRequest) *UpdateDefaultTopicInvoker {
	requestDef := GenReqDefForUpdateDefaultTopic()
	return &UpdateDefaultTopicInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowOpenApiCalledRecords 查询OpenApi调用记录
//
// 查询OpenApi调用记录
//
// Please refer to HUAWEI cloud API Explorer for details.
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
