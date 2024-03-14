package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RootDto 包含根的详细信息。根是组织层次结构中的顶级父节点，可以包含组织单元和账号。
type RootDto struct {

	// 根的唯一标识符（ID）。
	Id string `json:"id"`

	// 根的统一资源名称。
	Urn string `json:"urn"`

	// 根的名称。
	Name string `json:"name"`

	// 策略类型在当前根已启用，则该类型策略可以绑定到根或其组织单元或账号。
	PolicyTypes []PolicyTypeSummaryDto `json:"policy_types"`

	// 根的创建时间。
	CreatedAt *sdktime.SdkTime `json:"created_at"`
}

func (o RootDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RootDto struct{}"
	}

	return strings.Join([]string{"RootDto", string(data)}, " ")
}
