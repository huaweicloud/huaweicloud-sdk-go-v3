package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建秘钥对请求体。
type CreateKeypairRequestBody struct {
	// 密钥对名称。 新创建的密钥名称不能和已有密钥名称相同。

	Name string `json:"name"`
	// 导入的公钥信息。 建议导入的公钥长度不大于1024字节。 > 长度超过1024字节会导致边缘实例注入该密钥失败。

	PublicKey *string `json:"public_key,omitempty"`
}

func (o CreateKeypairRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateKeypairRequestBody struct{}"
	}

	return strings.Join([]string{"CreateKeypairRequestBody", string(data)}, " ")
}
