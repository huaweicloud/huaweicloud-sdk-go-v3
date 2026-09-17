package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BaseLineVo **参数解释**： 基线属性信息，指定基线或取消基线操作。 **约束限制**： 不涉及。
type BaseLineVo struct {

	// **参数解释**： 基线或取消基线操作类型。 **约束限制**： 不涉及。 **取值范围**： - baselined：基线发布/迭代计划 - unbaseline：取消基线，恢复发布/迭代计划 **默认取值**： 不涉及。
	Baseline *string `json:"baseline,omitempty"`
}

func (o BaseLineVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BaseLineVo struct{}"
	}

	return strings.Join([]string{"BaseLineVo", string(data)}, " ")
}
