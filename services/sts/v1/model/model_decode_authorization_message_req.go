package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DecodeAuthorizationMessageReq **参数解释**： 接口/v5/decode-authorization-message的Http请求体。  **约束限制**： 不涉及。  **取值范围**： 不涉及。  **默认取值**： 不涉及。
type DecodeAuthorizationMessageReq struct {

	// **参数解释**： 加密的鉴权失败原因。  **约束限制**： 长度范围为[1,10240]。  **取值范围**： 不涉及。  **默认取值**： 不涉及。
	EncodedMessage string `json:"encoded_message"`
}

func (o DecodeAuthorizationMessageReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DecodeAuthorizationMessageReq struct{}"
	}

	return strings.Join([]string{"DecodeAuthorizationMessageReq", string(data)}, " ")
}
