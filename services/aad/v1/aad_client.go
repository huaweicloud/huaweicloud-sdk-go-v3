package v1

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/aad/v1/model"
)

type AadClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewAadClient(hcClient *httpclient.HcHttpClient) *AadClient {
	return &AadClient{HcClient: hcClient}
}

func AadClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder().WithCredentialsType("global.Credentials")
	return builder
}

// ExecuteUnblockIp 解封IP
//
// 解封IP
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AadClient) ExecuteUnblockIp(request *model.ExecuteUnblockIpRequest) (*model.ExecuteUnblockIpResponse, error) {
	requestDef := GenReqDefForExecuteUnblockIp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteUnblockIpResponse), nil
	}
}

// ExecuteUnblockIpInvoker 解封IP
func (c *AadClient) ExecuteUnblockIpInvoker(request *model.ExecuteUnblockIpRequest) *ExecuteUnblockIpInvoker {
	requestDef := GenReqDefForExecuteUnblockIp()
	return &ExecuteUnblockIpInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListBlockIps 查询租户封堵列表
//
// 查询租户封堵列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AadClient) ListBlockIps(request *model.ListBlockIpsRequest) (*model.ListBlockIpsResponse, error) {
	requestDef := GenReqDefForListBlockIps()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBlockIpsResponse), nil
	}
}

// ListBlockIpsInvoker 查询租户封堵列表
func (c *AadClient) ListBlockIpsInvoker(request *model.ListBlockIpsRequest) *ListBlockIpsInvoker {
	requestDef := GenReqDefForListBlockIps()
	return &ListBlockIpsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListUnblockQuotaStatistics 查询解封配额
//
// 查询解封配额
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AadClient) ListUnblockQuotaStatistics(request *model.ListUnblockQuotaStatisticsRequest) (*model.ListUnblockQuotaStatisticsResponse, error) {
	requestDef := GenReqDefForListUnblockQuotaStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListUnblockQuotaStatisticsResponse), nil
	}
}

// ListUnblockQuotaStatisticsInvoker 查询解封配额
func (c *AadClient) ListUnblockQuotaStatisticsInvoker(request *model.ListUnblockQuotaStatisticsRequest) *ListUnblockQuotaStatisticsInvoker {
	requestDef := GenReqDefForListUnblockQuotaStatistics()
	return &ListUnblockQuotaStatisticsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBlockStatistics 查询封堵统计数据
//
// 查询封堵统计数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AadClient) ShowBlockStatistics(request *model.ShowBlockStatisticsRequest) (*model.ShowBlockStatisticsResponse, error) {
	requestDef := GenReqDefForShowBlockStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBlockStatisticsResponse), nil
	}
}

// ShowBlockStatisticsInvoker 查询封堵统计数据
func (c *AadClient) ShowBlockStatisticsInvoker(request *model.ShowBlockStatisticsRequest) *ShowBlockStatisticsInvoker {
	requestDef := GenReqDefForShowBlockStatistics()
	return &ShowBlockStatisticsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowUnblockRecord 查询租户解封记录
//
// 查询租户解封记录
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AadClient) ShowUnblockRecord(request *model.ShowUnblockRecordRequest) (*model.ShowUnblockRecordResponse, error) {
	requestDef := GenReqDefForShowUnblockRecord()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowUnblockRecordResponse), nil
	}
}

// ShowUnblockRecordInvoker 查询租户解封记录
func (c *AadClient) ShowUnblockRecordInvoker(request *model.ShowUnblockRecordRequest) *ShowUnblockRecordInvoker {
	requestDef := GenReqDefForShowUnblockRecord()
	return &ShowUnblockRecordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddBlackWhiteIpList 高防实例添加黑白名单
//
// 高防实例添加黑白名单
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AadClient) AddBlackWhiteIpList(request *model.AddBlackWhiteIpListRequest) (*model.AddBlackWhiteIpListResponse, error) {
	requestDef := GenReqDefForAddBlackWhiteIpList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddBlackWhiteIpListResponse), nil
	}
}

// AddBlackWhiteIpListInvoker 高防实例添加黑白名单
func (c *AadClient) AddBlackWhiteIpListInvoker(request *model.AddBlackWhiteIpListRequest) *AddBlackWhiteIpListInvoker {
	requestDef := GenReqDefForAddBlackWhiteIpList()
	return &AddBlackWhiteIpListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddPolicyBlackAndWhiteIpList 策略添加黑白名单
//
// 策略添加黑白名单
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AadClient) AddPolicyBlackAndWhiteIpList(request *model.AddPolicyBlackAndWhiteIpListRequest) (*model.AddPolicyBlackAndWhiteIpListResponse, error) {
	requestDef := GenReqDefForAddPolicyBlackAndWhiteIpList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddPolicyBlackAndWhiteIpListResponse), nil
	}
}

// AddPolicyBlackAndWhiteIpListInvoker 策略添加黑白名单
func (c *AadClient) AddPolicyBlackAndWhiteIpListInvoker(request *model.AddPolicyBlackAndWhiteIpListRequest) *AddPolicyBlackAndWhiteIpListInvoker {
	requestDef := GenReqDefForAddPolicyBlackAndWhiteIpList()
	return &AddPolicyBlackAndWhiteIpListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AssociateIpToPolicy 策略绑定防护对象
//
// 策略绑定防护对象
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AadClient) AssociateIpToPolicy(request *model.AssociateIpToPolicyRequest) (*model.AssociateIpToPolicyResponse, error) {
	requestDef := GenReqDefForAssociateIpToPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AssociateIpToPolicyResponse), nil
	}
}

// AssociateIpToPolicyInvoker 策略绑定防护对象
func (c *AadClient) AssociateIpToPolicyInvoker(request *model.AssociateIpToPolicyRequest) *AssociateIpToPolicyInvoker {
	requestDef := GenReqDefForAssociateIpToPolicy()
	return &AssociateIpToPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AssociateIpToPolicyAndPackage 策略和防护包绑定防护对象
//
// 策略和防护包绑定防护对象
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AadClient) AssociateIpToPolicyAndPackage(request *model.AssociateIpToPolicyAndPackageRequest) (*model.AssociateIpToPolicyAndPackageResponse, error) {
	requestDef := GenReqDefForAssociateIpToPolicyAndPackage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AssociateIpToPolicyAndPackageResponse), nil
	}
}

// AssociateIpToPolicyAndPackageInvoker 策略和防护包绑定防护对象
func (c *AadClient) AssociateIpToPolicyAndPackageInvoker(request *model.AssociateIpToPolicyAndPackageRequest) *AssociateIpToPolicyAndPackageInvoker {
	requestDef := GenReqDefForAssociateIpToPolicyAndPackage()
	return &AssociateIpToPolicyAndPackageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchCreateInstanceIpRule 批量创建高防实例IP的转发规则
//
// 批量创建高防实例IP的转发规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AadClient) BatchCreateInstanceIpRule(request *model.BatchCreateInstanceIpRuleRequest) (*model.BatchCreateInstanceIpRuleResponse, error) {
	requestDef := GenReqDefForBatchCreateInstanceIpRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateInstanceIpRuleResponse), nil
	}
}

// BatchCreateInstanceIpRuleInvoker 批量创建高防实例IP的转发规则
func (c *AadClient) BatchCreateInstanceIpRuleInvoker(request *model.BatchCreateInstanceIpRuleRequest) *BatchCreateInstanceIpRuleInvoker {
	requestDef := GenReqDefForBatchCreateInstanceIpRule()
	return &BatchCreateInstanceIpRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteInstanceIpRule 批量删除高防实例IP的转发规则
//
// 批量删除高防实例IP的转发规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AadClient) BatchDeleteInstanceIpRule(request *model.BatchDeleteInstanceIpRuleRequest) (*model.BatchDeleteInstanceIpRuleResponse, error) {
	requestDef := GenReqDefForBatchDeleteInstanceIpRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteInstanceIpRuleResponse), nil
	}
}

// BatchDeleteInstanceIpRuleInvoker 批量删除高防实例IP的转发规则
func (c *AadClient) BatchDeleteInstanceIpRuleInvoker(request *model.BatchDeleteInstanceIpRuleRequest) *BatchDeleteInstanceIpRuleInvoker {
	requestDef := GenReqDefForBatchDeleteInstanceIpRule()
	return &BatchDeleteInstanceIpRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAadDomain 创建防护域名
//
// 创建防护域名
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AadClient) CreateAadDomain(request *model.CreateAadDomainRequest) (*model.CreateAadDomainResponse, error) {
	requestDef := GenReqDefForCreateAadDomain()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAadDomainResponse), nil
	}
}

// CreateAadDomainInvoker 创建防护域名
func (c *AadClient) CreateAadDomainInvoker(request *model.CreateAadDomainRequest) *CreateAadDomainInvoker {
	requestDef := GenReqDefForCreateAadDomain()
	return &CreateAadDomainInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePolicy 创建策略
//
// 创建策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AadClient) CreatePolicy(request *model.CreatePolicyRequest) (*model.CreatePolicyResponse, error) {
	requestDef := GenReqDefForCreatePolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePolicyResponse), nil
	}
}

// CreatePolicyInvoker 创建策略
func (c *AadClient) CreatePolicyInvoker(request *model.CreatePolicyRequest) *CreatePolicyInvoker {
	requestDef := GenReqDefForCreatePolicy()
	return &CreatePolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAlarmConfig 删除告警配置
//
// 删除告警配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AadClient) DeleteAlarmConfig(request *model.DeleteAlarmConfigRequest) (*model.DeleteAlarmConfigResponse, error) {
	requestDef := GenReqDefForDeleteAlarmConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAlarmConfigResponse), nil
	}
}

// DeleteAlarmConfigInvoker 删除告警配置
func (c *AadClient) DeleteAlarmConfigInvoker(request *model.DeleteAlarmConfigRequest) *DeleteAlarmConfigInvoker {
	requestDef := GenReqDefForDeleteAlarmConfig()
	return &DeleteAlarmConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteBlackWhiteIpList 高防实例删除黑白名单
//
// 高防实例删除黑白名单
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AadClient) DeleteBlackWhiteIpList(request *model.DeleteBlackWhiteIpListRequest) (*model.DeleteBlackWhiteIpListResponse, error) {
	requestDef := GenReqDefForDeleteBlackWhiteIpList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteBlackWhiteIpListResponse), nil
	}
}

// DeleteBlackWhiteIpListInvoker 高防实例删除黑白名单
func (c *AadClient) DeleteBlackWhiteIpListInvoker(request *model.DeleteBlackWhiteIpListRequest) *DeleteBlackWhiteIpListInvoker {
	requestDef := GenReqDefForDeleteBlackWhiteIpList()
	return &DeleteBlackWhiteIpListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeletePolicy 删除策略
//
// 删除策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AadClient) DeletePolicy(request *model.DeletePolicyRequest) (*model.DeletePolicyResponse, error) {
	requestDef := GenReqDefForDeletePolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePolicyResponse), nil
	}
}

// DeletePolicyInvoker 删除策略
func (c *AadClient) DeletePolicyInvoker(request *model.DeletePolicyRequest) *DeletePolicyInvoker {
	requestDef := GenReqDefForDeletePolicy()
	return &DeletePolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeletePolicyBlackAndWhiteIpList 策略删除黑白名单
//
// 策略删除黑白名单
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AadClient) DeletePolicyBlackAndWhiteIpList(request *model.DeletePolicyBlackAndWhiteIpListRequest) (*model.DeletePolicyBlackAndWhiteIpListResponse, error) {
	requestDef := GenReqDefForDeletePolicyBlackAndWhiteIpList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePolicyBlackAndWhiteIpListResponse), nil
	}
}

// DeletePolicyBlackAndWhiteIpListInvoker 策略删除黑白名单
func (c *AadClient) DeletePolicyBlackAndWhiteIpListInvoker(request *model.DeletePolicyBlackAndWhiteIpListRequest) *DeletePolicyBlackAndWhiteIpListInvoker {
	requestDef := GenReqDefForDeletePolicyBlackAndWhiteIpList()
	return &DeletePolicyBlackAndWhiteIpListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DisassociateIpFromPolicy 策略解绑防护对象
//
// 策略解绑防护对象
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AadClient) DisassociateIpFromPolicy(request *model.DisassociateIpFromPolicyRequest) (*model.DisassociateIpFromPolicyResponse, error) {
	requestDef := GenReqDefForDisassociateIpFromPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisassociateIpFromPolicyResponse), nil
	}
}

// DisassociateIpFromPolicyInvoker 策略解绑防护对象
func (c *AadClient) DisassociateIpFromPolicyInvoker(request *model.DisassociateIpFromPolicyRequest) *DisassociateIpFromPolicyInvoker {
	requestDef := GenReqDefForDisassociateIpFromPolicy()
	return &DisassociateIpFromPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DisassociateIpFromPolicyAndPackage 策略和防护包解绑防护对象
//
// 策略和防护包解绑防护对象
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AadClient) DisassociateIpFromPolicyAndPackage(request *model.DisassociateIpFromPolicyAndPackageRequest) (*model.DisassociateIpFromPolicyAndPackageResponse, error) {
	requestDef := GenReqDefForDisassociateIpFromPolicyAndPackage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisassociateIpFromPolicyAndPackageResponse), nil
	}
}

// DisassociateIpFromPolicyAndPackageInvoker 策略和防护包解绑防护对象
func (c *AadClient) DisassociateIpFromPolicyAndPackageInvoker(request *model.DisassociateIpFromPolicyAndPackageRequest) *DisassociateIpFromPolicyAndPackageInvoker {
	requestDef := GenReqDefForDisassociateIpFromPolicyAndPackage()
	return &DisassociateIpFromPolicyAndPackageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDomain 查询域名列表
//
// 查询域名列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AadClient) ListDomain(request *model.ListDomainRequest) (*model.ListDomainResponse, error) {
	requestDef := GenReqDefForListDomain()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDomainResponse), nil
	}
}

// ListDomainInvoker 查询域名列表
func (c *AadClient) ListDomainInvoker(request *model.ListDomainRequest) *ListDomainInvoker {
	requestDef := GenReqDefForListDomain()
	return &ListDomainInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstance 查询高防实例列表
//
// 查询高防实例列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AadClient) ListInstance(request *model.ListInstanceRequest) (*model.ListInstanceResponse, error) {
	requestDef := GenReqDefForListInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstanceResponse), nil
	}
}

// ListInstanceInvoker 查询高防实例列表
func (c *AadClient) ListInstanceInvoker(request *model.ListInstanceRequest) *ListInstanceInvoker {
	requestDef := GenReqDefForListInstance()
	return &ListInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstanceId 查询域名关联的实例ID
//
// 查询域名关联的实例ID
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AadClient) ListInstanceId(request *model.ListInstanceIdRequest) (*model.ListInstanceIdResponse, error) {
	requestDef := GenReqDefForListInstanceId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstanceIdResponse), nil
	}
}

// ListInstanceIdInvoker 查询域名关联的实例ID
func (c *AadClient) ListInstanceIdInvoker(request *model.ListInstanceIdRequest) *ListInstanceIdInvoker {
	requestDef := GenReqDefForListInstanceId()
	return &ListInstanceIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstanceIpRule 查询高防实例IP的转发规则列表
//
// 查询高防实例IP的转发规则列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AadClient) ListInstanceIpRule(request *model.ListInstanceIpRuleRequest) (*model.ListInstanceIpRuleResponse, error) {
	requestDef := GenReqDefForListInstanceIpRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstanceIpRuleResponse), nil
	}
}

// ListInstanceIpRuleInvoker 查询高防实例IP的转发规则列表
func (c *AadClient) ListInstanceIpRuleInvoker(request *model.ListInstanceIpRuleRequest) *ListInstanceIpRuleInvoker {
	requestDef := GenReqDefForListInstanceIpRule()
	return &ListInstanceIpRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPackage 查询防护包列表
//
// 查询防护包列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AadClient) ListPackage(request *model.ListPackageRequest) (*model.ListPackageResponse, error) {
	requestDef := GenReqDefForListPackage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPackageResponse), nil
	}
}

// ListPackageInvoker 查询防护包列表
func (c *AadClient) ListPackageInvoker(request *model.ListPackageRequest) *ListPackageInvoker {
	requestDef := GenReqDefForListPackage()
	return &ListPackageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPeak 查询流量峰值、攻击计数
//
// 查询流量峰值、攻击计数
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AadClient) ListPeak(request *model.ListPeakRequest) (*model.ListPeakResponse, error) {
	requestDef := GenReqDefForListPeak()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPeakResponse), nil
	}
}

// ListPeakInvoker 查询流量峰值、攻击计数
func (c *AadClient) ListPeakInvoker(request *model.ListPeakRequest) *ListPeakInvoker {
	requestDef := GenReqDefForListPeak()
	return &ListPeakInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPolicy 查询策略列表
//
// 查询策略列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AadClient) ListPolicy(request *model.ListPolicyRequest) (*model.ListPolicyResponse, error) {
	requestDef := GenReqDefForListPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPolicyResponse), nil
	}
}

// ListPolicyInvoker 查询策略列表
func (c *AadClient) ListPolicyInvoker(request *model.ListPolicyRequest) *ListPolicyInvoker {
	requestDef := GenReqDefForListPolicy()
	return &ListPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProtectedIp 查询防护对象列表
//
// 查询防护对象列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AadClient) ListProtectedIp(request *model.ListProtectedIpRequest) (*model.ListProtectedIpResponse, error) {
	requestDef := GenReqDefForListProtectedIp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProtectedIpResponse), nil
	}
}

// ListProtectedIpInvoker 查询防护对象列表
func (c *AadClient) ListProtectedIpInvoker(request *model.ListProtectedIpRequest) *ListProtectedIpInvoker {
	requestDef := GenReqDefForListProtectedIp()
	return &ListProtectedIpInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSourceIps 查询高防回源IP段列表
//
// 查询高防回源IP段列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AadClient) ListSourceIps(request *model.ListSourceIpsRequest) (*model.ListSourceIpsResponse, error) {
	requestDef := GenReqDefForListSourceIps()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSourceIpsResponse), nil
	}
}

// ListSourceIpsInvoker 查询高防回源IP段列表
func (c *AadClient) ListSourceIpsInvoker(request *model.ListSourceIpsRequest) *ListSourceIpsInvoker {
	requestDef := GenReqDefForListSourceIps()
	return &ListSourceIpsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListUnboundProtectedIp 查询可绑定的防护对象列表
//
// 查询可绑定的防护对象列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AadClient) ListUnboundProtectedIp(request *model.ListUnboundProtectedIpRequest) (*model.ListUnboundProtectedIpResponse, error) {
	requestDef := GenReqDefForListUnboundProtectedIp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListUnboundProtectedIpResponse), nil
	}
}

// ListUnboundProtectedIpInvoker 查询可绑定的防护对象列表
func (c *AadClient) ListUnboundProtectedIpInvoker(request *model.ListUnboundProtectedIpRequest) *ListUnboundProtectedIpInvoker {
	requestDef := GenReqDefForListUnboundProtectedIp()
	return &ListUnboundProtectedIpInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ModifyDomainWebSwitch 修改域名WEB基础防护开关/CC防护开关
//
// 修改域名WEB基础防护开关/CC防护开关
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AadClient) ModifyDomainWebSwitch(request *model.ModifyDomainWebSwitchRequest) (*model.ModifyDomainWebSwitchResponse, error) {
	requestDef := GenReqDefForModifyDomainWebSwitch()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ModifyDomainWebSwitchResponse), nil
	}
}

// ModifyDomainWebSwitchInvoker 修改域名WEB基础防护开关/CC防护开关
func (c *AadClient) ModifyDomainWebSwitchInvoker(request *model.ModifyDomainWebSwitchRequest) *ModifyDomainWebSwitchInvoker {
	requestDef := GenReqDefForModifyDomainWebSwitch()
	return &ModifyDomainWebSwitchInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetCertForDomain 上传/修改域名对应证书
//
// 上传/修改域名对应证书
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AadClient) SetCertForDomain(request *model.SetCertForDomainRequest) (*model.SetCertForDomainResponse, error) {
	requestDef := GenReqDefForSetCertForDomain()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetCertForDomainResponse), nil
	}
}

// SetCertForDomainInvoker 上传/修改域名对应证书
func (c *AadClient) SetCertForDomainInvoker(request *model.SetCertForDomainRequest) *SetCertForDomainInvoker {
	requestDef := GenReqDefForSetCertForDomain()
	return &SetCertForDomainInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAlarmConfig 查询告警配置
//
// 查询告警配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AadClient) ShowAlarmConfig(request *model.ShowAlarmConfigRequest) (*model.ShowAlarmConfigResponse, error) {
	requestDef := GenReqDefForShowAlarmConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAlarmConfigResponse), nil
	}
}

// ShowAlarmConfigInvoker 查询告警配置
func (c *AadClient) ShowAlarmConfigInvoker(request *model.ShowAlarmConfigRequest) *ShowAlarmConfigInvoker {
	requestDef := GenReqDefForShowAlarmConfig()
	return &ShowAlarmConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPolicy 查询策略详情
//
// 查询策略详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AadClient) ShowPolicy(request *model.ShowPolicyRequest) (*model.ShowPolicyResponse, error) {
	requestDef := GenReqDefForShowPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPolicyResponse), nil
	}
}

// ShowPolicyInvoker 查询策略详情
func (c *AadClient) ShowPolicyInvoker(request *model.ShowPolicyRequest) *ShowPolicyInvoker {
	requestDef := GenReqDefForShowPolicy()
	return &ShowPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAlarmConfig 设置告警配置
//
// 设置告警配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AadClient) UpdateAlarmConfig(request *model.UpdateAlarmConfigRequest) (*model.UpdateAlarmConfigResponse, error) {
	requestDef := GenReqDefForUpdateAlarmConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAlarmConfigResponse), nil
	}
}

// UpdateAlarmConfigInvoker 设置告警配置
func (c *AadClient) UpdateAlarmConfigInvoker(request *model.UpdateAlarmConfigRequest) *UpdateAlarmConfigInvoker {
	requestDef := GenReqDefForUpdateAlarmConfig()
	return &UpdateAlarmConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateDomain 更新域名信息
//
// 更新域名源站配置信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AadClient) UpdateDomain(request *model.UpdateDomainRequest) (*model.UpdateDomainResponse, error) {
	requestDef := GenReqDefForUpdateDomain()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDomainResponse), nil
	}
}

// UpdateDomainInvoker 更新域名信息
func (c *AadClient) UpdateDomainInvoker(request *model.UpdateDomainRequest) *UpdateDomainInvoker {
	requestDef := GenReqDefForUpdateDomain()
	return &UpdateDomainInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateInstanceIpRule 修改高防实例转发配置的源站IP
//
// 修改高防实例转发配置的源站IP
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AadClient) UpdateInstanceIpRule(request *model.UpdateInstanceIpRuleRequest) (*model.UpdateInstanceIpRuleResponse, error) {
	requestDef := GenReqDefForUpdateInstanceIpRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInstanceIpRuleResponse), nil
	}
}

// UpdateInstanceIpRuleInvoker 修改高防实例转发配置的源站IP
func (c *AadClient) UpdateInstanceIpRuleInvoker(request *model.UpdateInstanceIpRuleRequest) *UpdateInstanceIpRuleInvoker {
	requestDef := GenReqDefForUpdateInstanceIpRule()
	return &UpdateInstanceIpRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePackageIp 更新防护包绑定的全量防护对象
//
// 更新防护包绑定的全量防护对象
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AadClient) UpdatePackageIp(request *model.UpdatePackageIpRequest) (*model.UpdatePackageIpResponse, error) {
	requestDef := GenReqDefForUpdatePackageIp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePackageIpResponse), nil
	}
}

// UpdatePackageIpInvoker 更新防护包绑定的全量防护对象
func (c *AadClient) UpdatePackageIpInvoker(request *model.UpdatePackageIpRequest) *UpdatePackageIpInvoker {
	requestDef := GenReqDefForUpdatePackageIp()
	return &UpdatePackageIpInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePackageName 更新防护包名字
//
// 更新防护包名字
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AadClient) UpdatePackageName(request *model.UpdatePackageNameRequest) (*model.UpdatePackageNameResponse, error) {
	requestDef := GenReqDefForUpdatePackageName()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePackageNameResponse), nil
	}
}

// UpdatePackageNameInvoker 更新防护包名字
func (c *AadClient) UpdatePackageNameInvoker(request *model.UpdatePackageNameRequest) *UpdatePackageNameInvoker {
	requestDef := GenReqDefForUpdatePackageName()
	return &UpdatePackageNameInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePolicy 更新策略
//
// 更新策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AadClient) UpdatePolicy(request *model.UpdatePolicyRequest) (*model.UpdatePolicyResponse, error) {
	requestDef := GenReqDefForUpdatePolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePolicyResponse), nil
	}
}

// UpdatePolicyInvoker 更新策略
func (c *AadClient) UpdatePolicyInvoker(request *model.UpdatePolicyRequest) *UpdatePolicyInvoker {
	requestDef := GenReqDefForUpdatePolicy()
	return &UpdatePolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateTagForProtectedIp 防护对象设置标签
//
// 防护对象设置标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AadClient) UpdateTagForProtectedIp(request *model.UpdateTagForProtectedIpRequest) (*model.UpdateTagForProtectedIpResponse, error) {
	requestDef := GenReqDefForUpdateTagForProtectedIp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTagForProtectedIpResponse), nil
	}
}

// UpdateTagForProtectedIpInvoker 防护对象设置标签
func (c *AadClient) UpdateTagForProtectedIpInvoker(request *model.UpdateTagForProtectedIpRequest) *UpdateTagForProtectedIpInvoker {
	requestDef := GenReqDefForUpdateTagForProtectedIp()
	return &UpdateTagForProtectedIpInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
