package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KeystoneUpdateProtocolRequest Request Object
type KeystoneUpdateProtocolRequest struct {

	// 身份提供商ID。
	IdpId string `json:"idp_id"`

	// 待更新的协议ID。
	ProtocolId string `json:"protocol_id"`

	Body *KeystoneUpdateProtocolRequestBody `json:"body,omitempty"`
}

func (o KeystoneUpdateProtocolRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneUpdateProtocolRequest struct{}"
	}

	return strings.Join([]string{"KeystoneUpdateProtocolRequest", string(data)}, " ")
}
