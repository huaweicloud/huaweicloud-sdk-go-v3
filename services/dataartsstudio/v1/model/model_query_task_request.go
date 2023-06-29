package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QueryTaskRequest struct {

	// 创建人
	UserName *string `json:"user_name,omitempty"`

	// 任务名
	Name *string `json:"name,omitempty"`

	// 数据源类型
	DataSourceType *string `json:"data_source_type,omitempty"`

	// 数据连接id
	DataConnectionId *string `json:"data_connection_id,omitempty"`

	// 开始时间
	StartTime *string `json:"start_time,omitempty"`

	// 结束时间
	EndTime *string `json:"end_time,omitempty"`

	// 目录id
	DirectoryId *string `json:"directory_id,omitempty"`

	// 桶名
	BucketName *string `json:"bucket_name,omitempty"`

	// 分页参数limit，默认值：10
	Limit *int32 `json:"limit,omitempty"`

	// 分页参数offset，默认值：0
	Offset *int32 `json:"offset,omitempty"`
}

func (o QueryTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryTaskRequest struct{}"
	}

	return strings.Join([]string{"QueryTaskRequest", string(data)}, " ")
}
