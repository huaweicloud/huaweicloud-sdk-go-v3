package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SlowLogDownloadInfo struct {

	// 慢日志ID
	Id *string `json:"id,omitempty"`

	// 实例ID
	InstanceId *string `json:"instance_id,omitempty"`

	// 节点ID
	NodeId *string `json:"node_id,omitempty"`

	// 工作流ID
	WorkflowId *string `json:"workflow_id,omitempty"`

	// 文件名
	FileName *string `json:"file_name,omitempty"`

	// 文件大小, 单位：Byte
	FileSize *string `json:"file_size,omitempty"`

	// 文件下载链接
	FileLink *string `json:"file_link,omitempty"`

	// 桶名称
	BucketName *string `json:"bucket_name,omitempty"`

	// 创建时间
	CreatedAt *int64 `json:"created_at,omitempty"`

	// 更新时间
	UpdatedAt *int64 `json:"updated_at,omitempty"`

	// 版本
	Version *string `json:"version,omitempty"`

	// 状态
	Status *string `json:"status,omitempty"`

	// 消息
	Message *string `json:"message,omitempty"`
}

func (o SlowLogDownloadInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SlowLogDownloadInfo struct{}"
	}

	return strings.Join([]string{"SlowLogDownloadInfo", string(data)}, " ")
}
