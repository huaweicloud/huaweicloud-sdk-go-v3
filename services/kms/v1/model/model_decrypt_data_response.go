package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DecryptDataResponse struct {

	// 密钥ID。
	KeyId *string `json:"key_id,omitempty" xml:"key_id"`

	// 明文。
	PlainText      *string `json:"plain_text,omitempty" xml:"plain_text"`
	HttpStatusCode int     `json:"-"`
}

func (o DecryptDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DecryptDataResponse struct{}"
	}

	return strings.Join([]string{"DecryptDataResponse", string(data)}, " ")
}
