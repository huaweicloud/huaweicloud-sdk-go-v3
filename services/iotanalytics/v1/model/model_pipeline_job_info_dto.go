package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PipelineJobInfoDto 管道作业基础信息，包括：管道ID、管道名称、管道类型等。
type PipelineJobInfoDto struct {

	// 管道ID
	PipelineId *string `json:"pipeline_id,omitempty"`

	// 管道名称
	PipelineName *string `json:"pipeline_name,omitempty"`

	// 数据源ID
	DataSourceId *string `json:"data_source_id,omitempty"`

	// 存储列表
	DataStoreList *[]DataStoreDto `json:"data_store_list,omitempty"`

	// 管道描述
	PipelineDescription *string `json:"pipeline_description,omitempty"`

	// 存储列表
	TagList *[]TagInfoDto `json:"tag_list,omitempty"`

	// 管道状态
	PipelineState *string `json:"pipeline_state,omitempty"`

	// 操作状态
	Status *string `json:"status,omitempty"`

	// 运行管道的RTU个数
	Rtu *int32 `json:"rtu,omitempty"`

	// 创建时间
	CreatedTime *string `json:"created_time,omitempty"`

	// 修改时间
	ModifiedTime *string `json:"modified_time,omitempty"`

	// 用户ID
	UserId *string `json:"user_id,omitempty"`

	// 已停止的管道作业是否有历史缓存数据
	HasSavepoint *bool `json:"has_savepoint,omitempty"`
}

func (o PipelineJobInfoDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PipelineJobInfoDto struct{}"
	}

	return strings.Join([]string{"PipelineJobInfoDto", string(data)}, " ")
}
