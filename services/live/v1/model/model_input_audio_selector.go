package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InputAudioSelector 音频选择器
type InputAudioSelector struct {

	// 音频选择器的名称。仅支持大小写字母、数字、中划线和下划线。  同一个频道中每个选择器的名称需要唯一。
	Name string `json:"name"`

	SelectorSettings *AudioSelectorSettings `json:"selector_settings,omitempty"`
}

func (o InputAudioSelector) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InputAudioSelector struct{}"
	}

	return strings.Join([]string{"InputAudioSelector", string(data)}, " ")
}
