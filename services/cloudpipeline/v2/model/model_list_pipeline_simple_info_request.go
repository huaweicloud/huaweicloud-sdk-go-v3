package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListPipelineSimpleInfoRequest struct {
	// 语言类型 中文:zh-cn 英文:en-us，默认en-us

	XLanguage *string `json:"X-Language,omitempty"`

	Body *ListPipelineSimpleInfoRequestBody `json:"body,omitempty"`
}

func (o ListPipelineSimpleInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPipelineSimpleInfoRequest struct{}"
	}

	return strings.Join([]string{"ListPipelineSimpleInfoRequest", string(data)}, " ")
}
