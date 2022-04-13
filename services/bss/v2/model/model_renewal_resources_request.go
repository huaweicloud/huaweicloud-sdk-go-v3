package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RenewalResourcesRequest struct {
	Body *RenewalResourcesReq `json:"body,omitempty"`
}

func (o RenewalResourcesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RenewalResourcesRequest struct{}"
	}

	return strings.Join([]string{"RenewalResourcesRequest", string(data)}, " ")
}
