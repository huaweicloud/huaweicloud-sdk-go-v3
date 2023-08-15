package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchQueryPrecheckResultReq 批量查询预检查结果请求体
type BatchQueryPrecheckResultReq struct {

	// 批量查询预检查结果请求列表。 约束：不能包含空对象。集合中的元素取值严格匹配UUID规则。任务id不能重复。
	Jobs []string `json:"jobs"`
}

func (o BatchQueryPrecheckResultReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchQueryPrecheckResultReq struct{}"
	}

	return strings.Join([]string{"BatchQueryPrecheckResultReq", string(data)}, " ")
}
