package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchUpdateHostResourceRequestBody struct {

	// **参数解释**: 是否全量处理所有主机Agent，false的时候host_ids有值 **约束限制**: 必填 **取值范围**: - true：是。 - false：否。  **默认取值**: 不涉及
	OperateAll bool `json:"operate_all"`

	// **参数解释**: 批量修改的HostId列表。operate_all参数为false时,需要填写此批量查询条件,operate_all参数为true时不处理host_ids参数 **约束限制**： 不涉及 **取值范围**: 最小值0，最大值200 **默认取值**： 不涉及
	HostIds *[]string `json:"host_ids,omitempty"`

	// **参数解释**： 资源限制类型：默认规则or自定义or自适应 **约束限制**： 不涉及 **取值范围**： - default：默认类型。 - customized：用户自定义类型。 - adaptive：自适应类型。  **默认取值**： 不涉及
	Mode string `json:"mode"`

	// **参数解释**: cpu最大值 **约束限制**: 不涉及 **取值范围**: 字符长度0-32位 **默认取值**: 不涉及
	CpuLimit string `json:"cpu_limit"`

	// **参数解释**: 内存最大值 **约束限制**: 不涉及 **取值范围**: 字符长度0-32位 **默认取值**: 不涉及
	MemLimit string `json:"mem_limit"`
}

func (o BatchUpdateHostResourceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateHostResourceRequestBody struct{}"
	}

	return strings.Join([]string{"BatchUpdateHostResourceRequestBody", string(data)}, " ")
}
