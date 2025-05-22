package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkloadResourceItem **参数解释**： 资源池资源项。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
type WorkloadResourceItem struct {

	// **参数解释**： 资源名称。 **约束限制**： 不涉及。 **取值范围**： cpu：占用CPU时间片的百分比。 cpu_limit：使用CPU物理核数的百分比。 memory：指每个数据节点上可用内存资源的百分比。 concurrency：并发数。 shortQueryConcurrencyNum：简单语句并发数。 weight：网络调度时权重值。 **默认取值**： 不涉及。
	ResourceName string `json:"resource_name"`

	// **参数解释**： 资源属性值。 **约束限制**： 不涉及。 **取值范围**： 根据配置不同，取值范围不同 cpu：取值范围为1~99的整数。 cpu_limit：取值范围为0~100的整数，0表示不限制。 memory：取值范围为0~100的整数，0表示不管控。 concurrency：取值范围为-1~2147483647的整数，-1/0表示不限制。 shortQueryConcurrencyNum：取值范围为-1~2147483647的整数，-1/0表示不管控。 weight：取值范围为1~2147483647的整数，默认配置为-1。 **默认取值**： 不涉及。
	ResourceValue int32 `json:"resource_value"`

	// **参数解释**： 资源属性单位。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ValueUnit *string `json:"value_unit,omitempty"`

	// **参数解释**： 资源附加描述。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ResourceDescription *string `json:"resource_description,omitempty"`
}

func (o WorkloadResourceItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkloadResourceItem struct{}"
	}

	return strings.Join([]string{"WorkloadResourceItem", string(data)}, " ")
}
