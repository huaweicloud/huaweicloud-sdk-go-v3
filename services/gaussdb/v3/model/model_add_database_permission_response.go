package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddDatabasePermissionResponse Response Object
type AddDatabasePermissionResponse struct {

	// 授予用户权限的任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AddDatabasePermissionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddDatabasePermissionResponse struct{}"
	}

	return strings.Join([]string{"AddDatabasePermissionResponse", string(data)}, " ")
}
