package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BatchRebootInstanceRequest struct {
	Body *BatchRebootInstanceRequestBody `json:"body,omitempty" xml:"body"`
}

func (o BatchRebootInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchRebootInstanceRequest struct{}"
	}

	return strings.Join([]string{"BatchRebootInstanceRequest", string(data)}, " ")
}
