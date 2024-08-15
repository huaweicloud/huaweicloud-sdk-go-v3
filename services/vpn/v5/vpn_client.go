package v5

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpn/v5/model"
)

type VpnClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewVpnClient(hcClient *httpclient.HcHttpClient) *VpnClient {
	return &VpnClient{HcClient: hcClient}
}

func VpnClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// CheckClientCaCertificate 校验客户端CA
//
// 创建服务端时，可以先调用客户端CA的预校验API，检查CA的合法性
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpnClient) CheckClientCaCertificate(request *model.CheckClientCaCertificateRequest) (*model.CheckClientCaCertificateResponse, error) {
	requestDef := GenReqDefForCheckClientCaCertificate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckClientCaCertificateResponse), nil
	}
}

// CheckClientCaCertificateInvoker 校验客户端CA
func (c *VpnClient) CheckClientCaCertificateInvoker(request *model.CheckClientCaCertificateRequest) *CheckClientCaCertificateInvoker {
	requestDef := GenReqDefForCheckClientCaCertificate()
	return &CheckClientCaCertificateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteClientCa 删除客户端的CA证书
//
// 删除客户端CA证书
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpnClient) DeleteClientCa(request *model.DeleteClientCaRequest) (*model.DeleteClientCaResponse, error) {
	requestDef := GenReqDefForDeleteClientCa()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteClientCaResponse), nil
	}
}

// DeleteClientCaInvoker 删除客户端的CA证书
func (c *VpnClient) DeleteClientCaInvoker(request *model.DeleteClientCaRequest) *DeleteClientCaInvoker {
	requestDef := GenReqDefForDeleteClientCa()
	return &DeleteClientCaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ImportClientCa 导入客户端 CA 证书
//
// 导入客户端 CA 证书
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpnClient) ImportClientCa(request *model.ImportClientCaRequest) (*model.ImportClientCaResponse, error) {
	requestDef := GenReqDefForImportClientCa()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ImportClientCaResponse), nil
	}
}

// ImportClientCaInvoker 导入客户端 CA 证书
func (c *VpnClient) ImportClientCaInvoker(request *model.ImportClientCaRequest) *ImportClientCaInvoker {
	requestDef := GenReqDefForImportClientCa()
	return &ImportClientCaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowClientCa 查询客户端的CA证书
//
// 查询客户端CA证书
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpnClient) ShowClientCa(request *model.ShowClientCaRequest) (*model.ShowClientCaResponse, error) {
	requestDef := GenReqDefForShowClientCa()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowClientCaResponse), nil
	}
}

// ShowClientCaInvoker 查询客户端的CA证书
func (c *VpnClient) ShowClientCaInvoker(request *model.ShowClientCaRequest) *ShowClientCaInvoker {
	requestDef := GenReqDefForShowClientCa()
	return &ShowClientCaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateClientCa 修改客户端的CA证书
//
// 修改客户端CA证书
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpnClient) UpdateClientCa(request *model.UpdateClientCaRequest) (*model.UpdateClientCaResponse, error) {
	requestDef := GenReqDefForUpdateClientCa()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateClientCaResponse), nil
	}
}

// UpdateClientCaInvoker 修改客户端的CA证书
func (c *VpnClient) UpdateClientCaInvoker(request *model.UpdateClientCaRequest) *UpdateClientCaInvoker {
	requestDef := GenReqDefForUpdateClientCa()
	return &UpdateClientCaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateConnectionMonitor 创建VPN连接监控
//
// 创建VPN连接监控
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpnClient) CreateConnectionMonitor(request *model.CreateConnectionMonitorRequest) (*model.CreateConnectionMonitorResponse, error) {
	requestDef := GenReqDefForCreateConnectionMonitor()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateConnectionMonitorResponse), nil
	}
}

// CreateConnectionMonitorInvoker 创建VPN连接监控
func (c *VpnClient) CreateConnectionMonitorInvoker(request *model.CreateConnectionMonitorRequest) *CreateConnectionMonitorInvoker {
	requestDef := GenReqDefForCreateConnectionMonitor()
	return &CreateConnectionMonitorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteConnectionMonitor 删除VPN连接监控
//
// 根据VPN连接监控的ID，删除指定的VPN连接监控
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpnClient) DeleteConnectionMonitor(request *model.DeleteConnectionMonitorRequest) (*model.DeleteConnectionMonitorResponse, error) {
	requestDef := GenReqDefForDeleteConnectionMonitor()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteConnectionMonitorResponse), nil
	}
}

// DeleteConnectionMonitorInvoker 删除VPN连接监控
func (c *VpnClient) DeleteConnectionMonitorInvoker(request *model.DeleteConnectionMonitorRequest) *DeleteConnectionMonitorInvoker {
	requestDef := GenReqDefForDeleteConnectionMonitor()
	return &DeleteConnectionMonitorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListConnectionMonitors 查询VPN连接监控列表
//
// 查询VPN连接监控列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpnClient) ListConnectionMonitors(request *model.ListConnectionMonitorsRequest) (*model.ListConnectionMonitorsResponse, error) {
	requestDef := GenReqDefForListConnectionMonitors()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListConnectionMonitorsResponse), nil
	}
}

// ListConnectionMonitorsInvoker 查询VPN连接监控列表
func (c *VpnClient) ListConnectionMonitorsInvoker(request *model.ListConnectionMonitorsRequest) *ListConnectionMonitorsInvoker {
	requestDef := GenReqDefForListConnectionMonitors()
	return &ListConnectionMonitorsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowConnectionMonitor 查询VPN连接监控
//
// 根据VPN连接监控的ID,查询指定的VPN连接监控
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpnClient) ShowConnectionMonitor(request *model.ShowConnectionMonitorRequest) (*model.ShowConnectionMonitorResponse, error) {
	requestDef := GenReqDefForShowConnectionMonitor()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowConnectionMonitorResponse), nil
	}
}

// ShowConnectionMonitorInvoker 查询VPN连接监控
func (c *VpnClient) ShowConnectionMonitorInvoker(request *model.ShowConnectionMonitorRequest) *ShowConnectionMonitorInvoker {
	requestDef := GenReqDefForShowConnectionMonitor()
	return &ShowConnectionMonitorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateCgw 创建对端网关
//
// 创建租户用于与VPN网关相连的对端网关
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpnClient) CreateCgw(request *model.CreateCgwRequest) (*model.CreateCgwResponse, error) {
	requestDef := GenReqDefForCreateCgw()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCgwResponse), nil
	}
}

// CreateCgwInvoker 创建对端网关
func (c *VpnClient) CreateCgwInvoker(request *model.CreateCgwRequest) *CreateCgwInvoker {
	requestDef := GenReqDefForCreateCgw()
	return &CreateCgwInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteCgw 删除对端网关
//
// 根据对端网关ID，删除指定的对端网关
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpnClient) DeleteCgw(request *model.DeleteCgwRequest) (*model.DeleteCgwResponse, error) {
	requestDef := GenReqDefForDeleteCgw()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteCgwResponse), nil
	}
}

// DeleteCgwInvoker 删除对端网关
func (c *VpnClient) DeleteCgwInvoker(request *model.DeleteCgwRequest) *DeleteCgwInvoker {
	requestDef := GenReqDefForDeleteCgw()
	return &DeleteCgwInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCgws 查询对端网关列表
//
// 查询对端网关列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpnClient) ListCgws(request *model.ListCgwsRequest) (*model.ListCgwsResponse, error) {
	requestDef := GenReqDefForListCgws()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCgwsResponse), nil
	}
}

// ListCgwsInvoker 查询对端网关列表
func (c *VpnClient) ListCgwsInvoker(request *model.ListCgwsRequest) *ListCgwsInvoker {
	requestDef := GenReqDefForListCgws()
	return &ListCgwsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCgw 查询对端网关
//
// 根据对端网关ID，查询指定的对端网关
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpnClient) ShowCgw(request *model.ShowCgwRequest) (*model.ShowCgwResponse, error) {
	requestDef := GenReqDefForShowCgw()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCgwResponse), nil
	}
}

// ShowCgwInvoker 查询对端网关
func (c *VpnClient) ShowCgwInvoker(request *model.ShowCgwRequest) *ShowCgwInvoker {
	requestDef := GenReqDefForShowCgw()
	return &ShowCgwInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateCgw 更新对端网关
//
// 根据对端网关ID，更新指定的对端网关
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpnClient) UpdateCgw(request *model.UpdateCgwRequest) (*model.UpdateCgwResponse, error) {
	requestDef := GenReqDefForUpdateCgw()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateCgwResponse), nil
	}
}

// UpdateCgwInvoker 更新对端网关
func (c *VpnClient) UpdateCgwInvoker(request *model.UpdateCgwRequest) *UpdateCgwInvoker {
	requestDef := GenReqDefForUpdateCgw()
	return &UpdateCgwInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListP2cVgwAvailabilityZones 查询P2C VPN网关可用区
//
// 查询P2C VPN网关可用区
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpnClient) ListP2cVgwAvailabilityZones(request *model.ListP2cVgwAvailabilityZonesRequest) (*model.ListP2cVgwAvailabilityZonesResponse, error) {
	requestDef := GenReqDefForListP2cVgwAvailabilityZones()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListP2cVgwAvailabilityZonesResponse), nil
	}
}

// ListP2cVgwAvailabilityZonesInvoker 查询P2C VPN网关可用区
func (c *VpnClient) ListP2cVgwAvailabilityZonesInvoker(request *model.ListP2cVgwAvailabilityZonesRequest) *ListP2cVgwAvailabilityZonesInvoker {
	requestDef := GenReqDefForListP2cVgwAvailabilityZones()
	return &ListP2cVgwAvailabilityZonesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListP2cVgwConnections 查询P2C VPN网关连接信息列表
//
// # List p2c vpn gateway connections
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpnClient) ListP2cVgwConnections(request *model.ListP2cVgwConnectionsRequest) (*model.ListP2cVgwConnectionsResponse, error) {
	requestDef := GenReqDefForListP2cVgwConnections()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListP2cVgwConnectionsResponse), nil
	}
}

// ListP2cVgwConnectionsInvoker 查询P2C VPN网关连接信息列表
func (c *VpnClient) ListP2cVgwConnectionsInvoker(request *model.ListP2cVgwConnectionsRequest) *ListP2cVgwConnectionsInvoker {
	requestDef := GenReqDefForListP2cVgwConnections()
	return &ListP2cVgwConnectionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListP2cVgws 查询P2C VPN网关列表
//
// 查询P2C VPN网关列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpnClient) ListP2cVgws(request *model.ListP2cVgwsRequest) (*model.ListP2cVgwsResponse, error) {
	requestDef := GenReqDefForListP2cVgws()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListP2cVgwsResponse), nil
	}
}

// ListP2cVgwsInvoker 查询P2C VPN网关列表
func (c *VpnClient) ListP2cVgwsInvoker(request *model.ListP2cVgwsRequest) *ListP2cVgwsInvoker {
	requestDef := GenReqDefForListP2cVgws()
	return &ListP2cVgwsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowP2cVgw 查询P2C VPN网关
//
// 根据P2C VPN网关ID，查询指定的VPN网关
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpnClient) ShowP2cVgw(request *model.ShowP2cVgwRequest) (*model.ShowP2cVgwResponse, error) {
	requestDef := GenReqDefForShowP2cVgw()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowP2cVgwResponse), nil
	}
}

// ShowP2cVgwInvoker 查询P2C VPN网关
func (c *VpnClient) ShowP2cVgwInvoker(request *model.ShowP2cVgwRequest) *ShowP2cVgwInvoker {
	requestDef := GenReqDefForShowP2cVgw()
	return &ShowP2cVgwInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateP2cVgw 更新P2C VPN网关
//
// 根据P2C VPN网关ID，更新指定的P2C VPN网关
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpnClient) UpdateP2cVgw(request *model.UpdateP2cVgwRequest) (*model.UpdateP2cVgwResponse, error) {
	requestDef := GenReqDefForUpdateP2cVgw()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateP2cVgwResponse), nil
	}
}

// UpdateP2cVgwInvoker 更新P2C VPN网关
func (c *VpnClient) UpdateP2cVgwInvoker(request *model.UpdateP2cVgwRequest) *UpdateP2cVgwInvoker {
	requestDef := GenReqDefForUpdateP2cVgw()
	return &UpdateP2cVgwInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchCreateResourceTags 批量添加资源标签
//
// 为指定实例批量添加标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpnClient) BatchCreateResourceTags(request *model.BatchCreateResourceTagsRequest) (*model.BatchCreateResourceTagsResponse, error) {
	requestDef := GenReqDefForBatchCreateResourceTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateResourceTagsResponse), nil
	}
}

// BatchCreateResourceTagsInvoker 批量添加资源标签
func (c *VpnClient) BatchCreateResourceTagsInvoker(request *model.BatchCreateResourceTagsRequest) *BatchCreateResourceTagsInvoker {
	requestDef := GenReqDefForBatchCreateResourceTags()
	return &BatchCreateResourceTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteResourceTags 批量删除资源标签
//
// 为指定实例批量删除标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpnClient) BatchDeleteResourceTags(request *model.BatchDeleteResourceTagsRequest) (*model.BatchDeleteResourceTagsResponse, error) {
	requestDef := GenReqDefForBatchDeleteResourceTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteResourceTagsResponse), nil
	}
}

// BatchDeleteResourceTagsInvoker 批量删除资源标签
func (c *VpnClient) BatchDeleteResourceTagsInvoker(request *model.BatchDeleteResourceTagsRequest) *BatchDeleteResourceTagsInvoker {
	requestDef := GenReqDefForBatchDeleteResourceTags()
	return &BatchDeleteResourceTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CountResourcesByTags 查询资源实例数量
//
// 根据标签查询资源实例数量
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpnClient) CountResourcesByTags(request *model.CountResourcesByTagsRequest) (*model.CountResourcesByTagsResponse, error) {
	requestDef := GenReqDefForCountResourcesByTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CountResourcesByTagsResponse), nil
	}
}

// CountResourcesByTagsInvoker 查询资源实例数量
func (c *VpnClient) CountResourcesByTagsInvoker(request *model.CountResourcesByTagsRequest) *CountResourcesByTagsInvoker {
	requestDef := GenReqDefForCountResourcesByTags()
	return &CountResourcesByTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProjectTags 查询项目标签
//
// 查询租户在指定项目中指定资源类型下的所有标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpnClient) ListProjectTags(request *model.ListProjectTagsRequest) (*model.ListProjectTagsResponse, error) {
	requestDef := GenReqDefForListProjectTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProjectTagsResponse), nil
	}
}

// ListProjectTagsInvoker 查询项目标签
func (c *VpnClient) ListProjectTagsInvoker(request *model.ListProjectTagsRequest) *ListProjectTagsInvoker {
	requestDef := GenReqDefForListProjectTags()
	return &ListProjectTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListResourcesByTags 查询资源实例列表
//
// 根据标签查询资源实例列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpnClient) ListResourcesByTags(request *model.ListResourcesByTagsRequest) (*model.ListResourcesByTagsResponse, error) {
	requestDef := GenReqDefForListResourcesByTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResourcesByTagsResponse), nil
	}
}

// ListResourcesByTagsInvoker 查询资源实例列表
func (c *VpnClient) ListResourcesByTagsInvoker(request *model.ListResourcesByTagsRequest) *ListResourcesByTagsInvoker {
	requestDef := GenReqDefForListResourcesByTags()
	return &ListResourcesByTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowResourceTags 查询资源标签
//
// 查询指定实例的标签信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpnClient) ShowResourceTags(request *model.ShowResourceTagsRequest) (*model.ShowResourceTagsResponse, error) {
	requestDef := GenReqDefForShowResourceTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowResourceTagsResponse), nil
	}
}

// ShowResourceTagsInvoker 查询资源标签
func (c *VpnClient) ShowResourceTagsInvoker(request *model.ShowResourceTagsRequest) *ShowResourceTagsInvoker {
	requestDef := GenReqDefForShowResourceTags()
	return &ShowResourceTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateVpnAccessPolicy 创建VPN访问策略
//
// 创建VPN访问策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpnClient) CreateVpnAccessPolicy(request *model.CreateVpnAccessPolicyRequest) (*model.CreateVpnAccessPolicyResponse, error) {
	requestDef := GenReqDefForCreateVpnAccessPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateVpnAccessPolicyResponse), nil
	}
}

// CreateVpnAccessPolicyInvoker 创建VPN访问策略
func (c *VpnClient) CreateVpnAccessPolicyInvoker(request *model.CreateVpnAccessPolicyRequest) *CreateVpnAccessPolicyInvoker {
	requestDef := GenReqDefForCreateVpnAccessPolicy()
	return &CreateVpnAccessPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteVpnAccessPolicy 删除VPN访问策略
//
// 删除VPN访问策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpnClient) DeleteVpnAccessPolicy(request *model.DeleteVpnAccessPolicyRequest) (*model.DeleteVpnAccessPolicyResponse, error) {
	requestDef := GenReqDefForDeleteVpnAccessPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteVpnAccessPolicyResponse), nil
	}
}

// DeleteVpnAccessPolicyInvoker 删除VPN访问策略
func (c *VpnClient) DeleteVpnAccessPolicyInvoker(request *model.DeleteVpnAccessPolicyRequest) *DeleteVpnAccessPolicyInvoker {
	requestDef := GenReqDefForDeleteVpnAccessPolicy()
	return &DeleteVpnAccessPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListVpnAccessPolicies 查询VPN访问策略列表
//
// 查询VPN访问策略列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpnClient) ListVpnAccessPolicies(request *model.ListVpnAccessPoliciesRequest) (*model.ListVpnAccessPoliciesResponse, error) {
	requestDef := GenReqDefForListVpnAccessPolicies()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVpnAccessPoliciesResponse), nil
	}
}

// ListVpnAccessPoliciesInvoker 查询VPN访问策略列表
func (c *VpnClient) ListVpnAccessPoliciesInvoker(request *model.ListVpnAccessPoliciesRequest) *ListVpnAccessPoliciesInvoker {
	requestDef := GenReqDefForListVpnAccessPolicies()
	return &ListVpnAccessPoliciesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowVpnAccessPolicy 查询VPN访问策略
//
// 查询VPN访问策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpnClient) ShowVpnAccessPolicy(request *model.ShowVpnAccessPolicyRequest) (*model.ShowVpnAccessPolicyResponse, error) {
	requestDef := GenReqDefForShowVpnAccessPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVpnAccessPolicyResponse), nil
	}
}

// ShowVpnAccessPolicyInvoker 查询VPN访问策略
func (c *VpnClient) ShowVpnAccessPolicyInvoker(request *model.ShowVpnAccessPolicyRequest) *ShowVpnAccessPolicyInvoker {
	requestDef := GenReqDefForShowVpnAccessPolicy()
	return &ShowVpnAccessPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateVpnAccessPolicy 修改VPN访问策略
//
// 修改VPN访问策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpnClient) UpdateVpnAccessPolicy(request *model.UpdateVpnAccessPolicyRequest) (*model.UpdateVpnAccessPolicyResponse, error) {
	requestDef := GenReqDefForUpdateVpnAccessPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateVpnAccessPolicyResponse), nil
	}
}

// UpdateVpnAccessPolicyInvoker 修改VPN访问策略
func (c *VpnClient) UpdateVpnAccessPolicyInvoker(request *model.UpdateVpnAccessPolicyRequest) *UpdateVpnAccessPolicyInvoker {
	requestDef := GenReqDefForUpdateVpnAccessPolicy()
	return &UpdateVpnAccessPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateVpnConnection 创建VPN连接
//
// 创建VPN连接，连接VPN网关与对端网关
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpnClient) CreateVpnConnection(request *model.CreateVpnConnectionRequest) (*model.CreateVpnConnectionResponse, error) {
	requestDef := GenReqDefForCreateVpnConnection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateVpnConnectionResponse), nil
	}
}

// CreateVpnConnectionInvoker 创建VPN连接
func (c *VpnClient) CreateVpnConnectionInvoker(request *model.CreateVpnConnectionRequest) *CreateVpnConnectionInvoker {
	requestDef := GenReqDefForCreateVpnConnection()
	return &CreateVpnConnectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteVpnConnection 删除VPN连接
//
// 根据连接ID，删除指定的VPN连接
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpnClient) DeleteVpnConnection(request *model.DeleteVpnConnectionRequest) (*model.DeleteVpnConnectionResponse, error) {
	requestDef := GenReqDefForDeleteVpnConnection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteVpnConnectionResponse), nil
	}
}

// DeleteVpnConnectionInvoker 删除VPN连接
func (c *VpnClient) DeleteVpnConnectionInvoker(request *model.DeleteVpnConnectionRequest) *DeleteVpnConnectionInvoker {
	requestDef := GenReqDefForDeleteVpnConnection()
	return &DeleteVpnConnectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListVpnConnections 查询VPN连接列表
//
// 查询VPN连接列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpnClient) ListVpnConnections(request *model.ListVpnConnectionsRequest) (*model.ListVpnConnectionsResponse, error) {
	requestDef := GenReqDefForListVpnConnections()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVpnConnectionsResponse), nil
	}
}

// ListVpnConnectionsInvoker 查询VPN连接列表
func (c *VpnClient) ListVpnConnectionsInvoker(request *model.ListVpnConnectionsRequest) *ListVpnConnectionsInvoker {
	requestDef := GenReqDefForListVpnConnections()
	return &ListVpnConnectionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowVpnConnection 查询VPN连接
//
// 根据连接ID，查询指定的VPN连接的参数
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpnClient) ShowVpnConnection(request *model.ShowVpnConnectionRequest) (*model.ShowVpnConnectionResponse, error) {
	requestDef := GenReqDefForShowVpnConnection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVpnConnectionResponse), nil
	}
}

// ShowVpnConnectionInvoker 查询VPN连接
func (c *VpnClient) ShowVpnConnectionInvoker(request *model.ShowVpnConnectionRequest) *ShowVpnConnectionInvoker {
	requestDef := GenReqDefForShowVpnConnection()
	return &ShowVpnConnectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateVpnConnection 更新VPN连接
//
// 根据连接ID，更新指定的VPN连接的参数
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpnClient) UpdateVpnConnection(request *model.UpdateVpnConnectionRequest) (*model.UpdateVpnConnectionResponse, error) {
	requestDef := GenReqDefForUpdateVpnConnection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateVpnConnectionResponse), nil
	}
}

// UpdateVpnConnectionInvoker 更新VPN连接
func (c *VpnClient) UpdateVpnConnectionInvoker(request *model.UpdateVpnConnectionRequest) *UpdateVpnConnectionInvoker {
	requestDef := GenReqDefForUpdateVpnConnection()
	return &UpdateVpnConnectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateVgw 创建VPN网关
//
// 创建一个VPN网关
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpnClient) CreateVgw(request *model.CreateVgwRequest) (*model.CreateVgwResponse, error) {
	requestDef := GenReqDefForCreateVgw()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateVgwResponse), nil
	}
}

// CreateVgwInvoker 创建VPN网关
func (c *VpnClient) CreateVgwInvoker(request *model.CreateVgwRequest) *CreateVgwInvoker {
	requestDef := GenReqDefForCreateVgw()
	return &CreateVgwInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteVgw 删除VPN网关
//
// 根据VPN网关ID，删除指定的VPN网关
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpnClient) DeleteVgw(request *model.DeleteVgwRequest) (*model.DeleteVgwResponse, error) {
	requestDef := GenReqDefForDeleteVgw()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteVgwResponse), nil
	}
}

// DeleteVgwInvoker 删除VPN网关
func (c *VpnClient) DeleteVgwInvoker(request *model.DeleteVgwRequest) *DeleteVgwInvoker {
	requestDef := GenReqDefForDeleteVgw()
	return &DeleteVgwInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAvailabilityZones 查询VPN网关可用区
//
// 查询VPN网关可用区
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpnClient) ListAvailabilityZones(request *model.ListAvailabilityZonesRequest) (*model.ListAvailabilityZonesResponse, error) {
	requestDef := GenReqDefForListAvailabilityZones()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAvailabilityZonesResponse), nil
	}
}

// ListAvailabilityZonesInvoker 查询VPN网关可用区
func (c *VpnClient) ListAvailabilityZonesInvoker(request *model.ListAvailabilityZonesRequest) *ListAvailabilityZonesInvoker {
	requestDef := GenReqDefForListAvailabilityZones()
	return &ListAvailabilityZonesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListVgws 查询VPN网关列表
//
// 查询VPN网关列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpnClient) ListVgws(request *model.ListVgwsRequest) (*model.ListVgwsResponse, error) {
	requestDef := GenReqDefForListVgws()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVgwsResponse), nil
	}
}

// ListVgwsInvoker 查询VPN网关列表
func (c *VpnClient) ListVgwsInvoker(request *model.ListVgwsRequest) *ListVgwsInvoker {
	requestDef := GenReqDefForListVgws()
	return &ListVgwsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowVgw 查询VPN网关
//
// 根据VPN网关ID，查询指定的VPN网关
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpnClient) ShowVgw(request *model.ShowVgwRequest) (*model.ShowVgwResponse, error) {
	requestDef := GenReqDefForShowVgw()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVgwResponse), nil
	}
}

// ShowVgwInvoker 查询VPN网关
func (c *VpnClient) ShowVgwInvoker(request *model.ShowVgwRequest) *ShowVgwInvoker {
	requestDef := GenReqDefForShowVgw()
	return &ShowVgwInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePostpaidVgwSpecification 修改网关规格
//
// 对单个网关规格进行修改，可以升配或降配
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpnClient) UpdatePostpaidVgwSpecification(request *model.UpdatePostpaidVgwSpecificationRequest) (*model.UpdatePostpaidVgwSpecificationResponse, error) {
	requestDef := GenReqDefForUpdatePostpaidVgwSpecification()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePostpaidVgwSpecificationResponse), nil
	}
}

// UpdatePostpaidVgwSpecificationInvoker 修改网关规格
func (c *VpnClient) UpdatePostpaidVgwSpecificationInvoker(request *model.UpdatePostpaidVgwSpecificationRequest) *UpdatePostpaidVgwSpecificationInvoker {
	requestDef := GenReqDefForUpdatePostpaidVgwSpecification()
	return &UpdatePostpaidVgwSpecificationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateVgw 更新VPN网关
//
// 根据VPN网关ID，更新指定的VPN网关
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpnClient) UpdateVgw(request *model.UpdateVgwRequest) (*model.UpdateVgwResponse, error) {
	requestDef := GenReqDefForUpdateVgw()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateVgwResponse), nil
	}
}

// UpdateVgwInvoker 更新VPN网关
func (c *VpnClient) UpdateVgwInvoker(request *model.UpdateVgwRequest) *UpdateVgwInvoker {
	requestDef := GenReqDefForUpdateVgw()
	return &UpdateVgwInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateVgwCertificate 导入VPN网关证书
//
// 导入租户VPN网关所使用的证书，包括签名证书、签名私钥、加密证书、加密私钥和CA证书链。当前只支持国密证书
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpnClient) CreateVgwCertificate(request *model.CreateVgwCertificateRequest) (*model.CreateVgwCertificateResponse, error) {
	requestDef := GenReqDefForCreateVgwCertificate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateVgwCertificateResponse), nil
	}
}

// CreateVgwCertificateInvoker 导入VPN网关证书
func (c *VpnClient) CreateVgwCertificateInvoker(request *model.CreateVgwCertificateRequest) *CreateVgwCertificateInvoker {
	requestDef := GenReqDefForCreateVgwCertificate()
	return &CreateVgwCertificateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowVpnGatewayCertificate 查询VPN网关证书
//
// 根据VPN网关ID，查询所指定的VPN网关证书
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpnClient) ShowVpnGatewayCertificate(request *model.ShowVpnGatewayCertificateRequest) (*model.ShowVpnGatewayCertificateResponse, error) {
	requestDef := GenReqDefForShowVpnGatewayCertificate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVpnGatewayCertificateResponse), nil
	}
}

// ShowVpnGatewayCertificateInvoker 查询VPN网关证书
func (c *VpnClient) ShowVpnGatewayCertificateInvoker(request *model.ShowVpnGatewayCertificateRequest) *ShowVpnGatewayCertificateInvoker {
	requestDef := GenReqDefForShowVpnGatewayCertificate()
	return &ShowVpnGatewayCertificateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateVgwCertificate 更新VPN网关证书
//
// 更新租户VPN网关所使用的证书，包括签名证书、签名私钥、加密证书、加密私钥和CA证书链。当前只支持国密证书
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpnClient) UpdateVgwCertificate(request *model.UpdateVgwCertificateRequest) (*model.UpdateVgwCertificateResponse, error) {
	requestDef := GenReqDefForUpdateVgwCertificate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateVgwCertificateResponse), nil
	}
}

// UpdateVgwCertificateInvoker 更新VPN网关证书
func (c *VpnClient) UpdateVgwCertificateInvoker(request *model.UpdateVgwCertificateRequest) *UpdateVgwCertificateInvoker {
	requestDef := GenReqDefForUpdateVgwCertificate()
	return &UpdateVgwCertificateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowQuotasInfo 查询指定租户配额
//
// 查询指定租户的配额
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpnClient) ShowQuotasInfo(request *model.ShowQuotasInfoRequest) (*model.ShowQuotasInfoResponse, error) {
	requestDef := GenReqDefForShowQuotasInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowQuotasInfoResponse), nil
	}
}

// ShowQuotasInfoInvoker 查询指定租户配额
func (c *VpnClient) ShowQuotasInfoInvoker(request *model.ShowQuotasInfoRequest) *ShowQuotasInfoInvoker {
	requestDef := GenReqDefForShowQuotasInfo()
	return &ShowQuotasInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateVpnServer 创建一个VPN 服务端
//
// 创建一个VPN 服务端
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpnClient) CreateVpnServer(request *model.CreateVpnServerRequest) (*model.CreateVpnServerResponse, error) {
	requestDef := GenReqDefForCreateVpnServer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateVpnServerResponse), nil
	}
}

// CreateVpnServerInvoker 创建一个VPN 服务端
func (c *VpnClient) CreateVpnServerInvoker(request *model.CreateVpnServerRequest) *CreateVpnServerInvoker {
	requestDef := GenReqDefForCreateVpnServer()
	return &CreateVpnServerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExportClientConfig 导出服务端对应的客户端配置信息
//
// 导出客户端配置信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpnClient) ExportClientConfig(request *model.ExportClientConfigRequest) (*model.ExportClientConfigResponse, error) {
	requestDef := GenReqDefForExportClientConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExportClientConfigResponse), nil
	}
}

// ExportClientConfigInvoker 导出服务端对应的客户端配置信息
func (c *VpnClient) ExportClientConfigInvoker(request *model.ExportClientConfigRequest) *ExportClientConfigInvoker {
	requestDef := GenReqDefForExportClientConfig()
	return &ExportClientConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListVpnServersByProject 查询租户下的所有服务端信息
//
// 查询租户下的所有服务端信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpnClient) ListVpnServersByProject(request *model.ListVpnServersByProjectRequest) (*model.ListVpnServersByProjectResponse, error) {
	requestDef := GenReqDefForListVpnServersByProject()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVpnServersByProjectResponse), nil
	}
}

// ListVpnServersByProjectInvoker 查询租户下的所有服务端信息
func (c *VpnClient) ListVpnServersByProjectInvoker(request *model.ListVpnServersByProjectRequest) *ListVpnServersByProjectInvoker {
	requestDef := GenReqDefForListVpnServersByProject()
	return &ListVpnServersByProjectInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListVpnServersByVgw 查询一个网关下的服务端信息
//
// 查询一个网关下的服务端信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpnClient) ListVpnServersByVgw(request *model.ListVpnServersByVgwRequest) (*model.ListVpnServersByVgwResponse, error) {
	requestDef := GenReqDefForListVpnServersByVgw()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVpnServersByVgwResponse), nil
	}
}

// ListVpnServersByVgwInvoker 查询一个网关下的服务端信息
func (c *VpnClient) ListVpnServersByVgwInvoker(request *model.ListVpnServersByVgwRequest) *ListVpnServersByVgwInvoker {
	requestDef := GenReqDefForListVpnServersByVgw()
	return &ListVpnServersByVgwInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateVpnServer 更新指定VPN 服务端
//
// 更新指定VPN 服务端
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpnClient) UpdateVpnServer(request *model.UpdateVpnServerRequest) (*model.UpdateVpnServerResponse, error) {
	requestDef := GenReqDefForUpdateVpnServer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateVpnServerResponse), nil
	}
}

// UpdateVpnServerInvoker 更新指定VPN 服务端
func (c *VpnClient) UpdateVpnServerInvoker(request *model.UpdateVpnServerRequest) *UpdateVpnServerInvoker {
	requestDef := GenReqDefForUpdateVpnServer()
	return &UpdateVpnServerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateVpnUser 创建VPN用户
//
// 创建VPN用户
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpnClient) CreateVpnUser(request *model.CreateVpnUserRequest) (*model.CreateVpnUserResponse, error) {
	requestDef := GenReqDefForCreateVpnUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateVpnUserResponse), nil
	}
}

// CreateVpnUserInvoker 创建VPN用户
func (c *VpnClient) CreateVpnUserInvoker(request *model.CreateVpnUserRequest) *CreateVpnUserInvoker {
	requestDef := GenReqDefForCreateVpnUser()
	return &CreateVpnUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteVpnUser 删除VPN用户
//
// 删除VPN用户
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpnClient) DeleteVpnUser(request *model.DeleteVpnUserRequest) (*model.DeleteVpnUserResponse, error) {
	requestDef := GenReqDefForDeleteVpnUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteVpnUserResponse), nil
	}
}

// DeleteVpnUserInvoker 删除VPN用户
func (c *VpnClient) DeleteVpnUserInvoker(request *model.DeleteVpnUserRequest) *DeleteVpnUserInvoker {
	requestDef := GenReqDefForDeleteVpnUser()
	return &DeleteVpnUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListVpnUsers 查询VPN用户列表
//
// 查询VPN用户列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpnClient) ListVpnUsers(request *model.ListVpnUsersRequest) (*model.ListVpnUsersResponse, error) {
	requestDef := GenReqDefForListVpnUsers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVpnUsersResponse), nil
	}
}

// ListVpnUsersInvoker 查询VPN用户列表
func (c *VpnClient) ListVpnUsersInvoker(request *model.ListVpnUsersRequest) *ListVpnUsersInvoker {
	requestDef := GenReqDefForListVpnUsers()
	return &ListVpnUsersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResetVpnUserPassword 重置VPN用户密码
//
// 重置VPN用户密码
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpnClient) ResetVpnUserPassword(request *model.ResetVpnUserPasswordRequest) (*model.ResetVpnUserPasswordResponse, error) {
	requestDef := GenReqDefForResetVpnUserPassword()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResetVpnUserPasswordResponse), nil
	}
}

// ResetVpnUserPasswordInvoker 重置VPN用户密码
func (c *VpnClient) ResetVpnUserPasswordInvoker(request *model.ResetVpnUserPasswordRequest) *ResetVpnUserPasswordInvoker {
	requestDef := GenReqDefForResetVpnUserPassword()
	return &ResetVpnUserPasswordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowVpnUser 查询VPN用户
//
// 查询VPN用户
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpnClient) ShowVpnUser(request *model.ShowVpnUserRequest) (*model.ShowVpnUserResponse, error) {
	requestDef := GenReqDefForShowVpnUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVpnUserResponse), nil
	}
}

// ShowVpnUserInvoker 查询VPN用户
func (c *VpnClient) ShowVpnUserInvoker(request *model.ShowVpnUserRequest) *ShowVpnUserInvoker {
	requestDef := GenReqDefForShowVpnUser()
	return &ShowVpnUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateVpnUser 修改VPN用户
//
// 修改VPN用户
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpnClient) UpdateVpnUser(request *model.UpdateVpnUserRequest) (*model.UpdateVpnUserResponse, error) {
	requestDef := GenReqDefForUpdateVpnUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateVpnUserResponse), nil
	}
}

// UpdateVpnUserInvoker 修改VPN用户
func (c *VpnClient) UpdateVpnUserInvoker(request *model.UpdateVpnUserRequest) *UpdateVpnUserInvoker {
	requestDef := GenReqDefForUpdateVpnUser()
	return &UpdateVpnUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateVpnUserPassword 修改VPN用户密码
//
// 修改VPN用户密码
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpnClient) UpdateVpnUserPassword(request *model.UpdateVpnUserPasswordRequest) (*model.UpdateVpnUserPasswordResponse, error) {
	requestDef := GenReqDefForUpdateVpnUserPassword()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateVpnUserPasswordResponse), nil
	}
}

// UpdateVpnUserPasswordInvoker 修改VPN用户密码
func (c *VpnClient) UpdateVpnUserPasswordInvoker(request *model.UpdateVpnUserPasswordRequest) *UpdateVpnUserPasswordInvoker {
	requestDef := GenReqDefForUpdateVpnUserPassword()
	return &UpdateVpnUserPasswordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddVpnUsersToGroup 添加VPN用户到组
//
// 添加VPN用户到组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpnClient) AddVpnUsersToGroup(request *model.AddVpnUsersToGroupRequest) (*model.AddVpnUsersToGroupResponse, error) {
	requestDef := GenReqDefForAddVpnUsersToGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddVpnUsersToGroupResponse), nil
	}
}

// AddVpnUsersToGroupInvoker 添加VPN用户到组
func (c *VpnClient) AddVpnUsersToGroupInvoker(request *model.AddVpnUsersToGroupRequest) *AddVpnUsersToGroupInvoker {
	requestDef := GenReqDefForAddVpnUsersToGroup()
	return &AddVpnUsersToGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateVpnUserGroup 创建VPN用户组
//
// 创建VPN用户组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpnClient) CreateVpnUserGroup(request *model.CreateVpnUserGroupRequest) (*model.CreateVpnUserGroupResponse, error) {
	requestDef := GenReqDefForCreateVpnUserGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateVpnUserGroupResponse), nil
	}
}

// CreateVpnUserGroupInvoker 创建VPN用户组
func (c *VpnClient) CreateVpnUserGroupInvoker(request *model.CreateVpnUserGroupRequest) *CreateVpnUserGroupInvoker {
	requestDef := GenReqDefForCreateVpnUserGroup()
	return &CreateVpnUserGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteVpnUserGroup 删除VPN用户组
//
// 删除VPN用户组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpnClient) DeleteVpnUserGroup(request *model.DeleteVpnUserGroupRequest) (*model.DeleteVpnUserGroupResponse, error) {
	requestDef := GenReqDefForDeleteVpnUserGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteVpnUserGroupResponse), nil
	}
}

// DeleteVpnUserGroupInvoker 删除VPN用户组
func (c *VpnClient) DeleteVpnUserGroupInvoker(request *model.DeleteVpnUserGroupRequest) *DeleteVpnUserGroupInvoker {
	requestDef := GenReqDefForDeleteVpnUserGroup()
	return &DeleteVpnUserGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListVpnUserGroups 查询VPN用户组列表
//
// 查询VPN用户组列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpnClient) ListVpnUserGroups(request *model.ListVpnUserGroupsRequest) (*model.ListVpnUserGroupsResponse, error) {
	requestDef := GenReqDefForListVpnUserGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVpnUserGroupsResponse), nil
	}
}

// ListVpnUserGroupsInvoker 查询VPN用户组列表
func (c *VpnClient) ListVpnUserGroupsInvoker(request *model.ListVpnUserGroupsRequest) *ListVpnUserGroupsInvoker {
	requestDef := GenReqDefForListVpnUserGroups()
	return &ListVpnUserGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListVpnUsersInGroup 查询组内VPN用户
//
// 查询组内VPN用户
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpnClient) ListVpnUsersInGroup(request *model.ListVpnUsersInGroupRequest) (*model.ListVpnUsersInGroupResponse, error) {
	requestDef := GenReqDefForListVpnUsersInGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVpnUsersInGroupResponse), nil
	}
}

// ListVpnUsersInGroupInvoker 查询组内VPN用户
func (c *VpnClient) ListVpnUsersInGroupInvoker(request *model.ListVpnUsersInGroupRequest) *ListVpnUsersInGroupInvoker {
	requestDef := GenReqDefForListVpnUsersInGroup()
	return &ListVpnUsersInGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RemoveVpnUsersFromGroup 删除组内VPN用户
//
// 删除组内VPN用户
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpnClient) RemoveVpnUsersFromGroup(request *model.RemoveVpnUsersFromGroupRequest) (*model.RemoveVpnUsersFromGroupResponse, error) {
	requestDef := GenReqDefForRemoveVpnUsersFromGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RemoveVpnUsersFromGroupResponse), nil
	}
}

// RemoveVpnUsersFromGroupInvoker 删除组内VPN用户
func (c *VpnClient) RemoveVpnUsersFromGroupInvoker(request *model.RemoveVpnUsersFromGroupRequest) *RemoveVpnUsersFromGroupInvoker {
	requestDef := GenReqDefForRemoveVpnUsersFromGroup()
	return &RemoveVpnUsersFromGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowVpnUserGroup 查询VPN用户组
//
// 查询VPN用户组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpnClient) ShowVpnUserGroup(request *model.ShowVpnUserGroupRequest) (*model.ShowVpnUserGroupResponse, error) {
	requestDef := GenReqDefForShowVpnUserGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVpnUserGroupResponse), nil
	}
}

// ShowVpnUserGroupInvoker 查询VPN用户组
func (c *VpnClient) ShowVpnUserGroupInvoker(request *model.ShowVpnUserGroupRequest) *ShowVpnUserGroupInvoker {
	requestDef := GenReqDefForShowVpnUserGroup()
	return &ShowVpnUserGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateVpnUserGroup 修改VPN用户组
//
// 修改VPN用户组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpnClient) UpdateVpnUserGroup(request *model.UpdateVpnUserGroupRequest) (*model.UpdateVpnUserGroupResponse, error) {
	requestDef := GenReqDefForUpdateVpnUserGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateVpnUserGroupResponse), nil
	}
}

// UpdateVpnUserGroupInvoker 修改VPN用户组
func (c *VpnClient) UpdateVpnUserGroupInvoker(request *model.UpdateVpnUserGroupRequest) *UpdateVpnUserGroupInvoker {
	requestDef := GenReqDefForUpdateVpnUserGroup()
	return &UpdateVpnUserGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
