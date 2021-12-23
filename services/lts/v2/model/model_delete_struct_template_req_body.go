package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 结构化配置
type DeleteStructTemplateReqBody struct {
	// 结构化规则ID

	StructTemplateId *string `json:"struct_template_id,omitempty"`
}

func (o DeleteStructTemplateReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteStructTemplateReqBody struct{}"
	}

	return strings.Join([]string{"DeleteStructTemplateReqBody", string(data)}, " ")
}
