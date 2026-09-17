package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchResultVo **参数解释**： 批量操作结果数据对象，包含成功数量、失败数量及详细结果列表。 **约束限制**： 不涉及。
type BatchResultVo struct {

	// **参数解释**： 批量操作成功的数量。 **取值范围**： 不涉及。
	SuccessNum float32 `json:"success_num,omitempty"`

	// **参数解释**： 批量操作失败的数量。 **取值范围**： 不涉及。
	FailNum float32 `json:"fail_num,omitempty"`

	// **参数解释**： 批量操作成功的计划列表。
	Success *[]BatchOperateInfo `json:"success,omitempty"`

	// **参数解释**： 批量操作失败的计划列表。
	Failed *[]BatchOperateInfo `json:"failed,omitempty"`
}

func (o BatchResultVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchResultVo struct{}"
	}

	return strings.Join([]string{"BatchResultVo", string(data)}, " ")
}
