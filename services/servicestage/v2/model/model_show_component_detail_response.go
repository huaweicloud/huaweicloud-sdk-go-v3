package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowComponentDetailResponse struct {
	// 应用组件ID。

	Id *string `json:"id,omitempty"`
	// 应用组件名称

	Name *string `json:"name,omitempty"`
	// 取值0或1。  0：表示正常状态。  1：表示正在删除。

	Status *int32 `json:"status,omitempty"`

	Runtime *RuntimeType `json:"runtime,omitempty"`

	Category *ComponentCategory `json:"category,omitempty"`

	SubCategory *ComponentSubCategory `json:"sub_category,omitempty"`
	// 描述。

	Description *string `json:"description,omitempty"`
	// 项目ID。

	ProjectId *string `json:"project_id,omitempty"`
	// 应用ID。

	ApplicationId *string `json:"application_id,omitempty"`

	Source *SourceObject `json:"source,omitempty"`

	Build *BuildInfo `json:"build,omitempty"`
	// 流水线Id列表，最多10个。

	PipelineIds *[]string `json:"pipeline_ids,omitempty"`
	// 创建时间。

	CreateTime *int64 `json:"create_time,omitempty"`
	// 修改时间。

	UpdateTime     *int64 `json:"update_time,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowComponentDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowComponentDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowComponentDetailResponse", string(data)}, " ")
}
