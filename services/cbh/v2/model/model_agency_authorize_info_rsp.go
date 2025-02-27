package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AgencyAuthorizeInfoRsp 云堡垒机委托权限返回对象。
type AgencyAuthorizeInfoRsp struct {

	// 凭据管理权限信息。
	Csms *bool `json:"csms,omitempty"`

	// 密钥管理权限信息。
	Kms *bool `json:"kms,omitempty"`
}

func (o AgencyAuthorizeInfoRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgencyAuthorizeInfoRsp struct{}"
	}

	return strings.Join([]string{"AgencyAuthorizeInfoRsp", string(data)}, " ")
}
