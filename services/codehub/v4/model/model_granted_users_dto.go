package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GrantedUsersDto struct {

	// **参数解释：** 用户id。
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 用户名称。 **取值范围：** 字符串长度不少于1，不超过1000。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 用户中文名称。 **取值范围：** 字符串长度不少于1，不超过1000。
	NameCn *string `json:"name_cn,omitempty"`

	// **参数解释：** 用户iam_id。 **取值范围：** 字符串长度不少于1，不超过1000。
	Username *string `json:"username,omitempty"`

	// **参数解释：** 用户昵称。 **取值范围：** 字符串长度不少于1，不超过1000。
	NickName *string `json:"nick_name,omitempty"`

	// **参数解释：** 租户名称。 **取值范围：** 字符串长度不少于1，不超过1000。
	TenantName *string `json:"tenant_name,omitempty"`

	// **参数解释：** 用户邮箱。 **取值范围：** 字符串长度不少于1，不超过1000。
	Email *string `json:"email,omitempty"`

	// **参数解释：** 用户iam_id。 **取值范围：** 字符串长度不少于1，不超过1000。
	IamId *string `json:"iam_id,omitempty"`

	// **参数解释：** license的状态。
	ServiceLicenseStatus *int32 `json:"service_license_status,omitempty"`
}

func (o GrantedUsersDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GrantedUsersDto struct{}"
	}

	return strings.Join([]string{"GrantedUsersDto", string(data)}, " ")
}
