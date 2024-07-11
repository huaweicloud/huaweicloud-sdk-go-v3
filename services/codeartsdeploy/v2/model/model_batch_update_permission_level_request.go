package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdatePermissionLevelRequest Request Object
type BatchUpdatePermissionLevelRequest struct {
	Body *BatchUpdatePermissionLevelRequestBody `json:"body,omitempty"`
}

func (o BatchUpdatePermissionLevelRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdatePermissionLevelRequest struct{}"
	}

	return strings.Join([]string{"BatchUpdatePermissionLevelRequest", string(data)}, " ")
}
