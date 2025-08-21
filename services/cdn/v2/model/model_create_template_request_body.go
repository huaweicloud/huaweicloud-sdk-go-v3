package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTemplateRequestBody **参数解释：** 创建域名模板配置 **约束限制：** 不涉及
type CreateTemplateRequestBody struct {

	// **参数解释：** 域名模板名称 **约束限制：** 不涉及 **取值范围：** - 1-100个字符 - 仅支持字母、数字、中文、下划线（_）、中横线（-） **默认取值：** 不涉及
	TmlName string `json:"tml_name"`

	// **参数解释：** 域名模板描述 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Remark *string `json:"remark,omitempty"`

	Configs *TemplateConfigs `json:"configs"`
}

func (o CreateTemplateRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTemplateRequestBody struct{}"
	}

	return strings.Join([]string{"CreateTemplateRequestBody", string(data)}, " ")
}
