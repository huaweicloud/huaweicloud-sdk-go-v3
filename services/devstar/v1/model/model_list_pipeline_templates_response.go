package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListPipelineTemplatesResponse struct {

	// 模板列表
	Templates *[]PipelineTemplate `json:"templates,omitempty" xml:"templates"`

	// 模板总数
	Count          *int32 `json:"count,omitempty" xml:"count"`
	HttpStatusCode int    `json:"-"`
}

func (o ListPipelineTemplatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPipelineTemplatesResponse struct{}"
	}

	return strings.Join([]string{"ListPipelineTemplatesResponse", string(data)}, " ")
}
