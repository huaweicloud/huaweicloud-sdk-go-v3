package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisablePolicyTypeRequest Request Object
type DisablePolicyTypeRequest struct {
	Body *PolicyTypeReqBody `json:"body,omitempty"`
}

func (o DisablePolicyTypeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisablePolicyTypeRequest struct{}"
	}

	return strings.Join([]string{"DisablePolicyTypeRequest", string(data)}, " ")
}
