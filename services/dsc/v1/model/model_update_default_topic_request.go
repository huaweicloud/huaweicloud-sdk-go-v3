package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDefaultTopicRequest Request Object
type UpdateDefaultTopicRequest struct {
	Body *DefaultTopicRequest `json:"body,omitempty"`
}

func (o UpdateDefaultTopicRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDefaultTopicRequest struct{}"
	}

	return strings.Join([]string{"UpdateDefaultTopicRequest", string(data)}, " ")
}
