package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OperateSprintReqVo **参数解释**： 批量基线/取消基线请求体，包含需要操作的计划ID列表和基线属性信息。 **约束限制**： 不涉及。
type OperateSprintReqVo struct {

	// **参数解释**： 发布/迭代计划ID列表，通过[发布/迭代计划列表查询](ListPlan.xml)接口获取，响应消息体中的**id**字段的值就是发布/迭代ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Ids []string `json:"ids"`

	Attribute *BaseLineVo `json:"attribute"`
}

func (o OperateSprintReqVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OperateSprintReqVo struct{}"
	}

	return strings.Join([]string{"OperateSprintReqVo", string(data)}, " ")
}
