package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InlineResponse200PrivateKeys struct {

	// **参数解释**： 解密 SAML 断言私钥的 ID。  **取值范围**： 不涉及。
	KeyId string `json:"key_id"`

	// **参数解释**： 上传解密 SAML 断言私钥的时间，符合 ISO 8601 格式。  **取值范围**： 不涉及。
	Timestamp *sdktime.SdkTime `json:"timestamp"`
}

func (o InlineResponse200PrivateKeys) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InlineResponse200PrivateKeys struct{}"
	}

	return strings.Join([]string{"InlineResponse200PrivateKeys", string(data)}, " ")
}
