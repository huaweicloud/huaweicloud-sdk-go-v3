package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HttpHeader drm加密信息,添加在请求头中的鉴权信息
type HttpHeader struct {

	// 请求头中key字段名
	Key string `json:"key"`

	// 请求头中key对应的value值
	Value string `json:"value"`
}

func (o HttpHeader) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HttpHeader struct{}"
	}

	return strings.Join([]string{"HttpHeader", string(data)}, " ")
}
