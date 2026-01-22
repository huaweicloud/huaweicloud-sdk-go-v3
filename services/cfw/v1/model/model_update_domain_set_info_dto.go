package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateDomainSetInfoDto struct {

	// **参数解释**： 域名组名称 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Name string `json:"name"`

	// **参数解释**： 域名组描述 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Description *string `json:"description,omitempty"`
}

func (o UpdateDomainSetInfoDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDomainSetInfoDto struct{}"
	}

	return strings.Join([]string{"UpdateDomainSetInfoDto", string(data)}, " ")
}
