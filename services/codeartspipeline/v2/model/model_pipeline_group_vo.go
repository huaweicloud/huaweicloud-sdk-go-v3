package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PipelineGroupVo PipelineGroupVo
type PipelineGroupVo struct {

	// 分组ID
	Id string `json:"id"`

	// 租户ID
	DomainId string `json:"domain_id"`

	// 项目ID
	ProjectId string `json:"project_id"`

	// 分组名
	Name string `json:"name"`

	// 父分组ID
	ParentId *string `json:"parent_id,omitempty"`

	// 分组路径ID
	PathId string `json:"path_id"`

	// 序号
	Ordinal *int32 `json:"ordinal,omitempty"`

	// 创建用户ID
	Creator string `json:"creator"`

	// 更新用户ID
	Updater *string `json:"updater,omitempty"`

	// 创建时间
	CreateTime *int64 `json:"create_time,omitempty"`

	// 更新时间
	UpdateTime *int64 `json:"update_time,omitempty"`

	// 子分组列表
	Children *[]PipelineGroupVo `json:"children,omitempty"`
}

func (o PipelineGroupVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PipelineGroupVo struct{}"
	}

	return strings.Join([]string{"PipelineGroupVo", string(data)}, " ")
}
