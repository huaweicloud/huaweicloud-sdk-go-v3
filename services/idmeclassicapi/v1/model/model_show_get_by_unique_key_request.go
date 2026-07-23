package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowGetByUniqueKeyRequest Request Object
type ShowGetByUniqueKeyRequest struct {

	// **参数解释：**  应用的唯一标识。  - 于2023年06月01日之前创建的应用，其唯一标识为该应用的名称。 - 于2023年06月01日之后创建的应用，其唯一标识为该应用的ID。 获取方法请参见[获取运行服务清单 - ListEnvs](https://support.huaweicloud.com/api-idme/ListApps.html)。  **约束限制：**  不涉及。  **取值范围：**  - 于2023年06月01日之前创建的应用：由英文字母和数字组成，长度为1-36个字符。 - 于2023年06月01日之后创建的应用：由英文字母和数字组成，且长度为32个字符。  **默认取值：**  不涉及。
	Identifier string `json:"identifier"`

	// **参数解释：**  数据模型的英文名称。  **约束限制：**  不涉及。  **取值范围：**  以大写字母开头，只能包含字母、数字、“_”，且长度为1-60个字符。  **默认取值：**  不涉及。
	ModelName string `json:"modelName"`

	// **参数解释：**  动态方法名称，格式为：getBy{uniqueName}。  uniqueName：表示“唯一键”为“是”的基本属性英文名称。 示例： - 如唯一键属性英文名称为uniqueKey，则方法名为getByUniqueKey。 - 如唯一键属性英文名称为code，则方法名为getByCode。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	GetUniqueFieldMethod string `json:"getUniqueFieldMethod"`

	Body *RdmParamVoPersistableModelUniqueKeyDto `json:"body,omitempty"`
}

func (o ShowGetByUniqueKeyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGetByUniqueKeyRequest struct{}"
	}

	return strings.Join([]string{"ShowGetByUniqueKeyRequest", string(data)}, " ")
}
