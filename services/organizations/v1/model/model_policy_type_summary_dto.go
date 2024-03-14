package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PolicyTypeSummaryDto 包含根中的一个策略类型和状态信息。
type PolicyTypeSummaryDto struct {

	// 与根关联的策略类型状态。要将指定类型的策略绑定到根或该根中的组织单元或账号，该策略必须在组织中可用，并在该根已启用。enabled：启用；pending_enable：启用中；disabled：禁用；pending_disable：禁用中。
	Status string `json:"status"`

	// 策略类型的名称，service_control_policy：服务控制策略；tag_policy：标签策略。
	Type string `json:"type"`
}

func (o PolicyTypeSummaryDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyTypeSummaryDto struct{}"
	}

	return strings.Join([]string{"PolicyTypeSummaryDto", string(data)}, " ")
}
