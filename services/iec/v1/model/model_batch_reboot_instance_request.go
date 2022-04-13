package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BatchRebootInstanceRequest struct {
	Body *BatchRebootInstanceRequestBody `json:"body,omitempty"`
}

func (o BatchRebootInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchRebootInstanceRequest struct{}"
	}

	return strings.Join([]string{"BatchRebootInstanceRequest", string(data)}, " ")
}
