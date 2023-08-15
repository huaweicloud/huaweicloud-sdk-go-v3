package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RelatedIntention
type RelatedIntention struct {

	// 意图名称。
	Intention string `json:"intention"`

	// 意图置信度。
	Confidence *float64 `json:"confidence,omitempty"`
}

func (o RelatedIntention) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RelatedIntention struct{}"
	}

	return strings.Join([]string{"RelatedIntention", string(data)}, " ")
}
