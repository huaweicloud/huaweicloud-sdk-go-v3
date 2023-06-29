package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListOnDemandResourceRatingsRequest Request Object
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
