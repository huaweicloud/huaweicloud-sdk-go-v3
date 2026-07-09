package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VersionInfo 服务实例支持功能版本信息。
type VersionInfo struct {

	// 支持EIP的实例版本。
	RequireEip *string `json:"require_eip,omitempty"`

	// 支持IAM登录的实例版本。
	IamLogin *string `json:"iam_login,omitempty"`

	// 支持管理员登录的实例版本。
	AdminLogin *string `json:"admin_login,omitempty"`

	// 支持浮动IPv6的实例版本。
	FloatIpv6 *string `json:"float_ipv6,omitempty"`
}

func (o VersionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VersionInfo struct{}"
	}

	return strings.Join([]string{"VersionInfo", string(data)}, " ")
}
