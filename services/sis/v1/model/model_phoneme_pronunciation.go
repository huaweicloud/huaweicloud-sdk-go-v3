package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PhonemePronunciation 音素的发音打分
type PhonemePronunciation struct {

	//
	Score float32 `json:"score"`

	//
	Gop float32 `json:"gop"`
}

func (o PhonemePronunciation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PhonemePronunciation struct{}"
	}

	return strings.Join([]string{"PhonemePronunciation", string(data)}, " ")
}
