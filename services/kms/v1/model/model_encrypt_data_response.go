/*
 * kms
 *
 * KMS v1.0 API, open API
 *
 */

package model

import (
	"encoding/json"
	"strings"
)

// Response Object
type EncryptDataResponse struct {
	// 密钥ID。
	KeyId string `json:"key_id,omitempty"`
	// DEK密文16进制，两位表示1byte。
	CipherText string `json:"cipher_text,omitempty"`
}

func (o EncryptDataResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"EncryptDataResponse", string(data)}, " ")
}
