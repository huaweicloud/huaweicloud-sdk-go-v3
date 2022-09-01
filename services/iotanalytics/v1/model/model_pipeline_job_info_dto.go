package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 管道作业基础信息，包括：管道ID、管道名称、管道类型等。
type PipelineJobInfoDto struct {

	// 管道ID
	PipelineId *string `json:"pipeline_id,omitempty" xml:"pipeline_id"`

	// 管道名称
	PipelineName *string `json:"pipeline_name,omitempty" xml:"pipeline_name"`

	// 数据源ID
	DataSourceId *string `json:"data_source_id,omitempty" xml:"data_source_id"`

	// 存储列表
	DataStoreList *[]DataStoreDto `json:"data_store_list,omitempty" xml:"data_store_list"`

	// 管道描述
	PipelineDescription *string `json:"pipeline_description,omitempty" xml:"pipeline_description"`

	// 存储列表
	TagList *[]TagInfoDto `json:"tag_list,omitempty" xml:"tag_list"`

	// 管道状态
	PipelineState *string `json:"pipeline_state,omitempty" xml:"pipeline_state"`

	// 操作状态
	Status *string `json:"status,omitempty" xml:"status"`

	// 运行管道的RTU个数
	Rtu *int32 `json:"rtu,omitempty" xml:"rtu"`

	// 创建时间
	CreatedTime *string `json:"created_time,omitempty" xml:"created_time"`

	// 修改时间
	ModifiedTime *string `json:"modified_time,omitempty" xml:"modified_time"`

	// 用户ID
	UserId *string `json:"user_id,omitempty" xml:"user_id"`

	// 已停止的管道作业是否有历史缓存数据
	HasSavepoint *bool `json:"has_savepoint,omitempty" xml:"has_savepoint"`
}

func (o PipelineJobInfoDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PipelineJobInfoDto struct{}"
	}

	return strings.Join([]string{"PipelineJobInfoDto", string(data)}, " ")
}
