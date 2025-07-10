package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateKeystorePermissionRequest Request Object
type UpdateKeystorePermissionRequest struct {
	Body *UpdateKeystorePermissionRequestBody `json:"body,omitempty"`
}

func (o UpdateKeystorePermissionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateKeystorePermissionRequest struct{}"
	}

	return strings.Join([]string{"UpdateKeystorePermissionRequest", string(data)}, " ")
}
