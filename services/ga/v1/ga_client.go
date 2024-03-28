package v1

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ga/v1/model"
)

type GaClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewGaClient(hcClient *httpclient.HcHttpClient) *GaClient {
	return &GaClient{HcClient: hcClient}
}

func GaClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// CreateAccelerator 创建全球加速器
//
// 创建全球加速器。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaClient) CreateAccelerator(request *model.CreateAcceleratorRequest) (*model.CreateAcceleratorResponse, error) {
	requestDef := GenReqDefForCreateAccelerator()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAcceleratorResponse), nil
	}
}

// CreateAcceleratorInvoker 创建全球加速器
func (c *GaClient) CreateAcceleratorInvoker(request *model.CreateAcceleratorRequest) *CreateAcceleratorInvoker {
	requestDef := GenReqDefForCreateAccelerator()
	return &CreateAcceleratorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAccelerator 删除全球加速器
//
// 删除全球加速器。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaClient) DeleteAccelerator(request *model.DeleteAcceleratorRequest) (*model.DeleteAcceleratorResponse, error) {
	requestDef := GenReqDefForDeleteAccelerator()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAcceleratorResponse), nil
	}
}

// DeleteAcceleratorInvoker 删除全球加速器
func (c *GaClient) DeleteAcceleratorInvoker(request *model.DeleteAcceleratorRequest) *DeleteAcceleratorInvoker {
	requestDef := GenReqDefForDeleteAccelerator()
	return &DeleteAcceleratorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAccelerators 查询全球加速器列表
//
// 查询全球加速器列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaClient) ListAccelerators(request *model.ListAcceleratorsRequest) (*model.ListAcceleratorsResponse, error) {
	requestDef := GenReqDefForListAccelerators()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAcceleratorsResponse), nil
	}
}

// ListAcceleratorsInvoker 查询全球加速器列表
func (c *GaClient) ListAcceleratorsInvoker(request *model.ListAcceleratorsRequest) *ListAcceleratorsInvoker {
	requestDef := GenReqDefForListAccelerators()
	return &ListAcceleratorsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAccelerator 查询全球加速器详情
//
// 查询全球加速器详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaClient) ShowAccelerator(request *model.ShowAcceleratorRequest) (*model.ShowAcceleratorResponse, error) {
	requestDef := GenReqDefForShowAccelerator()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAcceleratorResponse), nil
	}
}

// ShowAcceleratorInvoker 查询全球加速器详情
func (c *GaClient) ShowAcceleratorInvoker(request *model.ShowAcceleratorRequest) *ShowAcceleratorInvoker {
	requestDef := GenReqDefForShowAccelerator()
	return &ShowAcceleratorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAccelerator 更新全球加速器
//
// 更新全球加速器。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaClient) UpdateAccelerator(request *model.UpdateAcceleratorRequest) (*model.UpdateAcceleratorResponse, error) {
	requestDef := GenReqDefForUpdateAccelerator()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAcceleratorResponse), nil
	}
}

// UpdateAcceleratorInvoker 更新全球加速器
func (c *GaClient) UpdateAcceleratorInvoker(request *model.UpdateAcceleratorRequest) *UpdateAcceleratorInvoker {
	requestDef := GenReqDefForUpdateAccelerator()
	return &UpdateAcceleratorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateEndpoint 创建终端节点
//
// 创建终端节点。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaClient) CreateEndpoint(request *model.CreateEndpointRequest) (*model.CreateEndpointResponse, error) {
	requestDef := GenReqDefForCreateEndpoint()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateEndpointResponse), nil
	}
}

// CreateEndpointInvoker 创建终端节点
func (c *GaClient) CreateEndpointInvoker(request *model.CreateEndpointRequest) *CreateEndpointInvoker {
	requestDef := GenReqDefForCreateEndpoint()
	return &CreateEndpointInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteEndpoint 删除终端节点
//
// 删除终端节点。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaClient) DeleteEndpoint(request *model.DeleteEndpointRequest) (*model.DeleteEndpointResponse, error) {
	requestDef := GenReqDefForDeleteEndpoint()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteEndpointResponse), nil
	}
}

// DeleteEndpointInvoker 删除终端节点
func (c *GaClient) DeleteEndpointInvoker(request *model.DeleteEndpointRequest) *DeleteEndpointInvoker {
	requestDef := GenReqDefForDeleteEndpoint()
	return &DeleteEndpointInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEndpoints 查询终端节点组下终端节点列表
//
// 查询终端节点组下终端节点列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaClient) ListEndpoints(request *model.ListEndpointsRequest) (*model.ListEndpointsResponse, error) {
	requestDef := GenReqDefForListEndpoints()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEndpointsResponse), nil
	}
}

// ListEndpointsInvoker 查询终端节点组下终端节点列表
func (c *GaClient) ListEndpointsInvoker(request *model.ListEndpointsRequest) *ListEndpointsInvoker {
	requestDef := GenReqDefForListEndpoints()
	return &ListEndpointsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowEndpoint 查询终端节点详情
//
// 查询终端节点详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaClient) ShowEndpoint(request *model.ShowEndpointRequest) (*model.ShowEndpointResponse, error) {
	requestDef := GenReqDefForShowEndpoint()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEndpointResponse), nil
	}
}

// ShowEndpointInvoker 查询终端节点详情
func (c *GaClient) ShowEndpointInvoker(request *model.ShowEndpointRequest) *ShowEndpointInvoker {
	requestDef := GenReqDefForShowEndpoint()
	return &ShowEndpointInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateEndpoint 更新终端节点
//
// 更新终端节点。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaClient) UpdateEndpoint(request *model.UpdateEndpointRequest) (*model.UpdateEndpointResponse, error) {
	requestDef := GenReqDefForUpdateEndpoint()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateEndpointResponse), nil
	}
}

// UpdateEndpointInvoker 更新终端节点
func (c *GaClient) UpdateEndpointInvoker(request *model.UpdateEndpointRequest) *UpdateEndpointInvoker {
	requestDef := GenReqDefForUpdateEndpoint()
	return &UpdateEndpointInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateEndpointGroup 创建终端节点组
//
// 创建终端节点组。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaClient) CreateEndpointGroup(request *model.CreateEndpointGroupRequest) (*model.CreateEndpointGroupResponse, error) {
	requestDef := GenReqDefForCreateEndpointGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateEndpointGroupResponse), nil
	}
}

// CreateEndpointGroupInvoker 创建终端节点组
func (c *GaClient) CreateEndpointGroupInvoker(request *model.CreateEndpointGroupRequest) *CreateEndpointGroupInvoker {
	requestDef := GenReqDefForCreateEndpointGroup()
	return &CreateEndpointGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteEndpointGroup 删除终端节点组
//
// 删除终端节点组。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaClient) DeleteEndpointGroup(request *model.DeleteEndpointGroupRequest) (*model.DeleteEndpointGroupResponse, error) {
	requestDef := GenReqDefForDeleteEndpointGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteEndpointGroupResponse), nil
	}
}

// DeleteEndpointGroupInvoker 删除终端节点组
func (c *GaClient) DeleteEndpointGroupInvoker(request *model.DeleteEndpointGroupRequest) *DeleteEndpointGroupInvoker {
	requestDef := GenReqDefForDeleteEndpointGroup()
	return &DeleteEndpointGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEndpointGroups 查询终端节点组列表
//
// 查询终端节点组列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaClient) ListEndpointGroups(request *model.ListEndpointGroupsRequest) (*model.ListEndpointGroupsResponse, error) {
	requestDef := GenReqDefForListEndpointGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEndpointGroupsResponse), nil
	}
}

// ListEndpointGroupsInvoker 查询终端节点组列表
func (c *GaClient) ListEndpointGroupsInvoker(request *model.ListEndpointGroupsRequest) *ListEndpointGroupsInvoker {
	requestDef := GenReqDefForListEndpointGroups()
	return &ListEndpointGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowEndpointGroup 查询终端节点组详情
//
// 查询终端节点组详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaClient) ShowEndpointGroup(request *model.ShowEndpointGroupRequest) (*model.ShowEndpointGroupResponse, error) {
	requestDef := GenReqDefForShowEndpointGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEndpointGroupResponse), nil
	}
}

// ShowEndpointGroupInvoker 查询终端节点组详情
func (c *GaClient) ShowEndpointGroupInvoker(request *model.ShowEndpointGroupRequest) *ShowEndpointGroupInvoker {
	requestDef := GenReqDefForShowEndpointGroup()
	return &ShowEndpointGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateEndpointGroup 更新终端节点组
//
// 更新终端节点组。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaClient) UpdateEndpointGroup(request *model.UpdateEndpointGroupRequest) (*model.UpdateEndpointGroupResponse, error) {
	requestDef := GenReqDefForUpdateEndpointGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateEndpointGroupResponse), nil
	}
}

// UpdateEndpointGroupInvoker 更新终端节点组
func (c *GaClient) UpdateEndpointGroupInvoker(request *model.UpdateEndpointGroupRequest) *UpdateEndpointGroupInvoker {
	requestDef := GenReqDefForUpdateEndpointGroup()
	return &UpdateEndpointGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateHealthCheck 创建健康检查
//
// 创建健康检查。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaClient) CreateHealthCheck(request *model.CreateHealthCheckRequest) (*model.CreateHealthCheckResponse, error) {
	requestDef := GenReqDefForCreateHealthCheck()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateHealthCheckResponse), nil
	}
}

// CreateHealthCheckInvoker 创建健康检查
func (c *GaClient) CreateHealthCheckInvoker(request *model.CreateHealthCheckRequest) *CreateHealthCheckInvoker {
	requestDef := GenReqDefForCreateHealthCheck()
	return &CreateHealthCheckInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteHealthCheck 删除健康检查
//
// 删除健康检查。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaClient) DeleteHealthCheck(request *model.DeleteHealthCheckRequest) (*model.DeleteHealthCheckResponse, error) {
	requestDef := GenReqDefForDeleteHealthCheck()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteHealthCheckResponse), nil
	}
}

// DeleteHealthCheckInvoker 删除健康检查
func (c *GaClient) DeleteHealthCheckInvoker(request *model.DeleteHealthCheckRequest) *DeleteHealthCheckInvoker {
	requestDef := GenReqDefForDeleteHealthCheck()
	return &DeleteHealthCheckInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListHealthChecks 查询健康检查列表
//
// 查询健康检查列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaClient) ListHealthChecks(request *model.ListHealthChecksRequest) (*model.ListHealthChecksResponse, error) {
	requestDef := GenReqDefForListHealthChecks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHealthChecksResponse), nil
	}
}

// ListHealthChecksInvoker 查询健康检查列表
func (c *GaClient) ListHealthChecksInvoker(request *model.ListHealthChecksRequest) *ListHealthChecksInvoker {
	requestDef := GenReqDefForListHealthChecks()
	return &ListHealthChecksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowHealthCheck 查询健康检查详情
//
// 查询健康检查详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaClient) ShowHealthCheck(request *model.ShowHealthCheckRequest) (*model.ShowHealthCheckResponse, error) {
	requestDef := GenReqDefForShowHealthCheck()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowHealthCheckResponse), nil
	}
}

// ShowHealthCheckInvoker 查询健康检查详情
func (c *GaClient) ShowHealthCheckInvoker(request *model.ShowHealthCheckRequest) *ShowHealthCheckInvoker {
	requestDef := GenReqDefForShowHealthCheck()
	return &ShowHealthCheckInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateHealthCheck 更新健康检查
//
// 更新健康检查。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaClient) UpdateHealthCheck(request *model.UpdateHealthCheckRequest) (*model.UpdateHealthCheckResponse, error) {
	requestDef := GenReqDefForUpdateHealthCheck()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateHealthCheckResponse), nil
	}
}

// UpdateHealthCheckInvoker 更新健康检查
func (c *GaClient) UpdateHealthCheckInvoker(request *model.UpdateHealthCheckRequest) *UpdateHealthCheckInvoker {
	requestDef := GenReqDefForUpdateHealthCheck()
	return &UpdateHealthCheckInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddIpGroupIp 添加IP地址组中的IP网段
//
// 添加IP地址组中的IP网段。
// 该接口属于异步接口，接口返回后，后台的添加任务仍在执行；可以使用查询IP地址组详情接口查询状态，当IP地址组状态为ACTIVE时，表示条目添加完成。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaClient) AddIpGroupIp(request *model.AddIpGroupIpRequest) (*model.AddIpGroupIpResponse, error) {
	requestDef := GenReqDefForAddIpGroupIp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddIpGroupIpResponse), nil
	}
}

// AddIpGroupIpInvoker 添加IP地址组中的IP网段
func (c *GaClient) AddIpGroupIpInvoker(request *model.AddIpGroupIpRequest) *AddIpGroupIpInvoker {
	requestDef := GenReqDefForAddIpGroupIp()
	return &AddIpGroupIpInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AssociateListener 绑定IP地址组与监听器
//
// 绑定IP地址组与监听器。
// 该接口属于异步接口，接口返回后，后台的绑定任务仍在执行；可以使用查询IP地址组详情接口查询状态，当IP地址组状态为ACTIVE时，表示绑定完成。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaClient) AssociateListener(request *model.AssociateListenerRequest) (*model.AssociateListenerResponse, error) {
	requestDef := GenReqDefForAssociateListener()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AssociateListenerResponse), nil
	}
}

// AssociateListenerInvoker 绑定IP地址组与监听器
func (c *GaClient) AssociateListenerInvoker(request *model.AssociateListenerRequest) *AssociateListenerInvoker {
	requestDef := GenReqDefForAssociateListener()
	return &AssociateListenerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateIpGroup 创建IP地址组
//
// 创建IP地址组。
// 该接口属于异步接口，会先返回一个IP地址组ID，但后台的创建任务仍在执行；可以使用查询IP地址组详情接口查询状态，当IP地址组状态为ACTIVE时，表示IP地址组创建完成。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaClient) CreateIpGroup(request *model.CreateIpGroupRequest) (*model.CreateIpGroupResponse, error) {
	requestDef := GenReqDefForCreateIpGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateIpGroupResponse), nil
	}
}

// CreateIpGroupInvoker 创建IP地址组
func (c *GaClient) CreateIpGroupInvoker(request *model.CreateIpGroupRequest) *CreateIpGroupInvoker {
	requestDef := GenReqDefForCreateIpGroup()
	return &CreateIpGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteIpGroup 删除IP地址组
//
// 删除IP地址组。
// 该接口属于异步接口，接口返回后，后台的删除任务仍在执行；可以使用查询IP地址组详情接口查询状态，当查询不到该IP地址组时，表示删除完成；删除IP地址组时，若IP地址组已经绑定了监听器，则需要先解绑IP地址组与监听器，再进行删除。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaClient) DeleteIpGroup(request *model.DeleteIpGroupRequest) (*model.DeleteIpGroupResponse, error) {
	requestDef := GenReqDefForDeleteIpGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteIpGroupResponse), nil
	}
}

// DeleteIpGroupInvoker 删除IP地址组
func (c *GaClient) DeleteIpGroupInvoker(request *model.DeleteIpGroupRequest) *DeleteIpGroupInvoker {
	requestDef := GenReqDefForDeleteIpGroup()
	return &DeleteIpGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DisassociateListener 解绑IP地址组与监听器
//
// 解绑IP地址组与监听器。
// 该接口属于异步接口，接口返回后，后台的解绑任务仍在执行；可以使用查询IP地址组详情接口查询状态，当IP地址组状态为ACTIVE时，表示解绑完成。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaClient) DisassociateListener(request *model.DisassociateListenerRequest) (*model.DisassociateListenerResponse, error) {
	requestDef := GenReqDefForDisassociateListener()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisassociateListenerResponse), nil
	}
}

// DisassociateListenerInvoker 解绑IP地址组与监听器
func (c *GaClient) DisassociateListenerInvoker(request *model.DisassociateListenerRequest) *DisassociateListenerInvoker {
	requestDef := GenReqDefForDisassociateListener()
	return &DisassociateListenerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListIpGroups 查询IP地址组列表
//
// 查询IP地址组列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaClient) ListIpGroups(request *model.ListIpGroupsRequest) (*model.ListIpGroupsResponse, error) {
	requestDef := GenReqDefForListIpGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListIpGroupsResponse), nil
	}
}

// ListIpGroupsInvoker 查询IP地址组列表
func (c *GaClient) ListIpGroupsInvoker(request *model.ListIpGroupsRequest) *ListIpGroupsInvoker {
	requestDef := GenReqDefForListIpGroups()
	return &ListIpGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RemoveIpGroupIp 删除IP地址组中的IP网段
//
// 删除IP地址组中的IP网段。
// 该接口属于异步接口，接口返回后，后台的删除任务仍在执行；可以使用查询IP地址组详情接口查询状态，当IP地址组状态为ACTIVE时，表示条目删除完成。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaClient) RemoveIpGroupIp(request *model.RemoveIpGroupIpRequest) (*model.RemoveIpGroupIpResponse, error) {
	requestDef := GenReqDefForRemoveIpGroupIp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RemoveIpGroupIpResponse), nil
	}
}

// RemoveIpGroupIpInvoker 删除IP地址组中的IP网段
func (c *GaClient) RemoveIpGroupIpInvoker(request *model.RemoveIpGroupIpRequest) *RemoveIpGroupIpInvoker {
	requestDef := GenReqDefForRemoveIpGroupIp()
	return &RemoveIpGroupIpInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowIpGroup 查询IP地址组详情
//
// 查询IP地址组详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaClient) ShowIpGroup(request *model.ShowIpGroupRequest) (*model.ShowIpGroupResponse, error) {
	requestDef := GenReqDefForShowIpGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowIpGroupResponse), nil
	}
}

// ShowIpGroupInvoker 查询IP地址组详情
func (c *GaClient) ShowIpGroupInvoker(request *model.ShowIpGroupRequest) *ShowIpGroupInvoker {
	requestDef := GenReqDefForShowIpGroup()
	return &ShowIpGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateIpGroup 更新IP地址组
//
// 更新IP地址组。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaClient) UpdateIpGroup(request *model.UpdateIpGroupRequest) (*model.UpdateIpGroupResponse, error) {
	requestDef := GenReqDefForUpdateIpGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateIpGroupResponse), nil
	}
}

// UpdateIpGroupInvoker 更新IP地址组
func (c *GaClient) UpdateIpGroupInvoker(request *model.UpdateIpGroupRequest) *UpdateIpGroupInvoker {
	requestDef := GenReqDefForUpdateIpGroup()
	return &UpdateIpGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateListener 创建监听器
//
// 创建监听器。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaClient) CreateListener(request *model.CreateListenerRequest) (*model.CreateListenerResponse, error) {
	requestDef := GenReqDefForCreateListener()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateListenerResponse), nil
	}
}

// CreateListenerInvoker 创建监听器
func (c *GaClient) CreateListenerInvoker(request *model.CreateListenerRequest) *CreateListenerInvoker {
	requestDef := GenReqDefForCreateListener()
	return &CreateListenerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteListener 删除监听器
//
// 删除监听器。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaClient) DeleteListener(request *model.DeleteListenerRequest) (*model.DeleteListenerResponse, error) {
	requestDef := GenReqDefForDeleteListener()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteListenerResponse), nil
	}
}

// DeleteListenerInvoker 删除监听器
func (c *GaClient) DeleteListenerInvoker(request *model.DeleteListenerRequest) *DeleteListenerInvoker {
	requestDef := GenReqDefForDeleteListener()
	return &DeleteListenerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListListeners 查询监听器列表
//
// 查询监听器列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaClient) ListListeners(request *model.ListListenersRequest) (*model.ListListenersResponse, error) {
	requestDef := GenReqDefForListListeners()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListListenersResponse), nil
	}
}

// ListListenersInvoker 查询监听器列表
func (c *GaClient) ListListenersInvoker(request *model.ListListenersRequest) *ListListenersInvoker {
	requestDef := GenReqDefForListListeners()
	return &ListListenersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowListener 查询监听器详情
//
// 查询监听器详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaClient) ShowListener(request *model.ShowListenerRequest) (*model.ShowListenerResponse, error) {
	requestDef := GenReqDefForShowListener()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowListenerResponse), nil
	}
}

// ShowListenerInvoker 查询监听器详情
func (c *GaClient) ShowListenerInvoker(request *model.ShowListenerRequest) *ShowListenerInvoker {
	requestDef := GenReqDefForShowListener()
	return &ShowListenerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateListener 更新监听器
//
// 更新监听器。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaClient) UpdateListener(request *model.UpdateListenerRequest) (*model.UpdateListenerResponse, error) {
	requestDef := GenReqDefForUpdateListener()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateListenerResponse), nil
	}
}

// UpdateListenerInvoker 更新监听器
func (c *GaClient) UpdateListenerInvoker(request *model.UpdateListenerRequest) *UpdateListenerInvoker {
	requestDef := GenReqDefForUpdateListener()
	return &UpdateListenerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRegions 查询区域列表
//
// 查询区域列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaClient) ListRegions(request *model.ListRegionsRequest) (*model.ListRegionsResponse, error) {
	requestDef := GenReqDefForListRegions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRegionsResponse), nil
	}
}

// ListRegionsInvoker 查询区域列表
func (c *GaClient) ListRegionsInvoker(request *model.ListRegionsRequest) *ListRegionsInvoker {
	requestDef := GenReqDefForListRegions()
	return &ListRegionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CountResourcesByTag 通过标签查询资源实例数量
//
// 通过标签查询资源实例数量。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaClient) CountResourcesByTag(request *model.CountResourcesByTagRequest) (*model.CountResourcesByTagResponse, error) {
	requestDef := GenReqDefForCountResourcesByTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CountResourcesByTagResponse), nil
	}
}

// CountResourcesByTagInvoker 通过标签查询资源实例数量
func (c *GaClient) CountResourcesByTagInvoker(request *model.CountResourcesByTagRequest) *CountResourcesByTagInvoker {
	requestDef := GenReqDefForCountResourcesByTag()
	return &CountResourcesByTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateTags 创建资源标签
//
// 创建资源标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaClient) CreateTags(request *model.CreateTagsRequest) (*model.CreateTagsResponse, error) {
	requestDef := GenReqDefForCreateTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTagsResponse), nil
	}
}

// CreateTagsInvoker 创建资源标签
func (c *GaClient) CreateTagsInvoker(request *model.CreateTagsRequest) *CreateTagsInvoker {
	requestDef := GenReqDefForCreateTags()
	return &CreateTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteTags 删除资源标签
//
// 删除资源标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaClient) DeleteTags(request *model.DeleteTagsRequest) (*model.DeleteTagsResponse, error) {
	requestDef := GenReqDefForDeleteTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTagsResponse), nil
	}
}

// DeleteTagsInvoker 删除资源标签
func (c *GaClient) DeleteTagsInvoker(request *model.DeleteTagsRequest) *DeleteTagsInvoker {
	requestDef := GenReqDefForDeleteTags()
	return &DeleteTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListResourcesByTag 通过标签查询资源实例列表
//
// 通过标签查询资源实例列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaClient) ListResourcesByTag(request *model.ListResourcesByTagRequest) (*model.ListResourcesByTagResponse, error) {
	requestDef := GenReqDefForListResourcesByTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResourcesByTagResponse), nil
	}
}

// ListResourcesByTagInvoker 通过标签查询资源实例列表
func (c *GaClient) ListResourcesByTagInvoker(request *model.ListResourcesByTagRequest) *ListResourcesByTagInvoker {
	requestDef := GenReqDefForListResourcesByTag()
	return &ListResourcesByTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTags 查询标签列表
//
// 查询标签列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaClient) ListTags(request *model.ListTagsRequest) (*model.ListTagsResponse, error) {
	requestDef := GenReqDefForListTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTagsResponse), nil
	}
}

// ListTagsInvoker 查询标签列表
func (c *GaClient) ListTagsInvoker(request *model.ListTagsRequest) *ListTagsInvoker {
	requestDef := GenReqDefForListTags()
	return &ListTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowResourceTags 查询特定资源标签
//
// 查询特定资源标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaClient) ShowResourceTags(request *model.ShowResourceTagsRequest) (*model.ShowResourceTagsResponse, error) {
	requestDef := GenReqDefForShowResourceTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowResourceTagsResponse), nil
	}
}

// ShowResourceTagsInvoker 查询特定资源标签
func (c *GaClient) ShowResourceTagsInvoker(request *model.ShowResourceTagsRequest) *ShowResourceTagsInvoker {
	requestDef := GenReqDefForShowResourceTags()
	return &ShowResourceTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
