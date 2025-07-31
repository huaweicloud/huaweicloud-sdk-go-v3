package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCicdConfigurationResponse Response Object
type ShowCicdConfigurationResponse struct {

	// **参数解释** CI/CD标识 **约束限制** 不涉及 **取值范围** 字符长度0-128位  **默认取值** 不涉及
	CicdId *string `json:"cicd_id,omitempty"`

	// **参数解释** CI/CD名称 **约束限制** 不涉及 **取值范围** 字符长度0-128位  **默认取值** 不涉及
	CicdName *string `json:"cicd_name,omitempty"`

	// **参数解释** 关联镜像扫描数量 **约束限制** 不涉及 **取值范围** 最小值0，最大值2147483647  **默认取值** 不涉及
	AssociatedImagesNum *int32 `json:"associated_images_num,omitempty"`

	// **参数解释** 漏洞白名单 **约束限制** 不涉及 **取值范围** 白名单数量范围0-1000  **默认取值** 不涉及
	VulnerabilityWhitelist *[]string `json:"vulnerability_whitelist,omitempty"`

	// **参数解释** 漏洞黑名单 **约束限制** 不涉及 **取值范围** 黑名单数量范围0-1000  **默认取值** 不涉及
	VulnerabilityBlocklist *[]string `json:"vulnerability_blocklist,omitempty"`

	// **参数解释** 镜像白名单 **约束限制** 不涉及 **取值范围** 白名单数量范围0-1000  **默认取值** 不涉及
	ImageWhitelist *[]string `json:"image_whitelist,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowCicdConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCicdConfigurationResponse struct{}"
	}

	return strings.Join([]string{"ShowCicdConfigurationResponse", string(data)}, " ")
}
