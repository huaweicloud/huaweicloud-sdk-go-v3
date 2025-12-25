package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HwcDomainContact 联系人信息
type HwcDomainContact struct {

	// 邮箱地址
	Email string `json:"email"`

	// 域名所有者
	Register string `json:"register"`

	// 联系人
	ContactName string `json:"contact_name"`

	// 联系电话
	PhoneNum string `json:"phone_num"`

	// 省份
	Province string `json:"province"`

	// 城市
	City string `json:"city"`

	// 通讯地址
	Address string `json:"address"`

	// 邮编
	ZipCode string `json:"zip_code"`
}

func (o HwcDomainContact) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HwcDomainContact struct{}"
	}

	return strings.Join([]string{"HwcDomainContact", string(data)}, " ")
}
