package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CredentialsDto 临时安全凭证。
type CredentialsDto struct {

	// 临时安全凭证的AK。
	AccessKeyId string `json:"access_key_id"`

	// 临时安全凭证的失效时间。
	Expiration *sdktime.SdkTime `json:"expiration"`

	// 临时安全凭证的SK。
	SecretAccessKey string `json:"secret_access_key"`

	// 临时安全凭证的security_token。
	SecurityToken string `json:"security_token"`
}

func (o CredentialsDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CredentialsDto struct{}"
	}

	return strings.Join([]string{"CredentialsDto", string(data)}, " ")
}
