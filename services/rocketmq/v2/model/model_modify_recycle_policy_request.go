package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyRecyclePolicyRequest Request Object
type ModifyRecyclePolicyRequest struct {
	Body *ModifyRecyclePolicyReq `json:"body,omitempty"`
}

func (o ModifyRecyclePolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyRecyclePolicyRequest struct{}"
	}

	return strings.Join([]string{"ModifyRecyclePolicyRequest", string(data)}, " ")
}
