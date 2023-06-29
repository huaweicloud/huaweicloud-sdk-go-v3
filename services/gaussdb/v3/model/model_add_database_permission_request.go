package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddDatabasePermissionRequest Request Object
type AddDatabasePermissionRequest struct {

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *GrantDatabasePermissionRequestBody `json:"body,omitempty"`
}

func (o AddDatabasePermissionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddDatabasePermissionRequest struct{}"
	}

	return strings.Join([]string{"AddDatabasePermissionRequest", string(data)}, " ")
}
