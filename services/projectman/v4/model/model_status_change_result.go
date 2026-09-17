package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StatusChangeResult **参数解释**： 状态变更结果对象，包含不可完成的AR工作项列表。 **约束限制**： 不涉及。
type StatusChangeResult struct {

	// **参数解释**： 不可完成的AR工作项列表，当完成发布/迭代时，未完成的AR工作项会列出在此。
	CannotFinishAr *[]WorkItemVo `json:"cannot_finish_ar,omitempty"`
}

func (o StatusChangeResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StatusChangeResult struct{}"
	}

	return strings.Join([]string{"StatusChangeResult", string(data)}, " ")
}
