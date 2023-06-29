package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckPermissionResponse Response Object
type CheckPermissionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CheckPermissionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckPermissionResponse struct{}"
	}

	return strings.Join([]string{"CheckPermissionResponse", string(data)}, " ")
}
