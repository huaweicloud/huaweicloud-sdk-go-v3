package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyParameterReq 更新任务参数请求体。
type ModifyParameterReq struct {

	// 参数值对象，基于默认参数模板初始化的参数值。  key：参数名称，如“applier_thread_num”，“read_task_num”。为空时不修改参数值。 value：参数值，如“6”，“20”。key不为空时value也不可为空。
	Values []ParameterInfo `json:"values"`
}

func (o ModifyParameterReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyParameterReq struct{}"
	}

	return strings.Join([]string{"ModifyParameterReq", string(data)}, " ")
}
