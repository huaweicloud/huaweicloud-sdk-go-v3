package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PostAddressInfoIntl struct {

	// 收件人地址。
	Address *string `json:"address,omitempty" xml:"address"`

	// 收件人。
	Recipients *string `json:"recipients,omitempty" xml:"recipients"`

	// 收件所在地邮政编码。
	ZipCode *string `json:"zipCode,omitempty" xml:"zipCode"`

	// 收件人手机号码。
	MobilePhone *string `json:"mobilePhone,omitempty" xml:"mobilePhone"`
}

func (o PostAddressInfoIntl) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PostAddressInfoIntl struct{}"
	}

	return strings.Join([]string{"PostAddressInfoIntl", string(data)}, " ")
}
