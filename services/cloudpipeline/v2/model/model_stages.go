package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Stages 流水线阶段执行信息
type Stages struct {

	// 阶段执行结果。取值及含义：success：成功；error：失败；aborted：终止
	Result string `json:"result"`

	// 阶段执行状态。取值和含义:waiting:等待;running:执行中;verifying:待审核;suspending:挂起;completed:完成
	Status string `json:"status"`

	// 阶段名字
	Name string `json:"name"`

	// -
	Parameters *interface{} `json:"parameters"`

	// 阶段顺序
	Order int32 `json:"order"`

	// 阶段类型
	DslMethod string `json:"dsl_method"`

	// 阶段显示名称
	DisplayName string `json:"display_name"`
}

func (o Stages) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Stages struct{}"
	}

	return strings.Join([]string{"Stages", string(data)}, " ")
}
