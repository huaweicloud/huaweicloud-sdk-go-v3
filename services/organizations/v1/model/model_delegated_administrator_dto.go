package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DelegatedAdministratorDto 包含有关委托管理员的信息。
type DelegatedAdministratorDto struct {

	// 将帐号设置为委托管理员的日期。
	DelegationEnabledAt *sdktime.SdkTime `json:"delegation_enabled_at"`

	// 帐号的唯一标识符（ID）。
	AccountId string `json:"account_id"`

	// 帐号的统一资源名称。
	AccountUrn string `json:"account_urn"`

	// 帐号加入组织的方式,invited：邀请加入，created：创建加入。
	JoinMethod string `json:"join_method"`

	// 帐号成为组织一部分的日期。
	JoinedAt *sdktime.SdkTime `json:"joined_at"`

	// 帐号名称。
	AccountName string `json:"account_name"`
}

func (o DelegatedAdministratorDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DelegatedAdministratorDto struct{}"
	}

	return strings.Join([]string{"DelegatedAdministratorDto", string(data)}, " ")
}
