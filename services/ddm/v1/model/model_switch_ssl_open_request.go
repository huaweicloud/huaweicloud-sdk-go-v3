package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SwitchSslOpenRequest struct {

	// true:  打开 false: 关闭
	SslEnabled bool `json:"ssl_enabled"`
}

func (o SwitchSslOpenRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchSslOpenRequest struct{}"
	}

	return strings.Join([]string{"SwitchSslOpenRequest", string(data)}, " ")
}
