package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListAimSendTasksResponse struct {

	// 短信记录查询结果列表。
	AimSendTasks *[]AimSendTasksMode `json:"aim_send_tasks,omitempty"`

	PageInfo       *Page `json:"page_info,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ListAimSendTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAimSendTasksResponse struct{}"
	}

	return strings.Join([]string{"ListAimSendTasksResponse", string(data)}, " ")
}
