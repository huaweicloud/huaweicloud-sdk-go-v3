package v3

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/er/v3/model"
)

type ErClient struct {
	HcClient *http_client.HcHttpClient
}

func NewErClient(hcClient *http_client.HcHttpClient) *ErClient {
	return &ErClient{HcClient: hcClient}
}

func ErClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// AssociateRouteTable 创建路由关联
//
// 每个连接只能关联到一张路由表。通过创建关联将连接关联到路由表，从该连接收到的报文会用被关联的路由表进行路由。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ErClient) AssociateRouteTable(request *model.AssociateRouteTableRequest) (*model.AssociateRouteTableResponse, error) {
	requestDef := GenReqDefForAssociateRouteTable()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AssociateRouteTableResponse), nil
	}
}

// AssociateRouteTableInvoker 创建路由关联
func (c *ErClient) AssociateRouteTableInvoker(request *model.AssociateRouteTableRequest) *AssociateRouteTableInvoker {
	requestDef := GenReqDefForAssociateRouteTable()
	return &AssociateRouteTableInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DisassociateRouteTable 删除路由关联
//
// 解绑连接和路由表的关联关系。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ErClient) DisassociateRouteTable(request *model.DisassociateRouteTableRequest) (*model.DisassociateRouteTableResponse, error) {
	requestDef := GenReqDefForDisassociateRouteTable()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisassociateRouteTableResponse), nil
	}
}

// DisassociateRouteTableInvoker 删除路由关联
func (c *ErClient) DisassociateRouteTableInvoker(request *model.DisassociateRouteTableRequest) *DisassociateRouteTableInvoker {
	requestDef := GenReqDefForDisassociateRouteTable()
	return &DisassociateRouteTableInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAssociations 查询路由关联列表
//
// 查询路由关联列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ErClient) ListAssociations(request *model.ListAssociationsRequest) (*model.ListAssociationsResponse, error) {
	requestDef := GenReqDefForListAssociations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAssociationsResponse), nil
	}
}

// ListAssociationsInvoker 查询路由关联列表
func (c *ErClient) ListAssociationsInvoker(request *model.ListAssociationsRequest) *ListAssociationsInvoker {
	requestDef := GenReqDefForListAssociations()
	return &ListAssociationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAttachments 查询连接列表
//
// 查询企业路由器实例下的连接列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ErClient) ListAttachments(request *model.ListAttachmentsRequest) (*model.ListAttachmentsResponse, error) {
	requestDef := GenReqDefForListAttachments()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAttachmentsResponse), nil
	}
}

// ListAttachmentsInvoker 查询连接列表
func (c *ErClient) ListAttachmentsInvoker(request *model.ListAttachmentsRequest) *ListAttachmentsInvoker {
	requestDef := GenReqDefForListAttachments()
	return &ListAttachmentsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAttachment 查询连接详情
//
// 查询连接详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ErClient) ShowAttachment(request *model.ShowAttachmentRequest) (*model.ShowAttachmentResponse, error) {
	requestDef := GenReqDefForShowAttachment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAttachmentResponse), nil
	}
}

// ShowAttachmentInvoker 查询连接详情
func (c *ErClient) ShowAttachmentInvoker(request *model.ShowAttachmentRequest) *ShowAttachmentInvoker {
	requestDef := GenReqDefForShowAttachment()
	return &ShowAttachmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAttachment 更新连接基本信息
//
// 修改连接基本信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ErClient) UpdateAttachment(request *model.UpdateAttachmentRequest) (*model.UpdateAttachmentResponse, error) {
	requestDef := GenReqDefForUpdateAttachment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAttachmentResponse), nil
	}
}

// UpdateAttachmentInvoker 更新连接基本信息
func (c *ErClient) UpdateAttachmentInvoker(request *model.UpdateAttachmentRequest) *UpdateAttachmentInvoker {
	requestDef := GenReqDefForUpdateAttachment()
	return &UpdateAttachmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAvailabilityZone 查询可用区列表
//
// 查询支持创建企业路由器实例的可用区列表，当可用区状态为available时，表示可以创建企业路由器实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ErClient) ListAvailabilityZone(request *model.ListAvailabilityZoneRequest) (*model.ListAvailabilityZoneResponse, error) {
	requestDef := GenReqDefForListAvailabilityZone()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAvailabilityZoneResponse), nil
	}
}

// ListAvailabilityZoneInvoker 查询可用区列表
func (c *ErClient) ListAvailabilityZoneInvoker(request *model.ListAvailabilityZoneRequest) *ListAvailabilityZoneInvoker {
	requestDef := GenReqDefForListAvailabilityZone()
	return &ListAvailabilityZoneInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeAvailabilityZone 更新企业路由器的可用区信息
//
// 更新企业路由器的可用区信息，企业路由器实例状态为available的时候才能更新。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ErClient) ChangeAvailabilityZone(request *model.ChangeAvailabilityZoneRequest) (*model.ChangeAvailabilityZoneResponse, error) {
	requestDef := GenReqDefForChangeAvailabilityZone()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeAvailabilityZoneResponse), nil
	}
}

// ChangeAvailabilityZoneInvoker 更新企业路由器的可用区信息
func (c *ErClient) ChangeAvailabilityZoneInvoker(request *model.ChangeAvailabilityZoneRequest) *ChangeAvailabilityZoneInvoker {
	requestDef := GenReqDefForChangeAvailabilityZone()
	return &ChangeAvailabilityZoneInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateEnterpriseRouter 创建企业路由器
//
// 创建企业路由器实例，如果使能默认关联路由表或使能默认传递路由表，那么系统会默认创建一张路由表，作为默认关联路由表或默认传递路由表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ErClient) CreateEnterpriseRouter(request *model.CreateEnterpriseRouterRequest) (*model.CreateEnterpriseRouterResponse, error) {
	requestDef := GenReqDefForCreateEnterpriseRouter()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateEnterpriseRouterResponse), nil
	}
}

// CreateEnterpriseRouterInvoker 创建企业路由器
func (c *ErClient) CreateEnterpriseRouterInvoker(request *model.CreateEnterpriseRouterRequest) *CreateEnterpriseRouterInvoker {
	requestDef := GenReqDefForCreateEnterpriseRouter()
	return &CreateEnterpriseRouterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteEnterpriseRouter 删除企业路由器
//
// 删除企业路由器。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ErClient) DeleteEnterpriseRouter(request *model.DeleteEnterpriseRouterRequest) (*model.DeleteEnterpriseRouterResponse, error) {
	requestDef := GenReqDefForDeleteEnterpriseRouter()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteEnterpriseRouterResponse), nil
	}
}

// DeleteEnterpriseRouterInvoker 删除企业路由器
func (c *ErClient) DeleteEnterpriseRouterInvoker(request *model.DeleteEnterpriseRouterRequest) *DeleteEnterpriseRouterInvoker {
	requestDef := GenReqDefForDeleteEnterpriseRouter()
	return &DeleteEnterpriseRouterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEnterpriseRouters 查询企业路由器列表
//
// 查询企业路由器列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ErClient) ListEnterpriseRouters(request *model.ListEnterpriseRoutersRequest) (*model.ListEnterpriseRoutersResponse, error) {
	requestDef := GenReqDefForListEnterpriseRouters()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEnterpriseRoutersResponse), nil
	}
}

// ListEnterpriseRoutersInvoker 查询企业路由器列表
func (c *ErClient) ListEnterpriseRoutersInvoker(request *model.ListEnterpriseRoutersRequest) *ListEnterpriseRoutersInvoker {
	requestDef := GenReqDefForListEnterpriseRouters()
	return &ListEnterpriseRoutersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowEnterpriseRouter 查询企业路由器详情
//
// 查询企业路由器详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ErClient) ShowEnterpriseRouter(request *model.ShowEnterpriseRouterRequest) (*model.ShowEnterpriseRouterResponse, error) {
	requestDef := GenReqDefForShowEnterpriseRouter()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEnterpriseRouterResponse), nil
	}
}

// ShowEnterpriseRouterInvoker 查询企业路由器详情
func (c *ErClient) ShowEnterpriseRouterInvoker(request *model.ShowEnterpriseRouterRequest) *ShowEnterpriseRouterInvoker {
	requestDef := GenReqDefForShowEnterpriseRouter()
	return &ShowEnterpriseRouterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateEnterpriseRouter 更新企业路由器
//
// 更新企业路由器基本信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ErClient) UpdateEnterpriseRouter(request *model.UpdateEnterpriseRouterRequest) (*model.UpdateEnterpriseRouterResponse, error) {
	requestDef := GenReqDefForUpdateEnterpriseRouter()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateEnterpriseRouterResponse), nil
	}
}

// UpdateEnterpriseRouterInvoker 更新企业路由器
func (c *ErClient) UpdateEnterpriseRouterInvoker(request *model.UpdateEnterpriseRouterRequest) *UpdateEnterpriseRouterInvoker {
	requestDef := GenReqDefForUpdateEnterpriseRouter()
	return &UpdateEnterpriseRouterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateFlowLog 创建流日志
//
// 给ER实例创建流日志。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ErClient) CreateFlowLog(request *model.CreateFlowLogRequest) (*model.CreateFlowLogResponse, error) {
	requestDef := GenReqDefForCreateFlowLog()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateFlowLogResponse), nil
	}
}

// CreateFlowLogInvoker 创建流日志
func (c *ErClient) CreateFlowLogInvoker(request *model.CreateFlowLogRequest) *CreateFlowLogInvoker {
	requestDef := GenReqDefForCreateFlowLog()
	return &CreateFlowLogInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteFlowLog 删除流日志
//
// 删除流日志
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ErClient) DeleteFlowLog(request *model.DeleteFlowLogRequest) (*model.DeleteFlowLogResponse, error) {
	requestDef := GenReqDefForDeleteFlowLog()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteFlowLogResponse), nil
	}
}

// DeleteFlowLogInvoker 删除流日志
func (c *ErClient) DeleteFlowLogInvoker(request *model.DeleteFlowLogRequest) *DeleteFlowLogInvoker {
	requestDef := GenReqDefForDeleteFlowLog()
	return &DeleteFlowLogInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DisableFlowLog 关闭流日志
//
// 关闭流日志
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ErClient) DisableFlowLog(request *model.DisableFlowLogRequest) (*model.DisableFlowLogResponse, error) {
	requestDef := GenReqDefForDisableFlowLog()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisableFlowLogResponse), nil
	}
}

// DisableFlowLogInvoker 关闭流日志
func (c *ErClient) DisableFlowLogInvoker(request *model.DisableFlowLogRequest) *DisableFlowLogInvoker {
	requestDef := GenReqDefForDisableFlowLog()
	return &DisableFlowLogInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// EnableFlowLog 开启流日志
//
// 开启流日志
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ErClient) EnableFlowLog(request *model.EnableFlowLogRequest) (*model.EnableFlowLogResponse, error) {
	requestDef := GenReqDefForEnableFlowLog()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.EnableFlowLogResponse), nil
	}
}

// EnableFlowLogInvoker 开启流日志
func (c *ErClient) EnableFlowLogInvoker(request *model.EnableFlowLogRequest) *EnableFlowLogInvoker {
	requestDef := GenReqDefForEnableFlowLog()
	return &EnableFlowLogInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFlowLogs 查询流日志列表
//
// 查询企业路由器实例下的流日志列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ErClient) ListFlowLogs(request *model.ListFlowLogsRequest) (*model.ListFlowLogsResponse, error) {
	requestDef := GenReqDefForListFlowLogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFlowLogsResponse), nil
	}
}

// ListFlowLogsInvoker 查询流日志列表
func (c *ErClient) ListFlowLogsInvoker(request *model.ListFlowLogsRequest) *ListFlowLogsInvoker {
	requestDef := GenReqDefForListFlowLogs()
	return &ListFlowLogsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowFlowLog 查询流日志详情
//
// 查询流日志详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ErClient) ShowFlowLog(request *model.ShowFlowLogRequest) (*model.ShowFlowLogResponse, error) {
	requestDef := GenReqDefForShowFlowLog()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowFlowLogResponse), nil
	}
}

// ShowFlowLogInvoker 查询流日志详情
func (c *ErClient) ShowFlowLogInvoker(request *model.ShowFlowLogRequest) *ShowFlowLogInvoker {
	requestDef := GenReqDefForShowFlowLog()
	return &ShowFlowLogInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateFlowLog 更新流日志基本信息
//
// 更新流日志基本信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ErClient) UpdateFlowLog(request *model.UpdateFlowLogRequest) (*model.UpdateFlowLogResponse, error) {
	requestDef := GenReqDefForUpdateFlowLog()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateFlowLogResponse), nil
	}
}

// UpdateFlowLogInvoker 更新流日志基本信息
func (c *ErClient) UpdateFlowLogInvoker(request *model.UpdateFlowLogRequest) *UpdateFlowLogInvoker {
	requestDef := GenReqDefForUpdateFlowLog()
	return &UpdateFlowLogInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DisablePropagation 删除路由传播
//
// 解绑连接和路由表的传播关系。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ErClient) DisablePropagation(request *model.DisablePropagationRequest) (*model.DisablePropagationResponse, error) {
	requestDef := GenReqDefForDisablePropagation()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisablePropagationResponse), nil
	}
}

// DisablePropagationInvoker 删除路由传播
func (c *ErClient) DisablePropagationInvoker(request *model.DisablePropagationRequest) *DisablePropagationInvoker {
	requestDef := GenReqDefForDisablePropagation()
	return &DisablePropagationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// EnablePropagation 创建路由传播
//
// 每个连接可以和多个路由表建立传播关系，从该连接学习到的路由会应用到具有传播关系的路由表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ErClient) EnablePropagation(request *model.EnablePropagationRequest) (*model.EnablePropagationResponse, error) {
	requestDef := GenReqDefForEnablePropagation()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.EnablePropagationResponse), nil
	}
}

// EnablePropagationInvoker 创建路由传播
func (c *ErClient) EnablePropagationInvoker(request *model.EnablePropagationRequest) *EnablePropagationInvoker {
	requestDef := GenReqDefForEnablePropagation()
	return &EnablePropagationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPropagations 查询路由传播列表
//
// 查询路由传播列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ErClient) ListPropagations(request *model.ListPropagationsRequest) (*model.ListPropagationsResponse, error) {
	requestDef := GenReqDefForListPropagations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPropagationsResponse), nil
	}
}

// ListPropagationsInvoker 查询路由传播列表
func (c *ErClient) ListPropagationsInvoker(request *model.ListPropagationsRequest) *ListPropagationsInvoker {
	requestDef := GenReqDefForListPropagations()
	return &ListPropagationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowQuotas 查询配额
//
// 查询租户各类资源的使用情况，如企业路由器的使用量，VPC连接的使用量等。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ErClient) ShowQuotas(request *model.ShowQuotasRequest) (*model.ShowQuotasResponse, error) {
	requestDef := GenReqDefForShowQuotas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowQuotasResponse), nil
	}
}

// ShowQuotasInvoker 查询配额
func (c *ErClient) ShowQuotasInvoker(request *model.ShowQuotasRequest) *ShowQuotasInvoker {
	requestDef := GenReqDefForShowQuotas()
	return &ShowQuotasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateStaticRoute 创建静态路由
//
// 创建静态路由
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ErClient) CreateStaticRoute(request *model.CreateStaticRouteRequest) (*model.CreateStaticRouteResponse, error) {
	requestDef := GenReqDefForCreateStaticRoute()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateStaticRouteResponse), nil
	}
}

// CreateStaticRouteInvoker 创建静态路由
func (c *ErClient) CreateStaticRouteInvoker(request *model.CreateStaticRouteRequest) *CreateStaticRouteInvoker {
	requestDef := GenReqDefForCreateStaticRoute()
	return &CreateStaticRouteInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteStaticRoute 删除静态路由
//
// 删除静态路由
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ErClient) DeleteStaticRoute(request *model.DeleteStaticRouteRequest) (*model.DeleteStaticRouteResponse, error) {
	requestDef := GenReqDefForDeleteStaticRoute()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteStaticRouteResponse), nil
	}
}

// DeleteStaticRouteInvoker 删除静态路由
func (c *ErClient) DeleteStaticRouteInvoker(request *model.DeleteStaticRouteRequest) *DeleteStaticRouteInvoker {
	requestDef := GenReqDefForDeleteStaticRoute()
	return &DeleteStaticRouteInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEffectiveRoutes 查询有效路由列表
//
// 查询有效的路由列表，支持分页查询能力
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ErClient) ListEffectiveRoutes(request *model.ListEffectiveRoutesRequest) (*model.ListEffectiveRoutesResponse, error) {
	requestDef := GenReqDefForListEffectiveRoutes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEffectiveRoutesResponse), nil
	}
}

// ListEffectiveRoutesInvoker 查询有效路由列表
func (c *ErClient) ListEffectiveRoutesInvoker(request *model.ListEffectiveRoutesRequest) *ListEffectiveRoutesInvoker {
	requestDef := GenReqDefForListEffectiveRoutes()
	return &ListEffectiveRoutesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListStaticRoutes 查询静态路由列表
//
// 查询静态路由列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ErClient) ListStaticRoutes(request *model.ListStaticRoutesRequest) (*model.ListStaticRoutesResponse, error) {
	requestDef := GenReqDefForListStaticRoutes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListStaticRoutesResponse), nil
	}
}

// ListStaticRoutesInvoker 查询静态路由列表
func (c *ErClient) ListStaticRoutesInvoker(request *model.ListStaticRoutesRequest) *ListStaticRoutesInvoker {
	requestDef := GenReqDefForListStaticRoutes()
	return &ListStaticRoutesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowStaticRoute 查询静态路由详情
//
// 查询静态路由详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ErClient) ShowStaticRoute(request *model.ShowStaticRouteRequest) (*model.ShowStaticRouteResponse, error) {
	requestDef := GenReqDefForShowStaticRoute()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowStaticRouteResponse), nil
	}
}

// ShowStaticRouteInvoker 查询静态路由详情
func (c *ErClient) ShowStaticRouteInvoker(request *model.ShowStaticRouteRequest) *ShowStaticRouteInvoker {
	requestDef := GenReqDefForShowStaticRoute()
	return &ShowStaticRouteInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateStaticRoute 更新静态路由
//
// 更新静态路由
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ErClient) UpdateStaticRoute(request *model.UpdateStaticRouteRequest) (*model.UpdateStaticRouteResponse, error) {
	requestDef := GenReqDefForUpdateStaticRoute()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateStaticRouteResponse), nil
	}
}

// UpdateStaticRouteInvoker 更新静态路由
func (c *ErClient) UpdateStaticRouteInvoker(request *model.UpdateStaticRouteRequest) *UpdateStaticRouteInvoker {
	requestDef := GenReqDefForUpdateStaticRoute()
	return &UpdateStaticRouteInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateRouteTable 创建路由表
//
// 路由表是企业路由器收发报文的依据，包含了连接的关联关系，传播关系以及路由信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ErClient) CreateRouteTable(request *model.CreateRouteTableRequest) (*model.CreateRouteTableResponse, error) {
	requestDef := GenReqDefForCreateRouteTable()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRouteTableResponse), nil
	}
}

// CreateRouteTableInvoker 创建路由表
func (c *ErClient) CreateRouteTableInvoker(request *model.CreateRouteTableRequest) *CreateRouteTableInvoker {
	requestDef := GenReqDefForCreateRouteTable()
	return &CreateRouteTableInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteRouteTable 删除路由表
//
// 删除路由表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ErClient) DeleteRouteTable(request *model.DeleteRouteTableRequest) (*model.DeleteRouteTableResponse, error) {
	requestDef := GenReqDefForDeleteRouteTable()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteRouteTableResponse), nil
	}
}

// DeleteRouteTableInvoker 删除路由表
func (c *ErClient) DeleteRouteTableInvoker(request *model.DeleteRouteTableRequest) *DeleteRouteTableInvoker {
	requestDef := GenReqDefForDeleteRouteTable()
	return &DeleteRouteTableInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRouteTables 查询路由表列表
//
// 查询路由表列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ErClient) ListRouteTables(request *model.ListRouteTablesRequest) (*model.ListRouteTablesResponse, error) {
	requestDef := GenReqDefForListRouteTables()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRouteTablesResponse), nil
	}
}

// ListRouteTablesInvoker 查询路由表列表
func (c *ErClient) ListRouteTablesInvoker(request *model.ListRouteTablesRequest) *ListRouteTablesInvoker {
	requestDef := GenReqDefForListRouteTables()
	return &ListRouteTablesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRouteTable 查询路由表详情
//
// 查询路由表详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ErClient) ShowRouteTable(request *model.ShowRouteTableRequest) (*model.ShowRouteTableResponse, error) {
	requestDef := GenReqDefForShowRouteTable()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRouteTableResponse), nil
	}
}

// ShowRouteTableInvoker 查询路由表详情
func (c *ErClient) ShowRouteTableInvoker(request *model.ShowRouteTableRequest) *ShowRouteTableInvoker {
	requestDef := GenReqDefForShowRouteTable()
	return &ShowRouteTableInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateRouteTable 更新路由表信息
//
// 更新路由表基本信息，如名称，描述等。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ErClient) UpdateRouteTable(request *model.UpdateRouteTableRequest) (*model.UpdateRouteTableResponse, error) {
	requestDef := GenReqDefForUpdateRouteTable()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateRouteTableResponse), nil
	}
}

// UpdateRouteTableInvoker 更新路由表信息
func (c *ErClient) UpdateRouteTableInvoker(request *model.UpdateRouteTableRequest) *UpdateRouteTableInvoker {
	requestDef := GenReqDefForUpdateRouteTable()
	return &UpdateRouteTableInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchCreateResourceTags 批量添加删除资源标签
//
// - 为指定实例批量添加或删除标签
// - 标签管理服务需要使用该接口批量管理实例的标签。
// - 一个资源上最多有10个标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ErClient) BatchCreateResourceTags(request *model.BatchCreateResourceTagsRequest) (*model.BatchCreateResourceTagsResponse, error) {
	requestDef := GenReqDefForBatchCreateResourceTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateResourceTagsResponse), nil
	}
}

// BatchCreateResourceTagsInvoker 批量添加删除资源标签
func (c *ErClient) BatchCreateResourceTagsInvoker(request *model.BatchCreateResourceTagsRequest) *BatchCreateResourceTagsInvoker {
	requestDef := GenReqDefForBatchCreateResourceTags()
	return &BatchCreateResourceTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateResourceTag 创建资源标签
//
// 为特定类型的资源创建标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ErClient) CreateResourceTag(request *model.CreateResourceTagRequest) (*model.CreateResourceTagResponse, error) {
	requestDef := GenReqDefForCreateResourceTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateResourceTagResponse), nil
	}
}

// CreateResourceTagInvoker 创建资源标签
func (c *ErClient) CreateResourceTagInvoker(request *model.CreateResourceTagRequest) *CreateResourceTagInvoker {
	requestDef := GenReqDefForCreateResourceTag()
	return &CreateResourceTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteResourceTag 删除资源标签
//
// 删除特定类型资源的标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ErClient) DeleteResourceTag(request *model.DeleteResourceTagRequest) (*model.DeleteResourceTagResponse, error) {
	requestDef := GenReqDefForDeleteResourceTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteResourceTagResponse), nil
	}
}

// DeleteResourceTagInvoker 删除资源标签
func (c *ErClient) DeleteResourceTagInvoker(request *model.DeleteResourceTagRequest) *DeleteResourceTagInvoker {
	requestDef := GenReqDefForDeleteResourceTag()
	return &DeleteResourceTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProjectTags 查询项目标签
//
// 查询特定类型资源的标签集合。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ErClient) ListProjectTags(request *model.ListProjectTagsRequest) (*model.ListProjectTagsResponse, error) {
	requestDef := GenReqDefForListProjectTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProjectTagsResponse), nil
	}
}

// ListProjectTagsInvoker 查询项目标签
func (c *ErClient) ListProjectTagsInvoker(request *model.ListProjectTagsRequest) *ListProjectTagsInvoker {
	requestDef := GenReqDefForListProjectTags()
	return &ListProjectTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowResourceTag 查询资源标签
//
// 查询特定类型资源的标签信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ErClient) ShowResourceTag(request *model.ShowResourceTagRequest) (*model.ShowResourceTagResponse, error) {
	requestDef := GenReqDefForShowResourceTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowResourceTagResponse), nil
	}
}

// ShowResourceTagInvoker 查询资源标签
func (c *ErClient) ShowResourceTagInvoker(request *model.ShowResourceTagRequest) *ShowResourceTagInvoker {
	requestDef := GenReqDefForShowResourceTag()
	return &ShowResourceTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateVpcAttachment 创建VPC连接
//
// 给ER实例创建VPC连接。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ErClient) CreateVpcAttachment(request *model.CreateVpcAttachmentRequest) (*model.CreateVpcAttachmentResponse, error) {
	requestDef := GenReqDefForCreateVpcAttachment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateVpcAttachmentResponse), nil
	}
}

// CreateVpcAttachmentInvoker 创建VPC连接
func (c *ErClient) CreateVpcAttachmentInvoker(request *model.CreateVpcAttachmentRequest) *CreateVpcAttachmentInvoker {
	requestDef := GenReqDefForCreateVpcAttachment()
	return &CreateVpcAttachmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteVpcAttachment 删除VPC连接
//
// 删除VPC连接。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ErClient) DeleteVpcAttachment(request *model.DeleteVpcAttachmentRequest) (*model.DeleteVpcAttachmentResponse, error) {
	requestDef := GenReqDefForDeleteVpcAttachment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteVpcAttachmentResponse), nil
	}
}

// DeleteVpcAttachmentInvoker 删除VPC连接
func (c *ErClient) DeleteVpcAttachmentInvoker(request *model.DeleteVpcAttachmentRequest) *DeleteVpcAttachmentInvoker {
	requestDef := GenReqDefForDeleteVpcAttachment()
	return &DeleteVpcAttachmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListVpcAttachments 查询VPC连接列表
//
// 查询企业路由器实例下的VPC连接列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ErClient) ListVpcAttachments(request *model.ListVpcAttachmentsRequest) (*model.ListVpcAttachmentsResponse, error) {
	requestDef := GenReqDefForListVpcAttachments()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVpcAttachmentsResponse), nil
	}
}

// ListVpcAttachmentsInvoker 查询VPC连接列表
func (c *ErClient) ListVpcAttachmentsInvoker(request *model.ListVpcAttachmentsRequest) *ListVpcAttachmentsInvoker {
	requestDef := GenReqDefForListVpcAttachments()
	return &ListVpcAttachmentsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowVpcAttachment 查询VPC连接详情
//
// 查询VPC连接详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ErClient) ShowVpcAttachment(request *model.ShowVpcAttachmentRequest) (*model.ShowVpcAttachmentResponse, error) {
	requestDef := GenReqDefForShowVpcAttachment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVpcAttachmentResponse), nil
	}
}

// ShowVpcAttachmentInvoker 查询VPC连接详情
func (c *ErClient) ShowVpcAttachmentInvoker(request *model.ShowVpcAttachmentRequest) *ShowVpcAttachmentInvoker {
	requestDef := GenReqDefForShowVpcAttachment()
	return &ShowVpcAttachmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateVpcAttachment 更新VPC连接基本信息
//
// 修改VPC连接基本信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ErClient) UpdateVpcAttachment(request *model.UpdateVpcAttachmentRequest) (*model.UpdateVpcAttachmentResponse, error) {
	requestDef := GenReqDefForUpdateVpcAttachment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateVpcAttachmentResponse), nil
	}
}

// UpdateVpcAttachmentInvoker 更新VPC连接基本信息
func (c *ErClient) UpdateVpcAttachmentInvoker(request *model.UpdateVpcAttachmentRequest) *UpdateVpcAttachmentInvoker {
	requestDef := GenReqDefForUpdateVpcAttachment()
	return &UpdateVpcAttachmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
