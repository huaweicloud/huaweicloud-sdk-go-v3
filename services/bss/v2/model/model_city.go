package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type City struct {

	// 城市的编码。
	Code string `json:"code"`

	// 城市的名称，根据请求参数X-Language的取值返回对应语言的名称，目前仅支持中文。
	Name string `json:"name"`
}

func (o City) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "City struct{}"
	}

	return strings.Join([]string{"City", string(data)}, " ")
}
