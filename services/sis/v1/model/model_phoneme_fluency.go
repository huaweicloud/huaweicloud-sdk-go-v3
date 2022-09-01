package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//  音素的流利度打分
type PhonemeFluency struct {

	//
	Score float32 `json:"score" xml:"score"`

	//
	Rhythm float32 `json:"rhythm" xml:"rhythm"`
}

func (o PhonemeFluency) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PhonemeFluency struct{}"
	}

	return strings.Join([]string{"PhonemeFluency", string(data)}, " ")
}
