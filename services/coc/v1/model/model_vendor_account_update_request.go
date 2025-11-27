package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VendorAccountUpdateRequest struct {

	// **参数解释：** 账户名。 **约束限制：** 不涉及。 **取值范围：** 字符串，长度0到64个字符。 **默认取值：** 不涉及。
	AccountName *string `json:"account_name,omitempty"`

	// **参数解释：** 访问密钥id。 **约束限制：** 不涉及。 **取值范围：** 字符串，长度0到64个字符。 **默认取值：** 不涉及。
	Ak *string `json:"ak,omitempty"`

	// **参数解释：** 访问密钥密码。 **约束限制：** 不涉及。 **取值范围：** 字符串，长度0到64个字符。 **默认取值：** 不涉及。
	Sk *string `json:"sk,omitempty"`
}

func (o VendorAccountUpdateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VendorAccountUpdateRequest struct{}"
	}

	return strings.Join([]string{"VendorAccountUpdateRequest", string(data)}, " ")
}
