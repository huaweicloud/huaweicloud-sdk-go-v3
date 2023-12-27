package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTtsaDataResponse Response Object
type ListTtsaDataResponse struct {

	// 任务ID。
	JobId *string `json:"jobId,omitempty"`

	// 驱动任务开始时间，格式遵循：RFC 3339， 例 “2020-07-30T10:43:17Z”
	StartTime *string `json:"start_time,omitempty"`

	// 驱动任务结束时间，格式遵循：RFC 3339， 例 “2020-07-30T10:45:17Z”
	EndTime *string `json:"end_time,omitempty"`

	// 是否为尾部(任务数据已全部生成，后续没有新的数据)
	IsTail *bool `json:"is_tail,omitempty"`

	// 音频数据，Base64编码，1秒内的数据。
	Audio *string `json:"audio,omitempty"`

	// 语音驱动的表情基数据。
	Blendshapes *[]string `json:"blendshapes,omitempty"`

	// 手工指定的动作库动作数据。
	Animations *[]AnimationItem `json:"animations,omitempty"`

	// 语义驱动的智能动作数据。
	Motions *[]MotionItem `json:"motions,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListTtsaDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTtsaDataResponse struct{}"
	}

	return strings.Join([]string{"ListTtsaDataResponse", string(data)}, " ")
}
