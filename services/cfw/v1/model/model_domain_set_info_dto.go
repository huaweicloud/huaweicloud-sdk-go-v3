package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DomainSetInfoDto struct {

	// **参数解释**： 域名 **约束限制**： 不涉及 **取值范围**： 域名格式，如www.example.com **默认取值**： 不涉及
	DomainName string `json:"domain_name"`

	// **参数解释**： 域名描述 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Description *string `json:"description,omitempty"`
}

func (o DomainSetInfoDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DomainSetInfoDto struct{}"
	}

	return strings.Join([]string{"DomainSetInfoDto", string(data)}, " ")
}
