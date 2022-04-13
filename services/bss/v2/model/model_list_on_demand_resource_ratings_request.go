package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListOnDemandResourceRatingsRequest struct {
	Body *RateOnDemandReq `json:"body,omitempty"`
}

func (o ListOnDemandResourceRatingsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOnDemandResourceRatingsRequest struct{}"
	}

	return strings.Join([]string{"ListOnDemandResourceRatingsRequest", string(data)}, " ")
}
