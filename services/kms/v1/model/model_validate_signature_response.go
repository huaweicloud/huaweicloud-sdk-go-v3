package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ValidateSignatureResponse struct {
	// 密钥ID。

	KeyId *string `json:"key_id,omitempty"`
	// 签名验证合法性，“true”表示验证签名合法，“false”表示验证签名非法。

	SignatureVaild *bool `json:"signature_vaild,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ValidateSignatureResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ValidateSignatureResponse struct{}"
	}

	return strings.Join([]string{"ValidateSignatureResponse", string(data)}, " ")
}
