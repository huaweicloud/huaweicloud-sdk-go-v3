package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IoTdbConnectionInfoResp 创建外部推送通道请求结构体
type IoTdbConnectionInfoResp struct {

	// 鉴权用户名
	Username *string `json:"username,omitempty"`

	// 鉴权密码
	Password *string `json:"password,omitempty"`
}

func (o IoTdbConnectionInfoResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IoTdbConnectionInfoResp struct{}"
	}

	return strings.Join([]string{"IoTdbConnectionInfoResp", string(data)}, " ")
}
