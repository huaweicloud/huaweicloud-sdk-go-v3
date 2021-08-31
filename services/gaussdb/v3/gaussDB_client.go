package v3

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/gaussdb/v3/model"
)

type GaussDBClient struct {
	HcClient *http_client.HcHttpClient
}

func NewGaussDBClient(hcClient *http_client.HcHttpClient) *GaussDBClient {
	return &GaussDBClient{HcClient: hcClient}
}

func GaussDBClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

//变更数据库实例的规格。
func (c *GaussDBClient) ChangeMysqlInstanceSpecification(request *model.ChangeMysqlInstanceSpecificationRequest) (*model.ChangeMysqlInstanceSpecificationResponse, error) {
	requestDef := GenReqDefForChangeMysqlInstanceSpecification()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeMysqlInstanceSpecificationResponse), nil
	}
}

//创建手动备份
func (c *GaussDBClient) CreateMysqlBackup(request *model.CreateMysqlBackupRequest) (*model.CreateMysqlBackupResponse, error) {
	requestDef := GenReqDefForCreateMysqlBackup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateMysqlBackupResponse), nil
	}
}

//创建云数据库 GaussDB(for MySQL)实例。
func (c *GaussDBClient) CreateMysqlInstance(request *model.CreateMysqlInstanceRequest) (*model.CreateMysqlInstanceResponse, error) {
	requestDef := GenReqDefForCreateMysqlInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateMysqlInstanceResponse), nil
	}
}

//开启数据库代理，只支持ELB模式。
func (c *GaussDBClient) CreateMysqlProxy(request *model.CreateMysqlProxyRequest) (*model.CreateMysqlProxyResponse, error) {
	requestDef := GenReqDefForCreateMysqlProxy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateMysqlProxyResponse), nil
	}
}

//创建只读节点。
func (c *GaussDBClient) CreateMysqlReadonlyNode(request *model.CreateMysqlReadonlyNodeRequest) (*model.CreateMysqlReadonlyNodeResponse, error) {
	requestDef := GenReqDefForCreateMysqlReadonlyNode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateMysqlReadonlyNodeResponse), nil
	}
}

//删除数据库实例，不支持删除包周期实例。
func (c *GaussDBClient) DeleteMysqlInstance(request *model.DeleteMysqlInstanceRequest) (*model.DeleteMysqlInstanceResponse, error) {
	requestDef := GenReqDefForDeleteMysqlInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteMysqlInstanceResponse), nil
	}
}

//关闭数据库代理。
func (c *GaussDBClient) DeleteMysqlProxy(request *model.DeleteMysqlProxyRequest) (*model.DeleteMysqlProxyResponse, error) {
	requestDef := GenReqDefForDeleteMysqlProxy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteMysqlProxyResponse), nil
	}
}

//删除实例的只读节点。多可用区模式删除只读节点时，要保证删除后，剩余的只读节点和主节点在不同的可用区中，否则无法删除该只读节点。
func (c *GaussDBClient) DeleteMysqlReadonlyNode(request *model.DeleteMysqlReadonlyNodeRequest) (*model.DeleteMysqlReadonlyNodeResponse, error) {
	requestDef := GenReqDefForDeleteMysqlReadonlyNode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteMysqlReadonlyNodeResponse), nil
	}
}

//包周期存储扩容
func (c *GaussDBClient) ExpandMysqlInstanceVolume(request *model.ExpandMysqlInstanceVolumeRequest) (*model.ExpandMysqlInstanceVolumeResponse, error) {
	requestDef := GenReqDefForExpandMysqlInstanceVolume()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExpandMysqlInstanceVolumeResponse), nil
	}
}

//扩容数据库代理节点的数量。 DeC专属云账号暂不支持数据库代理。
func (c *GaussDBClient) ExpandMysqlProxy(request *model.ExpandMysqlProxyRequest) (*model.ExpandMysqlProxyResponse, error) {
	requestDef := GenReqDefForExpandMysqlProxy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExpandMysqlProxyResponse), nil
	}
}

//获取专属资源池列表，包括用户开通的所有专属资源池信息。
func (c *GaussDBClient) ListDedicatedResources(request *model.ListDedicatedResourcesRequest) (*model.ListDedicatedResourcesResponse, error) {
	requestDef := GenReqDefForListDedicatedResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDedicatedResourcesResponse), nil
	}
}

//获取参数模板列表，包括所有数据库的默认参数模板和用户创建的参数模板。
func (c *GaussDBClient) ListMysqlConfigurations(request *model.ListMysqlConfigurationsRequest) (*model.ListMysqlConfigurationsResponse, error) {
	requestDef := GenReqDefForListMysqlConfigurations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMysqlConfigurationsResponse), nil
	}
}

//查询数据库错误日志。
func (c *GaussDBClient) ListMysqlErrorLog(request *model.ListMysqlErrorLogRequest) (*model.ListMysqlErrorLogResponse, error) {
	requestDef := GenReqDefForListMysqlErrorLog()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMysqlErrorLogResponse), nil
	}
}

//查询数据库慢日志
func (c *GaussDBClient) ListMysqlSlowLog(request *model.ListMysqlSlowLogRequest) (*model.ListMysqlSlowLogResponse, error) {
	requestDef := GenReqDefForListMysqlSlowLog()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMysqlSlowLogResponse), nil
	}
}

//重置数据库密码
func (c *GaussDBClient) ResetMysqlPassword(request *model.ResetMysqlPasswordRequest) (*model.ResetMysqlPasswordResponse, error) {
	requestDef := GenReqDefForResetMysqlPassword()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResetMysqlPasswordResponse), nil
	}
}

//设置指定企业项目的资源配额。
func (c *GaussDBClient) SetMysqlQuotas(request *model.SetMysqlQuotasRequest) (*model.SetMysqlQuotasResponse, error) {
	requestDef := GenReqDefForSetMysqlQuotas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetMysqlQuotasResponse), nil
	}
}

//查询备份列表
func (c *GaussDBClient) ShowMysqlBackupList(request *model.ShowMysqlBackupListRequest) (*model.ShowMysqlBackupListResponse, error) {
	requestDef := GenReqDefForShowMysqlBackupList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMysqlBackupListResponse), nil
	}
}

//查询自动备份策略。
func (c *GaussDBClient) ShowMysqlBackupPolicy(request *model.ShowMysqlBackupPolicyRequest) (*model.ShowMysqlBackupPolicyResponse, error) {
	requestDef := GenReqDefForShowMysqlBackupPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMysqlBackupPolicyResponse), nil
	}
}

//获取指定数据库引擎对应的数据库版本信息。
func (c *GaussDBClient) ShowMysqlEngineVersion(request *model.ShowMysqlEngineVersionRequest) (*model.ShowMysqlEngineVersionResponse, error) {
	requestDef := GenReqDefForShowMysqlEngineVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMysqlEngineVersionResponse), nil
	}
}

//获取指定数据库引擎版本对应的规格信息。
func (c *GaussDBClient) ShowMysqlFlavors(request *model.ShowMysqlFlavorsRequest) (*model.ShowMysqlFlavorsResponse, error) {
	requestDef := GenReqDefForShowMysqlFlavors()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMysqlFlavorsResponse), nil
	}
}

//查询实例详情信息
func (c *GaussDBClient) ShowMysqlInstanceInfo(request *model.ShowMysqlInstanceInfoRequest) (*model.ShowMysqlInstanceInfoResponse, error) {
	requestDef := GenReqDefForShowMysqlInstanceInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMysqlInstanceInfoResponse), nil
	}
}

//根据指定条件查询实例列表。
func (c *GaussDBClient) ShowMysqlInstanceList(request *model.ShowMysqlInstanceListRequest) (*model.ShowMysqlInstanceListResponse, error) {
	requestDef := GenReqDefForShowMysqlInstanceList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMysqlInstanceListResponse), nil
	}
}

//获取指定ID的任务信息。
func (c *GaussDBClient) ShowMysqlJobInfo(request *model.ShowMysqlJobInfoRequest) (*model.ShowMysqlJobInfoResponse, error) {
	requestDef := GenReqDefForShowMysqlJobInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMysqlJobInfoResponse), nil
	}
}

//获取指定租户的资源配额。
func (c *GaussDBClient) ShowMysqlProjectQuotas(request *model.ShowMysqlProjectQuotasRequest) (*model.ShowMysqlProjectQuotasResponse, error) {
	requestDef := GenReqDefForShowMysqlProjectQuotas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMysqlProjectQuotasResponse), nil
	}
}

//查询数据库代理信息。
func (c *GaussDBClient) ShowMysqlProxy(request *model.ShowMysqlProxyRequest) (*model.ShowMysqlProxyResponse, error) {
	requestDef := GenReqDefForShowMysqlProxy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMysqlProxyResponse), nil
	}
}

//查询数据库代理规格信息。
func (c *GaussDBClient) ShowMysqlProxyFlavors(request *model.ShowMysqlProxyFlavorsRequest) (*model.ShowMysqlProxyFlavorsResponse, error) {
	requestDef := GenReqDefForShowMysqlProxyFlavors()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMysqlProxyFlavorsResponse), nil
	}
}

//获取指定企业项目的资源配额。
func (c *GaussDBClient) ShowMysqlQuotas(request *model.ShowMysqlQuotasRequest) (*model.ShowMysqlQuotasResponse, error) {
	requestDef := GenReqDefForShowMysqlQuotas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMysqlQuotasResponse), nil
	}
}

//修改备份策略
func (c *GaussDBClient) UpdateMysqlBackupPolicy(request *model.UpdateMysqlBackupPolicyRequest) (*model.UpdateMysqlBackupPolicyResponse, error) {
	requestDef := GenReqDefForUpdateMysqlBackupPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateMysqlBackupPolicyResponse), nil
	}
}

//修改实例名称
func (c *GaussDBClient) UpdateMysqlInstanceName(request *model.UpdateMysqlInstanceNameRequest) (*model.UpdateMysqlInstanceNameResponse, error) {
	requestDef := GenReqDefForUpdateMysqlInstanceName()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateMysqlInstanceNameResponse), nil
	}
}

//修改指定企业项目的资源配额。
func (c *GaussDBClient) UpdateMysqlQuotas(request *model.UpdateMysqlQuotasRequest) (*model.UpdateMysqlQuotasResponse, error) {
	requestDef := GenReqDefForUpdateMysqlQuotas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateMysqlQuotasResponse), nil
	}
}
