package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateScaleInPolicyRequest Request Object
type UpdateScaleInPolicyRequest struct {
	Body *UpdateScaleInPolicyReq `json:"body,omitempty"`
}

func (o UpdateScaleInPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateScaleInPolicyRequest struct{}"
	}

	return strings.Join([]string{"UpdateScaleInPolicyRequest", string(data)}, " ")
}
