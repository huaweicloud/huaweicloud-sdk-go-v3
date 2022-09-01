package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 单个音素的发音评测结果
type Phoneme struct {

	// 音标（ARPAbet音标系统）
	Arpa string `json:"arpa" xml:"arpa"`

	// 音标（国际音标系统）
	Ipa string `json:"ipa" xml:"ipa"`

	//
	StartTime float32 `json:"start_time" xml:"start_time"`

	//
	EndTime float32 `json:"end_time" xml:"end_time"`

	Fluency *PhonemeFluency `json:"fluency" xml:"fluency"`

	Pronunciation *PhonemePronunciation `json:"pronunciation" xml:"pronunciation"`
}

func (o Phoneme) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Phoneme struct{}"
	}

	return strings.Join([]string{"Phoneme", string(data)}, " ")
}
