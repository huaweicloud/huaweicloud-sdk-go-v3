package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateLogStreamIndexRequest Request Object
type CreateLogStreamIndexRequest struct {

	// '项目ID，账号ID，日志组ID、日志流ID，获取方式请参见：获取项目ID，获取账号ID，日志组ID、日志流ID'
	GroupId string `json:"group_id"`

	// '项目ID，账号ID，日志组ID、日志流ID，获取方式请参见：获取项目ID，获取账号ID，日志组ID、日志流ID'
	StreamId string `json:"stream_id"`

	Body *LtsIndexConfigInfo `json:"body,omitempty"`
}

func (o CreateLogStreamIndexRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLogStreamIndexRequest struct{}"
	}

	return strings.Join([]string{"CreateLogStreamIndexRequest", string(data)}, " ")
}
