package v2

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/secmaster/v2/model"
)

type SecMasterClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewSecMasterClient(hcClient *httpclient.HcHttpClient) *SecMasterClient {
	return &SecMasterClient{HcClient: hcClient}
}

func SecMasterClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// CreateAdhocQuery 创建adhoc查询
//
// 创建adhoc查询
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) CreateAdhocQuery(request *model.CreateAdhocQueryRequest) (*model.CreateAdhocQueryResponse, error) {
	requestDef := GenReqDefForCreateAdhocQuery()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAdhocQueryResponse), nil
	}
}

// CreateAdhocQueryInvoker 创建adhoc查询
func (c *SecMasterClient) CreateAdhocQueryInvoker(request *model.CreateAdhocQueryRequest) *CreateAdhocQueryInvoker {
	requestDef := GenReqDefForCreateAdhocQuery()
	return &CreateAdhocQueryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAlertRule 创建告警规则
//
// 创建告警规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) CreateAlertRule(request *model.CreateAlertRuleRequest) (*model.CreateAlertRuleResponse, error) {
	requestDef := GenReqDefForCreateAlertRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAlertRuleResponse), nil
	}
}

// CreateAlertRuleInvoker 创建告警规则
func (c *SecMasterClient) CreateAlertRuleInvoker(request *model.CreateAlertRuleRequest) *CreateAlertRuleInvoker {
	requestDef := GenReqDefForCreateAlertRule()
	return &CreateAlertRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAnalysisScript 创建分析脚本
//
// 创建分析脚本
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) CreateAnalysisScript(request *model.CreateAnalysisScriptRequest) (*model.CreateAnalysisScriptResponse, error) {
	requestDef := GenReqDefForCreateAnalysisScript()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAnalysisScriptResponse), nil
	}
}

// CreateAnalysisScriptInvoker 创建分析脚本
func (c *SecMasterClient) CreateAnalysisScriptInvoker(request *model.CreateAnalysisScriptRequest) *CreateAnalysisScriptInvoker {
	requestDef := GenReqDefForCreateAnalysisScript()
	return &CreateAnalysisScriptInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateCodeSegment 创建代码片段
//
// 创建代码片段
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) CreateCodeSegment(request *model.CreateCodeSegmentRequest) (*model.CreateCodeSegmentResponse, error) {
	requestDef := GenReqDefForCreateCodeSegment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCodeSegmentResponse), nil
	}
}

// CreateCodeSegmentInvoker 创建代码片段
func (c *SecMasterClient) CreateCodeSegmentInvoker(request *model.CreateCodeSegmentRequest) *CreateCodeSegmentInvoker {
	requestDef := GenReqDefForCreateCodeSegment()
	return &CreateCodeSegmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateCollectConfig 保存云服务采集配置
//
// 保存云服务采集配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) CreateCollectConfig(request *model.CreateCollectConfigRequest) (*model.CreateCollectConfigResponse, error) {
	requestDef := GenReqDefForCreateCollectConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCollectConfigResponse), nil
	}
}

// CreateCollectConfigInvoker 保存云服务采集配置
func (c *SecMasterClient) CreateCollectConfigInvoker(request *model.CreateCollectConfigRequest) *CreateCollectConfigInvoker {
	requestDef := GenReqDefForCreateCollectConfig()
	return &CreateCollectConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateCustomizedCheckitem 新增自定义检查项
//
// 新增自定义检查项
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) CreateCustomizedCheckitem(request *model.CreateCustomizedCheckitemRequest) (*model.CreateCustomizedCheckitemResponse, error) {
	requestDef := GenReqDefForCreateCustomizedCheckitem()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCustomizedCheckitemResponse), nil
	}
}

// CreateCustomizedCheckitemInvoker 新增自定义检查项
func (c *SecMasterClient) CreateCustomizedCheckitemInvoker(request *model.CreateCustomizedCheckitemRequest) *CreateCustomizedCheckitemInvoker {
	requestDef := GenReqDefForCreateCustomizedCheckitem()
	return &CreateCustomizedCheckitemInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateCustomizedCompliancePackage 新增自定义遵从包
//
// 新增自定义遵从包
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) CreateCustomizedCompliancePackage(request *model.CreateCustomizedCompliancePackageRequest) (*model.CreateCustomizedCompliancePackageResponse, error) {
	requestDef := GenReqDefForCreateCustomizedCompliancePackage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCustomizedCompliancePackageResponse), nil
	}
}

// CreateCustomizedCompliancePackageInvoker 新增自定义遵从包
func (c *SecMasterClient) CreateCustomizedCompliancePackageInvoker(request *model.CreateCustomizedCompliancePackageRequest) *CreateCustomizedCompliancePackageInvoker {
	requestDef := GenReqDefForCreateCustomizedCompliancePackage()
	return &CreateCustomizedCompliancePackageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDataTransformation 创建数据加工
//
// 创建数据加工
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) CreateDataTransformation(request *model.CreateDataTransformationRequest) (*model.CreateDataTransformationResponse, error) {
	requestDef := GenReqDefForCreateDataTransformation()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDataTransformationResponse), nil
	}
}

// CreateDataTransformationInvoker 创建数据加工
func (c *SecMasterClient) CreateDataTransformationInvoker(request *model.CreateDataTransformationRequest) *CreateDataTransformationInvoker {
	requestDef := GenReqDefForCreateDataTransformation()
	return &CreateDataTransformationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateLayoutField 创建布局字段
//
// 创建布局字段
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) CreateLayoutField(request *model.CreateLayoutFieldRequest) (*model.CreateLayoutFieldResponse, error) {
	requestDef := GenReqDefForCreateLayoutField()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateLayoutFieldResponse), nil
	}
}

// CreateLayoutFieldInvoker 创建布局字段
func (c *SecMasterClient) CreateLayoutFieldInvoker(request *model.CreateLayoutFieldRequest) *CreateLayoutFieldInvoker {
	requestDef := GenReqDefForCreateLayoutField()
	return &CreateLayoutFieldInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePipe 创建管道
//
// 创建管道
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) CreatePipe(request *model.CreatePipeRequest) (*model.CreatePipeResponse, error) {
	requestDef := GenReqDefForCreatePipe()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePipeResponse), nil
	}
}

// CreatePipeInvoker 创建管道
func (c *SecMasterClient) CreatePipeInvoker(request *model.CreatePipeRequest) *CreatePipeInvoker {
	requestDef := GenReqDefForCreatePipe()
	return &CreatePipeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateRetrieveScript 创建检索脚本
//
// 创建检索脚本
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) CreateRetrieveScript(request *model.CreateRetrieveScriptRequest) (*model.CreateRetrieveScriptResponse, error) {
	requestDef := GenReqDefForCreateRetrieveScript()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRetrieveScriptResponse), nil
	}
}

// CreateRetrieveScriptInvoker 创建检索脚本
func (c *SecMasterClient) CreateRetrieveScriptInvoker(request *model.CreateRetrieveScriptRequest) *CreateRetrieveScriptInvoker {
	requestDef := GenReqDefForCreateRetrieveScript()
	return &CreateRetrieveScriptInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateSqlRender Adhoc sql参数渲染
//
// Adhoc sql参数渲染
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) CreateSqlRender(request *model.CreateSqlRenderRequest) (*model.CreateSqlRenderResponse, error) {
	requestDef := GenReqDefForCreateSqlRender()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSqlRenderResponse), nil
	}
}

// CreateSqlRenderInvoker Adhoc sql参数渲染
func (c *SecMasterClient) CreateSqlRenderInvoker(request *model.CreateSqlRenderRequest) *CreateSqlRenderInvoker {
	requestDef := GenReqDefForCreateSqlRender()
	return &CreateSqlRenderInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateTable 创建表
//
// 创建表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) CreateTable(request *model.CreateTableRequest) (*model.CreateTableResponse, error) {
	requestDef := GenReqDefForCreateTable()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTableResponse), nil
	}
}

// CreateTableInvoker 创建表
func (c *SecMasterClient) CreateTableInvoker(request *model.CreateTableRequest) *CreateTableInvoker {
	requestDef := GenReqDefForCreateTable()
	return &CreateTableInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateTableAnalysis 创建安全分析查询
//
// 创建安全分析查询
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) CreateTableAnalysis(request *model.CreateTableAnalysisRequest) (*model.CreateTableAnalysisResponse, error) {
	requestDef := GenReqDefForCreateTableAnalysis()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTableAnalysisResponse), nil
	}
}

// CreateTableAnalysisInvoker 创建安全分析查询
func (c *SecMasterClient) CreateTableAnalysisInvoker(request *model.CreateTableAnalysisRequest) *CreateTableAnalysisInvoker {
	requestDef := GenReqDefForCreateTableAnalysis()
	return &CreateTableAnalysisInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAdhocQuery 关闭查询操作
//
// 关闭查询操作
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) DeleteAdhocQuery(request *model.DeleteAdhocQueryRequest) (*model.DeleteAdhocQueryResponse, error) {
	requestDef := GenReqDefForDeleteAdhocQuery()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAdhocQueryResponse), nil
	}
}

// DeleteAdhocQueryInvoker 关闭查询操作
func (c *SecMasterClient) DeleteAdhocQueryInvoker(request *model.DeleteAdhocQueryRequest) *DeleteAdhocQueryInvoker {
	requestDef := GenReqDefForDeleteAdhocQuery()
	return &DeleteAdhocQueryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAlertRule 删除告警规则
//
// 删除告警规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) DeleteAlertRule(request *model.DeleteAlertRuleRequest) (*model.DeleteAlertRuleResponse, error) {
	requestDef := GenReqDefForDeleteAlertRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAlertRuleResponse), nil
	}
}

// DeleteAlertRuleInvoker 删除告警规则
func (c *SecMasterClient) DeleteAlertRuleInvoker(request *model.DeleteAlertRuleRequest) *DeleteAlertRuleInvoker {
	requestDef := GenReqDefForDeleteAlertRule()
	return &DeleteAlertRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAnalysisScript 删除分析脚本
//
// 删除分析脚本
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) DeleteAnalysisScript(request *model.DeleteAnalysisScriptRequest) (*model.DeleteAnalysisScriptResponse, error) {
	requestDef := GenReqDefForDeleteAnalysisScript()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAnalysisScriptResponse), nil
	}
}

// DeleteAnalysisScriptInvoker 删除分析脚本
func (c *SecMasterClient) DeleteAnalysisScriptInvoker(request *model.DeleteAnalysisScriptRequest) *DeleteAnalysisScriptInvoker {
	requestDef := GenReqDefForDeleteAnalysisScript()
	return &DeleteAnalysisScriptInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteCodeSegment 删除代码片段
//
// 删除代码片段
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) DeleteCodeSegment(request *model.DeleteCodeSegmentRequest) (*model.DeleteCodeSegmentResponse, error) {
	requestDef := GenReqDefForDeleteCodeSegment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteCodeSegmentResponse), nil
	}
}

// DeleteCodeSegmentInvoker 删除代码片段
func (c *SecMasterClient) DeleteCodeSegmentInvoker(request *model.DeleteCodeSegmentRequest) *DeleteCodeSegmentInvoker {
	requestDef := GenReqDefForDeleteCodeSegment()
	return &DeleteCodeSegmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteCustomizedCheckitems 删除自定义检查项
//
// 删除自定义检查项
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) DeleteCustomizedCheckitems(request *model.DeleteCustomizedCheckitemsRequest) (*model.DeleteCustomizedCheckitemsResponse, error) {
	requestDef := GenReqDefForDeleteCustomizedCheckitems()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteCustomizedCheckitemsResponse), nil
	}
}

// DeleteCustomizedCheckitemsInvoker 删除自定义检查项
func (c *SecMasterClient) DeleteCustomizedCheckitemsInvoker(request *model.DeleteCustomizedCheckitemsRequest) *DeleteCustomizedCheckitemsInvoker {
	requestDef := GenReqDefForDeleteCustomizedCheckitems()
	return &DeleteCustomizedCheckitemsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteCustomizedCompliancePackages 删除自定义遵从包
//
// 删除自定义遵从包
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) DeleteCustomizedCompliancePackages(request *model.DeleteCustomizedCompliancePackagesRequest) (*model.DeleteCustomizedCompliancePackagesResponse, error) {
	requestDef := GenReqDefForDeleteCustomizedCompliancePackages()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteCustomizedCompliancePackagesResponse), nil
	}
}

// DeleteCustomizedCompliancePackagesInvoker 删除自定义遵从包
func (c *SecMasterClient) DeleteCustomizedCompliancePackagesInvoker(request *model.DeleteCustomizedCompliancePackagesRequest) *DeleteCustomizedCompliancePackagesInvoker {
	requestDef := GenReqDefForDeleteCustomizedCompliancePackages()
	return &DeleteCustomizedCompliancePackagesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDataTransformation 删除数据加工
//
// 删除数据加工
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) DeleteDataTransformation(request *model.DeleteDataTransformationRequest) (*model.DeleteDataTransformationResponse, error) {
	requestDef := GenReqDefForDeleteDataTransformation()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDataTransformationResponse), nil
	}
}

// DeleteDataTransformationInvoker 删除数据加工
func (c *SecMasterClient) DeleteDataTransformationInvoker(request *model.DeleteDataTransformationRequest) *DeleteDataTransformationInvoker {
	requestDef := GenReqDefForDeleteDataTransformation()
	return &DeleteDataTransformationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteLayoutField 批量删除布局字段
//
// 删除布局字段
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) DeleteLayoutField(request *model.DeleteLayoutFieldRequest) (*model.DeleteLayoutFieldResponse, error) {
	requestDef := GenReqDefForDeleteLayoutField()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteLayoutFieldResponse), nil
	}
}

// DeleteLayoutFieldInvoker 批量删除布局字段
func (c *SecMasterClient) DeleteLayoutFieldInvoker(request *model.DeleteLayoutFieldRequest) *DeleteLayoutFieldInvoker {
	requestDef := GenReqDefForDeleteLayoutField()
	return &DeleteLayoutFieldInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeletePipe 删除管道
//
// 删除管道
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) DeletePipe(request *model.DeletePipeRequest) (*model.DeletePipeResponse, error) {
	requestDef := GenReqDefForDeletePipe()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePipeResponse), nil
	}
}

// DeletePipeInvoker 删除管道
func (c *SecMasterClient) DeletePipeInvoker(request *model.DeletePipeRequest) *DeletePipeInvoker {
	requestDef := GenReqDefForDeletePipe()
	return &DeletePipeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteRetrieveScript 删除检索脚本
//
// 删除检索脚本
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) DeleteRetrieveScript(request *model.DeleteRetrieveScriptRequest) (*model.DeleteRetrieveScriptResponse, error) {
	requestDef := GenReqDefForDeleteRetrieveScript()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteRetrieveScriptResponse), nil
	}
}

// DeleteRetrieveScriptInvoker 删除检索脚本
func (c *SecMasterClient) DeleteRetrieveScriptInvoker(request *model.DeleteRetrieveScriptRequest) *DeleteRetrieveScriptInvoker {
	requestDef := GenReqDefForDeleteRetrieveScript()
	return &DeleteRetrieveScriptInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteTable 删除表
//
// 删除表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) DeleteTable(request *model.DeleteTableRequest) (*model.DeleteTableResponse, error) {
	requestDef := GenReqDefForDeleteTable()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTableResponse), nil
	}
}

// DeleteTableInvoker 删除表
func (c *SecMasterClient) DeleteTableInvoker(request *model.DeleteTableRequest) *DeleteTableInvoker {
	requestDef := GenReqDefForDeleteTable()
	return &DeleteTableInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DisableAlertRule 停用告警规则
//
// 停用告警规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) DisableAlertRule(request *model.DisableAlertRuleRequest) (*model.DisableAlertRuleResponse, error) {
	requestDef := GenReqDefForDisableAlertRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisableAlertRuleResponse), nil
	}
}

// DisableAlertRuleInvoker 停用告警规则
func (c *SecMasterClient) DisableAlertRuleInvoker(request *model.DisableAlertRuleRequest) *DisableAlertRuleInvoker {
	requestDef := GenReqDefForDisableAlertRule()
	return &DisableAlertRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DisableDataConsumption 关闭实时消费
//
// 关闭实时消费
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) DisableDataConsumption(request *model.DisableDataConsumptionRequest) (*model.DisableDataConsumptionResponse, error) {
	requestDef := GenReqDefForDisableDataConsumption()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisableDataConsumptionResponse), nil
	}
}

// DisableDataConsumptionInvoker 关闭实时消费
func (c *SecMasterClient) DisableDataConsumptionInvoker(request *model.DisableDataConsumptionRequest) *DisableDataConsumptionInvoker {
	requestDef := GenReqDefForDisableDataConsumption()
	return &DisableDataConsumptionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DisableDataTransformation 停用数据加工
//
// 停用数据加工
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) DisableDataTransformation(request *model.DisableDataTransformationRequest) (*model.DisableDataTransformationResponse, error) {
	requestDef := GenReqDefForDisableDataTransformation()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisableDataTransformationResponse), nil
	}
}

// DisableDataTransformationInvoker 停用数据加工
func (c *SecMasterClient) DisableDataTransformationInvoker(request *model.DisableDataTransformationRequest) *DisableDataTransformationInvoker {
	requestDef := GenReqDefForDisableDataTransformation()
	return &DisableDataTransformationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// EnableAlertRule 启用告警规则
//
// 启用告警规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) EnableAlertRule(request *model.EnableAlertRuleRequest) (*model.EnableAlertRuleResponse, error) {
	requestDef := GenReqDefForEnableAlertRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.EnableAlertRuleResponse), nil
	}
}

// EnableAlertRuleInvoker 启用告警规则
func (c *SecMasterClient) EnableAlertRuleInvoker(request *model.EnableAlertRuleRequest) *EnableAlertRuleInvoker {
	requestDef := GenReqDefForEnableAlertRule()
	return &EnableAlertRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// EnableDataConsumption 开启实时消费
//
// 开启实时消费
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) EnableDataConsumption(request *model.EnableDataConsumptionRequest) (*model.EnableDataConsumptionResponse, error) {
	requestDef := GenReqDefForEnableDataConsumption()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.EnableDataConsumptionResponse), nil
	}
}

// EnableDataConsumptionInvoker 开启实时消费
func (c *SecMasterClient) EnableDataConsumptionInvoker(request *model.EnableDataConsumptionRequest) *EnableDataConsumptionInvoker {
	requestDef := GenReqDefForEnableDataConsumption()
	return &EnableDataConsumptionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// EnableDataTransformation 启用数据加工
//
// 启用数据加工
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) EnableDataTransformation(request *model.EnableDataTransformationRequest) (*model.EnableDataTransformationResponse, error) {
	requestDef := GenReqDefForEnableDataTransformation()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.EnableDataTransformationResponse), nil
	}
}

// EnableDataTransformationInvoker 启用数据加工
func (c *SecMasterClient) EnableDataTransformationInvoker(request *model.EnableDataTransformationRequest) *EnableDataTransformationInvoker {
	requestDef := GenReqDefForEnableDataTransformation()
	return &EnableDataTransformationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAlertRuleMetrics 告警规则总览
//
// 告警规则总览
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListAlertRuleMetrics(request *model.ListAlertRuleMetricsRequest) (*model.ListAlertRuleMetricsResponse, error) {
	requestDef := GenReqDefForListAlertRuleMetrics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAlertRuleMetricsResponse), nil
	}
}

// ListAlertRuleMetricsInvoker 告警规则总览
func (c *SecMasterClient) ListAlertRuleMetricsInvoker(request *model.ListAlertRuleMetricsRequest) *ListAlertRuleMetricsInvoker {
	requestDef := GenReqDefForListAlertRuleMetrics()
	return &ListAlertRuleMetricsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAlertRuleTemplateMetrics 列出告警规则模板总览
//
// 列出告警规则模板总览
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListAlertRuleTemplateMetrics(request *model.ListAlertRuleTemplateMetricsRequest) (*model.ListAlertRuleTemplateMetricsResponse, error) {
	requestDef := GenReqDefForListAlertRuleTemplateMetrics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAlertRuleTemplateMetricsResponse), nil
	}
}

// ListAlertRuleTemplateMetricsInvoker 列出告警规则模板总览
func (c *SecMasterClient) ListAlertRuleTemplateMetricsInvoker(request *model.ListAlertRuleTemplateMetricsRequest) *ListAlertRuleTemplateMetricsInvoker {
	requestDef := GenReqDefForListAlertRuleTemplateMetrics()
	return &ListAlertRuleTemplateMetricsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAlertRuleTemplates 列出告警规则模板
//
// 列出告警规则模板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListAlertRuleTemplates(request *model.ListAlertRuleTemplatesRequest) (*model.ListAlertRuleTemplatesResponse, error) {
	requestDef := GenReqDefForListAlertRuleTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAlertRuleTemplatesResponse), nil
	}
}

// ListAlertRuleTemplatesInvoker 列出告警规则模板
func (c *SecMasterClient) ListAlertRuleTemplatesInvoker(request *model.ListAlertRuleTemplatesRequest) *ListAlertRuleTemplatesInvoker {
	requestDef := GenReqDefForListAlertRuleTemplates()
	return &ListAlertRuleTemplatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAlertRules 列出告警规则
//
// 列出告警规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListAlertRules(request *model.ListAlertRulesRequest) (*model.ListAlertRulesResponse, error) {
	requestDef := GenReqDefForListAlertRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAlertRulesResponse), nil
	}
}

// ListAlertRulesInvoker 列出告警规则
func (c *SecMasterClient) ListAlertRulesInvoker(request *model.ListAlertRulesRequest) *ListAlertRulesInvoker {
	requestDef := GenReqDefForListAlertRules()
	return &ListAlertRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAnalysisScripts 列出分析脚本
//
// 列出分析脚本
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListAnalysisScripts(request *model.ListAnalysisScriptsRequest) (*model.ListAnalysisScriptsResponse, error) {
	requestDef := GenReqDefForListAnalysisScripts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAnalysisScriptsResponse), nil
	}
}

// ListAnalysisScriptsInvoker 列出分析脚本
func (c *SecMasterClient) ListAnalysisScriptsInvoker(request *model.ListAnalysisScriptsRequest) *ListAnalysisScriptsInvoker {
	requestDef := GenReqDefForListAnalysisScripts()
	return &ListAnalysisScriptsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCodeSegments 列出代码片段
//
// 列出代码片段
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListCodeSegments(request *model.ListCodeSegmentsRequest) (*model.ListCodeSegmentsResponse, error) {
	requestDef := GenReqDefForListCodeSegments()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCodeSegmentsResponse), nil
	}
}

// ListCodeSegmentsInvoker 列出代码片段
func (c *SecMasterClient) ListCodeSegmentsInvoker(request *model.ListCodeSegmentsRequest) *ListCodeSegmentsInvoker {
	requestDef := GenReqDefForListCodeSegments()
	return &ListCodeSegmentsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCollectConfig 获取云服务采集配置
//
// 获取云服务采集配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListCollectConfig(request *model.ListCollectConfigRequest) (*model.ListCollectConfigResponse, error) {
	requestDef := GenReqDefForListCollectConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCollectConfigResponse), nil
	}
}

// ListCollectConfigInvoker 获取云服务采集配置
func (c *SecMasterClient) ListCollectConfigInvoker(request *model.ListCollectConfigRequest) *ListCollectConfigInvoker {
	requestDef := GenReqDefForListCollectConfig()
	return &ListCollectConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDataTransformationMetrics 数据加工总览
//
// 数据加工总览
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListDataTransformationMetrics(request *model.ListDataTransformationMetricsRequest) (*model.ListDataTransformationMetricsResponse, error) {
	requestDef := GenReqDefForListDataTransformationMetrics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDataTransformationMetricsResponse), nil
	}
}

// ListDataTransformationMetricsInvoker 数据加工总览
func (c *SecMasterClient) ListDataTransformationMetricsInvoker(request *model.ListDataTransformationMetricsRequest) *ListDataTransformationMetricsInvoker {
	requestDef := GenReqDefForListDataTransformationMetrics()
	return &ListDataTransformationMetricsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDataTransformations 列出数据加工
//
// 列出数据加工
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListDataTransformations(request *model.ListDataTransformationsRequest) (*model.ListDataTransformationsResponse, error) {
	requestDef := GenReqDefForListDataTransformations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDataTransformationsResponse), nil
	}
}

// ListDataTransformationsInvoker 列出数据加工
func (c *SecMasterClient) ListDataTransformationsInvoker(request *model.ListDataTransformationsRequest) *ListDataTransformationsInvoker {
	requestDef := GenReqDefForListDataTransformations()
	return &ListDataTransformationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDirectories 列出目录分组
//
// 列出目录分组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListDirectories(request *model.ListDirectoriesRequest) (*model.ListDirectoriesResponse, error) {
	requestDef := GenReqDefForListDirectories()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDirectoriesResponse), nil
	}
}

// ListDirectoriesInvoker 列出目录分组
func (c *SecMasterClient) ListDirectoriesInvoker(request *model.ListDirectoriesRequest) *ListDirectoriesInvoker {
	requestDef := GenReqDefForListDirectories()
	return &ListDirectoriesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListLayoutFieldAll 全部布局字段
//
// 查询布局字段列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListLayoutFieldAll(request *model.ListLayoutFieldAllRequest) (*model.ListLayoutFieldAllResponse, error) {
	requestDef := GenReqDefForListLayoutFieldAll()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLayoutFieldAllResponse), nil
	}
}

// ListLayoutFieldAllInvoker 全部布局字段
func (c *SecMasterClient) ListLayoutFieldAllInvoker(request *model.ListLayoutFieldAllRequest) *ListLayoutFieldAllInvoker {
	requestDef := GenReqDefForListLayoutFieldAll()
	return &ListLayoutFieldAllInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPipes 获取管道列表
//
// 获取管道列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListPipes(request *model.ListPipesRequest) (*model.ListPipesResponse, error) {
	requestDef := GenReqDefForListPipes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPipesResponse), nil
	}
}

// ListPipesInvoker 获取管道列表
func (c *SecMasterClient) ListPipesInvoker(request *model.ListPipesRequest) *ListPipesInvoker {
	requestDef := GenReqDefForListPipes()
	return &ListPipesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRetrieveScripts 列出检索脚本
//
// 列出检索脚本
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListRetrieveScripts(request *model.ListRetrieveScriptsRequest) (*model.ListRetrieveScriptsResponse, error) {
	requestDef := GenReqDefForListRetrieveScripts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRetrieveScriptsResponse), nil
	}
}

// ListRetrieveScriptsInvoker 列出检索脚本
func (c *SecMasterClient) ListRetrieveScriptsInvoker(request *model.ListRetrieveScriptsRequest) *ListRetrieveScriptsInvoker {
	requestDef := GenReqDefForListRetrieveScripts()
	return &ListRetrieveScriptsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTableHistograms 检索表直方图
//
// 检索表直方图
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListTableHistograms(request *model.ListTableHistogramsRequest) (*model.ListTableHistogramsResponse, error) {
	requestDef := GenReqDefForListTableHistograms()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTableHistogramsResponse), nil
	}
}

// ListTableHistogramsInvoker 检索表直方图
func (c *SecMasterClient) ListTableHistogramsInvoker(request *model.ListTableHistogramsRequest) *ListTableHistogramsInvoker {
	requestDef := GenReqDefForListTableHistograms()
	return &ListTableHistogramsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTableLogs 检索表日志
//
// 检索表日志
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListTableLogs(request *model.ListTableLogsRequest) (*model.ListTableLogsResponse, error) {
	requestDef := GenReqDefForListTableLogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTableLogsResponse), nil
	}
}

// ListTableLogsInvoker 检索表日志
func (c *SecMasterClient) ListTableLogsInvoker(request *model.ListTableLogsRequest) *ListTableLogsInvoker {
	requestDef := GenReqDefForListTableLogs()
	return &ListTableLogsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTables 获取表列表
//
// 获取表列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListTables(request *model.ListTablesRequest) (*model.ListTablesResponse, error) {
	requestDef := GenReqDefForListTables()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTablesResponse), nil
	}
}

// ListTablesInvoker 获取表列表
func (c *SecMasterClient) ListTablesInvoker(request *model.ListTablesRequest) *ListTablesInvoker {
	requestDef := GenReqDefForListTables()
	return &ListTablesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchBaseline 搜索基线检查结果列表
//
// 搜索基线检查结果列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) SearchBaseline(request *model.SearchBaselineRequest) (*model.SearchBaselineResponse, error) {
	requestDef := GenReqDefForSearchBaseline()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchBaselineResponse), nil
	}
}

// SearchBaselineInvoker 搜索基线检查结果列表
func (c *SecMasterClient) SearchBaselineInvoker(request *model.SearchBaselineRequest) *SearchBaselineInvoker {
	requestDef := GenReqDefForSearchBaseline()
	return &SearchBaselineInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchCheckitems 查询检查项列表
//
// 查询检查项列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) SearchCheckitems(request *model.SearchCheckitemsRequest) (*model.SearchCheckitemsResponse, error) {
	requestDef := GenReqDefForSearchCheckitems()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchCheckitemsResponse), nil
	}
}

// SearchCheckitemsInvoker 查询检查项列表
func (c *SecMasterClient) SearchCheckitemsInvoker(request *model.SearchCheckitemsRequest) *SearchCheckitemsInvoker {
	requestDef := GenReqDefForSearchCheckitems()
	return &SearchCheckitemsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchCompliancePackages 查询遵从包列表
//
// 查询遵从包列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) SearchCompliancePackages(request *model.SearchCompliancePackagesRequest) (*model.SearchCompliancePackagesResponse, error) {
	requestDef := GenReqDefForSearchCompliancePackages()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchCompliancePackagesResponse), nil
	}
}

// SearchCompliancePackagesInvoker 查询遵从包列表
func (c *SecMasterClient) SearchCompliancePackagesInvoker(request *model.SearchCompliancePackagesRequest) *SearchCompliancePackagesInvoker {
	requestDef := GenReqDefForSearchCompliancePackages()
	return &SearchCompliancePackagesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAdhocResult 获取adhoc查询结果
//
// 获取adhoc查询结果
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ShowAdhocResult(request *model.ShowAdhocResultRequest) (*model.ShowAdhocResultResponse, error) {
	requestDef := GenReqDefForShowAdhocResult()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAdhocResultResponse), nil
	}
}

// ShowAdhocResultInvoker 获取adhoc查询结果
func (c *SecMasterClient) ShowAdhocResultInvoker(request *model.ShowAdhocResultRequest) *ShowAdhocResultInvoker {
	requestDef := GenReqDefForShowAdhocResult()
	return &ShowAdhocResultInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAlertRule 查看告警规则
//
// 查看告警规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ShowAlertRule(request *model.ShowAlertRuleRequest) (*model.ShowAlertRuleResponse, error) {
	requestDef := GenReqDefForShowAlertRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAlertRuleResponse), nil
	}
}

// ShowAlertRuleInvoker 查看告警规则
func (c *SecMasterClient) ShowAlertRuleInvoker(request *model.ShowAlertRuleRequest) *ShowAlertRuleInvoker {
	requestDef := GenReqDefForShowAlertRule()
	return &ShowAlertRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAlertRuleTemplate 查看告警规则模板
//
// 查看告警规则模板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ShowAlertRuleTemplate(request *model.ShowAlertRuleTemplateRequest) (*model.ShowAlertRuleTemplateResponse, error) {
	requestDef := GenReqDefForShowAlertRuleTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAlertRuleTemplateResponse), nil
	}
}

// ShowAlertRuleTemplateInvoker 查看告警规则模板
func (c *SecMasterClient) ShowAlertRuleTemplateInvoker(request *model.ShowAlertRuleTemplateRequest) *ShowAlertRuleTemplateInvoker {
	requestDef := GenReqDefForShowAlertRuleTemplate()
	return &ShowAlertRuleTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAnalysisScript 查看分析脚本
//
// 查看分析脚本
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ShowAnalysisScript(request *model.ShowAnalysisScriptRequest) (*model.ShowAnalysisScriptResponse, error) {
	requestDef := GenReqDefForShowAnalysisScript()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAnalysisScriptResponse), nil
	}
}

// ShowAnalysisScriptInvoker 查看分析脚本
func (c *SecMasterClient) ShowAnalysisScriptInvoker(request *model.ShowAnalysisScriptRequest) *ShowAnalysisScriptInvoker {
	requestDef := GenReqDefForShowAnalysisScript()
	return &ShowAnalysisScriptInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCheckitemDetail 查询检查项详情
//
// 查询检查项详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ShowCheckitemDetail(request *model.ShowCheckitemDetailRequest) (*model.ShowCheckitemDetailResponse, error) {
	requestDef := GenReqDefForShowCheckitemDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCheckitemDetailResponse), nil
	}
}

// ShowCheckitemDetailInvoker 查询检查项详情
func (c *SecMasterClient) ShowCheckitemDetailInvoker(request *model.ShowCheckitemDetailRequest) *ShowCheckitemDetailInvoker {
	requestDef := GenReqDefForShowCheckitemDetail()
	return &ShowCheckitemDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCodeSegment 查看代码片段
//
// 查看代码片段
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ShowCodeSegment(request *model.ShowCodeSegmentRequest) (*model.ShowCodeSegmentResponse, error) {
	requestDef := GenReqDefForShowCodeSegment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCodeSegmentResponse), nil
	}
}

// ShowCodeSegmentInvoker 查看代码片段
func (c *SecMasterClient) ShowCodeSegmentInvoker(request *model.ShowCodeSegmentRequest) *ShowCodeSegmentInvoker {
	requestDef := GenReqDefForShowCodeSegment()
	return &ShowCodeSegmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCompliancePackageDetail 查询遵从包详情
//
// 查询遵从包详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ShowCompliancePackageDetail(request *model.ShowCompliancePackageDetailRequest) (*model.ShowCompliancePackageDetailResponse, error) {
	requestDef := GenReqDefForShowCompliancePackageDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCompliancePackageDetailResponse), nil
	}
}

// ShowCompliancePackageDetailInvoker 查询遵从包详情
func (c *SecMasterClient) ShowCompliancePackageDetailInvoker(request *model.ShowCompliancePackageDetailRequest) *ShowCompliancePackageDetailInvoker {
	requestDef := GenReqDefForShowCompliancePackageDetail()
	return &ShowCompliancePackageDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDataConsumption 获取实时消费配置
//
// 获取实时消费配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ShowDataConsumption(request *model.ShowDataConsumptionRequest) (*model.ShowDataConsumptionResponse, error) {
	requestDef := GenReqDefForShowDataConsumption()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDataConsumptionResponse), nil
	}
}

// ShowDataConsumptionInvoker 获取实时消费配置
func (c *SecMasterClient) ShowDataConsumptionInvoker(request *model.ShowDataConsumptionRequest) *ShowDataConsumptionInvoker {
	requestDef := GenReqDefForShowDataConsumption()
	return &ShowDataConsumptionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDataTransformation 查看数据加工
//
// 查看数据加工
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ShowDataTransformation(request *model.ShowDataTransformationRequest) (*model.ShowDataTransformationResponse, error) {
	requestDef := GenReqDefForShowDataTransformation()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDataTransformationResponse), nil
	}
}

// ShowDataTransformationInvoker 查看数据加工
func (c *SecMasterClient) ShowDataTransformationInvoker(request *model.ShowDataTransformationRequest) *ShowDataTransformationInvoker {
	requestDef := GenReqDefForShowDataTransformation()
	return &ShowDataTransformationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowLayoutFieldInfo 展示字段详情
//
// 查询布局字段详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ShowLayoutFieldInfo(request *model.ShowLayoutFieldInfoRequest) (*model.ShowLayoutFieldInfoResponse, error) {
	requestDef := GenReqDefForShowLayoutFieldInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLayoutFieldInfoResponse), nil
	}
}

// ShowLayoutFieldInfoInvoker 展示字段详情
func (c *SecMasterClient) ShowLayoutFieldInfoInvoker(request *model.ShowLayoutFieldInfoRequest) *ShowLayoutFieldInfoInvoker {
	requestDef := GenReqDefForShowLayoutFieldInfo()
	return &ShowLayoutFieldInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMonitorStats 获取监控统计信息
//
// 获取监控统计信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ShowMonitorStats(request *model.ShowMonitorStatsRequest) (*model.ShowMonitorStatsResponse, error) {
	requestDef := GenReqDefForShowMonitorStats()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMonitorStatsResponse), nil
	}
}

// ShowMonitorStatsInvoker 获取监控统计信息
func (c *SecMasterClient) ShowMonitorStatsInvoker(request *model.ShowMonitorStatsRequest) *ShowMonitorStatsInvoker {
	requestDef := GenReqDefForShowMonitorStats()
	return &ShowMonitorStatsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPipe 获取管道详情
//
// 获取管道详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ShowPipe(request *model.ShowPipeRequest) (*model.ShowPipeResponse, error) {
	requestDef := GenReqDefForShowPipe()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPipeResponse), nil
	}
}

// ShowPipeInvoker 获取管道详情
func (c *SecMasterClient) ShowPipeInvoker(request *model.ShowPipeRequest) *ShowPipeInvoker {
	requestDef := GenReqDefForShowPipe()
	return &ShowPipeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRetrieveScript 查看检索脚本
//
// 查看检索脚本
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ShowRetrieveScript(request *model.ShowRetrieveScriptRequest) (*model.ShowRetrieveScriptResponse, error) {
	requestDef := GenReqDefForShowRetrieveScript()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRetrieveScriptResponse), nil
	}
}

// ShowRetrieveScriptInvoker 查看检索脚本
func (c *SecMasterClient) ShowRetrieveScriptInvoker(request *model.ShowRetrieveScriptRequest) *ShowRetrieveScriptInvoker {
	requestDef := GenReqDefForShowRetrieveScript()
	return &ShowRetrieveScriptInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSubscriptionResources 获取订阅资源信息
//
// 获取订阅资源信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ShowSubscriptionResources(request *model.ShowSubscriptionResourcesRequest) (*model.ShowSubscriptionResourcesResponse, error) {
	requestDef := GenReqDefForShowSubscriptionResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSubscriptionResourcesResponse), nil
	}
}

// ShowSubscriptionResourcesInvoker 获取订阅资源信息
func (c *SecMasterClient) ShowSubscriptionResourcesInvoker(request *model.ShowSubscriptionResourcesRequest) *ShowSubscriptionResourcesInvoker {
	requestDef := GenReqDefForShowSubscriptionResources()
	return &ShowSubscriptionResourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTable 获取表详情
//
// 获取表详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ShowTable(request *model.ShowTableRequest) (*model.ShowTableResponse, error) {
	requestDef := GenReqDefForShowTable()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTableResponse), nil
	}
}

// ShowTableInvoker 获取表详情
func (c *SecMasterClient) ShowTableInvoker(request *model.ShowTableRequest) *ShowTableInvoker {
	requestDef := GenReqDefForShowTable()
	return &ShowTableInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowVersion 获取当前可用版本
//
// 获取当前可用版本
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ShowVersion(request *model.ShowVersionRequest) (*model.ShowVersionResponse, error) {
	requestDef := GenReqDefForShowVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVersionResponse), nil
	}
}

// ShowVersionInvoker 获取当前可用版本
func (c *SecMasterClient) ShowVersionInvoker(request *model.ShowVersionRequest) *ShowVersionInvoker {
	requestDef := GenReqDefForShowVersion()
	return &ShowVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAlertRule 更新告警规则
//
// 更新告警规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) UpdateAlertRule(request *model.UpdateAlertRuleRequest) (*model.UpdateAlertRuleResponse, error) {
	requestDef := GenReqDefForUpdateAlertRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAlertRuleResponse), nil
	}
}

// UpdateAlertRuleInvoker 更新告警规则
func (c *SecMasterClient) UpdateAlertRuleInvoker(request *model.UpdateAlertRuleRequest) *UpdateAlertRuleInvoker {
	requestDef := GenReqDefForUpdateAlertRule()
	return &UpdateAlertRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAnalysisScript 更新分析脚本
//
// 更新分析脚本
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) UpdateAnalysisScript(request *model.UpdateAnalysisScriptRequest) (*model.UpdateAnalysisScriptResponse, error) {
	requestDef := GenReqDefForUpdateAnalysisScript()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAnalysisScriptResponse), nil
	}
}

// UpdateAnalysisScriptInvoker 更新分析脚本
func (c *SecMasterClient) UpdateAnalysisScriptInvoker(request *model.UpdateAnalysisScriptRequest) *UpdateAnalysisScriptInvoker {
	requestDef := GenReqDefForUpdateAnalysisScript()
	return &UpdateAnalysisScriptInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateCheckitem 更新检查项
//
// 更新检查项
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) UpdateCheckitem(request *model.UpdateCheckitemRequest) (*model.UpdateCheckitemResponse, error) {
	requestDef := GenReqDefForUpdateCheckitem()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateCheckitemResponse), nil
	}
}

// UpdateCheckitemInvoker 更新检查项
func (c *SecMasterClient) UpdateCheckitemInvoker(request *model.UpdateCheckitemRequest) *UpdateCheckitemInvoker {
	requestDef := GenReqDefForUpdateCheckitem()
	return &UpdateCheckitemInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateCodeSegment 更新代码片段
//
// 更新代码片段
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) UpdateCodeSegment(request *model.UpdateCodeSegmentRequest) (*model.UpdateCodeSegmentResponse, error) {
	requestDef := GenReqDefForUpdateCodeSegment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateCodeSegmentResponse), nil
	}
}

// UpdateCodeSegmentInvoker 更新代码片段
func (c *SecMasterClient) UpdateCodeSegmentInvoker(request *model.UpdateCodeSegmentRequest) *UpdateCodeSegmentInvoker {
	requestDef := GenReqDefForUpdateCodeSegment()
	return &UpdateCodeSegmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateCompliancePackage 更新遵从包
//
// 更新遵从包
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) UpdateCompliancePackage(request *model.UpdateCompliancePackageRequest) (*model.UpdateCompliancePackageResponse, error) {
	requestDef := GenReqDefForUpdateCompliancePackage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateCompliancePackageResponse), nil
	}
}

// UpdateCompliancePackageInvoker 更新遵从包
func (c *SecMasterClient) UpdateCompliancePackageInvoker(request *model.UpdateCompliancePackageRequest) *UpdateCompliancePackageInvoker {
	requestDef := GenReqDefForUpdateCompliancePackage()
	return &UpdateCompliancePackageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateDataTransformation 更新数据加工
//
// 更新数据加工
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) UpdateDataTransformation(request *model.UpdateDataTransformationRequest) (*model.UpdateDataTransformationResponse, error) {
	requestDef := GenReqDefForUpdateDataTransformation()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDataTransformationResponse), nil
	}
}

// UpdateDataTransformationInvoker 更新数据加工
func (c *SecMasterClient) UpdateDataTransformationInvoker(request *model.UpdateDataTransformationRequest) *UpdateDataTransformationInvoker {
	requestDef := GenReqDefForUpdateDataTransformation()
	return &UpdateDataTransformationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateLayoutField 更新字段
//
// 更新布局字段
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) UpdateLayoutField(request *model.UpdateLayoutFieldRequest) (*model.UpdateLayoutFieldResponse, error) {
	requestDef := GenReqDefForUpdateLayoutField()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateLayoutFieldResponse), nil
	}
}

// UpdateLayoutFieldInvoker 更新字段
func (c *SecMasterClient) UpdateLayoutFieldInvoker(request *model.UpdateLayoutFieldRequest) *UpdateLayoutFieldInvoker {
	requestDef := GenReqDefForUpdateLayoutField()
	return &UpdateLayoutFieldInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePipe 更新管道
//
// 更新管道
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) UpdatePipe(request *model.UpdatePipeRequest) (*model.UpdatePipeResponse, error) {
	requestDef := GenReqDefForUpdatePipe()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePipeResponse), nil
	}
}

// UpdatePipeInvoker 更新管道
func (c *SecMasterClient) UpdatePipeInvoker(request *model.UpdatePipeRequest) *UpdatePipeInvoker {
	requestDef := GenReqDefForUpdatePipe()
	return &UpdatePipeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePipeSchema 更新管道结构
//
// 更新管道结构
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) UpdatePipeSchema(request *model.UpdatePipeSchemaRequest) (*model.UpdatePipeSchemaResponse, error) {
	requestDef := GenReqDefForUpdatePipeSchema()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePipeSchemaResponse), nil
	}
}

// UpdatePipeSchemaInvoker 更新管道结构
func (c *SecMasterClient) UpdatePipeSchemaInvoker(request *model.UpdatePipeSchemaRequest) *UpdatePipeSchemaInvoker {
	requestDef := GenReqDefForUpdatePipeSchema()
	return &UpdatePipeSchemaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateRetrieveScript 更新检索脚本
//
// 更新检索脚本
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) UpdateRetrieveScript(request *model.UpdateRetrieveScriptRequest) (*model.UpdateRetrieveScriptResponse, error) {
	requestDef := GenReqDefForUpdateRetrieveScript()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateRetrieveScriptResponse), nil
	}
}

// UpdateRetrieveScriptInvoker 更新检索脚本
func (c *SecMasterClient) UpdateRetrieveScriptInvoker(request *model.UpdateRetrieveScriptRequest) *UpdateRetrieveScriptInvoker {
	requestDef := GenReqDefForUpdateRetrieveScript()
	return &UpdateRetrieveScriptInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateTable 更改表详情
//
// 更改表详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) UpdateTable(request *model.UpdateTableRequest) (*model.UpdateTableResponse, error) {
	requestDef := GenReqDefForUpdateTable()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTableResponse), nil
	}
}

// UpdateTableInvoker 更改表详情
func (c *SecMasterClient) UpdateTableInvoker(request *model.UpdateTableRequest) *UpdateTableInvoker {
	requestDef := GenReqDefForUpdateTable()
	return &UpdateTableInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateTableSchema 更改表结构
//
// 更改表结构
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) UpdateTableSchema(request *model.UpdateTableSchemaRequest) (*model.UpdateTableSchemaResponse, error) {
	requestDef := GenReqDefForUpdateTableSchema()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTableSchemaResponse), nil
	}
}

// UpdateTableSchemaInvoker 更改表结构
func (c *SecMasterClient) UpdateTableSchemaInvoker(request *model.UpdateTableSchemaRequest) *UpdateTableSchemaInvoker {
	requestDef := GenReqDefForUpdateTableSchema()
	return &UpdateTableSchemaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateSqlValidation 创建SQL校验
//
// 创建SQL校验
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) CreateSqlValidation(request *model.CreateSqlValidationRequest) (*model.CreateSqlValidationResponse, error) {
	requestDef := GenReqDefForCreateSqlValidation()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSqlValidationResponse), nil
	}
}

// CreateSqlValidationInvoker 创建SQL校验
func (c *SecMasterClient) CreateSqlValidationInvoker(request *model.CreateSqlValidationRequest) *CreateSqlValidationInvoker {
	requestDef := GenReqDefForCreateSqlValidation()
	return &CreateSqlValidationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
