package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListPipelineTemplatesRequest struct {

	// 租户ID
	TenantId string `json:"tenant_id"`

	// 语言类型 中文:zh-cn 英文:en-us，默认en-us
	XLanguage *string `json:"X-Language,omitempty"`

	Body *ListPipelineTemplatesQuery `json:"body,omitempty"`
}

func (o ListPipelineTemplatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPipelineTemplatesRequest struct{}"
	}

	return strings.Join([]string{"ListPipelineTemplatesRequest", string(data)}, " ")
}
