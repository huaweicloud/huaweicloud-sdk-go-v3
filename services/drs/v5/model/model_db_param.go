package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 数据库参数名、目标数据库参数值。
type DbParam struct {

	// 数据库参数名。
	Key string `json:"key"`

	// 目标数据库参数值。
	TargetValue string `json:"target_value"`
}

func (o DbParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DbParam struct{}"
	}

	return strings.Join([]string{"DbParam", string(data)}, " ")
}
