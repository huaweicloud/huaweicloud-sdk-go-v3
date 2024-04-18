package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DecodeAuthorizationMessageResponse Response Object
type DecodeAuthorizationMessageResponse struct {

	// 鉴权失败原因的明文。
	DecodedMessage *string `json:"decoded_message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DecodeAuthorizationMessageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DecodeAuthorizationMessageResponse struct{}"
	}

	return strings.Join([]string{"DecodeAuthorizationMessageResponse", string(data)}, " ")
}
