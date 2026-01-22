package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DomainSetResponseData struct {

	// **参数解释**： 域名组id **取值范围**： 不涉及
	Id *string `json:"id,omitempty"`

	// **参数解释**： 域名组名称 **取值范围**： 不涉及
	Name *string `json:"name,omitempty"`
}

func (o DomainSetResponseData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DomainSetResponseData struct{}"
	}

	return strings.Join([]string{"DomainSetResponseData", string(data)}, " ")
}
