package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteDatabasePermissionResponse struct {

	// 删除数据库用户的数据库权限任务id。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteDatabasePermissionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDatabasePermissionResponse struct{}"
	}

	return strings.Join([]string{"DeleteDatabasePermissionResponse", string(data)}, " ")
}
