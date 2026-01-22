package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type LivePullStreamTask struct {

	// 任务执行区域，如果指定推流域名，则region需要与推流域名直播源站一致，缺省为租户归属region
	Region *string `json:"region,omitempty"`

	// 拉流源类型 PullLivePushLive：拉转推（拉直播推直播） PullVodPushLive：轮播（拉点播推直播）
	SourceType LivePullStreamTaskSourceType `json:"source_type"`

	// 拉流源URL 当 source_type 为 PullLivePushLive 时，只能填写一个URL 当 source_type 为 PullVodPushLive 时，可以指定多个轮播源文件URL
	SourceUrls []string `json:"source_urls"`

	// 推流域名
	Domain *string `json:"domain,omitempty"`

	// 推流app
	AppName *string `json:"app_name,omitempty"`

	// 推流流名
	StreamName *string `json:"stream_name,omitempty"`

	// 推流参数
	PushArgs *string `json:"push_args,omitempty"`

	// 完整的目标URL。 如果指定此参数，domain、app_name和stream_name需要传入空字符串或不传。
	DestinationUrl *string `json:"destination_url,omitempty"`

	// 任务启动时间，时间格式： \"2006-01-02T15:04:05Z\" 必须小于结束时间，缺省为当前时间
	StartTime *sdktime.SdkTime `json:"start_time,omitempty"`

	// 任务结束时间，时间格式： \"2006-01-02T15:04:05Z\" 必须大于开始时间，至多为开始时间+7天
	EndTime *sdktime.SdkTime `json:"end_time,omitempty"`

	// 轮播播放次数，仅当source_type为PullVodPushLive生效。缺省值”-1”。-1：无限轮播，以结束时间为准；N：按文件列表播放N轮，以播放完成和结束时间先到的为准。不传、为空（\"\"）时按缺省值生效
	VodLoopTime *string `json:"vod_loop_time,omitempty"`

	// 轮播文件播放形式，仅当source_type为PullVodPushLive生效。缺省值immediate_new_source immediate_new_source：立即播放新的文件源 continue_from_file_start：从上次断流URL文件重新播放（更新任务时有效） continue_from_break_point：从上次断流URL文件断流位置重新播放（更新任务时有效）
	VodRefreshType *LivePullStreamTaskVodRefreshType `json:"vod_refresh_type,omitempty"`

	// 指定播放文件的下标，仅当source_type为PullVodPushLive生效。缺省值 0，从0开始表示第一个文件，最大值 len(source_urls)-1
	VodStartVideoIndex *int32 `json:"vod_start_video_index,omitempty"`

	// 缺省值 0，指定从上述指定文件的第几秒开始播放，仅当source_type为PullVodPushLive生效
	VodStartVideoTime *int32 `json:"vod_start_video_time,omitempty"`

	// 备源的类型 - PullLivePushLive：直播 注意事项： 1. 仅当source_type为PullVodPushLive生效。 2. 主直播源拉流中断时，自动使用备源进行拉流。
	BackupSourceType *string `json:"backup_source_type,omitempty"`

	// 备源URL，仅当source_type为PullVodPushLive生效。
	BackupSourceUrls *[]string `json:"backup_source_urls,omitempty"`

	// 要回调的事件列表（不填则回调全部） - TaskStart：任务启动回调 - TaskExit：任务停止回调 - VodSourceFileStart：仅PullVodPushLive可用，文件启动切换 - VodSourceFileFinish：仅PullVodPushLive可用，文件播放完毕 - ResetTaskConfig：仅PullVodPushLive可用，任务更新 - TaskAlarm：用于告警事件通知
	CallbackEvents *[]LivePullStreamTaskCallbackEvents `json:"callback_events,omitempty"`

	// 回调接收地址
	CallbackUrl *string `json:"callback_url,omitempty"`
}

func (o LivePullStreamTask) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LivePullStreamTask struct{}"
	}

	return strings.Join([]string{"LivePullStreamTask", string(data)}, " ")
}

type LivePullStreamTaskSourceType struct {
	value string
}

type LivePullStreamTaskSourceTypeEnum struct {
	PULL_LIVE_PUSH_LIVE LivePullStreamTaskSourceType
	PULL_VOD_PUSH_LIVE  LivePullStreamTaskSourceType
}

func GetLivePullStreamTaskSourceTypeEnum() LivePullStreamTaskSourceTypeEnum {
	return LivePullStreamTaskSourceTypeEnum{
		PULL_LIVE_PUSH_LIVE: LivePullStreamTaskSourceType{
			value: "PullLivePushLive",
		},
		PULL_VOD_PUSH_LIVE: LivePullStreamTaskSourceType{
			value: "PullVodPushLive",
		},
	}
}

func (c LivePullStreamTaskSourceType) Value() string {
	return c.value
}

func (c LivePullStreamTaskSourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LivePullStreamTaskSourceType) UnmarshalJSON(b []byte) error {
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

type LivePullStreamTaskVodRefreshType struct {
	value string
}

type LivePullStreamTaskVodRefreshTypeEnum struct {
	IMMEDIATE_NEW_SOURCE      LivePullStreamTaskVodRefreshType
	CONTINUE_FROM_FILE_START  LivePullStreamTaskVodRefreshType
	CONTINUE_FROM_BREAK_POINT LivePullStreamTaskVodRefreshType
}

func GetLivePullStreamTaskVodRefreshTypeEnum() LivePullStreamTaskVodRefreshTypeEnum {
	return LivePullStreamTaskVodRefreshTypeEnum{
		IMMEDIATE_NEW_SOURCE: LivePullStreamTaskVodRefreshType{
			value: "immediate_new_source",
		},
		CONTINUE_FROM_FILE_START: LivePullStreamTaskVodRefreshType{
			value: "continue_from_file_start",
		},
		CONTINUE_FROM_BREAK_POINT: LivePullStreamTaskVodRefreshType{
			value: "continue_from_break_point",
		},
	}
}

func (c LivePullStreamTaskVodRefreshType) Value() string {
	return c.value
}

func (c LivePullStreamTaskVodRefreshType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LivePullStreamTaskVodRefreshType) UnmarshalJSON(b []byte) error {
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

type LivePullStreamTaskCallbackEvents struct {
	value string
}

type LivePullStreamTaskCallbackEventsEnum struct {
	TASK_START             LivePullStreamTaskCallbackEvents
	VOD_SOURCE_FILE_START  LivePullStreamTaskCallbackEvents
	VOD_SOURCE_FILE_FINISH LivePullStreamTaskCallbackEvents
	RESET_TASK_CONFIG      LivePullStreamTaskCallbackEvents
	TASK_EXIT              LivePullStreamTaskCallbackEvents
	TASK_ALARM             LivePullStreamTaskCallbackEvents
}

func GetLivePullStreamTaskCallbackEventsEnum() LivePullStreamTaskCallbackEventsEnum {
	return LivePullStreamTaskCallbackEventsEnum{
		TASK_START: LivePullStreamTaskCallbackEvents{
			value: "TaskStart",
		},
		VOD_SOURCE_FILE_START: LivePullStreamTaskCallbackEvents{
			value: "VodSourceFileStart",
		},
		VOD_SOURCE_FILE_FINISH: LivePullStreamTaskCallbackEvents{
			value: "VodSourceFileFinish",
		},
		RESET_TASK_CONFIG: LivePullStreamTaskCallbackEvents{
			value: "ResetTaskConfig",
		},
		TASK_EXIT: LivePullStreamTaskCallbackEvents{
			value: "TaskExit",
		},
		TASK_ALARM: LivePullStreamTaskCallbackEvents{
			value: "TaskAlarm",
		},
	}
}

func (c LivePullStreamTaskCallbackEvents) Value() string {
	return c.value
}

func (c LivePullStreamTaskCallbackEvents) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LivePullStreamTaskCallbackEvents) UnmarshalJSON(b []byte) error {
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
