package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BusinessCardTextConfig 用户输入的数字人名片信息。实际可以填写哪些字段取决于名片模板。
type BusinessCardTextConfig struct {

	// 姓名。
	Name *string `json:"name,omitempty"`

	// 企业或组织名称。
	Company *string `json:"company,omitempty"`

	// 职位名称。
	Title *string `json:"title,omitempty"`

	// 手机号码。
	MobilePhone *string `json:"mobile_phone,omitempty"`

	// 固话号码。
	Phone *string `json:"phone,omitempty"`

	// 电子邮件地址。
	Mail *string `json:"mail,omitempty"`

	// 地址。
	Address *string `json:"address,omitempty"`

	// 其他信息1。可填写一些公司广告语等
	Other1 *string `json:"other1,omitempty"`

	// 其他信息1。可填写一些公司广告语等
	Other2 *string `json:"other2,omitempty"`

	// 其他信息1。可填写一些公司广告语等
	Other3 *string `json:"other3,omitempty"`
}

func (o BusinessCardTextConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BusinessCardTextConfig struct{}"
	}

	return strings.Join([]string{"BusinessCardTextConfig", string(data)}, " ")
}
