package v3

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/gsl/v3/model"
)

type GslClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewGslClient(hcClient *httpclient.HcHttpClient) *GslClient {
	return &GslClient{HcClient: hcClient}
}

func GslClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// BatchSetAttributes 批量设置自定义属性接口
//
// 批量设置自定义属性接口
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GslClient) BatchSetAttributes(request *model.BatchSetAttributesRequest) (*model.BatchSetAttributesResponse, error) {
	requestDef := GenReqDefForBatchSetAttributes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchSetAttributesResponse), nil
	}
}

// BatchSetAttributesInvoker 批量设置自定义属性接口
func (c *GslClient) BatchSetAttributesInvoker(request *model.BatchSetAttributesRequest) *BatchSetAttributesInvoker {
	requestDef := GenReqDefForBatchSetAttributes()
	return &BatchSetAttributesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAttribute 用户新增自定义属性接口
//
// 用户新增自定义属性接口
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GslClient) CreateAttribute(request *model.CreateAttributeRequest) (*model.CreateAttributeResponse, error) {
	requestDef := GenReqDefForCreateAttribute()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAttributeResponse), nil
	}
}

// CreateAttributeInvoker 用户新增自定义属性接口
func (c *GslClient) CreateAttributeInvoker(request *model.CreateAttributeRequest) *CreateAttributeInvoker {
	requestDef := GenReqDefForCreateAttribute()
	return &CreateAttributeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DisableAttribute 停用自定义属性接口
//
// 停用自定义属性接口
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GslClient) DisableAttribute(request *model.DisableAttributeRequest) (*model.DisableAttributeResponse, error) {
	requestDef := GenReqDefForDisableAttribute()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisableAttributeResponse), nil
	}
}

// DisableAttributeInvoker 停用自定义属性接口
func (c *GslClient) DisableAttributeInvoker(request *model.DisableAttributeRequest) *DisableAttributeInvoker {
	requestDef := GenReqDefForDisableAttribute()
	return &DisableAttributeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// EnableAttribute 启用自定义属性接口
//
// 启用自定义属性接口
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GslClient) EnableAttribute(request *model.EnableAttributeRequest) (*model.EnableAttributeResponse, error) {
	requestDef := GenReqDefForEnableAttribute()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.EnableAttributeResponse), nil
	}
}

// EnableAttributeInvoker 启用自定义属性接口
func (c *GslClient) EnableAttributeInvoker(request *model.EnableAttributeRequest) *EnableAttributeInvoker {
	requestDef := GenReqDefForEnableAttribute()
	return &EnableAttributeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAttributes 查询自定义属性列表接口
//
// 查询自定义属性列表接口
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GslClient) ListAttributes(request *model.ListAttributesRequest) (*model.ListAttributesResponse, error) {
	requestDef := GenReqDefForListAttributes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAttributesResponse), nil
	}
}

// ListAttributesInvoker 查询自定义属性列表接口
func (c *GslClient) ListAttributesInvoker(request *model.ListAttributesRequest) *ListAttributesInvoker {
	requestDef := GenReqDefForListAttributes()
	return &ListAttributesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAttribute 修改自定义属性接口
//
// 修改自定义属性接口
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GslClient) UpdateAttribute(request *model.UpdateAttributeRequest) (*model.UpdateAttributeResponse, error) {
	requestDef := GenReqDefForUpdateAttribute()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAttributeResponse), nil
	}
}

// UpdateAttributeInvoker 修改自定义属性接口
func (c *GslClient) UpdateAttributeInvoker(request *model.UpdateAttributeRequest) *UpdateAttributeInvoker {
	requestDef := GenReqDefForUpdateAttribute()
	return &UpdateAttributeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListBackPoolMembers 查询后向流量池成员列表
//
// 查询后向流量池成员列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GslClient) ListBackPoolMembers(request *model.ListBackPoolMembersRequest) (*model.ListBackPoolMembersResponse, error) {
	requestDef := GenReqDefForListBackPoolMembers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBackPoolMembersResponse), nil
	}
}

// ListBackPoolMembersInvoker 查询后向流量池成员列表
func (c *GslClient) ListBackPoolMembersInvoker(request *model.ListBackPoolMembersRequest) *ListBackPoolMembersInvoker {
	requestDef := GenReqDefForListBackPoolMembers()
	return &ListBackPoolMembersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListBackPools 查询后向流量池列表
//
// 查询后向流量池列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GslClient) ListBackPools(request *model.ListBackPoolsRequest) (*model.ListBackPoolsResponse, error) {
	requestDef := GenReqDefForListBackPools()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBackPoolsResponse), nil
	}
}

// ListBackPoolsInvoker 查询后向流量池列表
func (c *GslClient) ListBackPoolsInvoker(request *model.ListBackPoolsRequest) *ListBackPoolsInvoker {
	requestDef := GenReqDefForListBackPools()
	return &ListBackPoolsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddNetworkSwitchPolicy 新增网络切换策略
//
// 新增网络切换策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GslClient) AddNetworkSwitchPolicy(request *model.AddNetworkSwitchPolicyRequest) (*model.AddNetworkSwitchPolicyResponse, error) {
	requestDef := GenReqDefForAddNetworkSwitchPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddNetworkSwitchPolicyResponse), nil
	}
}

// AddNetworkSwitchPolicyInvoker 新增网络切换策略
func (c *GslClient) AddNetworkSwitchPolicyInvoker(request *model.AddNetworkSwitchPolicyRequest) *AddNetworkSwitchPolicyInvoker {
	requestDef := GenReqDefForAddNetworkSwitchPolicy()
	return &AddNetworkSwitchPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListNetworkSwitchPolicies 查询策略列表
//
// 查询策略列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GslClient) ListNetworkSwitchPolicies(request *model.ListNetworkSwitchPoliciesRequest) (*model.ListNetworkSwitchPoliciesResponse, error) {
	requestDef := GenReqDefForListNetworkSwitchPolicies()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListNetworkSwitchPoliciesResponse), nil
	}
}

// ListNetworkSwitchPoliciesInvoker 查询策略列表
func (c *GslClient) ListNetworkSwitchPoliciesInvoker(request *model.ListNetworkSwitchPoliciesRequest) *ListNetworkSwitchPoliciesInvoker {
	requestDef := GenReqDefForListNetworkSwitchPolicies()
	return &ListNetworkSwitchPoliciesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProPricePlans 查询套餐列表信息
//
// 查询套餐列表信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GslClient) ListProPricePlans(request *model.ListProPricePlansRequest) (*model.ListProPricePlansResponse, error) {
	requestDef := GenReqDefForListProPricePlans()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProPricePlansResponse), nil
	}
}

// ListProPricePlansInvoker 查询套餐列表信息
func (c *GslClient) ListProPricePlansInvoker(request *model.ListProPricePlansRequest) *ListProPricePlansInvoker {
	requestDef := GenReqDefForListProPricePlans()
	return &ListProPricePlansInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteRealName 清除实名认证信息
//
// 清除实名认证信息，接口仅支持中国电信卡调用。注：由于运营商侧业务限制，建议您同一张SIM卡不要同时执行多种不同业务的操作。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GslClient) DeleteRealName(request *model.DeleteRealNameRequest) (*model.DeleteRealNameResponse, error) {
	requestDef := GenReqDefForDeleteRealName()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteRealNameResponse), nil
	}
}

// DeleteRealNameInvoker 清除实名认证信息
func (c *GslClient) DeleteRealNameInvoker(request *model.DeleteRealNameRequest) *DeleteRealNameInvoker {
	requestDef := GenReqDefForDeleteRealName()
	return &DeleteRealNameInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// EnableSimCard 激活实体卡
//
// 创建激活实体卡申请，返回业务受理单号。1~2个工作日完成激活操作。注：由于运营商侧业务限制，建议您同一张SIM卡不要同时执行多种不同业务的操作。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GslClient) EnableSimCard(request *model.EnableSimCardRequest) (*model.EnableSimCardResponse, error) {
	requestDef := GenReqDefForEnableSimCard()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.EnableSimCardResponse), nil
	}
}

// EnableSimCardInvoker 激活实体卡
func (c *GslClient) EnableSimCardInvoker(request *model.EnableSimCardRequest) *EnableSimCardInvoker {
	requestDef := GenReqDefForEnableSimCard()
	return &EnableSimCardInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSimCardFlowPerDay 批量查询SIM卡日用量
//
// 批量查询SIM卡日用量接口，支持按天或按月查询。SIM卡标识和容器ID只能选一个参数，天和月也只能选一个参数
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GslClient) ListSimCardFlowPerDay(request *model.ListSimCardFlowPerDayRequest) (*model.ListSimCardFlowPerDayResponse, error) {
	requestDef := GenReqDefForListSimCardFlowPerDay()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSimCardFlowPerDayResponse), nil
	}
}

// ListSimCardFlowPerDayInvoker 批量查询SIM卡日用量
func (c *GslClient) ListSimCardFlowPerDayInvoker(request *model.ListSimCardFlowPerDayRequest) *ListSimCardFlowPerDayInvoker {
	requestDef := GenReqDefForListSimCardFlowPerDay()
	return &ListSimCardFlowPerDayInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSimCards 查询SIM卡列表
//
// 查询SIM卡列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GslClient) ListSimCards(request *model.ListSimCardsRequest) (*model.ListSimCardsResponse, error) {
	requestDef := GenReqDefForListSimCards()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSimCardsResponse), nil
	}
}

// ListSimCardsInvoker 查询SIM卡列表
func (c *GslClient) ListSimCardsInvoker(request *model.ListSimCardsRequest) *ListSimCardsInvoker {
	requestDef := GenReqDefForListSimCards()
	return &ListSimCardsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RegisterImei SIM卡机卡重绑
//
// 支持固定机卡重绑(需要上传IMEI，将SIM卡绑定到指定IMEI的设备)和普通机卡重绑(会清除之前绑定的设备,将SIM卡绑定到正在使用的设备)，接口仅支持中国电信卡，中国移动卡调用。中国电信卡单卡每月只允许重绑2次，中国移动卡仅支持普通机卡重绑。注：由于运营商侧业务限制，建议您同一张SIM卡不要同时执行多种不同业务的操作。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GslClient) RegisterImei(request *model.RegisterImeiRequest) (*model.RegisterImeiResponse, error) {
	requestDef := GenReqDefForRegisterImei()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RegisterImeiResponse), nil
	}
}

// RegisterImeiInvoker SIM卡机卡重绑
func (c *GslClient) RegisterImeiInvoker(request *model.RegisterImeiRequest) *RegisterImeiInvoker {
	requestDef := GenReqDefForRegisterImei()
	return &RegisterImeiInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResetSimCard SIM卡单卡复机
//
// 创建复机申请，返回业务受理单号。1~2个工作日完成复机操作。注：由于运营商侧业务限制，建议您同一张SIM卡不要同时执行多种不同业务的操作。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GslClient) ResetSimCard(request *model.ResetSimCardRequest) (*model.ResetSimCardResponse, error) {
	requestDef := GenReqDefForResetSimCard()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResetSimCardResponse), nil
	}
}

// ResetSimCardInvoker SIM卡单卡复机
func (c *GslClient) ResetSimCardInvoker(request *model.ResetSimCardRequest) *ResetSimCardInvoker {
	requestDef := GenReqDefForResetSimCard()
	return &ResetSimCardInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetExceedCutNet SIM卡达量断网/取消达量断网
//
// SIM卡达量断网/取消达量断网，接口仅支持中国电信的卡以及中国联通、中国移动的组池卡调用。注：由于运营商侧业务限制，建议您同一张SIM卡不要同时执行多种不同业务的操作。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GslClient) SetExceedCutNet(request *model.SetExceedCutNetRequest) (*model.SetExceedCutNetResponse, error) {
	requestDef := GenReqDefForSetExceedCutNet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetExceedCutNetResponse), nil
	}
}

// SetExceedCutNetInvoker SIM卡达量断网/取消达量断网
func (c *GslClient) SetExceedCutNetInvoker(request *model.SetExceedCutNetRequest) *SetExceedCutNetInvoker {
	requestDef := GenReqDefForSetExceedCutNet()
	return &SetExceedCutNetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetSpeedValue 实体卡限速
//
// 实体卡限速接口，接口仅支持中国电信和中国联通实体卡调用。中国联通卡需要个人实名认证后才能使用限速功能。注：由于运营商侧业务限制，建议您同一张SIM卡不要同时执行多种不同业务的操作。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GslClient) SetSpeedValue(request *model.SetSpeedValueRequest) (*model.SetSpeedValueResponse, error) {
	requestDef := GenReqDefForSetSpeedValue()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetSpeedValueResponse), nil
	}
}

// SetSpeedValueInvoker 实体卡限速
func (c *GslClient) SetSpeedValueInvoker(request *model.SetSpeedValueRequest) *SetSpeedValueInvoker {
	requestDef := GenReqDefForSetSpeedValue()
	return &SetSpeedValueInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMonthUsages 月用量统计
//
// 设备月用量统计
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GslClient) ShowMonthUsages(request *model.ShowMonthUsagesRequest) (*model.ShowMonthUsagesResponse, error) {
	requestDef := GenReqDefForShowMonthUsages()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMonthUsagesResponse), nil
	}
}

// ShowMonthUsagesInvoker 月用量统计
func (c *GslClient) ShowMonthUsagesInvoker(request *model.ShowMonthUsagesRequest) *ShowMonthUsagesInvoker {
	requestDef := GenReqDefForShowMonthUsages()
	return &ShowMonthUsagesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRealNamed 查询SIM卡实名认证信息
//
// 实时查询SIM卡实名认证信息，接口仅支持查询中国大陆运营商卡片的实名认证信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GslClient) ShowRealNamed(request *model.ShowRealNamedRequest) (*model.ShowRealNamedResponse, error) {
	requestDef := GenReqDefForShowRealNamed()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRealNamedResponse), nil
	}
}

// ShowRealNamedInvoker 查询SIM卡实名认证信息
func (c *GslClient) ShowRealNamedInvoker(request *model.ShowRealNamedRequest) *ShowRealNamedInvoker {
	requestDef := GenReqDefForShowRealNamed()
	return &ShowRealNamedInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSimCard 查询SIM卡详情
//
// 查询SIM卡详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GslClient) ShowSimCard(request *model.ShowSimCardRequest) (*model.ShowSimCardResponse, error) {
	requestDef := GenReqDefForShowSimCard()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSimCardResponse), nil
	}
}

// ShowSimCardInvoker 查询SIM卡详情
func (c *GslClient) ShowSimCardInvoker(request *model.ShowSimCardRequest) *ShowSimCardInvoker {
	requestDef := GenReqDefForShowSimCard()
	return &ShowSimCardInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StartStopNet SIM卡申请断网/恢复在用
//
// SIM卡申请断网/恢复在用，接口仅支持中国电信卡调用。注：由于运营商侧业务限制，建议您同一张SIM卡不要同时执行多种不同业务的操作。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GslClient) StartStopNet(request *model.StartStopNetRequest) (*model.StartStopNetResponse, error) {
	requestDef := GenReqDefForStartStopNet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartStopNetResponse), nil
	}
}

// StartStopNetInvoker SIM卡申请断网/恢复在用
func (c *GslClient) StartStopNetInvoker(request *model.StartStopNetRequest) *StartStopNetInvoker {
	requestDef := GenReqDefForStartStopNet()
	return &StartStopNetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StopSimCard SIM卡单卡停机
//
// 创建停机申请，返回业务受理单号。1~2个工作日完成停机操作。注：由于运营商侧业务限制，建议您同一张SIM卡不要同时执行多种不同业务的操作。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GslClient) StopSimCard(request *model.StopSimCardRequest) (*model.StopSimCardResponse, error) {
	requestDef := GenReqDefForStopSimCard()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopSimCardResponse), nil
	}
}

// StopSimCardInvoker SIM卡单卡停机
func (c *GslClient) StopSimCardInvoker(request *model.StopSimCardRequest) *StopSimCardInvoker {
	requestDef := GenReqDefForStopSimCard()
	return &StopSimCardInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSimDeviceMultiply 查询三网卡列表
//
// 通过cid或全量查询三网卡列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GslClient) ListSimDeviceMultiply(request *model.ListSimDeviceMultiplyRequest) (*model.ListSimDeviceMultiplyResponse, error) {
	requestDef := GenReqDefForListSimDeviceMultiply()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSimDeviceMultiplyResponse), nil
	}
}

// ListSimDeviceMultiplyInvoker 查询三网卡列表
func (c *GslClient) ListSimDeviceMultiplyInvoker(request *model.ListSimDeviceMultiplyRequest) *ListSimDeviceMultiplyInvoker {
	requestDef := GenReqDefForListSimDeviceMultiply()
	return &ListSimDeviceMultiplyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetNetworkSwitchPolicy SIM卡设置网络切换策略
//
// SIM卡设置网络切换策略，接口仅支持三网卡调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GslClient) SetNetworkSwitchPolicy(request *model.SetNetworkSwitchPolicyRequest) (*model.SetNetworkSwitchPolicyResponse, error) {
	requestDef := GenReqDefForSetNetworkSwitchPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetNetworkSwitchPolicyResponse), nil
	}
}

// SetNetworkSwitchPolicyInvoker SIM卡设置网络切换策略
func (c *GslClient) SetNetworkSwitchPolicyInvoker(request *model.SetNetworkSwitchPolicyRequest) *SetNetworkSwitchPolicyInvoker {
	requestDef := GenReqDefForSetNetworkSwitchPolicy()
	return &SetNetworkSwitchPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SwitchNetwork 切换网络
//
// 切换网络
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GslClient) SwitchNetwork(request *model.SwitchNetworkRequest) (*model.SwitchNetworkResponse, error) {
	requestDef := GenReqDefForSwitchNetwork()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SwitchNetworkResponse), nil
	}
}

// SwitchNetworkInvoker 切换网络
func (c *GslClient) SwitchNetworkInvoker(request *model.SwitchNetworkRequest) *SwitchNetworkInvoker {
	requestDef := GenReqDefForSwitchNetwork()
	return &SwitchNetworkInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSimPoolMembers 查询流量池成员列表
//
// 查询流量池成员列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GslClient) ListSimPoolMembers(request *model.ListSimPoolMembersRequest) (*model.ListSimPoolMembersResponse, error) {
	requestDef := GenReqDefForListSimPoolMembers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSimPoolMembersResponse), nil
	}
}

// ListSimPoolMembersInvoker 查询流量池成员列表
func (c *GslClient) ListSimPoolMembersInvoker(request *model.ListSimPoolMembersRequest) *ListSimPoolMembersInvoker {
	requestDef := GenReqDefForListSimPoolMembers()
	return &ListSimPoolMembersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSimPools 查询流量池列表
//
// 查询流量池列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GslClient) ListSimPools(request *model.ListSimPoolsRequest) (*model.ListSimPoolsResponse, error) {
	requestDef := GenReqDefForListSimPools()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSimPoolsResponse), nil
	}
}

// ListSimPoolsInvoker 查询流量池列表
func (c *GslClient) ListSimPoolsInvoker(request *model.ListSimPoolsRequest) *ListSimPoolsInvoker {
	requestDef := GenReqDefForListSimPools()
	return &ListSimPoolsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFlowBySimCards 批量查询实体卡流量
//
// 批量查询实体卡流量
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GslClient) ListFlowBySimCards(request *model.ListFlowBySimCardsRequest) (*model.ListFlowBySimCardsResponse, error) {
	requestDef := GenReqDefForListFlowBySimCards()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFlowBySimCardsResponse), nil
	}
}

// ListFlowBySimCardsInvoker 批量查询实体卡流量
func (c *GslClient) ListFlowBySimCardsInvoker(request *model.ListFlowBySimCardsRequest) *ListFlowBySimCardsInvoker {
	requestDef := GenReqDefForListFlowBySimCards()
	return &ListFlowBySimCardsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSimPricePlans sim卡套餐列表查询
//
// SIM卡套餐列表查询，实体卡只会有一个套餐，eSIM/vSIM可能会有多个套餐
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GslClient) ListSimPricePlans(request *model.ListSimPricePlansRequest) (*model.ListSimPricePlansResponse, error) {
	requestDef := GenReqDefForListSimPricePlans()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSimPricePlansResponse), nil
	}
}

// ListSimPricePlansInvoker sim卡套餐列表查询
func (c *GslClient) ListSimPricePlansInvoker(request *model.ListSimPricePlansRequest) *ListSimPricePlansInvoker {
	requestDef := GenReqDefForListSimPricePlans()
	return &ListSimPricePlansInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSmsDetails 短信发送详情
//
// 短信发送详情，接口仅支持开通短信套餐的中国移动与中国电信卡调用
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GslClient) ListSmsDetails(request *model.ListSmsDetailsRequest) (*model.ListSmsDetailsResponse, error) {
	requestDef := GenReqDefForListSmsDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSmsDetailsResponse), nil
	}
}

// ListSmsDetailsInvoker 短信发送详情
func (c *GslClient) ListSmsDetailsInvoker(request *model.ListSmsDetailsRequest) *ListSmsDetailsInvoker {
	requestDef := GenReqDefForListSmsDetails()
	return &ListSmsDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SendSms 发送短信
//
// 发送短信，接口仅支持开通短信套餐的中国移动与中国电信卡调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GslClient) SendSms(request *model.SendSmsRequest) (*model.SendSmsResponse, error) {
	requestDef := GenReqDefForSendSms()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SendSmsResponse), nil
	}
}

// SendSmsInvoker 发送短信
func (c *GslClient) SendSmsInvoker(request *model.SendSmsRequest) *SendSmsInvoker {
	requestDef := GenReqDefForSendSms()
	return &SendSmsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchSetTags 批量设置/取消设置标签接口
//
// 批量设置/取消设置标签接口
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GslClient) BatchSetTags(request *model.BatchSetTagsRequest) (*model.BatchSetTagsResponse, error) {
	requestDef := GenReqDefForBatchSetTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchSetTagsResponse), nil
	}
}

// BatchSetTagsInvoker 批量设置/取消设置标签接口
func (c *GslClient) BatchSetTagsInvoker(request *model.BatchSetTagsRequest) *BatchSetTagsInvoker {
	requestDef := GenReqDefForBatchSetTags()
	return &BatchSetTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateTag 用户添加标签
//
// 添加标签接口
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GslClient) CreateTag(request *model.CreateTagRequest) (*model.CreateTagResponse, error) {
	requestDef := GenReqDefForCreateTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTagResponse), nil
	}
}

// CreateTagInvoker 用户添加标签
func (c *GslClient) CreateTagInvoker(request *model.CreateTagRequest) *CreateTagInvoker {
	requestDef := GenReqDefForCreateTag()
	return &CreateTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteTag 删除标签
//
// 删除标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GslClient) DeleteTag(request *model.DeleteTagRequest) (*model.DeleteTagResponse, error) {
	requestDef := GenReqDefForDeleteTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTagResponse), nil
	}
}

// DeleteTagInvoker 删除标签
func (c *GslClient) DeleteTagInvoker(request *model.DeleteTagRequest) *DeleteTagInvoker {
	requestDef := GenReqDefForDeleteTag()
	return &DeleteTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTags 查询标签列表
//
// 查询标签列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GslClient) ListTags(request *model.ListTagsRequest) (*model.ListTagsResponse, error) {
	requestDef := GenReqDefForListTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTagsResponse), nil
	}
}

// ListTagsInvoker 查询标签列表
func (c *GslClient) ListTagsInvoker(request *model.ListTagsRequest) *ListTagsInvoker {
	requestDef := GenReqDefForListTags()
	return &ListTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListWorkOrderDetails 分页查询业务受理明细
//
// 分页查询业务受理明细
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GslClient) ListWorkOrderDetails(request *model.ListWorkOrderDetailsRequest) (*model.ListWorkOrderDetailsResponse, error) {
	requestDef := GenReqDefForListWorkOrderDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListWorkOrderDetailsResponse), nil
	}
}

// ListWorkOrderDetailsInvoker 分页查询业务受理明细
func (c *GslClient) ListWorkOrderDetailsInvoker(request *model.ListWorkOrderDetailsRequest) *ListWorkOrderDetailsInvoker {
	requestDef := GenReqDefForListWorkOrderDetails()
	return &ListWorkOrderDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListWorkOrders 分页查询业务受理单
//
// 分页查询业务受理单
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GslClient) ListWorkOrders(request *model.ListWorkOrdersRequest) (*model.ListWorkOrdersResponse, error) {
	requestDef := GenReqDefForListWorkOrders()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListWorkOrdersResponse), nil
	}
}

// ListWorkOrdersInvoker 分页查询业务受理单
func (c *GslClient) ListWorkOrdersInvoker(request *model.ListWorkOrdersRequest) *ListWorkOrdersInvoker {
	requestDef := GenReqDefForListWorkOrders()
	return &ListWorkOrdersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
