package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AgencyAuthorizeInfo 设置云堡垒机委托授权凭据管理、密钥管理服务权限请求对象。（kms和csms必须传一个）
type AgencyAuthorizeInfo struct {

	// 凭据管理权限信息。
	Csms *bool `json:"csms,omitempty"`

	// 密钥管理权限信息。
	Kms *bool `json:"kms,omitempty"`
}

func (o AgencyAuthorizeInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgencyAuthorizeInfo struct{}"
	}

	return strings.Join([]string{"AgencyAuthorizeInfo", string(data)}, " ")
}
