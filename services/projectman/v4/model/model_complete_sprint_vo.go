package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CompleteSprintVo **参数解释**： 更新发布/迭代状态的请求体。 **约束限制**： 不涉及。
type CompleteSprintVo struct {

	// **参数解释**： 操作类型。 **约束限制**： 不涉及。 **取值范围**： - start：开始发布/迭代计划 - complete：完成发布/迭代计划 - reset：将计划状态设置为\"未开始\" - restart：重新开始发布/迭代计划 - move：将工作项移动到其他迭代 **默认取值**： 不涉及。
	Operate string `json:"operate"`

	// **参数解释**： 将工作项移动到指定迭代ID。operate参数值为complete时，需要将未完成的工作项移动到其他迭代。 **约束限制**： operate参数值为complete时需填写。空字符串表示将工作项移动到\"待规划\"。 **取值范围**： 长度为18~19个字符的数字字符串。 **默认取值**： 不涉及。
	MoveToSprintId *string `json:"move_to_sprint_id,omitempty"`
}

func (o CompleteSprintVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CompleteSprintVo struct{}"
	}

	return strings.Join([]string{"CompleteSprintVo", string(data)}, " ")
}
