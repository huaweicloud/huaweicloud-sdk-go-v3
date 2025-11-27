package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VendorAccountCreateRequest struct {

	// **参数解释：** 云广商信息。 **约束限制：** 不涉及。 **取值范围：** - RMS： 华为云。 - AWS：亚马逊。 - AZURE：微软。 - ALI：阿里云。 - VMWARE：VMware。 - OPENSTACK：openstack云平台。 - HCS：Huawei Cloud Stack。 - OTHER：其他云广商。 **默认取值：** 不涉及。
	Vendor string `json:"vendor"`

	// **参数解释：** 供应商的账户ID。 **约束限制：** 不涉及。 **取值范围：** 字符串，长度0到64个字符。 **默认取值：** 不涉及。
	AccountId string `json:"account_id"`

	// **参数解释：** 账户名。 **约束限制：** 不涉及。 **取值范围：** 字符串，长度0到64个字符。 **默认取值：** 不涉及。
	AccountName string `json:"account_name"`

	// **参数解释：** 访问密钥id。 **约束限制：** 不涉及。 **取值范围：** 字符串，长度0到64个字符。 **默认取值：** 不涉及。
	Ak string `json:"ak"`

	// **参数解释：** 访问密钥密码。 **约束限制：** 不涉及。 **取值范围：** 字符串，长度0到64个字符。 **默认取值：** 不涉及。
	Sk string `json:"sk"`
}

func (o VendorAccountCreateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VendorAccountCreateRequest struct{}"
	}

	return strings.Join([]string{"VendorAccountCreateRequest", string(data)}, " ")
}
