package v2

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cloudrtc/v2/model"
)

type CloudRTCClient struct {
	HcClient *http_client.HcHttpClient
}

func NewCloudRTCClient(hcClient *http_client.HcHttpClient) *CloudRTCClient {
	return &CloudRTCClient{HcClient: hcClient}
}

func CloudRTCClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// CreateApp 创建应用
//
// 调用此接口创建应用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudRTCClient) CreateApp(request *model.CreateAppRequest) (*model.CreateAppResponse, error) {
	requestDef := GenReqDefForCreateApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAppResponse), nil
	}
}

// CreateAppInvoker 创建应用
func (c *CloudRTCClient) CreateAppInvoker(request *model.CreateAppRequest) *CreateAppInvoker {
	requestDef := GenReqDefForCreateApp()
	return &CreateAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateIndividualStreamJob 启动单流任务
//
// 调用此接口接口启动单流任务。
//
// API触发单流录制流名规则：{jobtype}\\_{jobid}\\_{roomid}\\_{userid}
//
// jobtype取值为&#39;s&#39;代表单流录制。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudRTCClient) CreateIndividualStreamJob(request *model.CreateIndividualStreamJobRequest) (*model.CreateIndividualStreamJobResponse, error) {
	requestDef := GenReqDefForCreateIndividualStreamJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateIndividualStreamJobResponse), nil
	}
}

// CreateIndividualStreamJobInvoker 启动单流任务
func (c *CloudRTCClient) CreateIndividualStreamJobInvoker(request *model.CreateIndividualStreamJobRequest) *CreateIndividualStreamJobInvoker {
	requestDef := GenReqDefForCreateIndividualStreamJob()
	return &CreateIndividualStreamJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateMixJob 启动合流任务
//
// 调用此接口创建合流转码任务。
//
// 支持纯音频录制和音视频录制：
//
// - 纯音频录制
//
//   encode_template填audio_only，音频合流会动态选择最大三方的声音。
//
//   layout_template、layout_panes以及其他视频相关参数都不填，填就忽略。
//
// - 音视频录制（包括共享桌面）
//
//   encode_template非audio_only，layout_template、layout_panes必须非空。
//
//   音频合流会动态选择最大三方的声音。
//
//   API触发合流录制流名规则：{jobtype}\\_{jobid}\\_{roomid}，其中jobtype取值为&#39;m&#39;代表合流录制。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudRTCClient) CreateMixJob(request *model.CreateMixJobRequest) (*model.CreateMixJobResponse, error) {
	requestDef := GenReqDefForCreateMixJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateMixJobResponse), nil
	}
}

// CreateMixJobInvoker 启动合流任务
func (c *CloudRTCClient) CreateMixJobInvoker(request *model.CreateMixJobRequest) *CreateMixJobInvoker {
	requestDef := GenReqDefForCreateMixJob()
	return &CreateMixJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateRecordRule 创建或更新录制规则
//
// 调用此接口创建或更新录制规则。
//
// - 若当前app在请求的location中无录制规则，则会创建新的录制规则
// - 若当前app在请求的location中已有录制规则，则会更新原来的录制规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudRTCClient) CreateRecordRule(request *model.CreateRecordRuleRequest) (*model.CreateRecordRuleResponse, error) {
	requestDef := GenReqDefForCreateRecordRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRecordRuleResponse), nil
	}
}

// CreateRecordRuleInvoker 创建或更新录制规则
func (c *CloudRTCClient) CreateRecordRuleInvoker(request *model.CreateRecordRuleRequest) *CreateRecordRuleInvoker {
	requestDef := GenReqDefForCreateRecordRule()
	return &CreateRecordRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteApp 删除应用
//
// 调用此接口删除单个应用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudRTCClient) DeleteApp(request *model.DeleteAppRequest) (*model.DeleteAppResponse, error) {
	requestDef := GenReqDefForDeleteApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAppResponse), nil
	}
}

// DeleteAppInvoker 删除应用
func (c *CloudRTCClient) DeleteAppInvoker(request *model.DeleteAppRequest) *DeleteAppInvoker {
	requestDef := GenReqDefForDeleteApp()
	return &DeleteAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteRecordRule 删除录制规则
//
// 调用此接口删除录制规则，对于正在使用的录制规则，不允许删除。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudRTCClient) DeleteRecordRule(request *model.DeleteRecordRuleRequest) (*model.DeleteRecordRuleResponse, error) {
	requestDef := GenReqDefForDeleteRecordRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteRecordRuleResponse), nil
	}
}

// DeleteRecordRuleInvoker 删除录制规则
func (c *CloudRTCClient) DeleteRecordRuleInvoker(request *model.DeleteRecordRuleRequest) *DeleteRecordRuleInvoker {
	requestDef := GenReqDefForDeleteRecordRule()
	return &DeleteRecordRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListApps 查询应用列表
//
// 调用此接口查询应用列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudRTCClient) ListApps(request *model.ListAppsRequest) (*model.ListAppsResponse, error) {
	requestDef := GenReqDefForListApps()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAppsResponse), nil
	}
}

// ListAppsInvoker 查询应用列表
func (c *CloudRTCClient) ListAppsInvoker(request *model.ListAppsRequest) *ListAppsInvoker {
	requestDef := GenReqDefForListApps()
	return &ListAppsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRecordRules 查询录制规则列表
//
// 调用此接口查询录制规则列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudRTCClient) ListRecordRules(request *model.ListRecordRulesRequest) (*model.ListRecordRulesResponse, error) {
	requestDef := GenReqDefForListRecordRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRecordRulesResponse), nil
	}
}

// ListRecordRulesInvoker 查询录制规则列表
func (c *CloudRTCClient) ListRecordRulesInvoker(request *model.ListRecordRulesRequest) *ListRecordRulesInvoker {
	requestDef := GenReqDefForListRecordRules()
	return &ListRecordRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RemoveRoom 解散房间
//
// 调用此接口解散房间，将该房间中所有用户踢出房间。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudRTCClient) RemoveRoom(request *model.RemoveRoomRequest) (*model.RemoveRoomResponse, error) {
	requestDef := GenReqDefForRemoveRoom()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RemoveRoomResponse), nil
	}
}

// RemoveRoomInvoker 解散房间
func (c *CloudRTCClient) RemoveRoomInvoker(request *model.RemoveRoomRequest) *RemoveRoomInvoker {
	requestDef := GenReqDefForRemoveRoom()
	return &RemoveRoomInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RemoveUsers 踢除在线用户
//
// 调用此接口强制用户离开房间。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudRTCClient) RemoveUsers(request *model.RemoveUsersRequest) (*model.RemoveUsersResponse, error) {
	requestDef := GenReqDefForRemoveUsers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RemoveUsersResponse), nil
	}
}

// RemoveUsersInvoker 踢除在线用户
func (c *CloudRTCClient) RemoveUsersInvoker(request *model.RemoveUsersRequest) *RemoveUsersInvoker {
	requestDef := GenReqDefForRemoveUsers()
	return &RemoveUsersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowApp 查询单个应用
//
// 调用此接口查询单个应用详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudRTCClient) ShowApp(request *model.ShowAppRequest) (*model.ShowAppResponse, error) {
	requestDef := GenReqDefForShowApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAppResponse), nil
	}
}

// ShowAppInvoker 查询单个应用
func (c *CloudRTCClient) ShowAppInvoker(request *model.ShowAppRequest) *ShowAppInvoker {
	requestDef := GenReqDefForShowApp()
	return &ShowAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAutoRecord 查询自动录制配置
//
// 调用此接口查询自动录制配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudRTCClient) ShowAutoRecord(request *model.ShowAutoRecordRequest) (*model.ShowAutoRecordResponse, error) {
	requestDef := GenReqDefForShowAutoRecord()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAutoRecordResponse), nil
	}
}

// ShowAutoRecordInvoker 查询自动录制配置
func (c *CloudRTCClient) ShowAutoRecordInvoker(request *model.ShowAutoRecordRequest) *ShowAutoRecordInvoker {
	requestDef := GenReqDefForShowAutoRecord()
	return &ShowAutoRecordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowIndividualStreamJob 查询单流任务状态
//
// 调用此接口查询单流任务状态。
//
// 租户的OBS桶内的情况，暂不支持查询。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudRTCClient) ShowIndividualStreamJob(request *model.ShowIndividualStreamJobRequest) (*model.ShowIndividualStreamJobResponse, error) {
	requestDef := GenReqDefForShowIndividualStreamJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowIndividualStreamJobResponse), nil
	}
}

// ShowIndividualStreamJobInvoker 查询单流任务状态
func (c *CloudRTCClient) ShowIndividualStreamJobInvoker(request *model.ShowIndividualStreamJobRequest) *ShowIndividualStreamJobInvoker {
	requestDef := GenReqDefForShowIndividualStreamJob()
	return &ShowIndividualStreamJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMixJob 查询合流任务
//
// 调用此接口查询合流转码任务状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudRTCClient) ShowMixJob(request *model.ShowMixJobRequest) (*model.ShowMixJobResponse, error) {
	requestDef := GenReqDefForShowMixJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMixJobResponse), nil
	}
}

// ShowMixJobInvoker 查询合流任务
func (c *CloudRTCClient) ShowMixJobInvoker(request *model.ShowMixJobRequest) *ShowMixJobInvoker {
	requestDef := GenReqDefForShowMixJob()
	return &ShowMixJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRecordCallback 查询增值（录制）事件回调配置
//
// 调用此接口查询增值（录制）事件回调配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudRTCClient) ShowRecordCallback(request *model.ShowRecordCallbackRequest) (*model.ShowRecordCallbackResponse, error) {
	requestDef := GenReqDefForShowRecordCallback()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRecordCallbackResponse), nil
	}
}

// ShowRecordCallbackInvoker 查询增值（录制）事件回调配置
func (c *CloudRTCClient) ShowRecordCallbackInvoker(request *model.ShowRecordCallbackRequest) *ShowRecordCallbackInvoker {
	requestDef := GenReqDefForShowRecordCallback()
	return &ShowRecordCallbackInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRecordRule 查询录制规则
//
// 调用此接口查询指定录制规则。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudRTCClient) ShowRecordRule(request *model.ShowRecordRuleRequest) (*model.ShowRecordRuleResponse, error) {
	requestDef := GenReqDefForShowRecordRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRecordRuleResponse), nil
	}
}

// ShowRecordRuleInvoker 查询录制规则
func (c *CloudRTCClient) ShowRecordRuleInvoker(request *model.ShowRecordRuleRequest) *ShowRecordRuleInvoker {
	requestDef := GenReqDefForShowRecordRule()
	return &ShowRecordRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowUrlAuth 查询访问控制参数
//
// 查询应用鉴权配置参数
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudRTCClient) ShowUrlAuth(request *model.ShowUrlAuthRequest) (*model.ShowUrlAuthResponse, error) {
	requestDef := GenReqDefForShowUrlAuth()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowUrlAuthResponse), nil
	}
}

// ShowUrlAuthInvoker 查询访问控制参数
func (c *CloudRTCClient) ShowUrlAuthInvoker(request *model.ShowUrlAuthRequest) *ShowUrlAuthInvoker {
	requestDef := GenReqDefForShowUrlAuth()
	return &ShowUrlAuthInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StartApp 启用应用
//
// 调用此接口启用单个应用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudRTCClient) StartApp(request *model.StartAppRequest) (*model.StartAppResponse, error) {
	requestDef := GenReqDefForStartApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartAppResponse), nil
	}
}

// StartAppInvoker 启用应用
func (c *CloudRTCClient) StartAppInvoker(request *model.StartAppRequest) *StartAppInvoker {
	requestDef := GenReqDefForStartApp()
	return &StartAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StopApp 停用应用
//
// 调用此接口停用单个应用。
//
// 应用停用后，新房间无法新增和加入，已加入的房间可以继续使用。合流、录制功能等也不可用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudRTCClient) StopApp(request *model.StopAppRequest) (*model.StopAppResponse, error) {
	requestDef := GenReqDefForStopApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopAppResponse), nil
	}
}

// StopAppInvoker 停用应用
func (c *CloudRTCClient) StopAppInvoker(request *model.StopAppRequest) *StopAppInvoker {
	requestDef := GenReqDefForStopApp()
	return &StopAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StopIndividualStreamJob 停止单流任务
//
// 调用此接口停止单流任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudRTCClient) StopIndividualStreamJob(request *model.StopIndividualStreamJobRequest) (*model.StopIndividualStreamJobResponse, error) {
	requestDef := GenReqDefForStopIndividualStreamJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopIndividualStreamJobResponse), nil
	}
}

// StopIndividualStreamJobInvoker 停止单流任务
func (c *CloudRTCClient) StopIndividualStreamJobInvoker(request *model.StopIndividualStreamJobRequest) *StopIndividualStreamJobInvoker {
	requestDef := GenReqDefForStopIndividualStreamJob()
	return &StopIndividualStreamJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StopMixJob 停止合流任务
//
// 调用此接口停止已下发的合流转码任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudRTCClient) StopMixJob(request *model.StopMixJobRequest) (*model.StopMixJobResponse, error) {
	requestDef := GenReqDefForStopMixJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopMixJobResponse), nil
	}
}

// StopMixJobInvoker 停止合流任务
func (c *CloudRTCClient) StopMixJobInvoker(request *model.StopMixJobRequest) *StopMixJobInvoker {
	requestDef := GenReqDefForStopMixJob()
	return &StopMixJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAutoRecord 更新自动录制配置
//
// 更新自动录制配置，租户可以开启自动单流录制或者停用自动单流录制。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudRTCClient) UpdateAutoRecord(request *model.UpdateAutoRecordRequest) (*model.UpdateAutoRecordResponse, error) {
	requestDef := GenReqDefForUpdateAutoRecord()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAutoRecordResponse), nil
	}
}

// UpdateAutoRecordInvoker 更新自动录制配置
func (c *CloudRTCClient) UpdateAutoRecordInvoker(request *model.UpdateAutoRecordRequest) *UpdateAutoRecordInvoker {
	requestDef := GenReqDefForUpdateAutoRecord()
	return &UpdateAutoRecordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateIndividualStreamJob 更新单流任务
//
// 调用此接口修改单流任务。
//
// 仅部分场景支持修改。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudRTCClient) UpdateIndividualStreamJob(request *model.UpdateIndividualStreamJobRequest) (*model.UpdateIndividualStreamJobResponse, error) {
	requestDef := GenReqDefForUpdateIndividualStreamJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateIndividualStreamJobResponse), nil
	}
}

// UpdateIndividualStreamJobInvoker 更新单流任务
func (c *CloudRTCClient) UpdateIndividualStreamJobInvoker(request *model.UpdateIndividualStreamJobRequest) *UpdateIndividualStreamJobInvoker {
	requestDef := GenReqDefForUpdateIndividualStreamJob()
	return &UpdateIndividualStreamJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateMixJob 修改合流任务
//
// 调用此接口更新合流任务布局。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudRTCClient) UpdateMixJob(request *model.UpdateMixJobRequest) (*model.UpdateMixJobResponse, error) {
	requestDef := GenReqDefForUpdateMixJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateMixJobResponse), nil
	}
}

// UpdateMixJobInvoker 修改合流任务
func (c *CloudRTCClient) UpdateMixJobInvoker(request *model.UpdateMixJobRequest) *UpdateMixJobInvoker {
	requestDef := GenReqDefForUpdateMixJob()
	return &UpdateMixJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateRecordCallback RTC增值（录制）事件回调配置
//
// 调用此接口配置增值（录制）事件上报回调。
//
// 当任务发生订阅了的事件时，通过该接口配置的回调地址通知。
//
// 回调格式参考/customer-record-notify-url定义。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudRTCClient) UpdateRecordCallback(request *model.UpdateRecordCallbackRequest) (*model.UpdateRecordCallbackResponse, error) {
	requestDef := GenReqDefForUpdateRecordCallback()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateRecordCallbackResponse), nil
	}
}

// UpdateRecordCallbackInvoker RTC增值（录制）事件回调配置
func (c *CloudRTCClient) UpdateRecordCallbackInvoker(request *model.UpdateRecordCallbackRequest) *UpdateRecordCallbackInvoker {
	requestDef := GenReqDefForUpdateRecordCallback()
	return &UpdateRecordCallbackInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateRecordRule 更新录制规则
//
// 调用此接口更新录制规则。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudRTCClient) UpdateRecordRule(request *model.UpdateRecordRuleRequest) (*model.UpdateRecordRuleResponse, error) {
	requestDef := GenReqDefForUpdateRecordRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateRecordRuleResponse), nil
	}
}

// UpdateRecordRuleInvoker 更新录制规则
func (c *CloudRTCClient) UpdateRecordRuleInvoker(request *model.UpdateRecordRuleRequest) *UpdateRecordRuleInvoker {
	requestDef := GenReqDefForUpdateRecordRule()
	return &UpdateRecordRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateUrlAuth 开关访问控制
//
// 调用此接口开启或关闭URL鉴权。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudRTCClient) UpdateUrlAuth(request *model.UpdateUrlAuthRequest) (*model.UpdateUrlAuthResponse, error) {
	requestDef := GenReqDefForUpdateUrlAuth()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateUrlAuthResponse), nil
	}
}

// UpdateUrlAuthInvoker 开关访问控制
func (c *CloudRTCClient) UpdateUrlAuthInvoker(request *model.UpdateUrlAuthRequest) *UpdateUrlAuthInvoker {
	requestDef := GenReqDefForUpdateUrlAuth()
	return &UpdateUrlAuthInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
