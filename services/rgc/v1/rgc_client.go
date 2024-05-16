package v1

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/rgc/v1/model"
)

type RgcClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewRgcClient(hcClient *httpclient.HcHttpClient) *RgcClient {
	return &RgcClient{HcClient: hcClient}
}

func RgcClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// DisableControl 关闭控制策略
//
// 关闭组织下的某个单元的某个控制策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) DisableControl(request *model.DisableControlRequest) (*model.DisableControlResponse, error) {
	requestDef := GenReqDefForDisableControl()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisableControlResponse), nil
	}
}

// DisableControlInvoker 关闭控制策略
func (c *RgcClient) DisableControlInvoker(request *model.DisableControlRequest) *DisableControlInvoker {
	requestDef := GenReqDefForDisableControl()
	return &DisableControlInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// EnableControl 开启控制策略
//
// 给组织下的某个单元开启某个控制策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) EnableControl(request *model.EnableControlRequest) (*model.EnableControlResponse, error) {
	requestDef := GenReqDefForEnableControl()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.EnableControlResponse), nil
	}
}

// EnableControlInvoker 开启控制策略
func (c *RgcClient) EnableControlInvoker(request *model.EnableControlRequest) *EnableControlInvoker {
	requestDef := GenReqDefForEnableControl()
	return &EnableControlInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListControlsForOrganizationUnit 列出注册OU下开启的控制策略
//
// 列出组织里某个注册OU开启的所有控制策略信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) ListControlsForOrganizationUnit(request *model.ListControlsForOrganizationUnitRequest) (*model.ListControlsForOrganizationUnitResponse, error) {
	requestDef := GenReqDefForListControlsForOrganizationUnit()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListControlsForOrganizationUnitResponse), nil
	}
}

// ListControlsForOrganizationUnitInvoker 列出注册OU下开启的控制策略
func (c *RgcClient) ListControlsForOrganizationUnitInvoker(request *model.ListControlsForOrganizationUnitRequest) *ListControlsForOrganizationUnitInvoker {
	requestDef := GenReqDefForListControlsForOrganizationUnit()
	return &ListControlsForOrganizationUnitInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowControlOperate 查询控制策略操作状态
//
// 根据操作ID查询返回指定ID的操作状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) ShowControlOperate(request *model.ShowControlOperateRequest) (*model.ShowControlOperateResponse, error) {
	requestDef := GenReqDefForShowControlOperate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowControlOperateResponse), nil
	}
}

// ShowControlOperateInvoker 查询控制策略操作状态
func (c *RgcClient) ShowControlOperateInvoker(request *model.ShowControlOperateRequest) *ShowControlOperateInvoker {
	requestDef := GenReqDefForShowControlOperate()
	return &ShowControlOperateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
