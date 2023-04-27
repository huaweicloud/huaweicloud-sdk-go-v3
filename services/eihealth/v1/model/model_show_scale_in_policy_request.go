package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowScaleInPolicyRequest struct {
}

func (o ShowScaleInPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowScaleInPolicyRequest struct{}"
	}

	return strings.Join([]string{"ShowScaleInPolicyRequest", string(data)}, " ")
}
