package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DelegatedAdministratorDto 包含有关委托管理员的信息。
type DelegatedAdministratorDto struct {

	// 将账号设置为委托管理员的日期。
	DelegationEnabledAt *sdktime.SdkTime `json:"delegation_enabled_at"`

	// 账号的唯一标识符（ID）。
	AccountId string `json:"account_id"`

	// 账号的统一资源名称。
	AccountUrn string `json:"account_urn"`

	// 账号加入组织的方式,invited：邀请加入，created：创建加入。
	JoinMethod string `json:"join_method"`

	// 账号成为组织一部分的日期。
	JoinedAt *sdktime.SdkTime `json:"joined_at"`

	// 账号名称
	AccountName string `json:"account_name"`
}

func (o DelegatedAdministratorDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DelegatedAdministratorDto struct{}"
	}

	return strings.Join([]string{"DelegatedAdministratorDto", string(data)}, " ")
}
