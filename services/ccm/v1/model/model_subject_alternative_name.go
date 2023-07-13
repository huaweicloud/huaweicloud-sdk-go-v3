package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SubjectAlternativeName struct {

	// 备用名称类型，当前仅支持DNS、IP、EMAIL、URI。 - **DNS** - **IP** - **EMAIL** - **URI**
	Type string `json:"type"`

	// 对应备用名称类型的值。   - DNS类型，长度为不超过253；   - IP类型，长度不超过39，支持IPV4、IPV6；   - EMAIL类型，长度不超过256；   - URI类型，长度不超过253。
	Value string `json:"value"`
}

func (o SubjectAlternativeName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubjectAlternativeName struct{}"
	}

	return strings.Join([]string{"SubjectAlternativeName", string(data)}, " ")
}
