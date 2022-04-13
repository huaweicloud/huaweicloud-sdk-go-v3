package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListStructTemplateRequest struct {
	// 待查询模板id,非必填，不传时返回项目下所有自定义结构化模板

	Id *string `json:"id,omitempty"`
}

func (o ListStructTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStructTemplateRequest struct{}"
	}

	return strings.Join([]string{"ListStructTemplateRequest", string(data)}, " ")
}
