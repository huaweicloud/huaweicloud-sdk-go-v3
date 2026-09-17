package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIndexUsageDetailsRequest Request Object
type ListIndexUsageDetailsRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	Body *ListIndexUsageDetailsRequestBody `json:"body,omitempty"`
}

func (o ListIndexUsageDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIndexUsageDetailsRequest struct{}"
	}

	return strings.Join([]string{"ListIndexUsageDetailsRequest", string(data)}, " ")
}
