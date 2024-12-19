package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateExternalPeerLinkRequestBodyPeerLink struct {

	// 接入网关关联连接的名字
	Name *string `json:"name,omitempty"`

	// 描述信息
	Description *string `json:"description,omitempty"`
}

func (o UpdateExternalPeerLinkRequestBodyPeerLink) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateExternalPeerLinkRequestBodyPeerLink struct{}"
	}

	return strings.Join([]string{"UpdateExternalPeerLinkRequestBodyPeerLink", string(data)}, " ")
}
