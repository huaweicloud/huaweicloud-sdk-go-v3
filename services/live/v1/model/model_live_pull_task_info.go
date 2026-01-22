package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// LivePullTaskInfo 直播拉流转推任务信息
type LivePullTaskInfo struct {

	// 任务id
	TaskId *string `json:"task_id,omitempty"`

	// 任务状态
	Status *string `json:"status,omitempty"`

	// 创建时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 任务执行区域
	Region *string `json:"region,omitempty"`

	// 拉流源类型 PullLivePushLive：拉转推（拉直播推直播） PullVodPushLive：轮播（拉点播推直播）
	SourceType *string `json:"source_type,omitempty"`

	// 拉流源URL 当 source_type 为 PullLivePushLive 时，只能填写一个URL 当 source_type 为 PullVodPushLive 时，可以指定多个轮播源文件URL
	SourceUrls *[]string `json:"source_urls,omitempty"`

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

	// 任务启动时间，时间格式： \"2006-01-02T15:04:05Z\"
	StartTime *string `json:"start_time,omitempty"`

	// 任务结束时间，时间格式： \"2006-01-02T15:04:05Z\"  1. 结束时间必须大于开始时间； 2. 结束时间必须大于当前时间； 3. 结束时间 和 开始时间 间隔必须小于七天。
	EndTime *string `json:"end_time,omitempty"`

	// 要回调的事件列表（不填则回调全部） - TaskStart：任务启动回调 - TaskExit：任务停止回调 - VodSourceFileStart：仅PullVodPushLive可用，文件启动切换 - VodSourceFileFinish：仅PullVodPushLive可用，文件播放完毕 - ResetTaskConfig：仅PullVodPushLive可用，任务更新 - TaskAlarm：用于告警事件通知
	CallbackEvents *[]string `json:"callback_events,omitempty"`

	// 回调接收地址
	CallbackUrl *string `json:"callback_url,omitempty"`

	// 备源的类型 - PullLivePushLive：直播 注意事项： 1. 仅当source_type为PullVodPushLive生效。 2. 主直播源拉流中断时，自动使用备源进行拉流。
	BackupSourceType *string `json:"backup_source_type,omitempty"`

	// 备源URL，仅当source_type为PullVodPushLive生效。
	BackupSourceUrls *[]string `json:"backup_source_urls,omitempty"`

	// 缺省值”-1” -1：无限轮播，以结束时间为准；N：按文件列表播放N轮，以播放完成和结束时间先到为准。 不传、为空（\"\"）时按缺省值生效
	VodLoopTime *string `json:"vod_loop_time,omitempty"`

	// 指定播放文件的下标，从0开始表示第一个文件，最大值 len(source_urls)-1
	VodStartVideoIndex *int32 `json:"vod_start_video_index,omitempty"`

	// 指定从上述指定文件的第几秒开始播放
	VodStartVideoTime *int32 `json:"vod_start_video_time,omitempty"`

	// 修改任务时文件切换方式 immediate_new_source：立即播放新的拉流源内容 continue_from_file_start：从上次断流url文件重新开始推流 continue_from_break_point：从上次断流url文件断流位置开始推流（续上）
	VodRefreshType *LivePullTaskInfoVodRefreshType `json:"vod_refresh_type,omitempty"`
}

func (o LivePullTaskInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LivePullTaskInfo struct{}"
	}

	return strings.Join([]string{"LivePullTaskInfo", string(data)}, " ")
}

type LivePullTaskInfoVodRefreshType struct {
	value string
}

type LivePullTaskInfoVodRefreshTypeEnum struct {
	IMMEDIATE_NEW_SOURCE      LivePullTaskInfoVodRefreshType
	CONTINUE_FROM_FILE_START  LivePullTaskInfoVodRefreshType
	CONTINUE_FROM_BREAK_POINT LivePullTaskInfoVodRefreshType
}

func GetLivePullTaskInfoVodRefreshTypeEnum() LivePullTaskInfoVodRefreshTypeEnum {
	return LivePullTaskInfoVodRefreshTypeEnum{
		IMMEDIATE_NEW_SOURCE: LivePullTaskInfoVodRefreshType{
			value: "immediate_new_source",
		},
		CONTINUE_FROM_FILE_START: LivePullTaskInfoVodRefreshType{
			value: "continue_from_file_start",
		},
		CONTINUE_FROM_BREAK_POINT: LivePullTaskInfoVodRefreshType{
			value: "continue_from_break_point",
		},
	}
}

func (c LivePullTaskInfoVodRefreshType) Value() string {
	return c.value
}

func (c LivePullTaskInfoVodRefreshType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LivePullTaskInfoVodRefreshType) UnmarshalJSON(b []byte) error {
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
