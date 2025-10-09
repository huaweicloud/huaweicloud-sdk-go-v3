package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFeatureGatesResponse Response Object
type ShowFeatureGatesResponse struct {

	// **参数解释：** API类型。 **约束限制：** 该值不可修改 **取值范围：** - Configuration  **默认取值：** 不涉及
	Kind *string `json:"kind,omitempty"`

	// **参数解释：** API版本。 **约束限制：** 该值不可修改 **取值范围：** - v3.1  **默认取值：** 不涉及
	ApiVersion *string `json:"apiVersion,omitempty"`

	// **参数解释：** 特性开关详情 **取值范围**: 不涉及
	Spec           map[string]interface{} `json:"spec,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ShowFeatureGatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFeatureGatesResponse struct{}"
	}

	return strings.Join([]string{"ShowFeatureGatesResponse", string(data)}, " ")
}
