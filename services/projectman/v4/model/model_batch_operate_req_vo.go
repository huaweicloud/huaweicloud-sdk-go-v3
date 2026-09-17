package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchOperateReqVo **参数解释**： 批量操作请求体，包含需要操作的发布/迭代ID列表。 **约束限制**： 不涉及。
type BatchOperateReqVo struct {

	// **参数解释**： 批量删除的发布计划/迭代的ID列表。通过[发布/迭代计划列表查询](ListPlan.xml)接口获取，响应消息体中的**id**字段的值就是发布/迭代ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Ids []string `json:"ids"`
}

func (o BatchOperateReqVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchOperateReqVo struct{}"
	}

	return strings.Join([]string{"BatchOperateReqVo", string(data)}, " ")
}
