package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListScaleOutPolicyRequest Request Object
type ListScaleOutPolicyRequest struct {
}

func (o ListScaleOutPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScaleOutPolicyRequest struct{}"
	}

	return strings.Join([]string{"ListScaleOutPolicyRequest", string(data)}, " ")
}
