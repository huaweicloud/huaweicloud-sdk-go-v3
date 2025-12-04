package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListProductsRespHourly struct {

	// **参数解释**： 消息引擎的名称，该字段显示为kafka。 **取值范围**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 消息引擎的版本。 **取值范围**： - 1.1.0。 - 2.3.0。 - 2.7。
	Version *string `json:"version,omitempty"`

	// **参数解释**： 产品规格列表。 **取值范围**： 不涉及。
	Values *[]ListProductsRespValues `json:"values,omitempty"`
}

func (o ListProductsRespHourly) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProductsRespHourly struct{}"
	}

	return strings.Join([]string{"ListProductsRespHourly", string(data)}, " ")
}
