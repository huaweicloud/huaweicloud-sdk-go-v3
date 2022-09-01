package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 音素的发音打分
type PhonemePronunciation struct {

	//
	Score float32 `json:"score" xml:"score"`

	//
	Gop float32 `json:"gop" xml:"gop"`
}

func (o PhonemePronunciation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PhonemePronunciation struct{}"
	}

	return strings.Join([]string{"PhonemePronunciation", string(data)}, " ")
}
