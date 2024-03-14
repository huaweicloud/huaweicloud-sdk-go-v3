package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AccountDto 包含组织成员的账号的有关信息。
type AccountDto struct {

	// 账号的唯一标识符（ID）。
	Id string `json:"id"`

	// 账号的统一资源名称。
	Urn string `json:"urn"`

	// 账号加入组织的方式,invited：邀请加入，created：创建加入。
	JoinMethod string `json:"join_method"`

	// 账号当前的状态,ACTIVE 有效| SUSPENDED 静默| PENDING_CLOSURE 关闭中
	Status string `json:"status"`

	// 账号加入组织的日期。
	JoinedAt *sdktime.SdkTime `json:"joined_at"`

	// 账号名称
	Name string `json:"name"`
}

func (o AccountDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccountDto struct{}"
	}

	return strings.Join([]string{"AccountDto", string(data)}, " ")
}
