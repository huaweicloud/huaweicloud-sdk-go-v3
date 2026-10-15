package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchQueryPrecheckResultReq 批量查询预检查结果请求体
type BatchQueryPrecheckResultReq struct {

	// **参数解释**：  批量查询任务ID列表，通过该参数指定需要查询预检查结果的一个或多个任务。  **约束限制**：  - 不能包含空对象。 - 集合中的元素取值严格匹配UUID规则。 - 任务id不能重复。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Jobs []string `json:"jobs"`
}

func (o BatchQueryPrecheckResultReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchQueryPrecheckResultReq struct{}"
	}

	return strings.Join([]string{"BatchQueryPrecheckResultReq", string(data)}, " ")
}
