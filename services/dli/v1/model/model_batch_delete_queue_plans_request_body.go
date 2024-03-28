package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteQueuePlansRequestBody struct {

	// 表示SQL语句的类型
	PlanIds []int64 `json:"plan_ids"`
}

func (o BatchDeleteQueuePlansRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteQueuePlansRequestBody struct{}"
	}

	return strings.Join([]string{"BatchDeleteQueuePlansRequestBody", string(data)}, " ")
}
