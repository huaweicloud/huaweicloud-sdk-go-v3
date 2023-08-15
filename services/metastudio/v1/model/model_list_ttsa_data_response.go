package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTtsaDataResponse Response Object
type ListTtsaDataResponse struct {

	// 任务ID。
	JobId *string `json:"jobId,omitempty"`

	// 音频数据，Base64编码，1秒内的数据。
	Audio *string `json:"audio,omitempty"`

	// 语音驱动的表情基数据。
	Blendshapes *[]string `json:"blendshapes,omitempty"`

	// 手工指定的动作库动作数据。
	Animations *[]AnimationItem `json:"animations,omitempty"`

	// 语义驱动的智能动作数据。
	Motions        *[]MotionItem `json:"motions,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListTtsaDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTtsaDataResponse struct{}"
	}

	return strings.Join([]string{"ListTtsaDataResponse", string(data)}, " ")
}
