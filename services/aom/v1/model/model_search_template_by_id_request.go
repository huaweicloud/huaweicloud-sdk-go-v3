package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type SearchTemplateByIdRequest struct {

	// 方案id。
	TemplateId string `json:"template_id"`

	// 模板共享类型，默认为private。
	ShareType *string `json:"share_type,omitempty"`
}

func (o SearchTemplateByIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchTemplateByIdRequest struct{}"
	}

	return strings.Join([]string{"SearchTemplateByIdRequest", string(data)}, " ")
}
