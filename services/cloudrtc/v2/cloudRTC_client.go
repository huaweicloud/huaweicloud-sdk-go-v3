package v2

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"

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

// 创建应用
//
// 调用此接口创建应用。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudRTCClient) CreateApp(request *model.CreateAppRequest) (*model.CreateAppResponse, error) {
	requestDef := GenReqDefForCreateApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAppResponse), nil
	}
}

// 启动单流任务
//
// 调用此接口接口启动单流任务。
//
// API触发单流录制流名规则：{jobtype}\\_{jobid}\\_{roomid}\\_{userid}
//
// jobtype取值为&#39;s&#39;代表单流录制。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudRTCClient) CreateIndividualStreamJob(request *model.CreateIndividualStreamJobRequest) (*model.CreateIndividualStreamJobResponse, error) {
	requestDef := GenReqDefForCreateIndividualStreamJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateIndividualStreamJobResponse), nil
	}
}

// 启动合流任务
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
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudRTCClient) CreateMixJob(request *model.CreateMixJobRequest) (*model.CreateMixJobResponse, error) {
	requestDef := GenReqDefForCreateMixJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateMixJobResponse), nil
	}
}

// 创建或更新录制规则
//
// 调用此接口创建或更新录制规则。
//
// - 若当前app在请求的location中无录制规则，则会创建新的录制规则
// - 若当前app在请求的location中已有录制规则，则会更新原来的录制规则
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudRTCClient) CreateRecordRule(request *model.CreateRecordRuleRequest) (*model.CreateRecordRuleResponse, error) {
	requestDef := GenReqDefForCreateRecordRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRecordRuleResponse), nil
	}
}

// 删除应用
//
// 调用此接口删除单个应用。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudRTCClient) DeleteApp(request *model.DeleteAppRequest) (*model.DeleteAppResponse, error) {
	requestDef := GenReqDefForDeleteApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAppResponse), nil
	}
}

// 删除录制规则
//
// 调用此接口删除录制规则，对于正在使用的录制规则，不允许删除。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudRTCClient) DeleteRecordRule(request *model.DeleteRecordRuleRequest) (*model.DeleteRecordRuleResponse, error) {
	requestDef := GenReqDefForDeleteRecordRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteRecordRuleResponse), nil
	}
}

// 查询应用列表
//
// 调用此接口查询应用列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudRTCClient) ListApps(request *model.ListAppsRequest) (*model.ListAppsResponse, error) {
	requestDef := GenReqDefForListApps()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAppsResponse), nil
	}
}

// 查询录制规则列表
//
// 调用此接口查询录制规则列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudRTCClient) ListRecordRules(request *model.ListRecordRulesRequest) (*model.ListRecordRulesResponse, error) {
	requestDef := GenReqDefForListRecordRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRecordRulesResponse), nil
	}
}

// 查询单个应用
//
// 调用此接口查询单个应用详情。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudRTCClient) ShowApp(request *model.ShowAppRequest) (*model.ShowAppResponse, error) {
	requestDef := GenReqDefForShowApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAppResponse), nil
	}
}

// 查询自动录制配置
//
// 调用此接口查询自动录制配置
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudRTCClient) ShowAutoRecord(request *model.ShowAutoRecordRequest) (*model.ShowAutoRecordResponse, error) {
	requestDef := GenReqDefForShowAutoRecord()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAutoRecordResponse), nil
	}
}

// 查询单流任务状态
//
// 调用此接口查询单流任务状态。
//
// 租户的OBS桶内的情况，暂不支持查询。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudRTCClient) ShowIndividualStreamJob(request *model.ShowIndividualStreamJobRequest) (*model.ShowIndividualStreamJobResponse, error) {
	requestDef := GenReqDefForShowIndividualStreamJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowIndividualStreamJobResponse), nil
	}
}

// 查询合流任务
//
// 调用此接口查询合流转码任务状态。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudRTCClient) ShowMixJob(request *model.ShowMixJobRequest) (*model.ShowMixJobResponse, error) {
	requestDef := GenReqDefForShowMixJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMixJobResponse), nil
	}
}

// 查询增值（录制）事件回调配置
//
// 调用此接口查询增值（录制）事件回调配置
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudRTCClient) ShowRecordCallback(request *model.ShowRecordCallbackRequest) (*model.ShowRecordCallbackResponse, error) {
	requestDef := GenReqDefForShowRecordCallback()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRecordCallbackResponse), nil
	}
}

// 查询录制规则
//
// 调用此接口查询指定录制规则。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudRTCClient) ShowRecordRule(request *model.ShowRecordRuleRequest) (*model.ShowRecordRuleResponse, error) {
	requestDef := GenReqDefForShowRecordRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRecordRuleResponse), nil
	}
}

// 查询访问控制参数
//
// 查询应用鉴权配置参数
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudRTCClient) ShowUrlAuth(request *model.ShowUrlAuthRequest) (*model.ShowUrlAuthResponse, error) {
	requestDef := GenReqDefForShowUrlAuth()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowUrlAuthResponse), nil
	}
}

// 启用应用
//
// 调用此接口启用单个应用。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudRTCClient) StartApp(request *model.StartAppRequest) (*model.StartAppResponse, error) {
	requestDef := GenReqDefForStartApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartAppResponse), nil
	}
}

// 停用应用
//
// 调用此接口停用单个应用。
//
// 应用停用后，新房间无法新增和加入，已加入的房间可以继续使用。合流、录制功能等也不可用。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudRTCClient) StopApp(request *model.StopAppRequest) (*model.StopAppResponse, error) {
	requestDef := GenReqDefForStopApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopAppResponse), nil
	}
}

// 停止单流任务
//
// 调用此接口停止单流任务
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudRTCClient) StopIndividualStreamJob(request *model.StopIndividualStreamJobRequest) (*model.StopIndividualStreamJobResponse, error) {
	requestDef := GenReqDefForStopIndividualStreamJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopIndividualStreamJobResponse), nil
	}
}

// 停止合流任务
//
// 调用此接口停止已下发的合流转码任务。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudRTCClient) StopMixJob(request *model.StopMixJobRequest) (*model.StopMixJobResponse, error) {
	requestDef := GenReqDefForStopMixJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopMixJobResponse), nil
	}
}

// 更新自动录制配置
//
// 更新自动录制配置，租户可以开启自动单流录制或者停用自动单流录制。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudRTCClient) UpdateAutoRecord(request *model.UpdateAutoRecordRequest) (*model.UpdateAutoRecordResponse, error) {
	requestDef := GenReqDefForUpdateAutoRecord()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAutoRecordResponse), nil
	}
}

// 修改合流任务
//
// 调用此接口更新合流任务布局。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudRTCClient) UpdateMixJob(request *model.UpdateMixJobRequest) (*model.UpdateMixJobResponse, error) {
	requestDef := GenReqDefForUpdateMixJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateMixJobResponse), nil
	}
}

// RTC增值（录制）事件回调配置
//
// 调用此接口配置增值（录制）事件上报回调。
//
// 当任务发生订阅了的事件时，通过该接口配置的回调地址通知。
//
// 回调格式参考/customer-record-notify-url定义。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudRTCClient) UpdateRecordCallback(request *model.UpdateRecordCallbackRequest) (*model.UpdateRecordCallbackResponse, error) {
	requestDef := GenReqDefForUpdateRecordCallback()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateRecordCallbackResponse), nil
	}
}

// 更新录制规则
//
// 调用此接口更新录制规则。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudRTCClient) UpdateRecordRule(request *model.UpdateRecordRuleRequest) (*model.UpdateRecordRuleResponse, error) {
	requestDef := GenReqDefForUpdateRecordRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateRecordRuleResponse), nil
	}
}

// 开关访问控制
//
// 调用此接口开启或关闭URL鉴权。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudRTCClient) UpdateUrlAuth(request *model.UpdateUrlAuthRequest) (*model.UpdateUrlAuthResponse, error) {
	requestDef := GenReqDefForUpdateUrlAuth()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateUrlAuthResponse), nil
	}
}
