package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListPipelinesRequest struct {

	// 项目ID
	ProjectId string `json:"project_id"`

	// 语言类型 中文:zh-cn 英文:en-us，默认en-us
	XLanguage *string `json:"X-Language,omitempty"`

	Body *ListPipelineQuery `json:"body,omitempty"`
}

func (o ListPipelinesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPipelinesRequest struct{}"
	}

	return strings.Join([]string{"ListPipelinesRequest", string(data)}, " ")
}
