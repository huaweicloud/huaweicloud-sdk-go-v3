package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLogGroupsResponse Response Object
type ListLogGroupsResponse struct {

	// 日志组信息。
	LogGroups      *[]LogGroup `json:"log_groups,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ListLogGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLogGroupsResponse struct{}"
	}

	return strings.Join([]string{"ListLogGroupsResponse", string(data)}, " ")
}
