package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 修改app回调请求
type AppCallbackUrlReq struct {
	// 回调通知url地址，url必须以http://或https://开头，需要支持POST调用。

	Url string `json:"url"`
	// 回调秘钥，主要用于鉴权

	AuthKey *string `json:"auth_key,omitempty"`
}

func (o AppCallbackUrlReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppCallbackUrlReq struct{}"
	}

	return strings.Join([]string{"AppCallbackUrlReq", string(data)}, " ")
}
