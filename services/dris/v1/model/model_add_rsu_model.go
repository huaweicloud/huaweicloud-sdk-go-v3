package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddRsuModel 创建RSU型号信息结构体。
type AddRsuModel struct {

	// **参数说明**: RSU型号名称。  **取值范围**：长度不低于1不超过64，只允许中文、字母、数字、下划线（_）、问号（?）、反引号（'）、井号（#）、左小括号（(）、右小括号（)）、点（.）、逗号（,）、与（&）、百分号（%）、At（@）、感叹号（!）、连接符（-）的组合。
	Name string `json:"name"`

	// **参数说明**: RSU的厂商名称。  **取值范围**：长度不低于1不超过32，只允许中文、字母、数字、下划线（_）、问号（?）、反引号（'）、井号（#）、左小括号（(）、右小括号（)）、点（.）、逗号（,）、与（&）、百分号（%）、At（@）、感叹号（!）、连接符（-）的组合。
	ManufacturerName *string `json:"manufacturer_name,omitempty"`

	// **参数说明**: RSU型号的描述信息。  **取值范围**：长度不低于1不超过128，只允许中文、字母、数字、下划线（_）、问号（?）、反引号（'）、井号（#）、左小括号（(）、右小括号（)）、点（.）、逗号（,）、与（&）、百分号（%）、At（@）、感叹号（!）、连接符（-）的组合。
	Description *string `json:"description,omitempty"`
}

func (o AddRsuModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddRsuModel struct{}"
	}

	return strings.Join([]string{"AddRsuModel", string(data)}, " ")
}
