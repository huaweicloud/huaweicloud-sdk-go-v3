package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAggregationAuthorizationRequest Request Object
type CreateAggregationAuthorizationRequest struct {
	Body *AggregationAuthorizationRequest `json:"body,omitempty"`
}

func (o CreateAggregationAuthorizationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAggregationAuthorizationRequest struct{}"
	}

	return strings.Join([]string{"CreateAggregationAuthorizationRequest", string(data)}, " ")
}
