package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DataDiskCleanUpOption struct {

	// **参数解释：** 该参数用于控制腾挪节点时，是否擦除节点的除系统盘外的数据盘。 **约束限制：** 不涉及 **取值范围：** - false：腾挪节点时，不擦除节点的除系统盘外的数据盘。           - true：腾挪节点时，擦除节点的除系统盘外的数据盘。  **默认取值：** false
	Enable *bool `json:"enable,omitempty"`

	// **参数解释：** 该参数用于控制腾挪节点时，擦除节点的数据盘失败时的处理策略。 **约束限制：** 不涉及 **取值范围：** - ignore：表示清理数据盘失败时忽略错误，继续执行。 - abort：表示清理数据盘失败时立即停止，并向上报错。  **默认取值：** ignore
	OnFailure *string `json:"onFailure,omitempty"`
}

func (o DataDiskCleanUpOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataDiskCleanUpOption struct{}"
	}

	return strings.Join([]string{"DataDiskCleanUpOption", string(data)}, " ")
}
