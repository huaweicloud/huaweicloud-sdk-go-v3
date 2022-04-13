package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Compliance struct {
	// 检查项（检查规则）编号

	CheckitemId *string `json:"checkitem_id,omitempty"`
	// 检查点（检查结果）编号，检查项对同一个资源的检查结果

	CheckpointId *string `json:"checkpoint_id,omitempty"`
	// 检查规范编号，默认选第一个

	SpecId *string `json:"spec_id,omitempty"`
	// 合规检查结果，取值定义：PASSED、WARNING、FAILED、NOT_AVAILABLE。 说明： PASSED - 接受评估的所有资源都已通过安全检查。 WARNING - 某些信息缺失或配置不支持此检查。 FAILED - 至少有一个接受评估的资源未能通过安全检查。 NOT_AVAILABLE - 由于服务中断或 API 错误，无法执行检查。

	Status string `json:"status"`
	// 属性信息

	Properties *string `json:"properties,omitempty"`
}

func (o Compliance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Compliance struct{}"
	}

	return strings.Join([]string{"Compliance", string(data)}, " ")
}
