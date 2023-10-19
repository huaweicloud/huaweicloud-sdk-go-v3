package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCheckPermissionResponse Response Object
type BatchCheckPermissionResponse struct {
	Body           *[]CheckPermissionResult `json:"body,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o BatchCheckPermissionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCheckPermissionResponse struct{}"
	}

	return strings.Join([]string{"BatchCheckPermissionResponse", string(data)}, " ")
}
