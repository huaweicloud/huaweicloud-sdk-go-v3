package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PipelineTemplate struct {

	// 模板名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 模板id
	Id *string `json:"id,omitempty" xml:"id"`

	// 描述信息
	Description *string `json:"description,omitempty" xml:"description"`

	// 区域id
	RegionId *string `json:"region_id,omitempty" xml:"region_id"`

	// 预览链接
	Url *string `json:"url,omitempty" xml:"url"`
}

func (o PipelineTemplate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PipelineTemplate struct{}"
	}

	return strings.Join([]string{"PipelineTemplate", string(data)}, " ")
}
