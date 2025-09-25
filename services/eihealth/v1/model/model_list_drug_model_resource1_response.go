package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDrugModelResource1Response Response Object
type ListDrugModelResource1Response struct {

	// **参数解释**： 总数。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Count *int32 `json:"count,omitempty"`

	// **参数解释**： 模型列表。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Resources *[]PanguDrugModelResourceRsp `json:"resources,omitempty"`

	XResourceMappings *string `json:"X-Resource-Mappings,omitempty"`
	HttpStatusCode    int     `json:"-"`
}

func (o ListDrugModelResource1Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDrugModelResource1Response struct{}"
	}

	return strings.Join([]string{"ListDrugModelResource1Response", string(data)}, " ")
}
