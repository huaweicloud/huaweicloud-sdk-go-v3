package v1

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/mssi/v1/model"
)

type MssiClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewMssiClient(hcClient *httpclient.HcHttpClient) *MssiClient {
	return &MssiClient{HcClient: hcClient}
}

func MssiClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// CreateConnectionInfo 新建Connection
//
// 新建连接
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MssiClient) CreateConnectionInfo(request *model.CreateConnectionInfoRequest) (*model.CreateConnectionInfoResponse, error) {
	requestDef := GenReqDefForCreateConnectionInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateConnectionInfoResponse), nil
	}
}

// CreateConnectionInfoInvoker 新建Connection
func (c *MssiClient) CreateConnectionInfoInvoker(request *model.CreateConnectionInfoRequest) *CreateConnectionInfoInvoker {
	requestDef := GenReqDefForCreateConnectionInfo()
	return &CreateConnectionInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateCustomConnectorFromOpenapi 新建自定义连接器(导入swagger方式)
//
// 新建自定义连接器(导入swagger方式)
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MssiClient) CreateCustomConnectorFromOpenapi(request *model.CreateCustomConnectorFromOpenapiRequest) (*model.CreateCustomConnectorFromOpenapiResponse, error) {
	requestDef := GenReqDefForCreateCustomConnectorFromOpenapi()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCustomConnectorFromOpenapiResponse), nil
	}
}

// CreateCustomConnectorFromOpenapiInvoker 新建自定义连接器(导入swagger方式)
func (c *MssiClient) CreateCustomConnectorFromOpenapiInvoker(request *model.CreateCustomConnectorFromOpenapiRequest) *CreateCustomConnectorFromOpenapiInvoker {
	requestDef := GenReqDefForCreateCustomConnectorFromOpenapi()
	return &CreateCustomConnectorFromOpenapiInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateFlow 创建flow
//
// 创建flow
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MssiClient) CreateFlow(request *model.CreateFlowRequest) (*model.CreateFlowResponse, error) {
	requestDef := GenReqDefForCreateFlow()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateFlowResponse), nil
	}
}

// CreateFlowInvoker 创建flow
func (c *MssiClient) CreateFlowInvoker(request *model.CreateFlowRequest) *CreateFlowInvoker {
	requestDef := GenReqDefForCreateFlow()
	return &CreateFlowInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateFlowTemplateFromFlow 根据流创建Flow模板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MssiClient) CreateFlowTemplateFromFlow(request *model.CreateFlowTemplateFromFlowRequest) (*model.CreateFlowTemplateFromFlowResponse, error) {
	requestDef := GenReqDefForCreateFlowTemplateFromFlow()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateFlowTemplateFromFlowResponse), nil
	}
}

// CreateFlowTemplateFromFlowInvoker 根据流创建Flow模板
func (c *MssiClient) CreateFlowTemplateFromFlowInvoker(request *model.CreateFlowTemplateFromFlowRequest) *CreateFlowTemplateFromFlowInvoker {
	requestDef := GenReqDefForCreateFlowTemplateFromFlow()
	return &CreateFlowTemplateFromFlowInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteConnectionInfo 删除Connection
//
// 删除Connection
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MssiClient) DeleteConnectionInfo(request *model.DeleteConnectionInfoRequest) (*model.DeleteConnectionInfoResponse, error) {
	requestDef := GenReqDefForDeleteConnectionInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteConnectionInfoResponse), nil
	}
}

// DeleteConnectionInfoInvoker 删除Connection
func (c *MssiClient) DeleteConnectionInfoInvoker(request *model.DeleteConnectionInfoRequest) *DeleteConnectionInfoInvoker {
	requestDef := GenReqDefForDeleteConnectionInfo()
	return &DeleteConnectionInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteCustomConnector 删除自定义连接器
//
// 删除自定义连接器
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MssiClient) DeleteCustomConnector(request *model.DeleteCustomConnectorRequest) (*model.DeleteCustomConnectorResponse, error) {
	requestDef := GenReqDefForDeleteCustomConnector()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteCustomConnectorResponse), nil
	}
}

// DeleteCustomConnectorInvoker 删除自定义连接器
func (c *MssiClient) DeleteCustomConnectorInvoker(request *model.DeleteCustomConnectorRequest) *DeleteCustomConnectorInvoker {
	requestDef := GenReqDefForDeleteCustomConnector()
	return &DeleteCustomConnectorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteFlow 删除Flow
//
// 删除Flow
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MssiClient) DeleteFlow(request *model.DeleteFlowRequest) (*model.DeleteFlowResponse, error) {
	requestDef := GenReqDefForDeleteFlow()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteFlowResponse), nil
	}
}

// DeleteFlowInvoker 删除Flow
func (c *MssiClient) DeleteFlowInvoker(request *model.DeleteFlowRequest) *DeleteFlowInvoker {
	requestDef := GenReqDefForDeleteFlow()
	return &DeleteFlowInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchFlowById 查询特定flow
//
// 查询特定flow
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MssiClient) SearchFlowById(request *model.SearchFlowByIdRequest) (*model.SearchFlowByIdResponse, error) {
	requestDef := GenReqDefForSearchFlowById()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchFlowByIdResponse), nil
	}
}

// SearchFlowByIdInvoker 查询特定flow
func (c *MssiClient) SearchFlowByIdInvoker(request *model.SearchFlowByIdRequest) *SearchFlowByIdInvoker {
	requestDef := GenReqDefForSearchFlowById()
	return &SearchFlowByIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAllConnections 查询Connection列表
//
// 查询所有连接
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MssiClient) ShowAllConnections(request *model.ShowAllConnectionsRequest) (*model.ShowAllConnectionsResponse, error) {
	requestDef := GenReqDefForShowAllConnections()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAllConnectionsResponse), nil
	}
}

// ShowAllConnectionsInvoker 查询Connection列表
func (c *MssiClient) ShowAllConnectionsInvoker(request *model.ShowAllConnectionsRequest) *ShowAllConnectionsInvoker {
	requestDef := GenReqDefForShowAllConnections()
	return &ShowAllConnectionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAllFlows 查询所有Flow
//
// 查询所有Flow
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MssiClient) ShowAllFlows(request *model.ShowAllFlowsRequest) (*model.ShowAllFlowsResponse, error) {
	requestDef := GenReqDefForShowAllFlows()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAllFlowsResponse), nil
	}
}

// ShowAllFlowsInvoker 查询所有Flow
func (c *MssiClient) ShowAllFlowsInvoker(request *model.ShowAllFlowsRequest) *ShowAllFlowsInvoker {
	requestDef := GenReqDefForShowAllFlows()
	return &ShowAllFlowsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowConnectors 查询Connector列表
//
// 查询Connector列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MssiClient) ShowConnectors(request *model.ShowConnectorsRequest) (*model.ShowConnectorsResponse, error) {
	requestDef := GenReqDefForShowConnectors()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowConnectorsResponse), nil
	}
}

// ShowConnectorsInvoker 查询Connector列表
func (c *MssiClient) ShowConnectorsInvoker(request *model.ShowConnectorsRequest) *ShowConnectorsInvoker {
	requestDef := GenReqDefForShowConnectors()
	return &ShowConnectorsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCustomConnector 发布连接器
//
// 发布连接器
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MssiClient) ShowCustomConnector(request *model.ShowCustomConnectorRequest) (*model.ShowCustomConnectorResponse, error) {
	requestDef := GenReqDefForShowCustomConnector()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCustomConnectorResponse), nil
	}
}

// ShowCustomConnectorInvoker 发布连接器
func (c *MssiClient) ShowCustomConnectorInvoker(request *model.ShowCustomConnectorRequest) *ShowCustomConnectorInvoker {
	requestDef := GenReqDefForShowCustomConnector()
	return &ShowCustomConnectorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCustomConnectors 查询CustomConnector列表
//
// 查询CustomConnector列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MssiClient) ShowCustomConnectors(request *model.ShowCustomConnectorsRequest) (*model.ShowCustomConnectorsResponse, error) {
	requestDef := GenReqDefForShowCustomConnectors()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCustomConnectorsResponse), nil
	}
}

// ShowCustomConnectorsInvoker 查询CustomConnector列表
func (c *MssiClient) ShowCustomConnectorsInvoker(request *model.ShowCustomConnectorsRequest) *ShowCustomConnectorsInvoker {
	requestDef := GenReqDefForShowCustomConnectors()
	return &ShowCustomConnectorsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSingleConnection 查询单个Connection
//
// 查询单个连接
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MssiClient) ShowSingleConnection(request *model.ShowSingleConnectionRequest) (*model.ShowSingleConnectionResponse, error) {
	requestDef := GenReqDefForShowSingleConnection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSingleConnectionResponse), nil
	}
}

// ShowSingleConnectionInvoker 查询单个Connection
func (c *MssiClient) ShowSingleConnectionInvoker(request *model.ShowSingleConnectionRequest) *ShowSingleConnectionInvoker {
	requestDef := GenReqDefForShowSingleConnection()
	return &ShowSingleConnectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateConnectionInfo 修改连接配置内容
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MssiClient) UpdateConnectionInfo(request *model.UpdateConnectionInfoRequest) (*model.UpdateConnectionInfoResponse, error) {
	requestDef := GenReqDefForUpdateConnectionInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateConnectionInfoResponse), nil
	}
}

// UpdateConnectionInfoInvoker 修改连接配置内容
func (c *MssiClient) UpdateConnectionInfoInvoker(request *model.UpdateConnectionInfoRequest) *UpdateConnectionInfoInvoker {
	requestDef := GenReqDefForUpdateConnectionInfo()
	return &UpdateConnectionInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateFlow 更新flow
//
// 更新flow
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MssiClient) UpdateFlow(request *model.UpdateFlowRequest) (*model.UpdateFlowResponse, error) {
	requestDef := GenReqDefForUpdateFlow()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateFlowResponse), nil
	}
}

// UpdateFlowInvoker 更新flow
func (c *MssiClient) UpdateFlowInvoker(request *model.UpdateFlowRequest) *UpdateFlowInvoker {
	requestDef := GenReqDefForUpdateFlow()
	return &UpdateFlowInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
