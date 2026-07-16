package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DomainSetId struct {

	// **参数解释**： 域名组名称 **取值范围**： 不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释**： 域名组ID **取值范围**： 不涉及
	Id *string `json:"id,omitempty"`
}

func (o DomainSetId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DomainSetId struct{}"
	}

	return strings.Join([]string{"DomainSetId", string(data)}, " ")
}
