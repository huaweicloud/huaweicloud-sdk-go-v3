package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CredentialsDto **参数解释**： 临时安全凭证。  **取值范围**： 不涉及。
type CredentialsDto struct {

	// **参数解释**： 临时安全凭证的AK。  **取值范围**： 不涉及。
	AccessKeyId string `json:"access_key_id"`

	// **参数解释**： 临时安全凭证的失效时间。  **取值范围**： 不涉及。
	Expiration *sdktime.SdkTime `json:"expiration"`

	// **参数解释**： 临时安全凭证的SK。  **取值范围**： 不涉及。
	SecretAccessKey string `json:"secret_access_key"`

	// **参数解释**： 临时安全凭证的security_token。  **取值范围**： 不涉及。
	SecurityToken string `json:"security_token"`
}

func (o CredentialsDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CredentialsDto struct{}"
	}

	return strings.Join([]string{"CredentialsDto", string(data)}, " ")
}
