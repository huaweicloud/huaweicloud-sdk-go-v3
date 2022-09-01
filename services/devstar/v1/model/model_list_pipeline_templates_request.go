package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListPipelineTemplatesRequest struct {

	// 语言类型 中文:zh-cn 英文:en-us
	XLanguage *string `json:"X-Language,omitempty" xml:"X-Language"`

	// 区域id
	RegionId string `json:"region_id" xml:"region_id"`

	// 偏移量，表示从此偏移量开始查询，默认0
	Offset *int32 `json:"offset,omitempty" xml:"offset"`

	// 每页显示的条目数量，默认10
	Limit *int32 `json:"limit,omitempty" xml:"limit"`
}

func (o ListPipelineTemplatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPipelineTemplatesRequest struct{}"
	}

	return strings.Join([]string{"ListPipelineTemplatesRequest", string(data)}, " ")
}
