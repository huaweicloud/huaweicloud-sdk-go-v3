package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetAutoEnlargePoliciesRequest Request Object
type SetAutoEnlargePoliciesRequest struct {
	Body *SetAutoEnlargePoliciesRequestBody `json:"body,omitempty"`
}

func (o SetAutoEnlargePoliciesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetAutoEnlargePoliciesRequest struct{}"
	}

	return strings.Join([]string{"SetAutoEnlargePoliciesRequest", string(data)}, " ")
}
