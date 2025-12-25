package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePipeResponse Response Object
type CreatePipeResponse struct {

	// 创建者
	CreateBy *string `json:"create_by,omitempty"`

	// 创建时间
	CreateTime *int64 `json:"create_time,omitempty"`

	// 数据空间id
	DataspaceId *string `json:"dataspace_id,omitempty"`

	// 数据空间名称
	DataspaceName *string `json:"dataspace_name,omitempty"`

	// 描述信息
	Description *string `json:"description,omitempty"`

	// 用户domainId
	DomainId *string `json:"domain_id,omitempty"`

	// 管道ID
	PipeId *string `json:"pipe_id,omitempty"`

	// 管道名称
	PipeName *string `json:"pipe_name,omitempty"`

	// 管道类型(system-defined，系统预定义)、1(user-defined，用户自定义)
	PipeType *string `json:"pipe_type,omitempty"`

	// 项目id
	ProjectId *string `json:"project_id,omitempty"`

	// 索引分片数量
	Shards *int32 `json:"shards,omitempty"`

	// 索引存储天数
	StoragePeriod *int32 `json:"storage_period,omitempty"`

	// 更新者
	UpdateBy *string `json:"update_by,omitempty"`

	// 更新时间
	UpdateTime     *int32 `json:"update_time,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CreatePipeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePipeResponse struct{}"
	}

	return strings.Join([]string{"CreatePipeResponse", string(data)}, " ")
}
