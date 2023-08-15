package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateLogStreamRequest Request Object
type UpdateLogStreamRequest struct {

	// 日志组ID，获取方式请参见：获取帐号ID、项目ID、日志组ID、日志流ID。  缺省值：None 最小长度：36 最大长度：36
	LogGroupId string `json:"log_group_id"`

	// 日志流ID，获取方式请参见：获取帐号ID、项目ID、日志组ID、日志流ID。  缺省值：None 最小长度：36 最大长度：36
	LogStreamId string `json:"log_stream_id"`

	Body *UpdateLogStreamParams `json:"body,omitempty"`
}

func (o UpdateLogStreamRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLogStreamRequest struct{}"
	}

	return strings.Join([]string{"UpdateLogStreamRequest", string(data)}, " ")
}
