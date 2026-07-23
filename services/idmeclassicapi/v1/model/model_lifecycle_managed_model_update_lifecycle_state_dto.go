package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LifecycleManagedModelUpdateLifecycleStateDto struct {

	// **参数解释：**  数据实例ID，用于指定待修改生命周期状态的数据实例。 获取方法请参见[分页查询实例 - ShowFindUsingPost](https://support.huaweicloud.com/api-idme/ShowFindUsingPost.html)。  **约束限制：**  不涉及。  **取值范围：**  -9223372036854775808到9223372036854775807的整数。  **默认取值：**  不涉及。
	Id string `json:"id"`

	LifecycleState *ObjectReferenceParamDto `json:"lifecycleState"`

	// **参数解释：**  更新者账号，用于记录执行设置生命周期状态操作的用户标识。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Modifier *string `json:"modifier,omitempty"`
}

func (o LifecycleManagedModelUpdateLifecycleStateDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LifecycleManagedModelUpdateLifecycleStateDto struct{}"
	}

	return strings.Join([]string{"LifecycleManagedModelUpdateLifecycleStateDto", string(data)}, " ")
}
