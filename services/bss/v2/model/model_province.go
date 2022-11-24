package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Province struct {

	// 省份的编码。
	Code string `json:"code"`

	// 省份的名称，根据请求的语言会传递回对应的语言的名称，目前仅支持中文。
	Name string `json:"name"`
}

func (o Province) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Province struct{}"
	}

	return strings.Join([]string{"Province", string(data)}, " ")
}
