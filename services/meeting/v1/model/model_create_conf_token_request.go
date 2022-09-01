package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateConfTokenRequest struct {

	// 会议ID。
	ConferenceID string `json:"conferenceID" xml:"conferenceID"`

	// 如果携带该值，则表示是保活消息，如果会话已过期并且请求中携带了密码，则进行重新鉴权并回复新的会话标识。 该头域统一为BASE64编码。
	XConferenceAuthorization *string `json:"X-Conference-Authorization,omitempty" xml:"X-Conference-Authorization"`

	// 会议的主持人密码。 从创建会议的返回响应参数获取。 对于会控Token保活场景，可以不携带会议密码。
	XPassword string `json:"X-Password" xml:"X-Password"`

	// 请求类型。 - 1: 业务固定为1。
	XLoginType int32 `json:"X-Login-Type" xml:"X-Login-Type"`

	// 用户临时nonce token。
	XNonce *string `json:"X-Nonce,omitempty" xml:"X-Nonce"`
}

func (o CreateConfTokenRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateConfTokenRequest struct{}"
	}

	return strings.Join([]string{"CreateConfTokenRequest", string(data)}, " ")
}
