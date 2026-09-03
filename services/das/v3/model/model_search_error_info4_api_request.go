package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchErrorInfo4ApiRequest Request Object
type SearchErrorInfo4ApiRequest struct {

	// 连接ID
	ConnectionId string `json:"connection_id"`

	// 解析任务ID
	TaskId int64 `json:"task_id"`

	// 文件名称
	FileName *string `json:"file_name,omitempty"`

	// 数据库名称
	DbName *string `json:"db_name,omitempty"`

	// 表名称
	TableName *string `json:"table_name,omitempty"`
}

func (o SearchErrorInfo4ApiRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchErrorInfo4ApiRequest struct{}"
	}

	return strings.Join([]string{"SearchErrorInfo4ApiRequest", string(data)}, " ")
}
