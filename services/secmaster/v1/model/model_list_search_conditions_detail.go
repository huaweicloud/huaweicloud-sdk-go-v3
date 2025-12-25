package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListSearchConditionsDetail struct {

	// 检索条件ID
	ConditionId string `json:"condition_id"`

	// 检索条件名称
	ConditionName string `json:"condition_name"`
}

func (o ListSearchConditionsDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSearchConditionsDetail struct{}"
	}

	return strings.Join([]string{"ListSearchConditionsDetail", string(data)}, " ")
}
