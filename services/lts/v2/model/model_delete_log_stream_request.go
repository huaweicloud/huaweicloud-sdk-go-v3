package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteLogStreamRequest Request Object
type DeleteLogStreamRequest struct {

	// 租户想删除的日志流所在的日志组的groupid，一般为36位字符串。
	LogGroupId string `json:"log_group_id"`

	// 需要删除的日志流ID，获取方式请参见：获取账号ID、项目ID、日志组ID、日志流ID（https://support.huaweicloud.com/api-lts/lts_api_0006.html）。
	LogStreamId string `json:"log_stream_id"`
}

func (o DeleteLogStreamRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLogStreamRequest struct{}"
	}

	return strings.Join([]string{"DeleteLogStreamRequest", string(data)}, " ")
}
