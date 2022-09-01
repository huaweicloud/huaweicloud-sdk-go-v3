package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// user详细信息。
type OsFederationInfo struct {
	IdentityProvider *IdpIdInfo `json:"identity_provider" xml:"identity_provider"`

	Protocol *ProtocolIdInfo `json:"protocol" xml:"protocol"`

	// 用户组信息。
	Groups []interface{} `json:"groups" xml:"groups"`
}

func (o OsFederationInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OsFederationInfo struct{}"
	}

	return strings.Join([]string{"OsFederationInfo", string(data)}, " ")
}
