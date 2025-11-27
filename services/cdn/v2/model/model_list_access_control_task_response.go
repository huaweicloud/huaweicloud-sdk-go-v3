package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAccessControlTaskResponse Response Object
type ListAccessControlTaskResponse struct {

	// 查询结果总数。
	Total *int32 `json:"total,omitempty"`

	// 封禁或解禁URL任务信息。
	AccessControlTasks *[]AccessControlTask `json:"access_control_tasks,omitempty"`
	HttpStatusCode     int                  `json:"-"`
}

func (o ListAccessControlTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAccessControlTaskResponse struct{}"
	}

	return strings.Join([]string{"ListAccessControlTaskResponse", string(data)}, " ")
}
