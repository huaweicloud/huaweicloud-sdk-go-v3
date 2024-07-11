package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateApplicationPermissionsRequest Request Object
type BatchUpdateApplicationPermissionsRequest struct {
	Body *BatchUpdateApplicationPermissionsRequestBody `json:"body,omitempty"`
}

func (o BatchUpdateApplicationPermissionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateApplicationPermissionsRequest struct{}"
	}

	return strings.Join([]string{"BatchUpdateApplicationPermissionsRequest", string(data)}, " ")
}
