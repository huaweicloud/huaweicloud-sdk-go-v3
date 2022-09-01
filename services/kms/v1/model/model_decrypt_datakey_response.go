package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DecryptDatakeyResponse struct {

	// DEK明文的16进制字符串。
	DataKey *string `json:"data_key,omitempty" xml:"data_key"`

	// DEK明文字节长度。
	DatakeyLength *string `json:"datakey_length,omitempty" xml:"datakey_length"`

	// DEK明文的SHA256值对应的16进制字符串。
	DatakeyDgst    *string `json:"datakey_dgst,omitempty" xml:"datakey_dgst"`
	HttpStatusCode int     `json:"-"`
}

func (o DecryptDatakeyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DecryptDatakeyResponse struct{}"
	}

	return strings.Join([]string{"DecryptDatakeyResponse", string(data)}, " ")
}
