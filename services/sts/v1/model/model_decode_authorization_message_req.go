package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DecodeAuthorizationMessageReq 接口/v5/decode-authorization-message的Http请求体。
type DecodeAuthorizationMessageReq struct {

	// 加密的鉴权失败原因。
	EncodedMessage string `json:"encoded_message"`
}

func (o DecodeAuthorizationMessageReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DecodeAuthorizationMessageReq struct{}"
	}

	return strings.Join([]string{"DecodeAuthorizationMessageReq", string(data)}, " ")
}
