package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPipelineTemplatesResponse Response Object
type ListPipelineTemplatesResponse struct {

	// 起始偏移
	Offset *int32 `json:"offset,omitempty"`

	// 每页大小
	Limit *int32 `json:"limit,omitempty"`

	// 总数
	Total *int32 `json:"total,omitempty"`

	Templates      *[]PipelineTemplateSimpleVo `json:"templates,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o ListPipelineTemplatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPipelineTemplatesResponse struct{}"
	}

	return strings.Join([]string{"ListPipelineTemplatesResponse", string(data)}, " ")
}
