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

//变更实例规格
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

//
func (c *GaussDBClient) CreateMysqlInstance(request *model.CreateMysqlInstanceRequest) (*model.CreateMysqlInstanceResponse, error) {
	requestDef := GenReqDefForCreateMysqlInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateMysqlInstanceResponse), nil
	}
}

//
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

//
func (c *GaussDBClient) DeleteMysqlInstance(request *model.DeleteMysqlInstanceRequest) (*model.DeleteMysqlInstanceResponse, error) {
	requestDef := GenReqDefForDeleteMysqlInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteMysqlInstanceResponse), nil
	}
}

//
func (c *GaussDBClient) DeleteMysqlProxy(request *model.DeleteMysqlProxyRequest) (*model.DeleteMysqlProxyResponse, error) {
	requestDef := GenReqDefForDeleteMysqlProxy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteMysqlProxyResponse), nil
	}
}

//
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

//
func (c *GaussDBClient) ExpandMysqlProxy(request *model.ExpandMysqlProxyRequest) (*model.ExpandMysqlProxyResponse, error) {
	requestDef := GenReqDefForExpandMysqlProxy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExpandMysqlProxyResponse), nil
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

//
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

//
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

//
func (c *GaussDBClient) ShowMysqlBackupPolicy(request *model.ShowMysqlBackupPolicyRequest) (*model.ShowMysqlBackupPolicyResponse, error) {
	requestDef := GenReqDefForShowMysqlBackupPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMysqlBackupPolicyResponse), nil
	}
}

//
func (c *GaussDBClient) ShowMysqlEngineVersion(request *model.ShowMysqlEngineVersionRequest) (*model.ShowMysqlEngineVersionResponse, error) {
	requestDef := GenReqDefForShowMysqlEngineVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMysqlEngineVersionResponse), nil
	}
}

//
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

//
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

//
func (c *GaussDBClient) ShowMysqlProjectQuotas(request *model.ShowMysqlProjectQuotasRequest) (*model.ShowMysqlProjectQuotasResponse, error) {
	requestDef := GenReqDefForShowMysqlProjectQuotas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMysqlProjectQuotasResponse), nil
	}
}

//
func (c *GaussDBClient) ShowMysqlProxy(request *model.ShowMysqlProxyRequest) (*model.ShowMysqlProxyResponse, error) {
	requestDef := GenReqDefForShowMysqlProxy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMysqlProxyResponse), nil
	}
}

//
func (c *GaussDBClient) ShowMysqlProxyFlavors(request *model.ShowMysqlProxyFlavorsRequest) (*model.ShowMysqlProxyFlavorsResponse, error) {
	requestDef := GenReqDefForShowMysqlProxyFlavors()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMysqlProxyFlavorsResponse), nil
	}
}

//
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

//
func (c *GaussDBClient) UpdateMysqlQuotas(request *model.UpdateMysqlQuotasRequest) (*model.UpdateMysqlQuotasResponse, error) {
	requestDef := GenReqDefForUpdateMysqlQuotas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateMysqlQuotasResponse), nil
	}
}
