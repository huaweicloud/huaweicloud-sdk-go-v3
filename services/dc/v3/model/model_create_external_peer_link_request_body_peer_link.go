package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateExternalPeerLinkRequestBodyPeerLink 创建关联连接请求Body体
type CreateExternalPeerLinkRequestBodyPeerLink struct {

	// 关联连接的名字
	Name string `json:"name"`

	// 描述信息
	Description *string `json:"description,omitempty"`

	PeerSite *CreateExternalPeerLinkRequestBodyPeerLinkPeerSite `json:"peer_site"`
}

func (o CreateExternalPeerLinkRequestBodyPeerLink) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateExternalPeerLinkRequestBodyPeerLink struct{}"
	}

	return strings.Join([]string{"CreateExternalPeerLinkRequestBodyPeerLink", string(data)}, " ")
}
