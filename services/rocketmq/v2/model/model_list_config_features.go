package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListConfigFeatures struct {

	// **参数解释**： 特性ID。     **取值范围**： 不涉及。
	FeatureId *string `json:"featureId,omitempty"`

	// **参数解释**： 状态。  **取值范围**： - 1：特性开启。 - 0：特性关闭。
	Status *int32 `json:"status,omitempty"`

	// **参数解释**： 特性描述。  **取值范围**： 不涉及。
	Description *string `json:"description,omitempty"`
}

func (o ListConfigFeatures) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConfigFeatures struct{}"
	}

	return strings.Join([]string{"ListConfigFeatures", string(data)}, " ")
}
