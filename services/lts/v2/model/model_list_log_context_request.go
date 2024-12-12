package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLogContextRequest Request Object
type ListLogContextRequest struct {

	// 日志组ID，获取方式请参见：[获取项目ID，获取账号ID，日志组ID、日志流ID](lts_api_0006.xml)
	LogGroupId string `json:"log_group_id"`

	// 日志流ID，获取方式请参见：[获取项目ID，获取账号ID，日志组ID、日志流ID](lts_api_0006.xml)
	LogStreamId string `json:"log_stream_id"`

	Body *ListLogContextRequestBody `json:"body,omitempty"`
}

func (o ListLogContextRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLogContextRequest struct{}"
	}

	return strings.Join([]string{"ListLogContextRequest", string(data)}, " ")
}
