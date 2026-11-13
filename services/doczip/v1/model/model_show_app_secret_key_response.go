package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAppSecretKeyResponse Response Object
type ShowAppSecretKeyResponse struct {

	// **参数解释**： 应用请求密钥。 **取值范围**： [1-512]字符。
	Key *string `json:"key,omitempty"`

	XTraceId       *string `json:"X-Trace-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowAppSecretKeyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAppSecretKeyResponse struct{}"
	}

	return strings.Join([]string{"ShowAppSecretKeyResponse", string(data)}, " ")
}
