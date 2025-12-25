package v3

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dc/v3/model"
)

type DcClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewDcClient(hcClient *httpclient.HcHttpClient) *DcClient {
	return &DcClient{HcClient: hcClient}
}

func DcClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// BindGlobalEips 绑定GEIP操作
//
// 绑定GEIP操作
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcClient) BindGlobalEips(request *model.BindGlobalEipsRequest) (*model.BindGlobalEipsResponse, error) {
	requestDef := GenReqDefForBindGlobalEips()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BindGlobalEipsResponse), nil
	}
}

// BindGlobalEipsInvoker 绑定GEIP操作
func (c *DcClient) BindGlobalEipsInvoker(request *model.BindGlobalEipsRequest) *BindGlobalEipsInvoker {
	requestDef := GenReqDefForBindGlobalEips()
	return &BindGlobalEipsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGlobalEips 查询已经绑定的GEIP列表
//
// 查询已经绑定的GEIP列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcClient) ListGlobalEips(request *model.ListGlobalEipsRequest) (*model.ListGlobalEipsResponse, error) {
	requestDef := GenReqDefForListGlobalEips()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGlobalEipsResponse), nil
	}
}

// ListGlobalEipsInvoker 查询已经绑定的GEIP列表
func (c *DcClient) ListGlobalEipsInvoker(request *model.ListGlobalEipsRequest) *ListGlobalEipsInvoker {
	requestDef := GenReqDefForListGlobalEips()
	return &ListGlobalEipsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UnbindGlobalEips 解绑GEIP
//
// 解绑GEIP
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcClient) UnbindGlobalEips(request *model.UnbindGlobalEipsRequest) (*model.UnbindGlobalEipsResponse, error) {
	requestDef := GenReqDefForUnbindGlobalEips()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UnbindGlobalEipsResponse), nil
	}
}

// UnbindGlobalEipsInvoker 解绑GEIP
func (c *DcClient) UnbindGlobalEipsInvoker(request *model.UnbindGlobalEipsRequest) *UnbindGlobalEipsInvoker {
	requestDef := GenReqDefForUnbindGlobalEips()
	return &UnbindGlobalEipsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateConnectGateway 创建互联网关
//
// 创建互联网关
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcClient) CreateConnectGateway(request *model.CreateConnectGatewayRequest) (*model.CreateConnectGatewayResponse, error) {
	requestDef := GenReqDefForCreateConnectGateway()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateConnectGatewayResponse), nil
	}
}

// CreateConnectGatewayInvoker 创建互联网关
func (c *DcClient) CreateConnectGatewayInvoker(request *model.CreateConnectGatewayRequest) *CreateConnectGatewayInvoker {
	requestDef := GenReqDefForCreateConnectGateway()
	return &CreateConnectGatewayInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteConnectGateway 删除互联网关
//
// 删除互联网关
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcClient) DeleteConnectGateway(request *model.DeleteConnectGatewayRequest) (*model.DeleteConnectGatewayResponse, error) {
	requestDef := GenReqDefForDeleteConnectGateway()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteConnectGatewayResponse), nil
	}
}

// DeleteConnectGatewayInvoker 删除互联网关
func (c *DcClient) DeleteConnectGatewayInvoker(request *model.DeleteConnectGatewayRequest) *DeleteConnectGatewayInvoker {
	requestDef := GenReqDefForDeleteConnectGateway()
	return &DeleteConnectGatewayInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListConnectGateways 查询互联网关列表信息
//
// 查询互联网关列表信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcClient) ListConnectGateways(request *model.ListConnectGatewaysRequest) (*model.ListConnectGatewaysResponse, error) {
	requestDef := GenReqDefForListConnectGateways()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListConnectGatewaysResponse), nil
	}
}

// ListConnectGatewaysInvoker 查询互联网关列表信息
func (c *DcClient) ListConnectGatewaysInvoker(request *model.ListConnectGatewaysRequest) *ListConnectGatewaysInvoker {
	requestDef := GenReqDefForListConnectGateways()
	return &ListConnectGatewaysInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowConnectGateway 查询互联网关详细信息
//
// 查询互联网关详细信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcClient) ShowConnectGateway(request *model.ShowConnectGatewayRequest) (*model.ShowConnectGatewayResponse, error) {
	requestDef := GenReqDefForShowConnectGateway()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowConnectGatewayResponse), nil
	}
}

// ShowConnectGatewayInvoker 查询互联网关详细信息
func (c *DcClient) ShowConnectGatewayInvoker(request *model.ShowConnectGatewayRequest) *ShowConnectGatewayInvoker {
	requestDef := GenReqDefForShowConnectGateway()
	return &ShowConnectGatewayInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateConnectGateway 更新互联网关
//
// 更新互联网关
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcClient) UpdateConnectGateway(request *model.UpdateConnectGatewayRequest) (*model.UpdateConnectGatewayResponse, error) {
	requestDef := GenReqDefForUpdateConnectGateway()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateConnectGatewayResponse), nil
	}
}

// UpdateConnectGatewayInvoker 更新互联网关
func (c *DcClient) UpdateConnectGatewayInvoker(request *model.UpdateConnectGatewayRequest) *UpdateConnectGatewayInvoker {
	requestDef := GenReqDefForUpdateConnectGateway()
	return &UpdateConnectGatewayInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateHostedDirectConnect 创建托管专线连接
//
// 用于合作伙伴用户最终租户创建托管专线
// 创建者必须拥有合作伙伴资质，并且已经构建好运营(hosting)专线
//
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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

// ShowDirectConnect 查询物理连接详情
//
// 查询物理连接详细信息.
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcClient) ShowDirectConnect(request *model.ShowDirectConnectRequest) (*model.ShowDirectConnectResponse, error) {
	requestDef := GenReqDefForShowDirectConnect()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDirectConnectResponse), nil
	}
}

// ShowDirectConnectInvoker 查询物理连接详情
func (c *DcClient) ShowDirectConnectInvoker(request *model.ShowDirectConnectRequest) *ShowDirectConnectInvoker {
	requestDef := GenReqDefForShowDirectConnect()
	return &ShowDirectConnectInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowHostedDirectConnect 查询租户的托管专线详情
//
// 查询合法作伙伴的Hosted专线类型.
//
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// 合作伙伴更新托管专线.
//
// Please refer to HUAWEI cloud API Explorer for details.
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

// ListDirectConnectLocations 查询专线接入点位置列表
//
// 查询本区域下所有专线的接入点的信息，分页查询使用的参数为marker、limit。marker和limit一起使用时才会生效，单独使用无效
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcClient) ListDirectConnectLocations(request *model.ListDirectConnectLocationsRequest) (*model.ListDirectConnectLocationsResponse, error) {
	requestDef := GenReqDefForListDirectConnectLocations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDirectConnectLocationsResponse), nil
	}
}

// ListDirectConnectLocationsInvoker 查询专线接入点位置列表
func (c *DcClient) ListDirectConnectLocationsInvoker(request *model.ListDirectConnectLocationsRequest) *ListDirectConnectLocationsInvoker {
	requestDef := GenReqDefForListDirectConnectLocations()
	return &ListDirectConnectLocationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDirectConnectLocation 查询指定专线接入点详情
//
// 查询指定的专线接入点详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcClient) ShowDirectConnectLocation(request *model.ShowDirectConnectLocationRequest) (*model.ShowDirectConnectLocationResponse, error) {
	requestDef := GenReqDefForShowDirectConnectLocation()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDirectConnectLocationResponse), nil
	}
}

// ShowDirectConnectLocationInvoker 查询指定专线接入点详情
func (c *DcClient) ShowDirectConnectLocationInvoker(request *model.ShowDirectConnectLocationRequest) *ShowDirectConnectLocationInvoker {
	requestDef := GenReqDefForShowDirectConnectLocation()
	return &ShowDirectConnectLocationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGdgwRouteTables 查询全域接入网关路由表
//
// 查询全域接入网关路由表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcClient) ListGdgwRouteTables(request *model.ListGdgwRouteTablesRequest) (*model.ListGdgwRouteTablesResponse, error) {
	requestDef := GenReqDefForListGdgwRouteTables()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGdgwRouteTablesResponse), nil
	}
}

// ListGdgwRouteTablesInvoker 查询全域接入网关路由表
func (c *DcClient) ListGdgwRouteTablesInvoker(request *model.ListGdgwRouteTablesRequest) *ListGdgwRouteTablesInvoker {
	requestDef := GenReqDefForListGdgwRouteTables()
	return &ListGdgwRouteTablesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateGdgwRouteTable 修改全域接入网关路由表
//
// 支持的修改操作：新增、删除、修改
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcClient) UpdateGdgwRouteTable(request *model.UpdateGdgwRouteTableRequest) (*model.UpdateGdgwRouteTableResponse, error) {
	requestDef := GenReqDefForUpdateGdgwRouteTable()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateGdgwRouteTableResponse), nil
	}
}

// UpdateGdgwRouteTableInvoker 修改全域接入网关路由表
func (c *DcClient) UpdateGdgwRouteTableInvoker(request *model.UpdateGdgwRouteTableRequest) *UpdateGdgwRouteTableInvoker {
	requestDef := GenReqDefForUpdateGdgwRouteTable()
	return &UpdateGdgwRouteTableInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateGlobalDcGateway 创建专线全域接入网关
//
// 创建专线全域接入网关实例(global-dc-gateway)，用于接入全球的ER实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcClient) CreateGlobalDcGateway(request *model.CreateGlobalDcGatewayRequest) (*model.CreateGlobalDcGatewayResponse, error) {
	requestDef := GenReqDefForCreateGlobalDcGateway()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateGlobalDcGatewayResponse), nil
	}
}

// CreateGlobalDcGatewayInvoker 创建专线全域接入网关
func (c *DcClient) CreateGlobalDcGatewayInvoker(request *model.CreateGlobalDcGatewayRequest) *CreateGlobalDcGatewayInvoker {
	requestDef := GenReqDefForCreateGlobalDcGateway()
	return &CreateGlobalDcGatewayInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteGlobalDcGateway 删除专线全域接入网关
//
// 删除专线全域接入网关global-dc-gateway实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcClient) DeleteGlobalDcGateway(request *model.DeleteGlobalDcGatewayRequest) (*model.DeleteGlobalDcGatewayResponse, error) {
	requestDef := GenReqDefForDeleteGlobalDcGateway()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteGlobalDcGatewayResponse), nil
	}
}

// DeleteGlobalDcGatewayInvoker 删除专线全域接入网关
func (c *DcClient) DeleteGlobalDcGatewayInvoker(request *model.DeleteGlobalDcGatewayRequest) *DeleteGlobalDcGatewayInvoker {
	requestDef := GenReqDefForDeleteGlobalDcGateway()
	return &DeleteGlobalDcGatewayInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGlobalDcGateways 查询专线全域接入网关列表
//
// 查询专线全域接入网关列表，建议使用分页查询 分页查询使用的参数为marker、limit。marker和limit一起使用时才会生效，单独使用无效
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcClient) ListGlobalDcGateways(request *model.ListGlobalDcGatewaysRequest) (*model.ListGlobalDcGatewaysResponse, error) {
	requestDef := GenReqDefForListGlobalDcGateways()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGlobalDcGatewaysResponse), nil
	}
}

// ListGlobalDcGatewaysInvoker 查询专线全域接入网关列表
func (c *DcClient) ListGlobalDcGatewaysInvoker(request *model.ListGlobalDcGatewaysRequest) *ListGlobalDcGatewaysInvoker {
	requestDef := GenReqDefForListGlobalDcGateways()
	return &ListGlobalDcGatewaysInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowGlobalDcGateway 查询专线全域接入网关详情
//
// 查询专线全域接入网关实例详情信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcClient) ShowGlobalDcGateway(request *model.ShowGlobalDcGatewayRequest) (*model.ShowGlobalDcGatewayResponse, error) {
	requestDef := GenReqDefForShowGlobalDcGateway()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowGlobalDcGatewayResponse), nil
	}
}

// ShowGlobalDcGatewayInvoker 查询专线全域接入网关详情
func (c *DcClient) ShowGlobalDcGatewayInvoker(request *model.ShowGlobalDcGatewayRequest) *ShowGlobalDcGatewayInvoker {
	requestDef := GenReqDefForShowGlobalDcGateway()
	return &ShowGlobalDcGatewayInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateGlobalDcGateway 更新专线全域接入网关
//
// 更新专线全域接入网关global-dc-gateway的名字，描述等信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcClient) UpdateGlobalDcGateway(request *model.UpdateGlobalDcGatewayRequest) (*model.UpdateGlobalDcGatewayResponse, error) {
	requestDef := GenReqDefForUpdateGlobalDcGateway()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateGlobalDcGatewayResponse), nil
	}
}

// UpdateGlobalDcGatewayInvoker 更新专线全域接入网关
func (c *DcClient) UpdateGlobalDcGatewayInvoker(request *model.UpdateGlobalDcGatewayRequest) *UpdateGlobalDcGatewayInvoker {
	requestDef := GenReqDefForUpdateGlobalDcGateway()
	return &UpdateGlobalDcGatewayInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePeerLink 创建专线关联连接
//
// 创建专线全域接入网关的关联连接peer-link对象，用于连接企业路由器或者其他接入网关
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcClient) CreatePeerLink(request *model.CreatePeerLinkRequest) (*model.CreatePeerLinkResponse, error) {
	requestDef := GenReqDefForCreatePeerLink()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePeerLinkResponse), nil
	}
}

// CreatePeerLinkInvoker 创建专线关联连接
func (c *DcClient) CreatePeerLinkInvoker(request *model.CreatePeerLinkRequest) *CreatePeerLinkInvoker {
	requestDef := GenReqDefForCreatePeerLink()
	return &CreatePeerLinkInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeletePeerLink 删除专线关联连接
//
// 删除全域接入网关与ER的关联连接peer-link
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcClient) DeletePeerLink(request *model.DeletePeerLinkRequest) (*model.DeletePeerLinkResponse, error) {
	requestDef := GenReqDefForDeletePeerLink()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePeerLinkResponse), nil
	}
}

// DeletePeerLinkInvoker 删除专线关联连接
func (c *DcClient) DeletePeerLinkInvoker(request *model.DeletePeerLinkRequest) *DeletePeerLinkInvoker {
	requestDef := GenReqDefForDeletePeerLink()
	return &DeletePeerLinkInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPeerLinks 查询专线关联连接列表
//
// 查询全域接入网关与ER等对象的关联连接列表，分页查询使用的参数为marker、limit。marker和limit一起使用时才会生效，单独使用无效
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcClient) ListPeerLinks(request *model.ListPeerLinksRequest) (*model.ListPeerLinksResponse, error) {
	requestDef := GenReqDefForListPeerLinks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPeerLinksResponse), nil
	}
}

// ListPeerLinksInvoker 查询专线关联连接列表
func (c *DcClient) ListPeerLinksInvoker(request *model.ListPeerLinksRequest) *ListPeerLinksInvoker {
	requestDef := GenReqDefForListPeerLinks()
	return &ListPeerLinksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPeerLink 查询专线关联连接详情
//
// 查询指定接入网关的指定的关联连接(peer link)详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcClient) ShowPeerLink(request *model.ShowPeerLinkRequest) (*model.ShowPeerLinkResponse, error) {
	requestDef := GenReqDefForShowPeerLink()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPeerLinkResponse), nil
	}
}

// ShowPeerLinkInvoker 查询专线关联连接详情
func (c *DcClient) ShowPeerLinkInvoker(request *model.ShowPeerLinkRequest) *ShowPeerLinkInvoker {
	requestDef := GenReqDefForShowPeerLink()
	return &ShowPeerLinkInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePeerLink 更新专线关联连接
//
// 更新接入网关与ER对接的关联连接peer-link
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcClient) UpdatePeerLink(request *model.UpdatePeerLinkRequest) (*model.UpdatePeerLinkResponse, error) {
	requestDef := GenReqDefForUpdatePeerLink()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePeerLinkResponse), nil
	}
}

// UpdatePeerLinkInvoker 更新专线关联连接
func (c *DcClient) UpdatePeerLinkInvoker(request *model.UpdatePeerLinkRequest) *UpdatePeerLinkInvoker {
	requestDef := GenReqDefForUpdatePeerLink()
	return &UpdatePeerLinkInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowQuotas 查询配额
//
// 查询租户各类资源的使用情况，如Directconnect的使用量，虚拟接口的使用量等。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcClient) ShowQuotas(request *model.ShowQuotasRequest) (*model.ShowQuotasResponse, error) {
	requestDef := GenReqDefForShowQuotas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowQuotasResponse), nil
	}
}

// ShowQuotasInvoker 查询配额
func (c *DcClient) ShowQuotasInvoker(request *model.ShowQuotasRequest) *ShowQuotasInvoker {
	requestDef := GenReqDefForShowQuotas()
	return &ShowQuotasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchCreateResourceTags 批量添加删除资源标签
//
// - 为指定实例批量添加或删除标签
// - 标签管理服务需要使用该接口批量管理实例的标签。
// - 一个资源上最多有10个标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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

// CreateVifPeerDetection 创建虚拟接口对等体连通性探测实例
//
// 当您想对虚拟接口对等体的远端网关的连通性进行探测时，可以通过调用此接口创建一个虚拟接口对等体连通性探测实例来实现。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcClient) CreateVifPeerDetection(request *model.CreateVifPeerDetectionRequest) (*model.CreateVifPeerDetectionResponse, error) {
	requestDef := GenReqDefForCreateVifPeerDetection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateVifPeerDetectionResponse), nil
	}
}

// CreateVifPeerDetectionInvoker 创建虚拟接口对等体连通性探测实例
func (c *DcClient) CreateVifPeerDetectionInvoker(request *model.CreateVifPeerDetectionRequest) *CreateVifPeerDetectionInvoker {
	requestDef := GenReqDefForCreateVifPeerDetection()
	return &CreateVifPeerDetectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteVifPeerDetection 删除虚拟接口对等体连通性探测实例
//
// 当您想不再保留虚拟接口对等体连通性探测实例时，您可以通过调用此接口将其删除。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcClient) DeleteVifPeerDetection(request *model.DeleteVifPeerDetectionRequest) (*model.DeleteVifPeerDetectionResponse, error) {
	requestDef := GenReqDefForDeleteVifPeerDetection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteVifPeerDetectionResponse), nil
	}
}

// DeleteVifPeerDetectionInvoker 删除虚拟接口对等体连通性探测实例
func (c *DcClient) DeleteVifPeerDetectionInvoker(request *model.DeleteVifPeerDetectionRequest) *DeleteVifPeerDetectionInvoker {
	requestDef := GenReqDefForDeleteVifPeerDetection()
	return &DeleteVifPeerDetectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListVifPeerDetections 查询虚拟接口对等体连通性探测实例列表
//
// 当您的对虚拟接口对等体发起连通性探测后，您可以通过此接口查询多次探测的信息，包括ID、探测开始时间、探测结束时间、探测状态、丢包率等。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcClient) ListVifPeerDetections(request *model.ListVifPeerDetectionsRequest) (*model.ListVifPeerDetectionsResponse, error) {
	requestDef := GenReqDefForListVifPeerDetections()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVifPeerDetectionsResponse), nil
	}
}

// ListVifPeerDetectionsInvoker 查询虚拟接口对等体连通性探测实例列表
func (c *DcClient) ListVifPeerDetectionsInvoker(request *model.ListVifPeerDetectionsRequest) *ListVifPeerDetectionsInvoker {
	requestDef := GenReqDefForListVifPeerDetections()
	return &ListVifPeerDetectionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowVifPeerDetection 查询虚拟接口对等体连通性探测实例
//
// 当您的对虚拟接口对等体发起连通性探测后，您可以通过此接口查询单次探测的信息，包括ID、探测开始时间、探测结束时间、探测状态、丢包率等。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcClient) ShowVifPeerDetection(request *model.ShowVifPeerDetectionRequest) (*model.ShowVifPeerDetectionResponse, error) {
	requestDef := GenReqDefForShowVifPeerDetection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVifPeerDetectionResponse), nil
	}
}

// ShowVifPeerDetectionInvoker 查询虚拟接口对等体连通性探测实例
func (c *DcClient) ShowVifPeerDetectionInvoker(request *model.ShowVifPeerDetectionRequest) *ShowVifPeerDetectionInvoker {
	requestDef := GenReqDefForShowVifPeerDetection()
	return &ShowVifPeerDetectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateVirtualGateway 创建虚拟网关
//
// 创建虚拟网关
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcClient) CreateVirtualGateway(request *model.CreateVirtualGatewayRequest) (*model.CreateVirtualGatewayResponse, error) {
	requestDef := GenReqDefForCreateVirtualGateway()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateVirtualGatewayResponse), nil
	}
}

// CreateVirtualGatewayInvoker 创建虚拟网关
func (c *DcClient) CreateVirtualGatewayInvoker(request *model.CreateVirtualGatewayRequest) *CreateVirtualGatewayInvoker {
	requestDef := GenReqDefForCreateVirtualGateway()
	return &CreateVirtualGatewayInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteVirtualGateway 删除虚拟网关
//
// 删除指定的虚拟网关
//
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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

// UpdateVirtualGateway 更新虚拟网关信息
//
// 更新虚拟网关的信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcClient) UpdateVirtualGateway(request *model.UpdateVirtualGatewayRequest) (*model.UpdateVirtualGatewayResponse, error) {
	requestDef := GenReqDefForUpdateVirtualGateway()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateVirtualGatewayResponse), nil
	}
}

// UpdateVirtualGatewayInvoker 更新虚拟网关信息
func (c *DcClient) UpdateVirtualGatewayInvoker(request *model.UpdateVirtualGatewayRequest) *UpdateVirtualGatewayInvoker {
	requestDef := GenReqDefForUpdateVirtualGateway()
	return &UpdateVirtualGatewayInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateVifPeer 创建虚拟接口对等体
//
// 每个虚拟接口可支持两个对等体，IPv4和IPv6对等体，在创建虚拟接口时默认创建IPv4对等体， 本接口一般用于增加ipv6对等体。创建虚拟接口对接体之后，可以通过虚拟接口查询配置结果 本接口只在支持Ipv6的区域开放，如需要使用请联系客服
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcClient) CreateVifPeer(request *model.CreateVifPeerRequest) (*model.CreateVifPeerResponse, error) {
	requestDef := GenReqDefForCreateVifPeer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateVifPeerResponse), nil
	}
}

// CreateVifPeerInvoker 创建虚拟接口对等体
func (c *DcClient) CreateVifPeerInvoker(request *model.CreateVifPeerRequest) *CreateVifPeerInvoker {
	requestDef := GenReqDefForCreateVifPeer()
	return &CreateVifPeerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateVirtualInterface 创建虚拟接口
//
// 虚拟接口配置物理专线上与客户互联的IP和路由等相关信息
//
// Please refer to HUAWEI cloud API Explorer for details.
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

// DeleteVifPeer 删除虚拟接口对应的对等体
//
// 删除虚拟接口对等体信息,虚拟接口至少要含一个对等体，最后一个对等体不能删除 本接口只在支持Ipv6的区域开放，如需要使用请联系客服
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcClient) DeleteVifPeer(request *model.DeleteVifPeerRequest) (*model.DeleteVifPeerResponse, error) {
	requestDef := GenReqDefForDeleteVifPeer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteVifPeerResponse), nil
	}
}

// DeleteVifPeerInvoker 删除虚拟接口对应的对等体
func (c *DcClient) DeleteVifPeerInvoker(request *model.DeleteVifPeerRequest) *DeleteVifPeerInvoker {
	requestDef := GenReqDefForDeleteVifPeer()
	return &DeleteVifPeerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteVirtualInterface 删除虚拟接口
//
// 删除虚拟接口
//
// Please refer to HUAWEI cloud API Explorer for details.
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

// ListSwitchoverTestRecords 查询虚拟接口倒换测试记录列表
//
// 查询倒换测试记录列表，只展示operate_status为COMPELTE的记录。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcClient) ListSwitchoverTestRecords(request *model.ListSwitchoverTestRecordsRequest) (*model.ListSwitchoverTestRecordsResponse, error) {
	requestDef := GenReqDefForListSwitchoverTestRecords()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSwitchoverTestRecordsResponse), nil
	}
}

// ListSwitchoverTestRecordsInvoker 查询虚拟接口倒换测试记录列表
func (c *DcClient) ListSwitchoverTestRecordsInvoker(request *model.ListSwitchoverTestRecordsRequest) *ListSwitchoverTestRecordsInvoker {
	requestDef := GenReqDefForListSwitchoverTestRecords()
	return &ListSwitchoverTestRecordsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListVirtualInterfaces 查询虚拟接口列表
//
// 查询租户所有的虚拟接口列表
//
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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

// SwitchoverTest 执行虚拟接口倒换测试
//
// 客户双专线接入，需要支持双线自动倒换，方便进行功能测试。 虚拟接口进行倒换测试会导致接口关闭，业务流量中断。 对于虚拟接口，支持“关闭接口”和“开放接口”两种操作。 1、关闭接口：下发shutdown命令，关闭接口； 2、开放接口：下发undo_shutdown命令，使能接口。 倒换测试选择shutdown时，虚拟接口的状态为ADMIN_SHUTDOWN，此状态不允许虚拟接口的其他操作。 倒换测试选择undo_shutdown时，虚拟接口的状态为ACTIVE。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcClient) SwitchoverTest(request *model.SwitchoverTestRequest) (*model.SwitchoverTestResponse, error) {
	requestDef := GenReqDefForSwitchoverTest()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SwitchoverTestResponse), nil
	}
}

// SwitchoverTestInvoker 执行虚拟接口倒换测试
func (c *DcClient) SwitchoverTestInvoker(request *model.SwitchoverTestRequest) *SwitchoverTestInvoker {
	requestDef := GenReqDefForSwitchoverTest()
	return &SwitchoverTestInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateExtendAttribute 修改虚拟接口可靠性检测的扩展参数
//
// 虚拟接口有bfd与nqa两种可靠性检测方式，您可以通过调用此接口修改可靠性检测的参数，例如检测报文最小发送间隔、检测报文最大发送间隔、检测周期等信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcClient) UpdateExtendAttribute(request *model.UpdateExtendAttributeRequest) (*model.UpdateExtendAttributeResponse, error) {
	requestDef := GenReqDefForUpdateExtendAttribute()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateExtendAttributeResponse), nil
	}
}

// UpdateExtendAttributeInvoker 修改虚拟接口可靠性检测的扩展参数
func (c *DcClient) UpdateExtendAttributeInvoker(request *model.UpdateExtendAttributeRequest) *UpdateExtendAttributeInvoker {
	requestDef := GenReqDefForUpdateExtendAttribute()
	return &UpdateExtendAttributeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateVifPeer 更新虚拟接口对等体
//
// 更新虚拟接口对等体信息,包括远端子网，名字和描述等。 本接口只在支持Ipv6的区域开放，如需要使用请联系客服
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcClient) UpdateVifPeer(request *model.UpdateVifPeerRequest) (*model.UpdateVifPeerResponse, error) {
	requestDef := GenReqDefForUpdateVifPeer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateVifPeerResponse), nil
	}
}

// UpdateVifPeerInvoker 更新虚拟接口对等体
func (c *DcClient) UpdateVifPeerInvoker(request *model.UpdateVifPeerRequest) *UpdateVifPeerInvoker {
	requestDef := GenReqDefForUpdateVifPeer()
	return &UpdateVifPeerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateVirtualInterface 更新虚拟接口
//
// 更新虚拟接口的详细信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcClient) UpdateVirtualInterface(request *model.UpdateVirtualInterfaceRequest) (*model.UpdateVirtualInterfaceResponse, error) {
	requestDef := GenReqDefForUpdateVirtualInterface()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateVirtualInterfaceResponse), nil
	}
}

// UpdateVirtualInterfaceInvoker 更新虚拟接口
func (c *DcClient) UpdateVirtualInterfaceInvoker(request *model.UpdateVirtualInterfaceRequest) *UpdateVirtualInterfaceInvoker {
	requestDef := GenReqDefForUpdateVirtualInterface()
	return &UpdateVirtualInterfaceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
