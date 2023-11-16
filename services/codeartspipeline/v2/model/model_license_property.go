package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LicenseProperty 许可证规则详情
type LicenseProperty struct {

	// 是否开启
	Enable *bool `json:"enable,omitempty"`

	// 规则列表
	Rules *[]LicensePropertyRules `json:"rules,omitempty"`
}

func (o LicenseProperty) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LicenseProperty struct{}"
	}

	return strings.Join([]string{"LicenseProperty", string(data)}, " ")
}
