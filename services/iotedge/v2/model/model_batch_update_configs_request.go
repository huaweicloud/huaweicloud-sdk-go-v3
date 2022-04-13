package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BatchUpdateConfigsRequest struct {
	Body *BatchUpdateConfigs `json:"body,omitempty"`
}

func (o BatchUpdateConfigsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateConfigsRequest struct{}"
	}

	return strings.Join([]string{"BatchUpdateConfigsRequest", string(data)}, " ")
}
