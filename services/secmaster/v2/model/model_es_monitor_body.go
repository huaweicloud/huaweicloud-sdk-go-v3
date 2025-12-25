package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EsMonitorBody 监视器结果
type EsMonitorBody struct {

	// 租户ID
	DomainId *string `json:"domain_id,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 工作空间ID
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// 数据空间ID
	DataspaceId *string `json:"dataspace_id,omitempty"`

	// 管道ID
	PipeId *string `json:"pipe_id,omitempty"`

	// 存储字节数
	StorageBytes *int64 `json:"storage_bytes,omitempty"`

	// 存储条数
	StorageCount *int64 `json:"storage_count,omitempty"`
}

func (o EsMonitorBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EsMonitorBody struct{}"
	}

	return strings.Join([]string{"EsMonitorBody", string(data)}, " ")
}
