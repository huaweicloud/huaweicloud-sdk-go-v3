package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetRecyclePolicyRequest Request Object
type SetRecyclePolicyRequest struct {
	Body *RecyclePolicyRequestBody `json:"body,omitempty"`
}

func (o SetRecyclePolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetRecyclePolicyRequest struct{}"
	}

	return strings.Join([]string{"SetRecyclePolicyRequest", string(data)}, " ")
}
