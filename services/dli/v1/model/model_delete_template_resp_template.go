package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTemplateRespTemplate 删除模板信息。
type DeleteTemplateRespTemplate struct {

	// 模板ID。
	TemplateId *int64 `json:"template_id,omitempty"`
}

func (o DeleteTemplateRespTemplate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTemplateRespTemplate struct{}"
	}

	return strings.Join([]string{"DeleteTemplateRespTemplate", string(data)}, " ")
}
