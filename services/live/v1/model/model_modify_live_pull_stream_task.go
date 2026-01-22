package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type ModifyLivePullStreamTask struct {

	// 任务id
	TaskId string `json:"task_id"`

	// pause停用/enable启用
	Status *ModifyLivePullStreamTaskStatus `json:"status,omitempty"`

	// 拉流源URL
	SourceUrls *[]string `json:"source_urls,omitempty"`

	// 缺省值”-1”。-1：无限轮播，以结束时间为准；N：按文件列表播放N轮，以播放完成和结束时间先到的为准。不传、为空（\"\"）时按缺省值生效
	VodLoopTime *string `json:"vod_loop_time,omitempty"`

	// 缺省值 immediate_new_source，可选 immediate_new_source、continue_from_file_start、continue_from_break_point
	VodRefreshType *ModifyLivePullStreamTaskVodRefreshType `json:"vod_refresh_type,omitempty"`

	// 缺省值 0，指定播放文件的下标，从0开始表示第一个文件，最大值 len(source_urls)-1
	VodStartVideoIndex *int32 `json:"vod_start_video_index,omitempty"`

	// 任务启动时间，时间格式： \"2006-01-02T15:04:05Z\" 必须小于结束时间，缺省为当前时间
	StartTime *sdktime.SdkTime `json:"start_time,omitempty"`

	// 任务结束时间，时间格式： \"2006-01-02T15:04:05Z\" 必须大于开始时间，至多为开始时间+7天
	EndTime *sdktime.SdkTime `json:"end_time,omitempty"`

	// 缺省值 0，指定从上述指定文件的第几秒开始播放
	VodStartVideoTime *int32 `json:"vod_start_video_time,omitempty"`

	// 要回调的事件列表（不填则回调全部） - TaskStart：任务启动回调 - TaskExit：任务停止回调 - VodSourceFileStart 仅PullVodPushLive可用，文件启动切换 - VodSourceFileFinish 仅PullVodPushLive可用，文件播放完毕 - ResetTaskConfig 仅PullVodPushLive可用，任务更新 - TaskAlarm: 用于告警事件通知
	CallbackEvents *[]ModifyLivePullStreamTaskCallbackEvents `json:"callback_events,omitempty"`
}

func (o ModifyLivePullStreamTask) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyLivePullStreamTask struct{}"
	}

	return strings.Join([]string{"ModifyLivePullStreamTask", string(data)}, " ")
}

type ModifyLivePullStreamTaskStatus struct {
	value string
}

type ModifyLivePullStreamTaskStatusEnum struct {
	PAUSE  ModifyLivePullStreamTaskStatus
	ENABLE ModifyLivePullStreamTaskStatus
}

func GetModifyLivePullStreamTaskStatusEnum() ModifyLivePullStreamTaskStatusEnum {
	return ModifyLivePullStreamTaskStatusEnum{
		PAUSE: ModifyLivePullStreamTaskStatus{
			value: "pause",
		},
		ENABLE: ModifyLivePullStreamTaskStatus{
			value: "enable",
		},
	}
}

func (c ModifyLivePullStreamTaskStatus) Value() string {
	return c.value
}

func (c ModifyLivePullStreamTaskStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ModifyLivePullStreamTaskStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}

type ModifyLivePullStreamTaskVodRefreshType struct {
	value string
}

type ModifyLivePullStreamTaskVodRefreshTypeEnum struct {
	IMMEDIATE_NEW_SOURCE      ModifyLivePullStreamTaskVodRefreshType
	CONTINUE_FROM_FILE_START  ModifyLivePullStreamTaskVodRefreshType
	CONTINUE_FROM_BREAK_POINT ModifyLivePullStreamTaskVodRefreshType
}

func GetModifyLivePullStreamTaskVodRefreshTypeEnum() ModifyLivePullStreamTaskVodRefreshTypeEnum {
	return ModifyLivePullStreamTaskVodRefreshTypeEnum{
		IMMEDIATE_NEW_SOURCE: ModifyLivePullStreamTaskVodRefreshType{
			value: "immediate_new_source",
		},
		CONTINUE_FROM_FILE_START: ModifyLivePullStreamTaskVodRefreshType{
			value: "continue_from_file_start",
		},
		CONTINUE_FROM_BREAK_POINT: ModifyLivePullStreamTaskVodRefreshType{
			value: "continue_from_break_point",
		},
	}
}

func (c ModifyLivePullStreamTaskVodRefreshType) Value() string {
	return c.value
}

func (c ModifyLivePullStreamTaskVodRefreshType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ModifyLivePullStreamTaskVodRefreshType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}

type ModifyLivePullStreamTaskCallbackEvents struct {
	value string
}

type ModifyLivePullStreamTaskCallbackEventsEnum struct {
	TASK_START             ModifyLivePullStreamTaskCallbackEvents
	VOD_SOURCE_FILE_START  ModifyLivePullStreamTaskCallbackEvents
	VOD_SOURCE_FILE_FINISH ModifyLivePullStreamTaskCallbackEvents
	RESET_TASK_CONFIG      ModifyLivePullStreamTaskCallbackEvents
	TASK_EXIT              ModifyLivePullStreamTaskCallbackEvents
	TASK_ALARM             ModifyLivePullStreamTaskCallbackEvents
}

func GetModifyLivePullStreamTaskCallbackEventsEnum() ModifyLivePullStreamTaskCallbackEventsEnum {
	return ModifyLivePullStreamTaskCallbackEventsEnum{
		TASK_START: ModifyLivePullStreamTaskCallbackEvents{
			value: "TaskStart",
		},
		VOD_SOURCE_FILE_START: ModifyLivePullStreamTaskCallbackEvents{
			value: "VodSourceFileStart",
		},
		VOD_SOURCE_FILE_FINISH: ModifyLivePullStreamTaskCallbackEvents{
			value: "VodSourceFileFinish",
		},
		RESET_TASK_CONFIG: ModifyLivePullStreamTaskCallbackEvents{
			value: "ResetTaskConfig",
		},
		TASK_EXIT: ModifyLivePullStreamTaskCallbackEvents{
			value: "TaskExit",
		},
		TASK_ALARM: ModifyLivePullStreamTaskCallbackEvents{
			value: "TaskAlarm",
		},
	}
}

func (c ModifyLivePullStreamTaskCallbackEvents) Value() string {
	return c.value
}

func (c ModifyLivePullStreamTaskCallbackEvents) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ModifyLivePullStreamTaskCallbackEvents) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
