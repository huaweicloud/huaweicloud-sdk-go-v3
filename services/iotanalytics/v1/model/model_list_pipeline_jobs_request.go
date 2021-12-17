package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListPipelineJobsRequest struct {
	// 数据存储Id

	DataStoreId *string `json:"data_store_id,omitempty"`
	// 存储组Id

	DataStoreGroupId *string `json:"data_store_group_id,omitempty"`
	// 数据源Id

	DataSourceId *string `json:"data_source_id,omitempty"`
	// 管道名称

	PipelineName *string `json:"pipeline_name,omitempty"`
	// 包含的管道类名

	OperatorClassName *string `json:"operator_class_name,omitempty"`
	// 偏移量，表示从此偏移量开始查询，offset大于等于0

	Offset *int64 `json:"offset,omitempty"`
	// 每页显示的条目数量

	Limit *int32 `json:"limit,omitempty"`
	// 立即同步作业状态，默认是false

	SyncStatus *bool `json:"sync_status,omitempty"`
}

func (o ListPipelineJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPipelineJobsRequest struct{}"
	}

	return strings.Join([]string{"ListPipelineJobsRequest", string(data)}, " ")
}
