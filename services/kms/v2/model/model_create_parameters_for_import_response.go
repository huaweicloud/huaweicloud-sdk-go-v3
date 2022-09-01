package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateParametersForImportResponse struct {

	// 密钥ID。
	KeyId *string `json:"key_id,omitempty" xml:"key_id"`

	// 密钥导入令牌。
	ImportToken *string `json:"import_token,omitempty" xml:"import_token"`

	// 导入参数到期时间，时间戳，即从1970年1月1日至该时间的总秒数。
	ExpirationTime *int64 `json:"expiration_time,omitempty" xml:"expiration_time"`

	// 加密密钥材料的公钥，base64格式。
	PublicKey      *string `json:"public_key,omitempty" xml:"public_key"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateParametersForImportResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateParametersForImportResponse struct{}"
	}

	return strings.Join([]string{"CreateParametersForImportResponse", string(data)}, " ")
}
