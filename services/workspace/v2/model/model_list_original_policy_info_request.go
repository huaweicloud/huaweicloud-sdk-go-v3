package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListOriginalPolicyInfoRequest Request Object
type ListOriginalPolicyInfoRequest struct {
}

func (o ListOriginalPolicyInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOriginalPolicyInfoRequest struct{}"
	}

	return strings.Join([]string{"ListOriginalPolicyInfoRequest", string(data)}, " ")
}
