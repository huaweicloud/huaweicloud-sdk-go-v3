package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAggregateResourceConfigRequest Request Object
type ShowAggregateResourceConfigRequest struct {
	Body *AggregateResourceConfigRequest `json:"body,omitempty"`
}

func (o ShowAggregateResourceConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAggregateResourceConfigRequest struct{}"
	}

	return strings.Join([]string{"ShowAggregateResourceConfigRequest", string(data)}, " ")
}
