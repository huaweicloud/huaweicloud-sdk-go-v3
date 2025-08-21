package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TemplateApplyRequestBody **参数解释：** 应用模板配置请求体 **约束限制：** 不涉及
type TemplateApplyRequestBody struct {

	// **参数解释：** 域名列表，多个域名使用逗号分隔 **约束限制：** 不涉及 **取值范围：** 单次最多应用50个域名 **默认取值：** 不涉及
	Resources *string `json:"resources,omitempty"`
}

func (o TemplateApplyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TemplateApplyRequestBody struct{}"
	}

	return strings.Join([]string{"TemplateApplyRequestBody", string(data)}, " ")
}
