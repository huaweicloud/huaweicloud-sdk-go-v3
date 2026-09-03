package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchErrorInfoSource4ApiRequest Request Object
type SearchErrorInfoSource4ApiRequest struct {

	// 连接ID
	ConnectionId string `json:"connection_id"`

	// 解析任务ID
	TaskId int64 `json:"task_id"`

	// 文件名称
	FileName *string `json:"file_name,omitempty"`

	// 数据库名称
	DbName *string `json:"db_name,omitempty"`
}

func (o SearchErrorInfoSource4ApiRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchErrorInfoSource4ApiRequest struct{}"
	}

	return strings.Join([]string{"SearchErrorInfoSource4ApiRequest", string(data)}, " ")
}
