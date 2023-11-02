package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DataobjectSearchConditionLogics struct {

	// 表达式名称
	Name *string `json:"name,omitempty"`
}

func (o DataobjectSearchConditionLogics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataobjectSearchConditionLogics struct{}"
	}

	return strings.Join([]string{"DataobjectSearchConditionLogics", string(data)}, " ")
}
