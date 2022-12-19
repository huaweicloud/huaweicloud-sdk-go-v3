package v1

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cbh/v1/model"
)

type CbhClient struct {
	HcClient *http_client.HcHttpClient
}

func NewCbhClient(hcClient *http_client.HcHttpClient) *CbhClient {
	return &CbhClient{HcClient: hcClient}
}

func CbhClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// ChangeInstanceNetwork 修改CBH实例网络
//
// 修改CBH实例网络
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CbhClient) ChangeInstanceNetwork(request *model.ChangeInstanceNetworkRequest) (*model.ChangeInstanceNetworkResponse, error) {
	requestDef := GenReqDefForChangeInstanceNetwork()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeInstanceNetworkResponse), nil
	}
}

// ChangeInstanceNetworkInvoker 修改CBH实例网络
func (c *CbhClient) ChangeInstanceNetworkInvoker(request *model.ChangeInstanceNetworkRequest) *ChangeInstanceNetworkInvoker {
	requestDef := GenReqDefForChangeInstanceNetwork()
	return &ChangeInstanceNetworkInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateInstance 创建CBH实例
//
// 创建CBH实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CbhClient) CreateInstance(request *model.CreateInstanceRequest) (*model.CreateInstanceResponse, error) {
	requestDef := GenReqDefForCreateInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateInstanceResponse), nil
	}
}

// CreateInstanceInvoker 创建CBH实例
func (c *CbhClient) CreateInstanceInvoker(request *model.CreateInstanceRequest) *CreateInstanceInvoker {
	requestDef := GenReqDefForCreateInstance()
	return &CreateInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateInstanceOrder 创建CBH实例订单
//
// 创建CBH实例订单
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CbhClient) CreateInstanceOrder(request *model.CreateInstanceOrderRequest) (*model.CreateInstanceOrderResponse, error) {
	requestDef := GenReqDefForCreateInstanceOrder()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateInstanceOrderResponse), nil
	}
}

// CreateInstanceOrderInvoker 创建CBH实例订单
func (c *CbhClient) CreateInstanceOrderInvoker(request *model.CreateInstanceOrderRequest) *CreateInstanceOrderInvoker {
	requestDef := GenReqDefForCreateInstanceOrder()
	return &CreateInstanceOrderInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCbhInstance 获取CBH实例列表
//
// 获取CBH实例列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CbhClient) ListCbhInstance(request *model.ListCbhInstanceRequest) (*model.ListCbhInstanceResponse, error) {
	requestDef := GenReqDefForListCbhInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCbhInstanceResponse), nil
	}
}

// ListCbhInstanceInvoker 获取CBH实例列表
func (c *CbhClient) ListCbhInstanceInvoker(request *model.ListCbhInstanceRequest) *ListCbhInstanceInvoker {
	requestDef := GenReqDefForListCbhInstance()
	return &ListCbhInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RestartCbhInstance 重启CBH实例
//
// 重启CBH实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CbhClient) RestartCbhInstance(request *model.RestartCbhInstanceRequest) (*model.RestartCbhInstanceResponse, error) {
	requestDef := GenReqDefForRestartCbhInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RestartCbhInstanceResponse), nil
	}
}

// RestartCbhInstanceInvoker 重启CBH实例
func (c *CbhClient) RestartCbhInstanceInvoker(request *model.RestartCbhInstanceRequest) *RestartCbhInstanceInvoker {
	requestDef := GenReqDefForRestartCbhInstance()
	return &RestartCbhInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchQuota 查询堡垒机配额
//
// 查询堡垒机配额
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CbhClient) SearchQuota(request *model.SearchQuotaRequest) (*model.SearchQuotaResponse, error) {
	requestDef := GenReqDefForSearchQuota()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchQuotaResponse), nil
	}
}

// SearchQuotaInvoker 查询堡垒机配额
func (c *CbhClient) SearchQuotaInvoker(request *model.SearchQuotaRequest) *SearchQuotaInvoker {
	requestDef := GenReqDefForSearchQuota()
	return &SearchQuotaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAvailableZoneInfo 获取CBH服务可用分区信息
//
// 获取CBH服务可用分区信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CbhClient) ShowAvailableZoneInfo(request *model.ShowAvailableZoneInfoRequest) (*model.ShowAvailableZoneInfoResponse, error) {
	requestDef := GenReqDefForShowAvailableZoneInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAvailableZoneInfoResponse), nil
	}
}

// ShowAvailableZoneInfoInvoker 获取CBH服务可用分区信息
func (c *CbhClient) ShowAvailableZoneInfoInvoker(request *model.ShowAvailableZoneInfoRequest) *ShowAvailableZoneInfoInvoker {
	requestDef := GenReqDefForShowAvailableZoneInfo()
	return &ShowAvailableZoneInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowNetworkConfiguration 检查网络接口
//
// 检查网络接口
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CbhClient) ShowNetworkConfiguration(request *model.ShowNetworkConfigurationRequest) (*model.ShowNetworkConfigurationResponse, error) {
	requestDef := GenReqDefForShowNetworkConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowNetworkConfigurationResponse), nil
	}
}

// ShowNetworkConfigurationInvoker 检查网络接口
func (c *CbhClient) ShowNetworkConfigurationInvoker(request *model.ShowNetworkConfigurationRequest) *ShowNetworkConfigurationInvoker {
	requestDef := GenReqDefForShowNetworkConfiguration()
	return &ShowNetworkConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StartCbhInstance 启动CBH实例
//
// 启动CBH实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CbhClient) StartCbhInstance(request *model.StartCbhInstanceRequest) (*model.StartCbhInstanceResponse, error) {
	requestDef := GenReqDefForStartCbhInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartCbhInstanceResponse), nil
	}
}

// StartCbhInstanceInvoker 启动CBH实例
func (c *CbhClient) StartCbhInstanceInvoker(request *model.StartCbhInstanceRequest) *StartCbhInstanceInvoker {
	requestDef := GenReqDefForStartCbhInstance()
	return &StartCbhInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpgradeCbhInstance 升级CBH实例
//
// 升级CBH实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CbhClient) UpgradeCbhInstance(request *model.UpgradeCbhInstanceRequest) (*model.UpgradeCbhInstanceResponse, error) {
	requestDef := GenReqDefForUpgradeCbhInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpgradeCbhInstanceResponse), nil
	}
}

// UpgradeCbhInstanceInvoker 升级CBH实例
func (c *CbhClient) UpgradeCbhInstanceInvoker(request *model.UpgradeCbhInstanceRequest) *UpgradeCbhInstanceInvoker {
	requestDef := GenReqDefForUpgradeCbhInstance()
	return &UpgradeCbhInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResetLoginMethod 修改admin用户多因子认证方式
//
// 修改admin用户多因子认证方式
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CbhClient) ResetLoginMethod(request *model.ResetLoginMethodRequest) (*model.ResetLoginMethodResponse, error) {
	requestDef := GenReqDefForResetLoginMethod()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResetLoginMethodResponse), nil
	}
}

// ResetLoginMethodInvoker 修改admin用户多因子认证方式
func (c *CbhClient) ResetLoginMethodInvoker(request *model.ResetLoginMethodRequest) *ResetLoginMethodInvoker {
	requestDef := GenReqDefForResetLoginMethod()
	return &ResetLoginMethodInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResetPassword 修改admin密码
//
// 修改admin密码
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CbhClient) ResetPassword(request *model.ResetPasswordRequest) (*model.ResetPasswordResponse, error) {
	requestDef := GenReqDefForResetPassword()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResetPasswordResponse), nil
	}
}

// ResetPasswordInvoker 修改admin密码
func (c *CbhClient) ResetPasswordInvoker(request *model.ResetPasswordRequest) *ResetPasswordInvoker {
	requestDef := GenReqDefForResetPassword()
	return &ResetPasswordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeInstanceOrder 创建变更CBH订单
//
// 创建变更CBH订单
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CbhClient) ChangeInstanceOrder(request *model.ChangeInstanceOrderRequest) (*model.ChangeInstanceOrderResponse, error) {
	requestDef := GenReqDefForChangeInstanceOrder()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeInstanceOrderResponse), nil
	}
}

// ChangeInstanceOrderInvoker 创建变更CBH订单
func (c *CbhClient) ChangeInstanceOrderInvoker(request *model.ChangeInstanceOrderRequest) *ChangeInstanceOrderInvoker {
	requestDef := GenReqDefForChangeInstanceOrder()
	return &ChangeInstanceOrderInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StopCbhInstance 关闭CBH实例
//
// 关闭CBH实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CbhClient) StopCbhInstance(request *model.StopCbhInstanceRequest) (*model.StopCbhInstanceResponse, error) {
	requestDef := GenReqDefForStopCbhInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopCbhInstanceResponse), nil
	}
}

// StopCbhInstanceInvoker 关闭CBH实例
func (c *CbhClient) StopCbhInstanceInvoker(request *model.StopCbhInstanceRequest) *StopCbhInstanceInvoker {
	requestDef := GenReqDefForStopCbhInstance()
	return &StopCbhInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListQuotaState 获取ECS配额
//
// 获取当前租户所选择的可用分区、性能规格所对应的ECS flavor是否可用
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CbhClient) ListQuotaState(request *model.ListQuotaStateRequest) (*model.ListQuotaStateResponse, error) {
	requestDef := GenReqDefForListQuotaState()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListQuotaStateResponse), nil
	}
}

// ListQuotaStateInvoker 获取ECS配额
func (c *CbhClient) ListQuotaStateInvoker(request *model.ListQuotaStateRequest) *ListQuotaStateInvoker {
	requestDef := GenReqDefForListQuotaState()
	return &ListQuotaStateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// InstallInstanceEip 绑定CBH实例Eip
//
// 绑定CBH实例Eip
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CbhClient) InstallInstanceEip(request *model.InstallInstanceEipRequest) (*model.InstallInstanceEipResponse, error) {
	requestDef := GenReqDefForInstallInstanceEip()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.InstallInstanceEipResponse), nil
	}
}

// InstallInstanceEipInvoker 绑定CBH实例Eip
func (c *CbhClient) InstallInstanceEipInvoker(request *model.InstallInstanceEipRequest) *InstallInstanceEipInvoker {
	requestDef := GenReqDefForInstallInstanceEip()
	return &InstallInstanceEipInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UninstallInstanceEip 解绑CBH实例Eip
//
// 解绑CBH实例Eip
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CbhClient) UninstallInstanceEip(request *model.UninstallInstanceEipRequest) (*model.UninstallInstanceEipResponse, error) {
	requestDef := GenReqDefForUninstallInstanceEip()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UninstallInstanceEipResponse), nil
	}
}

// UninstallInstanceEipInvoker 解绑CBH实例Eip
func (c *CbhClient) UninstallInstanceEipInvoker(request *model.UninstallInstanceEipRequest) *UninstallInstanceEipInvoker {
	requestDef := GenReqDefForUninstallInstanceEip()
	return &UninstallInstanceEipInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
