package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateSearchConditionRequestBody struct {

	// 检索条件名称
	ConditionName string `json:"condition_name"`

	// 查询语句
	Query string `json:"query"`
}

func (o UpdateSearchConditionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSearchConditionRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateSearchConditionRequestBody", string(data)}, " ")
}
