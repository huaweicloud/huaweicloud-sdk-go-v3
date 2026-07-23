package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RdmParamVoListLifecycleManagedModelUpdateLifecycleStateDto struct {

	// **参数解释：**  应用ID，用于指定实例所属的应用。 当URL路径中已能明确应用上下文时，此参数可不传。 获取方法请参见[获取租户下的应用清单 - ListApps](https://support.huaweicloud.com/api-idme/ListApps.html)。  **约束限制：**  不涉及。  **取值范围：**  由英文字母和数字组成，且长度固定为1-32个字符。  **默认取值：**  不涉及。
	ApplicationId *string `json:"applicationId,omitempty"`

	// **参数解释：**  请求参数对象数组，每个元素包含一个待修改生命周期状态的数据实例信息。支持批量操作。  **约束限制：**  单次请求不超过1000条。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Params []LifecycleManagedModelUpdateLifecycleStateDto `json:"params"`
}

func (o RdmParamVoListLifecycleManagedModelUpdateLifecycleStateDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RdmParamVoListLifecycleManagedModelUpdateLifecycleStateDto struct{}"
	}

	return strings.Join([]string{"RdmParamVoListLifecycleManagedModelUpdateLifecycleStateDto", string(data)}, " ")
}
