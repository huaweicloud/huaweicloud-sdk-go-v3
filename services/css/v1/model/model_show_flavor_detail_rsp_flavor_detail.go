package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowFlavorDetailRspFlavorDetail struct {

	// **参数解释**： 属性名称。 **取值范围**：   - cpu：规格CPU属性。   - mem：规格内存属性。   - diskrange：规格磁盘取值范围属性。   - flavorTypeCn：中文规格分类属性。   - flavorTypeEn：英文规格分类属性。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 属性值。 **取值范围**： 不涉及
	Value *string `json:"value,omitempty"`
}

func (o ShowFlavorDetailRspFlavorDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFlavorDetailRspFlavorDetail struct{}"
	}

	return strings.Join([]string{"ShowFlavorDetailRspFlavorDetail", string(data)}, " ")
}
