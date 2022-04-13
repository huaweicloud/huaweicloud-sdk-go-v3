package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateLoadbalancerRequest struct {
	Body *CreateLoadbalancerRequestBody `json:"body,omitempty"`
}

func (o CreateLoadbalancerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLoadbalancerRequest struct{}"
	}

	return strings.Join([]string{"CreateLoadbalancerRequest", string(data)}, " ")
}
