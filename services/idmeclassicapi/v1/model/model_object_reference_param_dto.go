package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ObjectReferenceParamDto struct {

	// **参数解释：**  数据实例ID，用于唯一标识引用的数据实例。 获取方法请参见[分页查询实例 - ShowFindUsingPost](https://support.huaweicloud.com/api-idme/ShowFindUsingPost.html)。  **约束限制：**  不涉及。  **取值范围：**  -9223372036854775808到9223372036854775807的整数。  **默认取值：**  不涉及。
	Id string `json:"id"`

	// **参数解释：**  类名，用于指定引用对象的类类型。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Clazz *string `json:"clazz,omitempty"`
}

func (o ObjectReferenceParamDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ObjectReferenceParamDto struct{}"
	}

	return strings.Join([]string{"ObjectReferenceParamDto", string(data)}, " ")
}
