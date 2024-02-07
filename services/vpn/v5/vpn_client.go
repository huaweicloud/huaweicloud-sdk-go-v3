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
