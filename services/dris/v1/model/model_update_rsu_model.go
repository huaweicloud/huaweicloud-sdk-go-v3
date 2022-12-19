package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 修改RSU型号信息结构体。
type UpdateRsuModel struct {

	// **参数说明**: RSU的厂商名称。  **取值范围**：长度不低于1不超过32，只允许中文、字母、数字、下划线（_）、问号（?）、反引号（'）、井号（#）、左小括号（(）、右小括号（)）、点（.）、逗号（,）、与（&）、百分号（%）、At（@）、感叹号（!）、连接符（-）的组合。
	ManufacturerName *string `json:"manufacturer_name,omitempty"`

	// **参数说明**: RSU型号的描述信息。  **取值范围**：长度不低于1不超过128，只允许中文、字母、数字、下划线（_）、问号（?）、反引号（'）、井号（#）、左小括号（(）、右小括号（)）、点（.）、逗号（,）、与（&）、百分号（%）、At（@）、感叹号（!）、连接符（-）的组合。
	Description *string `json:"description,omitempty"`
}

func (o UpdateRsuModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRsuModel struct{}"
	}

	return strings.Join([]string{"UpdateRsuModel", string(data)}, " ")
}
