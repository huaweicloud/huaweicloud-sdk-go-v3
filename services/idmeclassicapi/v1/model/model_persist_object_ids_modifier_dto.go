package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PersistObjectIdsModifierDto struct {

	// **参数解释：**  数据实例ID列表，用于指定待删除的多个实例。 获取方法请参见[分页查询实例 - ShowFindUsingPost](ShowFindUsingPost.xml)。  **约束限制：**  单次请求不超过1000个。  **取值范围：**  -9223372036854775808到9223372036854775807的整数。  **默认取值：**  不涉及。
	Ids []string `json:"ids"`

	// **参数解释：**  更新者账号，用于记录执行删除操作的用户。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Modifier *string `json:"modifier,omitempty"`
}

func (o PersistObjectIdsModifierDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PersistObjectIdsModifierDto struct{}"
	}

	return strings.Join([]string{"PersistObjectIdsModifierDto", string(data)}, " ")
}
