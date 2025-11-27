package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVendorAccountRequest Request Object
type ListVendorAccountRequest struct {

	// **参数解释：** 分页查询每页显示的条目数量。 **约束限制：** 不涉及。 **取值范围：** 自定义，在1-500范围。 **默认取值：** 不涉及。
	Limit int32 `json:"limit"`

	// **参数解释：** 分页查询偏移量，表示从此偏移量开始查询。 **约束限制：** 不涉及。 **取值范围：** 0-2147483647。 **默认取值：** 0。
	Offset *string `json:"offset,omitempty"`

	// **参数解释：** 分页参数，上一页请求最后一个id。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Marker *string `json:"marker,omitempty"`

	// **参数解释：** 供应商。 **约束限制：** 不涉及。 **取值范围：** - RMS： 华为云。 - AWS：亚马逊。 - AZURE：微软。 - ALI：阿里云。 - VMWARE：VMware。 - OPENSTACK：openstack云平台。 - HCS：Huawei Cloud Stack。 - OTHER：其他云广商。 **默认取值：** 不涉及。
	Vendor *string `json:"vendor,omitempty"`

	// **参数解释：** 供应商的账户ID。 **约束限制：** 不涉及。 **取值范围：** 字符串，长度0到64个字符。 **默认取值：** 不涉及。
	AccountId *string `json:"account_id,omitempty"`

	// **参数解释：** 账户名。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	AccountName *string `json:"account_name,omitempty"`
}

func (o ListVendorAccountRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVendorAccountRequest struct{}"
	}

	return strings.Join([]string{"ListVendorAccountRequest", string(data)}, " ")
}
