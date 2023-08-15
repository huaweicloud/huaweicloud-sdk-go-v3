package v1

import (
	http_client "github.com/dysodeng/huaweicloud-sdk-go-v3/core"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/services/cbh/v1/model"
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

// ChangeInstanceNetwork 修改实例网络
//
// 修改云堡垒机实例网络。
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

// ChangeInstanceNetworkInvoker 修改实例网络
func (c *CbhClient) ChangeInstanceNetworkInvoker(request *model.ChangeInstanceNetworkRequest) *ChangeInstanceNetworkInvoker {
	requestDef := GenReqDefForChangeInstanceNetwork()
	return &ChangeInstanceNetworkInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeInstanceOrder 创建变更云堡垒机实例订单
//
// 创建变更云堡垒机实例订单。（调用此接口前先调用创建变更云堡垒机实例任务接口，当前接口未开放）
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

// ChangeInstanceOrderInvoker 创建变更云堡垒机实例订单
func (c *CbhClient) ChangeInstanceOrderInvoker(request *model.ChangeInstanceOrderRequest) *ChangeInstanceOrderInvoker {
	requestDef := GenReqDefForChangeInstanceOrder()
	return &ChangeInstanceOrderInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateInstance 创建云堡垒机实例
//
// 创建云堡垒机实例。（创建云堡垒机实例订单前，先调用此接口）
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

// CreateInstanceInvoker 创建云堡垒机实例
func (c *CbhClient) CreateInstanceInvoker(request *model.CreateInstanceRequest) *CreateInstanceInvoker {
	requestDef := GenReqDefForCreateInstance()
	return &CreateInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateInstanceOrder 创建云堡垒机实例订单
//
// 创建云堡垒机实例订单。(调用此接口前先调用创建云堡垒机实例接口)
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

// CreateInstanceOrderInvoker 创建云堡垒机实例订单
func (c *CbhClient) CreateInstanceOrderInvoker(request *model.CreateInstanceOrderRequest) *CreateInstanceOrderInvoker {
	requestDef := GenReqDefForCreateInstanceOrder()
	return &CreateInstanceOrderInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// InstallInstanceEip 绑定弹性公网IP
//
// 云堡垒机实例绑定弹性公网IP
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

// InstallInstanceEipInvoker 绑定弹性公网IP
func (c *CbhClient) InstallInstanceEipInvoker(request *model.InstallInstanceEipRequest) *InstallInstanceEipInvoker {
	requestDef := GenReqDefForInstallInstanceEip()
	return &InstallInstanceEipInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCbhInstance 获取CBH实例列表
//
// 获取当前租户下的云堡垒机实例列表。
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

// ListQuotaStatus 获取弹性云服务器配额
//
// 获取当前租户所选择的可用分区、性能规格所对应的弹性云服务器是否可用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CbhClient) ListQuotaStatus(request *model.ListQuotaStatusRequest) (*model.ListQuotaStatusResponse, error) {
	requestDef := GenReqDefForListQuotaStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListQuotaStatusResponse), nil
	}
}

// ListQuotaStatusInvoker 获取弹性云服务器配额
func (c *CbhClient) ListQuotaStatusInvoker(request *model.ListQuotaStatusRequest) *ListQuotaStatusInvoker {
	requestDef := GenReqDefForListQuotaStatus()
	return &ListQuotaStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResetLoginMethod 重置admin用户多因子认证方式
//
// 重置admin用户多因子认证方式。
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

// ResetLoginMethodInvoker 重置admin用户多因子认证方式
func (c *CbhClient) ResetLoginMethodInvoker(request *model.ResetLoginMethodRequest) *ResetLoginMethodInvoker {
	requestDef := GenReqDefForResetLoginMethod()
	return &ResetLoginMethodInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResetPassword 修改admin用户密码
//
// 修改云堡垒机实例web登录admin用户密码。
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

// ResetPasswordInvoker 修改admin用户密码
func (c *CbhClient) ResetPasswordInvoker(request *model.ResetPasswordRequest) *ResetPasswordInvoker {
	requestDef := GenReqDefForResetPassword()
	return &ResetPasswordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RestartCbhInstance 重启云堡垒机实例
//
// 重启云堡垒机实例。
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

// RestartCbhInstanceInvoker 重启云堡垒机实例
func (c *CbhClient) RestartCbhInstanceInvoker(request *model.RestartCbhInstanceRequest) *RestartCbhInstanceInvoker {
	requestDef := GenReqDefForRestartCbhInstance()
	return &RestartCbhInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchQuota 查询堡垒机配额
//
// 查询云堡垒机配额信息。
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

// ShowAvailableZoneInfo 获取可用用分区信息
//
// 获取云堡垒机服务可用分区信息。
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

// ShowAvailableZoneInfoInvoker 获取可用用分区信息
func (c *CbhClient) ShowAvailableZoneInfoInvoker(request *model.ShowAvailableZoneInfoRequest) *ShowAvailableZoneInfoInvoker {
	requestDef := GenReqDefForShowAvailableZoneInfo()
	return &ShowAvailableZoneInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowNetworkConfiguration 检查云堡垒机网络
//
// 检查云堡垒机实例网络信息。
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

// ShowNetworkConfigurationInvoker 检查云堡垒机网络
func (c *CbhClient) ShowNetworkConfigurationInvoker(request *model.ShowNetworkConfigurationRequest) *ShowNetworkConfigurationInvoker {
	requestDef := GenReqDefForShowNetworkConfiguration()
	return &ShowNetworkConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StartCbhInstance 启动云堡垒机实例
//
// 启动云堡垒机实例。
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

// StartCbhInstanceInvoker 启动云堡垒机实例
func (c *CbhClient) StartCbhInstanceInvoker(request *model.StartCbhInstanceRequest) *StartCbhInstanceInvoker {
	requestDef := GenReqDefForStartCbhInstance()
	return &StartCbhInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StopCbhInstance 关闭云堡垒机实例
//
// 关闭云堡垒机实例。
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

// StopCbhInstanceInvoker 关闭云堡垒机实例
func (c *CbhClient) StopCbhInstanceInvoker(request *model.StopCbhInstanceRequest) *StopCbhInstanceInvoker {
	requestDef := GenReqDefForStopCbhInstance()
	return &StopCbhInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UninstallInstanceEip 解绑弹性公网IP
//
// 云堡垒机实例解绑弹性公网IP。
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

// UninstallInstanceEipInvoker 解绑弹性公网IP
func (c *CbhClient) UninstallInstanceEipInvoker(request *model.UninstallInstanceEipRequest) *UninstallInstanceEipInvoker {
	requestDef := GenReqDefForUninstallInstanceEip()
	return &UninstallInstanceEipInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpgradeCbhInstance 升级云堡垒机实例
//
// 升级云堡垒机实例
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

// UpgradeCbhInstanceInvoker 升级云堡垒机实例
func (c *CbhClient) UpgradeCbhInstanceInvoker(request *model.UpgradeCbhInstanceRequest) *UpgradeCbhInstanceInvoker {
	requestDef := GenReqDefForUpgradeCbhInstance()
	return &UpgradeCbhInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
