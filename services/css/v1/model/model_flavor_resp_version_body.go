package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FlavorRespVersionBody struct {

	// **参数解释**： 版本id。 **取值范围**： 不涉及
	Id *string `json:"id,omitempty"`

	// **参数解释**： 版本名称。 **取值范围**： 不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释**： 规格信息。 **取值范围**： 不涉及
	Flavors *[]Flavor `json:"flavors,omitempty"`
}

func (o FlavorRespVersionBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlavorRespVersionBody struct{}"
	}

	return strings.Join([]string{"FlavorRespVersionBody", string(data)}, " ")
}
