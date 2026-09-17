package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IoTdbConnectionInfo 创建外部推送通道请求结构体
type IoTdbConnectionInfo struct {

	// 鉴权用户名
	Username string `json:"username"`

	// 鉴权密码
	Password string `json:"password"`
}

func (o IoTdbConnectionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IoTdbConnectionInfo struct{}"
	}

	return strings.Join([]string{"IoTdbConnectionInfo", string(data)}, " ")
}
