package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateConfTokenRequest Request Object
type CreateConfTokenRequest struct {

	// 会议ID。 > 创建会议时返回的conferenceID。不是vmrConferenceID。
	ConferenceID string `json:"conferenceID"`

	// 会控Token。 > * 仅会控Token保活场景需要携带 > * 如果会话已过期并且请求中携带了密码，则进行重新鉴权并回复新的会控Token
	XConferenceAuthorization *string `json:"X-Conference-Authorization,omitempty"`

	// 会议的主持人密码。 > 对于会控Token保活场景，不对主持人密码鉴权。
	XPassword string `json:"X-Password"`

	// 请求类型。 - 1: 业务固定为1。
	XLoginType int32 `json:"X-Login-Type"`

	// 用户临时nonce token。
	XNonce *string `json:"X-Nonce,omitempty"`
}

func (o CreateConfTokenRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateConfTokenRequest struct{}"
	}

	return strings.Join([]string{"CreateConfTokenRequest", string(data)}, " ")
}
