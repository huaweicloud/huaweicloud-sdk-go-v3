package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RenewalResourcesRequest Request Object
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
