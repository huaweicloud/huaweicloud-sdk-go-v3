package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CycleSubConf struct {

	// 子会议UUID。
	CycleSubConfID string `json:"cycleSubConfID"`

	// 会议ID。
	ConferenceID *string `json:"conferenceID,omitempty"`

	// 会议的媒体类型。 * Voice：语音 * Video：标清视频 * HDVideo：高清视频 * Data：数据
	MediaType *string `json:"mediaType,omitempty"`

	// 会议起始时间(格式：YYYY-MM-DD HH:MM)。
	StartTime *string `json:"startTime,omitempty"`

	// 会议结束时间(格式：YYYY-MM-DD HH:MM)。
	EndTime *string `json:"endTime,omitempty"`

	// 是否自动开启云录制。 - 0: 不自动启动 - 1: 自动启动
	IsAutoRecord *int32 `json:"isAutoRecord,omitempty"`

	ConfConfigInfo *CycleSubConfConfigDto `json:"confConfigInfo,omitempty"`

	// 观看/下载录播的鉴权方式。 - 0: 可通过链接观看/下载 - 1: 企业用户可观看/下载 - 2: 与会者可观看/下载
	RecordAuthType *int32 `json:"recordAuthType,omitempty"`

	// 会议描述。长度限制为200个字符。
	Description *string `json:"description,omitempty"`
}

func (o CycleSubConf) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CycleSubConf struct{}"
	}

	return strings.Join([]string{"CycleSubConf", string(data)}, " ")
}
