package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PolicySummaryDto 包含有关策略的信息，但不包括内容。
type PolicySummaryDto struct {

	// 一个布尔值，指示指定的策略是否为系统策略。如果为true，即为系统策略，则可以将策略附加到根、组织单元或账号，但不能编辑它。
	IsBuiltin bool `json:"is_builtin"`

	// 策略说明。
	Description string `json:"description"`

	// 策略的唯一标识符（ID）。
	Id string `json:"id"`

	// 策略的统一资源名称。
	Urn string `json:"urn"`

	// 策略的名称。
	Name string `json:"name"`

	// 策略的类型,service_control_policy：服务控制策略；tag_policy：标签策略。
	Type string `json:"type"`
}

func (o PolicySummaryDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicySummaryDto struct{}"
	}

	return strings.Join([]string{"PolicySummaryDto", string(data)}, " ")
}
