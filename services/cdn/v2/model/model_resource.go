package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Resource **参数解释：** 应用模板配置信息 **约束限制：** 不涉及
type Resource struct {

	// **参数解释：** 应用模板状态（域名粒度） **约束限制：** 不涉及 **取值范围：** - success: 应用模板成功 - fail: 应用模板失败 **默认取值：** 不涉及
	Status *string `json:"status,omitempty"`

	// **参数解释：** 域名 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	DomainName *string `json:"domain_name,omitempty"`

	// **参数解释：** 错误信息 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o Resource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Resource struct{}"
	}

	return strings.Join([]string{"Resource", string(data)}, " ")
}
