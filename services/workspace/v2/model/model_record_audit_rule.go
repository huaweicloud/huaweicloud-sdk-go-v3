package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RecordAuditRule struct {

	// 录制类型。取值为： whole：表示全程录屏。 interval：表示间隔录屏。 userOperations：表示用户操作录屏。 sessionMonitoring：监听会话生命周期录屏。
	RecordType *string `json:"record_type,omitempty"`

	// 间隔录制开始时间，仅录制类型为interval时有效 \"hh:mm\"。
	IntervalRecordStartTime *string `json:"interval_record_start_time,omitempty"`

	// 间隔录制结束时间，仅录制类型为interval时有效格式 \"hh:mm\"。
	IntervalRecordEndTime *string `json:"interval_record_end_time,omitempty"`

	// 操作触发类型，仅录制类型为userOperations时有效。取值为： input：表示用户输入内容即启动录屏。 filecopy：表示用户拷贝文件即启动录屏。
	OpType *string `json:"op_type,omitempty"`

	// 是否开启音频录制开关。取值为： false：表示关闭。 true：表示开启。
	AudioRecord *bool `json:"audio_record,omitempty"`

	// 录制帧率。取值为：2/5/10/15。
	Fps *string `json:"fps,omitempty"`

	// 录制视频单文件时长（分钟）。取值为：10/20/30/60。
	Duration *string `json:"duration,omitempty"`

	// 分辨率设置。取值为：720P/1080P/original。
	Resolution *string `json:"resolution,omitempty"`

	// 是否启动关键事件审计。取值为： false：表示关闭。 true：表示开启。
	EventEnable *bool `json:"event_enable,omitempty"`

	// 文件后缀，多个用\"|\"分隔。
	FileSuffix *string `json:"file_suffix,omitempty"`

	// 注册表路径，多个用\"|\"分隔。
	RegitPaths *string `json:"regit_paths,omitempty"`

	// 应用过滤类型，black（黑名单）或者white（白名单）二选一。
	AppFilterType *string `json:"app_filter_type,omitempty"`

	// APP开启/关闭白名单，仅监控配置的白名单应用列表。
	AppWhiteList *string `json:"app_white_list,omitempty"`

	// APP开启/关闭黑名单，忽略黑名单里面的应用列表。
	AppBlackList *string `json:"app_black_list,omitempty"`
}

func (o RecordAuditRule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecordAuditRule struct{}"
	}

	return strings.Join([]string{"RecordAuditRule", string(data)}, " ")
}
