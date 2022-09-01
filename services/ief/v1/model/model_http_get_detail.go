package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 执行http探测
type HttpGetDetail struct {

	// 必须要以/开头，构造结果为：协议类型://主机地址:端口路径
	Path string `json:"path" xml:"path"`

	// 探测的http端口，1到65535之间的整数
	Port int32 `json:"port" xml:"port"`

	// 请求的主机地址，默认为容器IP
	Host *string `json:"host,omitempty" xml:"host"`

	// 协议类型，HTTP或HTTPS，默认HTTP
	Scheme *string `json:"scheme,omitempty" xml:"scheme"`
}

func (o HttpGetDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HttpGetDetail struct{}"
	}

	return strings.Join([]string{"HttpGetDetail", string(data)}, " ")
}
