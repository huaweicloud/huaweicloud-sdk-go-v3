package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type County struct {

	// 区县的编码。
	Code string `json:"code"`

	// 区县的名称，根据请求的语言会传递回对应的语言的名称，目前仅支持中文。
	Name string `json:"name"`
}

func (o County) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "County struct{}"
	}

	return strings.Join([]string{"County", string(data)}, " ")
}
