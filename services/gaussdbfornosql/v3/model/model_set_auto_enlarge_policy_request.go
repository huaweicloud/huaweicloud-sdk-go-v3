package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetAutoEnlargePolicyRequest Request Object
type SetAutoEnlargePolicyRequest struct {
	Body *SetAutoPolicyRequestBody `json:"body,omitempty"`
}

func (o SetAutoEnlargePolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetAutoEnlargePolicyRequest struct{}"
	}

	return strings.Join([]string{"SetAutoEnlargePolicyRequest", string(data)}, " ")
}
