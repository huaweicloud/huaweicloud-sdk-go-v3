package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEngineSupportFeaturesEntity **参数解释**： 实例支持的功能特性。 **取值范围**： 不涉及。
type ListEngineSupportFeaturesEntity struct {

	// **参数解释**： 功能名称。 **取值范围**： 不涉及。
	Name *string `json:"name,omitempty"`

	Properties *ListEngineSupportFeaturesPropertiesEntity `json:"properties,omitempty"`
}

func (o ListEngineSupportFeaturesEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEngineSupportFeaturesEntity struct{}"
	}

	return strings.Join([]string{"ListEngineSupportFeaturesEntity", string(data)}, " ")
}
