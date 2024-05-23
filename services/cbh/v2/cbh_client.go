package v2

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cbh/v2/model"
)

type CbhClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewCbhClient(hcClient *httpclient.HcHttpClient) *CbhClient {
	return &CbhClient{HcClient: hcClient}
}

func CbhClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// BatchCreateInstanceTag 操作堡垒机实例资源标签
//
// 操作堡垒机实例资源标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CbhClient) BatchCreateInstanceTag(request *model.BatchCreateInstanceTagRequest) (*model.BatchCreateInstanceTagResponse, error) {
	requestDef := GenReqDefForBatchCreateInstanceTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateInstanceTagResponse), nil
	}
}

// BatchCreateInstanceTagInvoker 操作堡垒机实例资源标签
func (c *CbhClient) BatchCreateInstanceTagInvoker(request *model.BatchCreateInstanceTagRequest) *BatchCreateInstanceTagInvoker {
	requestDef := GenReqDefForBatchCreateInstanceTag()
	return &BatchCreateInstanceTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeInstanceType 修改单机堡垒机实例类型
//
// 修改单机堡垒机实例类型。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CbhClient) ChangeInstanceType(request *model.ChangeInstanceTypeRequest) (*model.ChangeInstanceTypeResponse, error) {
	requestDef := GenReqDefForChangeInstanceType()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeInstanceTypeResponse), nil
	}
}

// ChangeInstanceTypeInvoker 修改单机堡垒机实例类型
func (c *CbhClient) ChangeInstanceTypeInvoker(request *model.ChangeInstanceTypeRequest) *ChangeInstanceTypeInvoker {
	requestDef := GenReqDefForChangeInstanceType()
	return &ChangeInstanceTypeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CountInstancesByTag 统计符合标签条件的实例数量
//
// 统计符合标签条件的实例数量。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CbhClient) CountInstancesByTag(request *model.CountInstancesByTagRequest) (*model.CountInstancesByTagResponse, error) {
	requestDef := GenReqDefForCountInstancesByTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CountInstancesByTagResponse), nil
	}
}

// CountInstancesByTagInvoker 统计符合标签条件的实例数量
func (c *CbhClient) CountInstancesByTagInvoker(request *model.CountInstancesByTagRequest) *CountInstancesByTagInvoker {
	requestDef := GenReqDefForCountInstancesByTag()
	return &CountInstancesByTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateInstance 创建堡垒机实例
//
// 创建云堡垒机实例。
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

// CreateInstanceInvoker 创建堡垒机实例
func (c *CbhClient) CreateInstanceInvoker(request *model.CreateInstanceRequest) *CreateInstanceInvoker {
	requestDef := GenReqDefForCreateInstance()
	return &CreateInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteInstance 删除故障云堡垒机实例
//
// 删除云堡垒机故障实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CbhClient) DeleteInstance(request *model.DeleteInstanceRequest) (*model.DeleteInstanceResponse, error) {
	requestDef := GenReqDefForDeleteInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteInstanceResponse), nil
	}
}

// DeleteInstanceInvoker 删除故障云堡垒机实例
func (c *CbhClient) DeleteInstanceInvoker(request *model.DeleteInstanceRequest) *DeleteInstanceInvoker {
	requestDef := GenReqDefForDeleteInstance()
	return &DeleteInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// InstallInstanceEip 堡垒机实例绑定弹性公网IP
//
// 云堡垒机实例绑定弹性公网IP。
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

// InstallInstanceEipInvoker 堡垒机实例绑定弹性公网IP
func (c *CbhClient) InstallInstanceEipInvoker(request *model.InstallInstanceEipRequest) *InstallInstanceEipInvoker {
	requestDef := GenReqDefForInstallInstanceEip()
	return &InstallInstanceEipInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAvailableZones 获取服务可用区信息
//
// 获取云堡垒机服务可用区信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CbhClient) ListAvailableZones(request *model.ListAvailableZonesRequest) (*model.ListAvailableZonesResponse, error) {
	requestDef := GenReqDefForListAvailableZones()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAvailableZonesResponse), nil
	}
}

// ListAvailableZonesInvoker 获取服务可用区信息
func (c *CbhClient) ListAvailableZonesInvoker(request *model.ListAvailableZonesRequest) *ListAvailableZonesInvoker {
	requestDef := GenReqDefForListAvailableZones()
	return &ListAvailableZonesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstances 获取堡垒机实例列表
//
// 获取当前租户下的堡垒机实例列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CbhClient) ListInstances(request *model.ListInstancesRequest) (*model.ListInstancesResponse, error) {
	requestDef := GenReqDefForListInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstancesResponse), nil
	}
}

// ListInstancesInvoker 获取堡垒机实例列表
func (c *CbhClient) ListInstancesInvoker(request *model.ListInstancesRequest) *ListInstancesInvoker {
	requestDef := GenReqDefForListInstances()
	return &ListInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstancesByTag 使用标签过滤实例
//
// 使用标签过滤实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CbhClient) ListInstancesByTag(request *model.ListInstancesByTagRequest) (*model.ListInstancesByTagResponse, error) {
	requestDef := GenReqDefForListInstancesByTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstancesByTagResponse), nil
	}
}

// ListInstancesByTagInvoker 使用标签过滤实例
func (c *CbhClient) ListInstancesByTagInvoker(request *model.ListInstancesByTagRequest) *ListInstancesByTagInvoker {
	requestDef := GenReqDefForListInstancesByTag()
	return &ListInstancesByTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSpecifications 查询云堡垒机规格信息
//
// 查询云堡垒机规格信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CbhClient) ListSpecifications(request *model.ListSpecificationsRequest) (*model.ListSpecificationsResponse, error) {
	requestDef := GenReqDefForListSpecifications()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSpecificationsResponse), nil
	}
}

// ListSpecificationsInvoker 查询云堡垒机规格信息
func (c *CbhClient) ListSpecificationsInvoker(request *model.ListSpecificationsRequest) *ListSpecificationsInvoker {
	requestDef := GenReqDefForListSpecifications()
	return &ListSpecificationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTags 查询租户在项目中的资源标签集合
//
// 查询租户在项目中的资源标签集合。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CbhClient) ListTags(request *model.ListTagsRequest) (*model.ListTagsResponse, error) {
	requestDef := GenReqDefForListTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTagsResponse), nil
	}
}

// ListTagsInvoker 查询租户在项目中的资源标签集合
func (c *CbhClient) ListTagsInvoker(request *model.ListTagsRequest) *ListTagsInvoker {
	requestDef := GenReqDefForListTags()
	return &ListTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// LoginInstance IAM用户登录堡垒机实例console
//
// IAM用户登录堡垒机实例console。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CbhClient) LoginInstance(request *model.LoginInstanceRequest) (*model.LoginInstanceResponse, error) {
	requestDef := GenReqDefForLoginInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.LoginInstanceResponse), nil
	}
}

// LoginInstanceInvoker IAM用户登录堡垒机实例console
func (c *CbhClient) LoginInstanceInvoker(request *model.LoginInstanceRequest) *LoginInstanceInvoker {
	requestDef := GenReqDefForLoginInstance()
	return &LoginInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// LoginInstanceAdmin 用户登录堡垒机实例admin的console
//
// 用户登录堡垒机实例admin的console。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CbhClient) LoginInstanceAdmin(request *model.LoginInstanceAdminRequest) (*model.LoginInstanceAdminResponse, error) {
	requestDef := GenReqDefForLoginInstanceAdmin()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.LoginInstanceAdminResponse), nil
	}
}

// LoginInstanceAdminInvoker 用户登录堡垒机实例admin的console
func (c *CbhClient) LoginInstanceAdminInvoker(request *model.LoginInstanceAdminRequest) *LoginInstanceAdminInvoker {
	requestDef := GenReqDefForLoginInstanceAdmin()
	return &LoginInstanceAdminInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RebootInstance 重启堡垒机实例
//
// 重启云堡垒机实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CbhClient) RebootInstance(request *model.RebootInstanceRequest) (*model.RebootInstanceResponse, error) {
	requestDef := GenReqDefForRebootInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RebootInstanceResponse), nil
	}
}

// RebootInstanceInvoker 重启堡垒机实例
func (c *CbhClient) RebootInstanceInvoker(request *model.RebootInstanceRequest) *RebootInstanceInvoker {
	requestDef := GenReqDefForRebootInstance()
	return &RebootInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RegisterAuthorization 租户创建或取消云堡垒机服务的委托授权
//
// 租户创建或取消云堡垒机服务的委托授权。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CbhClient) RegisterAuthorization(request *model.RegisterAuthorizationRequest) (*model.RegisterAuthorizationResponse, error) {
	requestDef := GenReqDefForRegisterAuthorization()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RegisterAuthorizationResponse), nil
	}
}

// RegisterAuthorizationInvoker 租户创建或取消云堡垒机服务的委托授权
func (c *CbhClient) RegisterAuthorizationInvoker(request *model.RegisterAuthorizationRequest) *RegisterAuthorizationInvoker {
	requestDef := GenReqDefForRegisterAuthorization()
	return &RegisterAuthorizationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResetInstanceLoginMethod 重置堡垒机实例admin登录方式
//
// 重置堡垒机实例admin用户登录方式。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CbhClient) ResetInstanceLoginMethod(request *model.ResetInstanceLoginMethodRequest) (*model.ResetInstanceLoginMethodResponse, error) {
	requestDef := GenReqDefForResetInstanceLoginMethod()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResetInstanceLoginMethodResponse), nil
	}
}

// ResetInstanceLoginMethodInvoker 重置堡垒机实例admin登录方式
func (c *CbhClient) ResetInstanceLoginMethodInvoker(request *model.ResetInstanceLoginMethodRequest) *ResetInstanceLoginMethodInvoker {
	requestDef := GenReqDefForResetInstanceLoginMethod()
	return &ResetInstanceLoginMethodInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResetInstancePassword 重置堡垒机实例admin密码
//
// 重置云堡垒机实例web登录admin用户密码。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CbhClient) ResetInstancePassword(request *model.ResetInstancePasswordRequest) (*model.ResetInstancePasswordResponse, error) {
	requestDef := GenReqDefForResetInstancePassword()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResetInstancePasswordResponse), nil
	}
}

// ResetInstancePasswordInvoker 重置堡垒机实例admin密码
func (c *CbhClient) ResetInstancePasswordInvoker(request *model.ResetInstancePasswordRequest) *ResetInstancePasswordInvoker {
	requestDef := GenReqDefForResetInstancePassword()
	return &ResetInstancePasswordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResizeInstance 变更堡垒机实例
//
// 变更云堡垒机实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CbhClient) ResizeInstance(request *model.ResizeInstanceRequest) (*model.ResizeInstanceResponse, error) {
	requestDef := GenReqDefForResizeInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResizeInstanceResponse), nil
	}
}

// ResizeInstanceInvoker 变更堡垒机实例
func (c *CbhClient) ResizeInstanceInvoker(request *model.ResizeInstanceRequest) *ResizeInstanceInvoker {
	requestDef := GenReqDefForResizeInstance()
	return &ResizeInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RollbackInstance 回退升级的堡垒机实例
//
// 回退升级的云堡垒机实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CbhClient) RollbackInstance(request *model.RollbackInstanceRequest) (*model.RollbackInstanceResponse, error) {
	requestDef := GenReqDefForRollbackInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RollbackInstanceResponse), nil
	}
}

// RollbackInstanceInvoker 回退升级的堡垒机实例
func (c *CbhClient) RollbackInstanceInvoker(request *model.RollbackInstanceRequest) *RollbackInstanceInvoker {
	requestDef := GenReqDefForRollbackInstance()
	return &RollbackInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAuthorization 获取租户给云堡垒机服务委托授权信息
//
// 获取租户给云堡垒机服务委托授权信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CbhClient) ShowAuthorization(request *model.ShowAuthorizationRequest) (*model.ShowAuthorizationResponse, error) {
	requestDef := GenReqDefForShowAuthorization()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAuthorizationResponse), nil
	}
}

// ShowAuthorizationInvoker 获取租户给云堡垒机服务委托授权信息
func (c *CbhClient) ShowAuthorizationInvoker(request *model.ShowAuthorizationRequest) *ShowAuthorizationInvoker {
	requestDef := GenReqDefForShowAuthorization()
	return &ShowAuthorizationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowEcsQuota 获取创建堡垒机实例所需ECS资源配额
//
// 获取当前租户所选择的可用分区里的堡垒机ECS规格是否可用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CbhClient) ShowEcsQuota(request *model.ShowEcsQuotaRequest) (*model.ShowEcsQuotaResponse, error) {
	requestDef := GenReqDefForShowEcsQuota()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEcsQuotaResponse), nil
	}
}

// ShowEcsQuotaInvoker 获取创建堡垒机实例所需ECS资源配额
func (c *CbhClient) ShowEcsQuotaInvoker(request *model.ShowEcsQuotaRequest) *ShowEcsQuotaInvoker {
	requestDef := GenReqDefForShowEcsQuota()
	return &ShowEcsQuotaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowInstanceStatus 获取堡垒机实例状态信息
//
// 获取堡垒机实例状态信息（未删除实例）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CbhClient) ShowInstanceStatus(request *model.ShowInstanceStatusRequest) (*model.ShowInstanceStatusResponse, error) {
	requestDef := GenReqDefForShowInstanceStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceStatusResponse), nil
	}
}

// ShowInstanceStatusInvoker 获取堡垒机实例状态信息
func (c *CbhClient) ShowInstanceStatusInvoker(request *model.ShowInstanceStatusRequest) *ShowInstanceStatusInvoker {
	requestDef := GenReqDefForShowInstanceStatus()
	return &ShowInstanceStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowInstanceTags 查询堡垒机实例资源的标签信息
//
// 查询堡垒机实例资源的标签信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CbhClient) ShowInstanceTags(request *model.ShowInstanceTagsRequest) (*model.ShowInstanceTagsResponse, error) {
	requestDef := GenReqDefForShowInstanceTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceTagsResponse), nil
	}
}

// ShowInstanceTagsInvoker 查询堡垒机实例资源的标签信息
func (c *CbhClient) ShowInstanceTagsInvoker(request *model.ShowInstanceTagsRequest) *ShowInstanceTagsInvoker {
	requestDef := GenReqDefForShowInstanceTags()
	return &ShowInstanceTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowOmUrl 获取运维链接
//
// 获取运维链接
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CbhClient) ShowOmUrl(request *model.ShowOmUrlRequest) (*model.ShowOmUrlResponse, error) {
	requestDef := GenReqDefForShowOmUrl()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowOmUrlResponse), nil
	}
}

// ShowOmUrlInvoker 获取运维链接
func (c *CbhClient) ShowOmUrlInvoker(request *model.ShowOmUrlRequest) *ShowOmUrlInvoker {
	requestDef := GenReqDefForShowOmUrl()
	return &ShowOmUrlInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowQuota 获取堡垒机实例配额
//
// 获取堡垒机实例配额信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CbhClient) ShowQuota(request *model.ShowQuotaRequest) (*model.ShowQuotaResponse, error) {
	requestDef := GenReqDefForShowQuota()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowQuotaResponse), nil
	}
}

// ShowQuotaInvoker 获取堡垒机实例配额
func (c *CbhClient) ShowQuotaInvoker(request *model.ShowQuotaRequest) *ShowQuotaInvoker {
	requestDef := GenReqDefForShowQuota()
	return &ShowQuotaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StartInstance 启动堡垒机实例
//
// 启动云堡垒机实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CbhClient) StartInstance(request *model.StartInstanceRequest) (*model.StartInstanceResponse, error) {
	requestDef := GenReqDefForStartInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartInstanceResponse), nil
	}
}

// StartInstanceInvoker 启动堡垒机实例
func (c *CbhClient) StartInstanceInvoker(request *model.StartInstanceRequest) *StartInstanceInvoker {
	requestDef := GenReqDefForStartInstance()
	return &StartInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StopInstance 关闭堡垒机实例
//
// 关闭云堡垒机实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CbhClient) StopInstance(request *model.StopInstanceRequest) (*model.StopInstanceResponse, error) {
	requestDef := GenReqDefForStopInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopInstanceResponse), nil
	}
}

// StopInstanceInvoker 关闭堡垒机实例
func (c *CbhClient) StopInstanceInvoker(request *model.StopInstanceRequest) *StopInstanceInvoker {
	requestDef := GenReqDefForStopInstance()
	return &StopInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SwitchInstanceVpc 切换堡垒机虚拟私有云
//
// 切换堡垒机虚拟私有云
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CbhClient) SwitchInstanceVpc(request *model.SwitchInstanceVpcRequest) (*model.SwitchInstanceVpcResponse, error) {
	requestDef := GenReqDefForSwitchInstanceVpc()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SwitchInstanceVpcResponse), nil
	}
}

// SwitchInstanceVpcInvoker 切换堡垒机虚拟私有云
func (c *CbhClient) SwitchInstanceVpcInvoker(request *model.SwitchInstanceVpcRequest) *SwitchInstanceVpcInvoker {
	requestDef := GenReqDefForSwitchInstanceVpc()
	return &SwitchInstanceVpcInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UninstallInstanceEip 堡垒机实例解绑弹性公网IP
//
// 为云堡垒机实例解绑弹性公网IP。
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

// UninstallInstanceEipInvoker 堡垒机实例解绑弹性公网IP
func (c *CbhClient) UninstallInstanceEipInvoker(request *model.UninstallInstanceEipRequest) *UninstallInstanceEipInvoker {
	requestDef := GenReqDefForUninstallInstanceEip()
	return &UninstallInstanceEipInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateInstanceSecurityGroup 修改堡垒机实例安全组
//
// 修改堡垒机实例安全组。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CbhClient) UpdateInstanceSecurityGroup(request *model.UpdateInstanceSecurityGroupRequest) (*model.UpdateInstanceSecurityGroupResponse, error) {
	requestDef := GenReqDefForUpdateInstanceSecurityGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInstanceSecurityGroupResponse), nil
	}
}

// UpdateInstanceSecurityGroupInvoker 修改堡垒机实例安全组
func (c *CbhClient) UpdateInstanceSecurityGroupInvoker(request *model.UpdateInstanceSecurityGroupRequest) *UpdateInstanceSecurityGroupInvoker {
	requestDef := GenReqDefForUpdateInstanceSecurityGroup()
	return &UpdateInstanceSecurityGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpgradeInstance 升级堡垒机实例
//
// 升级云堡垒机实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CbhClient) UpgradeInstance(request *model.UpgradeInstanceRequest) (*model.UpgradeInstanceResponse, error) {
	requestDef := GenReqDefForUpgradeInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpgradeInstanceResponse), nil
	}
}

// UpgradeInstanceInvoker 升级堡垒机实例
func (c *CbhClient) UpgradeInstanceInvoker(request *model.UpgradeInstanceRequest) *UpgradeInstanceInvoker {
	requestDef := GenReqDefForUpgradeInstance()
	return &UpgradeInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
