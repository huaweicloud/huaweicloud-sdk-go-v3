package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PublicKey struct {

	// 密钥id
	Id string `json:"id" xml:"id"`

	// 密钥
	Key string `json:"key" xml:"key"`

	// 密钥名称
	Title string `json:"title" xml:"title"`
}

func (o PublicKey) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublicKey struct{}"
	}

	return strings.Join([]string{"PublicKey", string(data)}, " ")
}
