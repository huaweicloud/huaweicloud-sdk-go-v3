package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchOperateInfo **参数解释**： 批量操作结果详情，包含操作对象ID和操作人ID。
type BatchOperateInfo struct {

	// **参数解释**： 发布/迭代计划ID。 **取值范围**： 长度为18~19个字符的数字字符串。
	Id *string `json:"id,omitempty"`

	// **参数解释**： 最近更新人ID。 **取值范围**： 不涉及。
	ModifiedBy *string `json:"modified_by,omitempty"`
}

func (o BatchOperateInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchOperateInfo struct{}"
	}

	return strings.Join([]string{"BatchOperateInfo", string(data)}, " ")
}
