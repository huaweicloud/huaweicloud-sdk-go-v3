package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RemoveTagRequest Request Object
type RemoveTagRequest struct {

	// **参数解释：**  应用的唯一标识。  - 于2023年06月01日之前创建的应用，其唯一标识为该应用的名称。 - 于2023年06月01日之后创建的应用，其唯一标识为该应用的ID。 获取方法请参见[获取运行服务清单 - ListEnvs](https://support.huaweicloud.com/api-idme/ListApps.html)。  **约束限制：**  不涉及。  **取值范围：**  - 于2023年06月01日之前创建的应用：由英文字母和数字组成，长度为1-36个字符。 - 于2023年06月01日之后创建的应用：由英文字母和数字组成，且长度为32个字符。  **默认取值：**  不涉及。
	Identifier string `json:"identifier"`

	// **参数解释：**  数据模型的英文名称。  **约束限制：**  不涉及。  **取值范围：**  以大写字母开头，只能包含字母、数字、“_”，且长度为1-60个字符。  **默认取值：**  不涉及。
	ModelName string `json:"modelName"`

	Body *RdmParamVoTagOperationDto `json:"body,omitempty"`
}

func (o RemoveTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoveTagRequest struct{}"
	}

	return strings.Join([]string{"RemoveTagRequest", string(data)}, " ")
}
