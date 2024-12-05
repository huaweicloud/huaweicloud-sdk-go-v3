package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdatePeerLinkRequestBodyPeerLink struct {

	// 名称
	Name *string `json:"name,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`
}

func (o UpdatePeerLinkRequestBodyPeerLink) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePeerLinkRequestBodyPeerLink struct{}"
	}

	return strings.Join([]string{"UpdatePeerLinkRequestBodyPeerLink", string(data)}, " ")
}
