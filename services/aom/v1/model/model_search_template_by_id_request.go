package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchTemplateByIdRequest Request Object
type SearchTemplateByIdRequest struct {

	// 方案id。
	TemplateId string `json:"template_id"`

	// 模板共享类型，默认为private。可选public private
	ShareType string `json:"share_type"`
}

func (o SearchTemplateByIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchTemplateByIdRequest struct{}"
	}

	return strings.Join([]string{"SearchTemplateByIdRequest", string(data)}, " ")
}
