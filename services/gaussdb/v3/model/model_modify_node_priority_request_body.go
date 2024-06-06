package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ModifyNodePriorityRequestBody struct {

	// 故障倒换优先级。  故障倒换优先级的取值范围为1~16以及-1。取正数时数字越小，优先级越大，即故障倒换时，主节点会优先倒换到优先级高的只读节点上，优先级相同的只读节点选为主节点的概率相同。取-1时表示节点不参与故障倒换，当单可用区实例超过两个只读节点，或者多可用区实例修改后的可用区多于1个时可以设置成-1。
	Priority string `json:"priority"`
}

func (o ModifyNodePriorityRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyNodePriorityRequestBody struct{}"
	}

	return strings.Join([]string{"ModifyNodePriorityRequestBody", string(data)}, " ")
}
