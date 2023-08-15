package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RestScheduleCycleConfDto struct {

	// 子会议UUID。
	CycleSubConfID string `json:"cycleSubConfID"`

	// 会议的媒体类型。 * Voice：语音会议 * HDVideo：视频会议
	MediaTypes string `json:"mediaTypes"`

	// 会议开始时间（UTC时间）。格式：yyyy-MM-dd HH:mm。 > * 如果没有指定开始时间或填空串，则表示会议马上开始 > * 时间是UTC时间，即0时区的时间
	StartTime string `json:"startTime"`

	// 会议持续时长，单位分钟。默认30分钟。最大1440分钟（24小时），最小15分钟。
	Length int32 `json:"length"`

	// 会议是否自动启动录制，在录播类型为：录播、录播+直播时才生效。默认为不自动启动。 * 1：自动启动录制 * 0：不自动启动录制
	IsAutoRecord *int32 `json:"isAutoRecord,omitempty"`

	ConfConfigInfo *CycleSubConfConfigDto `json:"confConfigInfo,omitempty"`

	// 录播观看鉴权方式，在录播类型为:录播、直播+录播时有效。 * 0：可通过链接观看/下载 * 1：企业用户可观看/下载 * 2：与会者可观看/下载
	RecordAuthType *int32 `json:"recordAuthType,omitempty"`

	// 会议描述，长度限制为200个字符。
	Description *string `json:"description,omitempty"`
}

func (o RestScheduleCycleConfDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestScheduleCycleConfDto struct{}"
	}

	return strings.Join([]string{"RestScheduleCycleConfDto", string(data)}, " ")
}
