package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AccountDto 包含组织成员的帐号的有关信息。
type AccountDto struct {

	// 帐号的唯一标识符（ID）。
	Id string `json:"id"`

	// 帐号的统一资源名称。
	Urn string `json:"urn"`

	// 帐号加入组织的方式,invited：邀请加入，created：创建加入。
	JoinMethod string `json:"join_method"`

	// 帐号加入组织的日期。
	JoinedAt *sdktime.SdkTime `json:"joined_at"`

	// 帐号名称。
	Name string `json:"name"`
}

func (o AccountDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccountDto struct{}"
	}

	return strings.Join([]string{"AccountDto", string(data)}, " ")
}
