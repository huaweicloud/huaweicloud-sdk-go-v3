package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportBinlogRequestBody 导出binlog解析结果请求体
type ExportBinlogRequestBody struct {

	// OBS桶名称
	BucketName string `json:"bucket_name"`

	// binlog解析任务ID
	TaskId int64 `json:"task_id"`

	Info *ExportFilterInfo `json:"info,omitempty"`
}

func (o ExportBinlogRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportBinlogRequestBody struct{}"
	}

	return strings.Join([]string{"ExportBinlogRequestBody", string(data)}, " ")
}
