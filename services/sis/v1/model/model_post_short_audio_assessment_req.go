package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PostShortAudioAssessmentReq
type PostShortAudioAssessmentReq struct {
	Config *AudioConfig `json:"config"`

	// 语音数据，Base64编码，要求Base64编码后大小不超过1M。  注意评测接口使用次数定义为：每8秒的音频作为一次，不足8秒按一次计算。例如传入4秒或8秒的音频，都算作使用一次，传入9秒的音频则视为调用2次。
	AudioData string `json:"audio_data"`

	// 评测文本
	RefText string `json:"ref_text"`
}

func (o PostShortAudioAssessmentReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PostShortAudioAssessmentReq struct{}"
	}

	return strings.Join([]string{"PostShortAudioAssessmentReq", string(data)}, " ")
}
