package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AgencyCredentials 分配给用户的指定委托或信任委托的凭证
type AgencyCredentials struct {

	// 用于临时安全凭证的标识符
	AccessKeyId *string `json:"access_key_id,omitempty"`

	// 临时安全凭证到期的日期
	Expiration *int64 `json:"expiration,omitempty"`

	// 用于对请求进行签名的密钥
	SecretAccessKey *string `json:"secret_access_key,omitempty"`

	// 用于临时凭证的令牌
	SessionToken *string `json:"session_token,omitempty"`
}

func (o AgencyCredentials) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgencyCredentials struct{}"
	}

	return strings.Join([]string{"AgencyCredentials", string(data)}, " ")
}
