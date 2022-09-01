package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type RelatedIntention struct {

	// 意图名称。
	Intention string `json:"intention" xml:"intention"`

	// 意图置信度。
	Confidence *float64 `json:"confidence,omitempty" xml:"confidence"`
}

func (o RelatedIntention) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RelatedIntention struct{}"
	}

	return strings.Join([]string{"RelatedIntention", string(data)}, " ")
}
