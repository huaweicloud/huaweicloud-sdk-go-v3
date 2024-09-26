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

// ListControlsForOrganizationalUnit 列出注册OU下开启的控制策略
//
// 列出组织里某个注册OU开启的所有控制策略信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) ListControlsForOrganizationalUnit(request *model.ListControlsForOrganizationalUnitRequest) (*model.ListControlsForOrganizationalUnitResponse, error) {
	requestDef := GenReqDefForListControlsForOrganizationalUnit()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListControlsForOrganizationalUnitResponse), nil
	}
}

// ListControlsForOrganizationalUnitInvoker 列出注册OU下开启的控制策略
func (c *RgcClient) ListControlsForOrganizationalUnitInvoker(request *model.ListControlsForOrganizationalUnitRequest) *ListControlsForOrganizationalUnitInvoker {
	requestDef := GenReqDefForListControlsForOrganizationalUnit()
	return &ListControlsForOrganizationalUnitInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEnabledControls 列出开启的控制策略
//
// 列出组织里开启的所有控制策略信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) ListEnabledControls(request *model.ListEnabledControlsRequest) (*model.ListEnabledControlsResponse, error) {
	requestDef := GenReqDefForListEnabledControls()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEnabledControlsResponse), nil
	}
}

// ListEnabledControlsInvoker 列出开启的控制策略
func (c *RgcClient) ListEnabledControlsInvoker(request *model.ListEnabledControlsRequest) *ListEnabledControlsInvoker {
	requestDef := GenReqDefForListEnabledControls()
	return &ListEnabledControlsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// CreateAccount 创建账号
//
// 在组织里的某个注册OU下创建账号。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) CreateAccount(request *model.CreateAccountRequest) (*model.CreateAccountResponse, error) {
	requestDef := GenReqDefForCreateAccount()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAccountResponse), nil
	}
}

// CreateAccountInvoker 创建账号
func (c *RgcClient) CreateAccountInvoker(request *model.CreateAccountRequest) *CreateAccountInvoker {
	requestDef := GenReqDefForCreateAccount()
	return &CreateAccountInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RegisterOrganizationalUnit 注册OU
//
// 将组织里的某个OU注册到RGC服务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) RegisterOrganizationalUnit(request *model.RegisterOrganizationalUnitRequest) (*model.RegisterOrganizationalUnitResponse, error) {
	requestDef := GenReqDefForRegisterOrganizationalUnit()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RegisterOrganizationalUnitResponse), nil
	}
}

// RegisterOrganizationalUnitInvoker 注册OU
func (c *RgcClient) RegisterOrganizationalUnitInvoker(request *model.RegisterOrganizationalUnitRequest) *RegisterOrganizationalUnitInvoker {
	requestDef := GenReqDefForRegisterOrganizationalUnit()
	return &RegisterOrganizationalUnitInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowManagedAccount 查询纳管账号信息
//
// 查询组织里某个纳管账号信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) ShowManagedAccount(request *model.ShowManagedAccountRequest) (*model.ShowManagedAccountResponse, error) {
	requestDef := GenReqDefForShowManagedAccount()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowManagedAccountResponse), nil
	}
}

// ShowManagedAccountInvoker 查询纳管账号信息
func (c *RgcClient) ShowManagedAccountInvoker(request *model.ShowManagedAccountRequest) *ShowManagedAccountInvoker {
	requestDef := GenReqDefForShowManagedAccount()
	return &ShowManagedAccountInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowOperation 查询注册过程信息
//
// 查询在RGC服务里注册/取消注册的过程信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) ShowOperation(request *model.ShowOperationRequest) (*model.ShowOperationResponse, error) {
	requestDef := GenReqDefForShowOperation()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowOperationResponse), nil
	}
}

// ShowOperationInvoker 查询注册过程信息
func (c *RgcClient) ShowOperationInvoker(request *model.ShowOperationRequest) *ShowOperationInvoker {
	requestDef := GenReqDefForShowOperation()
	return &ShowOperationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
