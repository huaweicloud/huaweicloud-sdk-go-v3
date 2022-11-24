package v3

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dc/v3/model"
)

type DcClient struct {
	HcClient *http_client.HcHttpClient
}

func NewDcClient(hcClient *http_client.HcHttpClient) *DcClient {
	return &DcClient{HcClient: hcClient}
}

func DcClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// CreateHostedDirectConnect 创建托管专线连接
//
// 用于合作伙伴用户最终租户创建托管专线
// 创建者必须拥有合作伙伴资质，并且已经构建好运营(hosting)专线
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcClient) CreateHostedDirectConnect(request *model.CreateHostedDirectConnectRequest) (*model.CreateHostedDirectConnectResponse, error) {
	requestDef := GenReqDefForCreateHostedDirectConnect()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateHostedDirectConnectResponse), nil
	}
}

// CreateHostedDirectConnectInvoker 创建托管专线连接
func (c *DcClient) CreateHostedDirectConnectInvoker(request *model.CreateHostedDirectConnectRequest) *CreateHostedDirectConnectInvoker {
	requestDef := GenReqDefForCreateHostedDirectConnect()
	return &CreateHostedDirectConnectInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDirectConnect 删除物理连接
//
// 删除物理连接，本接口只适用于按需计费物理专线，对于包周期购买的专线通过订单退订的方式删除物理连接
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcClient) DeleteDirectConnect(request *model.DeleteDirectConnectRequest) (*model.DeleteDirectConnectResponse, error) {
	requestDef := GenReqDefForDeleteDirectConnect()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDirectConnectResponse), nil
	}
}

// DeleteDirectConnectInvoker 删除物理连接
func (c *DcClient) DeleteDirectConnectInvoker(request *model.DeleteDirectConnectRequest) *DeleteDirectConnectInvoker {
	requestDef := GenReqDefForDeleteDirectConnect()
	return &DeleteDirectConnectInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteHostedDirectConnect 删除托管专线连接
//
// 合作伙伴删除托管专线
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcClient) DeleteHostedDirectConnect(request *model.DeleteHostedDirectConnectRequest) (*model.DeleteHostedDirectConnectResponse, error) {
	requestDef := GenReqDefForDeleteHostedDirectConnect()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteHostedDirectConnectResponse), nil
	}
}

// DeleteHostedDirectConnectInvoker 删除托管专线连接
func (c *DcClient) DeleteHostedDirectConnectInvoker(request *model.DeleteHostedDirectConnectRequest) *DeleteHostedDirectConnectInvoker {
	requestDef := GenReqDefForDeleteHostedDirectConnect()
	return &DeleteHostedDirectConnectInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDirectConnects 查询物理连接列表
//
// 查询租户创建的所有的direct connect对象.
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcClient) ListDirectConnects(request *model.ListDirectConnectsRequest) (*model.ListDirectConnectsResponse, error) {
	requestDef := GenReqDefForListDirectConnects()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDirectConnectsResponse), nil
	}
}

// ListDirectConnectsInvoker 查询物理连接列表
func (c *DcClient) ListDirectConnectsInvoker(request *model.ListDirectConnectsRequest) *ListDirectConnectsInvoker {
	requestDef := GenReqDefForListDirectConnects()
	return &ListDirectConnectsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListHostedDirectConnects 查询租户的托管专线列表
//
// 查询合作伙伴创建的托管专线连接列表.
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcClient) ListHostedDirectConnects(request *model.ListHostedDirectConnectsRequest) (*model.ListHostedDirectConnectsResponse, error) {
	requestDef := GenReqDefForListHostedDirectConnects()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHostedDirectConnectsResponse), nil
	}
}

// ListHostedDirectConnectsInvoker 查询租户的托管专线列表
func (c *DcClient) ListHostedDirectConnectsInvoker(request *model.ListHostedDirectConnectsRequest) *ListHostedDirectConnectsInvoker {
	requestDef := GenReqDefForListHostedDirectConnects()
	return &ListHostedDirectConnectsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDirectConnect 查询物理连接详细信息
//
// 查询物理连接详细信息.
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcClient) ShowDirectConnect(request *model.ShowDirectConnectRequest) (*model.ShowDirectConnectResponse, error) {
	requestDef := GenReqDefForShowDirectConnect()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDirectConnectResponse), nil
	}
}

// ShowDirectConnectInvoker 查询物理连接详细信息
func (c *DcClient) ShowDirectConnectInvoker(request *model.ShowDirectConnectRequest) *ShowDirectConnectInvoker {
	requestDef := GenReqDefForShowDirectConnect()
	return &ShowDirectConnectInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowHostedDirectConnect 查询租户的托管专线详情
//
// 查询合法作伙伴的Hosted专线类型.
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcClient) ShowHostedDirectConnect(request *model.ShowHostedDirectConnectRequest) (*model.ShowHostedDirectConnectResponse, error) {
	requestDef := GenReqDefForShowHostedDirectConnect()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowHostedDirectConnectResponse), nil
	}
}

// ShowHostedDirectConnectInvoker 查询租户的托管专线详情
func (c *DcClient) ShowHostedDirectConnectInvoker(request *model.ShowHostedDirectConnectRequest) *ShowHostedDirectConnectInvoker {
	requestDef := GenReqDefForShowHostedDirectConnect()
	return &ShowHostedDirectConnectInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateDirectConnect 更新物理连接信息
//
// 更新物理连接信息，包括名字,描述等信息
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcClient) UpdateDirectConnect(request *model.UpdateDirectConnectRequest) (*model.UpdateDirectConnectResponse, error) {
	requestDef := GenReqDefForUpdateDirectConnect()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDirectConnectResponse), nil
	}
}

// UpdateDirectConnectInvoker 更新物理连接信息
func (c *DcClient) UpdateDirectConnectInvoker(request *model.UpdateDirectConnectRequest) *UpdateDirectConnectInvoker {
	requestDef := GenReqDefForUpdateDirectConnect()
	return &UpdateDirectConnectInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateHostedDirectConnect 更新托管专线连接
//
// 合作伙伴创建托管专线.
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcClient) UpdateHostedDirectConnect(request *model.UpdateHostedDirectConnectRequest) (*model.UpdateHostedDirectConnectResponse, error) {
	requestDef := GenReqDefForUpdateHostedDirectConnect()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateHostedDirectConnectResponse), nil
	}
}

// UpdateHostedDirectConnectInvoker 更新托管专线连接
func (c *DcClient) UpdateHostedDirectConnectInvoker(request *model.UpdateHostedDirectConnectRequest) *UpdateHostedDirectConnectInvoker {
	requestDef := GenReqDefForUpdateHostedDirectConnect()
	return &UpdateHostedDirectConnectInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchCreateResourceTags 批量添加删除资源标签
//
// - 为指定实例批量添加或删除标签
// - 标签管理服务需要使用该接口批量管理实例的标签。
// - 一个资源上最多有10个标签。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcClient) BatchCreateResourceTags(request *model.BatchCreateResourceTagsRequest) (*model.BatchCreateResourceTagsResponse, error) {
	requestDef := GenReqDefForBatchCreateResourceTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateResourceTagsResponse), nil
	}
}

// BatchCreateResourceTagsInvoker 批量添加删除资源标签
func (c *DcClient) BatchCreateResourceTagsInvoker(request *model.BatchCreateResourceTagsRequest) *BatchCreateResourceTagsInvoker {
	requestDef := GenReqDefForBatchCreateResourceTags()
	return &BatchCreateResourceTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateResourceTag 添加资源标签
//
// - 一个资源上最多有10个标签。
// - 此接口为幂等接口：
// - 创建时，如果创建的标签已经存在（key相同），则覆盖。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcClient) CreateResourceTag(request *model.CreateResourceTagRequest) (*model.CreateResourceTagResponse, error) {
	requestDef := GenReqDefForCreateResourceTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateResourceTagResponse), nil
	}
}

// CreateResourceTagInvoker 添加资源标签
func (c *DcClient) CreateResourceTagInvoker(request *model.CreateResourceTagRequest) *CreateResourceTagInvoker {
	requestDef := GenReqDefForCreateResourceTag()
	return &CreateResourceTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteResourceTag 删除资源标签
//
// 删除时,不对标签字符集做校验，调用接口前必须要做encodeURI，服务端需要对接口uri做decodeURI。删除的key不存在报404，Key不能为空或者空字符串。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcClient) DeleteResourceTag(request *model.DeleteResourceTagRequest) (*model.DeleteResourceTagResponse, error) {
	requestDef := GenReqDefForDeleteResourceTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteResourceTagResponse), nil
	}
}

// DeleteResourceTagInvoker 删除资源标签
func (c *DcClient) DeleteResourceTagInvoker(request *model.DeleteResourceTagRequest) *DeleteResourceTagInvoker {
	requestDef := GenReqDefForDeleteResourceTag()
	return &DeleteResourceTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProjectTags 查询项目标签
//
// - 查询租户在指定Project中实例类型的所有资源标签集合。
// - 标签管理服务需要能够列出当前租户全部已使用的资源标签集合，为各服务打资源标签和过滤实例时提供标签联想功能。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcClient) ListProjectTags(request *model.ListProjectTagsRequest) (*model.ListProjectTagsResponse, error) {
	requestDef := GenReqDefForListProjectTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProjectTagsResponse), nil
	}
}

// ListProjectTagsInvoker 查询项目标签
func (c *DcClient) ListProjectTagsInvoker(request *model.ListProjectTagsRequest) *ListProjectTagsInvoker {
	requestDef := GenReqDefForListProjectTags()
	return &ListProjectTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTagResourceInstances 通过标签查询资源实例
//
// 通过标签查询资源实例
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcClient) ListTagResourceInstances(request *model.ListTagResourceInstancesRequest) (*model.ListTagResourceInstancesResponse, error) {
	requestDef := GenReqDefForListTagResourceInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTagResourceInstancesResponse), nil
	}
}

// ListTagResourceInstancesInvoker 通过标签查询资源实例
func (c *DcClient) ListTagResourceInstancesInvoker(request *model.ListTagResourceInstancesRequest) *ListTagResourceInstancesInvoker {
	requestDef := GenReqDefForListTagResourceInstances()
	return &ListTagResourceInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowResourceTag 查询资源标签
//
// 查询资源标签
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcClient) ShowResourceTag(request *model.ShowResourceTagRequest) (*model.ShowResourceTagResponse, error) {
	requestDef := GenReqDefForShowResourceTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowResourceTagResponse), nil
	}
}

// ShowResourceTagInvoker 查询资源标签
func (c *DcClient) ShowResourceTagInvoker(request *model.ShowResourceTagRequest) *ShowResourceTagInvoker {
	requestDef := GenReqDefForShowResourceTag()
	return &ShowResourceTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateVirtualGateway 创建虑拟网关
//
// 创建虑拟网关
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcClient) CreateVirtualGateway(request *model.CreateVirtualGatewayRequest) (*model.CreateVirtualGatewayResponse, error) {
	requestDef := GenReqDefForCreateVirtualGateway()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateVirtualGatewayResponse), nil
	}
}

// CreateVirtualGatewayInvoker 创建虑拟网关
func (c *DcClient) CreateVirtualGatewayInvoker(request *model.CreateVirtualGatewayRequest) *CreateVirtualGatewayInvoker {
	requestDef := GenReqDefForCreateVirtualGateway()
	return &CreateVirtualGatewayInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteVirtualGateway 删除虚拟网关
//
// 删除指定的虚拟网关
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcClient) DeleteVirtualGateway(request *model.DeleteVirtualGatewayRequest) (*model.DeleteVirtualGatewayResponse, error) {
	requestDef := GenReqDefForDeleteVirtualGateway()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteVirtualGatewayResponse), nil
	}
}

// DeleteVirtualGatewayInvoker 删除虚拟网关
func (c *DcClient) DeleteVirtualGatewayInvoker(request *model.DeleteVirtualGatewayRequest) *DeleteVirtualGatewayInvoker {
	requestDef := GenReqDefForDeleteVirtualGateway()
	return &DeleteVirtualGatewayInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListVirtualGateways 查询虚拟网关列表
//
// 查询虚拟网关列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcClient) ListVirtualGateways(request *model.ListVirtualGatewaysRequest) (*model.ListVirtualGatewaysResponse, error) {
	requestDef := GenReqDefForListVirtualGateways()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVirtualGatewaysResponse), nil
	}
}

// ListVirtualGatewaysInvoker 查询虚拟网关列表
func (c *DcClient) ListVirtualGatewaysInvoker(request *model.ListVirtualGatewaysRequest) *ListVirtualGatewaysInvoker {
	requestDef := GenReqDefForListVirtualGateways()
	return &ListVirtualGatewaysInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowVirtualGateway 查询虚拟网关详情
//
// 查询指定虚拟网关的详细信息
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcClient) ShowVirtualGateway(request *model.ShowVirtualGatewayRequest) (*model.ShowVirtualGatewayResponse, error) {
	requestDef := GenReqDefForShowVirtualGateway()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVirtualGatewayResponse), nil
	}
}

// ShowVirtualGatewayInvoker 查询虚拟网关详情
func (c *DcClient) ShowVirtualGatewayInvoker(request *model.ShowVirtualGatewayRequest) *ShowVirtualGatewayInvoker {
	requestDef := GenReqDefForShowVirtualGateway()
	return &ShowVirtualGatewayInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateVirtualGateway 修改虚拟网关信息
//
// 修改虚拟网关的信息
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcClient) UpdateVirtualGateway(request *model.UpdateVirtualGatewayRequest) (*model.UpdateVirtualGatewayResponse, error) {
	requestDef := GenReqDefForUpdateVirtualGateway()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateVirtualGatewayResponse), nil
	}
}

// UpdateVirtualGatewayInvoker 修改虚拟网关信息
func (c *DcClient) UpdateVirtualGatewayInvoker(request *model.UpdateVirtualGatewayRequest) *UpdateVirtualGatewayInvoker {
	requestDef := GenReqDefForUpdateVirtualGateway()
	return &UpdateVirtualGatewayInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateVirtualInterface 创建虚拟接口
//
// 虚拟接口配置物理专线上与客户互联的IP和路由等相关信息
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcClient) CreateVirtualInterface(request *model.CreateVirtualInterfaceRequest) (*model.CreateVirtualInterfaceResponse, error) {
	requestDef := GenReqDefForCreateVirtualInterface()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateVirtualInterfaceResponse), nil
	}
}

// CreateVirtualInterfaceInvoker 创建虚拟接口
func (c *DcClient) CreateVirtualInterfaceInvoker(request *model.CreateVirtualInterfaceRequest) *CreateVirtualInterfaceInvoker {
	requestDef := GenReqDefForCreateVirtualInterface()
	return &CreateVirtualInterfaceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteVirtualInterface 删除虚拟接口
//
// 删除虚拟接口
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcClient) DeleteVirtualInterface(request *model.DeleteVirtualInterfaceRequest) (*model.DeleteVirtualInterfaceResponse, error) {
	requestDef := GenReqDefForDeleteVirtualInterface()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteVirtualInterfaceResponse), nil
	}
}

// DeleteVirtualInterfaceInvoker 删除虚拟接口
func (c *DcClient) DeleteVirtualInterfaceInvoker(request *model.DeleteVirtualInterfaceRequest) *DeleteVirtualInterfaceInvoker {
	requestDef := GenReqDefForDeleteVirtualInterface()
	return &DeleteVirtualInterfaceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListVirtualInterfaces 查询虚拟接口列表
//
// 查询租户所有的虚拟接口列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcClient) ListVirtualInterfaces(request *model.ListVirtualInterfacesRequest) (*model.ListVirtualInterfacesResponse, error) {
	requestDef := GenReqDefForListVirtualInterfaces()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVirtualInterfacesResponse), nil
	}
}

// ListVirtualInterfacesInvoker 查询虚拟接口列表
func (c *DcClient) ListVirtualInterfacesInvoker(request *model.ListVirtualInterfacesRequest) *ListVirtualInterfacesInvoker {
	requestDef := GenReqDefForListVirtualInterfaces()
	return &ListVirtualInterfacesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowVirtualInterface 查询虚拟接口详情
//
// 查询虚拟接口详细信息
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcClient) ShowVirtualInterface(request *model.ShowVirtualInterfaceRequest) (*model.ShowVirtualInterfaceResponse, error) {
	requestDef := GenReqDefForShowVirtualInterface()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVirtualInterfaceResponse), nil
	}
}

// ShowVirtualInterfaceInvoker 查询虚拟接口详情
func (c *DcClient) ShowVirtualInterfaceInvoker(request *model.ShowVirtualInterfaceRequest) *ShowVirtualInterfaceInvoker {
	requestDef := GenReqDefForShowVirtualInterface()
	return &ShowVirtualInterfaceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateVirtualInterface 修改虚拟接口virtual_interface
//
// 修改虚拟接口的详细信息
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcClient) UpdateVirtualInterface(request *model.UpdateVirtualInterfaceRequest) (*model.UpdateVirtualInterfaceResponse, error) {
	requestDef := GenReqDefForUpdateVirtualInterface()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateVirtualInterfaceResponse), nil
	}
}

// UpdateVirtualInterfaceInvoker 修改虚拟接口virtual_interface
func (c *DcClient) UpdateVirtualInterfaceInvoker(request *model.UpdateVirtualInterfaceRequest) *UpdateVirtualInterfaceInvoker {
	requestDef := GenReqDefForUpdateVirtualInterface()
	return &UpdateVirtualInterfaceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
