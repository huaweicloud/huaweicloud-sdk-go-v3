package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLogStreamRequest Request Object
type ListLogStreamRequest struct {

	// 租户想查询的日志流所在的日志组的groupid，一般为36位字符串。
	LogGroupId string `json:"log_group_id"`
}

func (o ListLogStreamRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLogStreamRequest struct{}"
	}

	return strings.Join([]string{"ListLogStreamRequest", string(data)}, " ")
}
