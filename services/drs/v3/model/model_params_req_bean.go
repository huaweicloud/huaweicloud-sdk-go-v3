package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 参数信息
type ParamsReqBean struct {
	// 数据库参数名

	Key string `json:"key"`
	// 目标数据库参数值

	TargetValue string `json:"target_value"`
}

func (o ParamsReqBean) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ParamsReqBean struct{}"
	}

	return strings.Join([]string{"ParamsReqBean", string(data)}, " ")
}
