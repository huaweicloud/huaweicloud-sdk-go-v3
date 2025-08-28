package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLogStreamIndexRequest Request Object
type ListLogStreamIndexRequest struct {

	// 日志组ID，获取方式请参见：[获取项目ID，获取账号ID，日志组ID、日志流ID](lts_api_0006.xml)
	GroupId string `json:"group_id"`

	// 日志流ID，获取方式请参见：[获取项目ID，获取账号ID，日志组ID、日志流ID](lts_api_0006.xml)
	StreamId string `json:"stream_id"`
}

func (o ListLogStreamIndexRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLogStreamIndexRequest struct{}"
	}

	return strings.Join([]string{"ListLogStreamIndexRequest", string(data)}, " ")
}
