package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateNet2CloudPhoneServerRequestBodyPorts struct {

	// 应用端口名称，不超过16个字节，系统关键服务名称不能使用\"adb\"和\"vnc\"
	Name *string `json:"name,omitempty" xml:"name"`

	// 端口号，大于等于10000，小于等于50000
	ListenPort *int32 `json:"listen_port,omitempty" xml:"listen_port"`

	// 为\"true\"则映射出公网访问（忽略大小写）。 为其他则不映射
	InternetAccessible *string `json:"internet_accessible,omitempty" xml:"internet_accessible"`
}

func (o CreateNet2CloudPhoneServerRequestBodyPorts) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNet2CloudPhoneServerRequestBodyPorts struct{}"
	}

	return strings.Join([]string{"CreateNet2CloudPhoneServerRequestBodyPorts", string(data)}, " ")
}
