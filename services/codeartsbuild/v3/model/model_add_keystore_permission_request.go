package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddKeystorePermissionRequest Request Object
type AddKeystorePermissionRequest struct {
	Body *AddKeystorePermissionRequestBody `json:"body,omitempty"`
}

func (o AddKeystorePermissionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddKeystorePermissionRequest struct{}"
	}

	return strings.Join([]string{"AddKeystorePermissionRequest", string(data)}, " ")
}
