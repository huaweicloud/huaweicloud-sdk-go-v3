package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMissingIndexDetailsRequest Request Object
type ListMissingIndexDetailsRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	Body *ListMissingIndexDetailsRequestBody `json:"body,omitempty"`
}

func (o ListMissingIndexDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMissingIndexDetailsRequest struct{}"
	}

	return strings.Join([]string{"ListMissingIndexDetailsRequest", string(data)}, " ")
}
