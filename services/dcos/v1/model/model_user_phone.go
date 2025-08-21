package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UserPhone 客户电话号码
type UserPhone struct {

	// 唯一标识
	Id *string `json:"id,omitempty"`

	// 国家码
	CountryCode *string `json:"country_code,omitempty"`

	// 电话号码
	Phone *string `json:"phone,omitempty"`

	// 是否默认
	DefaultOption *bool `json:"default_option,omitempty"`
}

func (o UserPhone) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserPhone struct{}"
	}

	return strings.Join([]string{"UserPhone", string(data)}, " ")
}
