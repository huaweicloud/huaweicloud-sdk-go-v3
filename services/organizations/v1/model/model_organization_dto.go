package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OrganizationDto 包含有关组织的详细信息。组织是帐号集合，使用合并计费集中管理，由组织单元构成的层次结构，并通过策略控制。
type OrganizationDto struct {

	// 组织的唯一标识符（ID）。
	Id string `json:"id"`

	// 组织的统一资源名称。
	Urn string `json:"urn"`

	// 组织管理帐号的唯一标识符（ID）。
	ManagementAccountId string `json:"management_account_id"`

	// 组织的管理帐号的名称。
	ManagementAccountName string `json:"management_account_name"`

	// 组织的创建时间。
	CreatedAt *sdktime.SdkTime `json:"created_at"`
}

func (o OrganizationDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OrganizationDto struct{}"
	}

	return strings.Join([]string{"OrganizationDto", string(data)}, " ")
}
