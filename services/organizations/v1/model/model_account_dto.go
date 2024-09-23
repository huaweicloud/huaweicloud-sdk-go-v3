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

	// 账号加入组织的方式。invited：邀请加入，created：创建加入。
	JoinMethod string `json:"join_method"`

	// 账号当前的状态。active：有效； suspended：已关闭； pending_closure：关闭中。
	Status string `json:"status"`

	// 账号加入组织的日期。
	JoinedAt *sdktime.SdkTime `json:"joined_at"`

	// 账号名称
	Name string `json:"name"`

	// 手机号码
	MobilePhone *string `json:"mobile_phone,omitempty"`

	// 手机号前缀。
	IntlNumberPrefix *string `json:"intl_number_prefix,omitempty"`

	// 与此账号关联的电子邮件地址。
	Email *string `json:"email,omitempty"`

	// 描述信息。
	Description *string `json:"description,omitempty"`
}

func (o AccountDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccountDto struct{}"
	}

	return strings.Join([]string{"AccountDto", string(data)}, " ")
}
