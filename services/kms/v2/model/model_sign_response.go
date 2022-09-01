package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type SignResponse struct {

	// 密钥ID。
	KeyId *string `json:"key_id,omitempty" xml:"key_id"`

	// 签名值，使用base64编码。
	Signature      *string `json:"signature,omitempty" xml:"signature"`
	HttpStatusCode int     `json:"-"`
}

func (o SignResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SignResponse struct{}"
	}

	return strings.Join([]string{"SignResponse", string(data)}, " ")
}
