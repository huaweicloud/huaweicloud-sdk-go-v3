package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListModelVendorsResponse Response Object
type ListModelVendorsResponse struct {

	// **参数解释**： 模型供应商列表。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	ModelVendors *[]ModelVendor `json:"model_vendors,omitempty"`

	// **参数解释**： 模型供应商个数。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListModelVendorsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListModelVendorsResponse struct{}"
	}

	return strings.Join([]string{"ListModelVendorsResponse", string(data)}, " ")
}
