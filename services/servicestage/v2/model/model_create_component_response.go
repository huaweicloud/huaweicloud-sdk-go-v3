package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateComponentResponse struct {

	// 应用组件ID。
	Id *string `json:"id,omitempty" xml:"id"`

	// 应用组件名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 取值0或1。  0：表示正常状态。  1：表示正在删除。
	Status *int32 `json:"status,omitempty" xml:"status"`

	Runtime *RuntimeType `json:"runtime,omitempty" xml:"runtime"`

	Category *ComponentCategory `json:"category,omitempty" xml:"category"`

	SubCategory *ComponentSubCategory `json:"sub_category,omitempty" xml:"sub_category"`

	// 描述。
	Description *string `json:"description,omitempty" xml:"description"`

	// 项目ID。
	ProjectId *string `json:"project_id,omitempty" xml:"project_id"`

	// 应用ID。
	ApplicationId *string `json:"application_id,omitempty" xml:"application_id"`

	Source *SourceObject `json:"source,omitempty" xml:"source"`

	Build *BuildInfo `json:"build,omitempty" xml:"build"`

	// 流水线Id列表，最多10个。
	PipelineIds *[]string `json:"pipeline_ids,omitempty" xml:"pipeline_ids"`

	// 创建时间。
	CreateTime *int64 `json:"create_time,omitempty" xml:"create_time"`

	// 修改时间。
	UpdateTime     *int64 `json:"update_time,omitempty" xml:"update_time"`
	HttpStatusCode int    `json:"-"`
}

func (o CreateComponentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateComponentResponse struct{}"
	}

	return strings.Join([]string{"CreateComponentResponse", string(data)}, " ")
}
