package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InitialConfigurationReq struct {

	// 加密URL的密码。由8-12位字符，包含数字、大写字母、小写字母，每种字符至少出现一次
	Password string `json:"password"`
}

func (o InitialConfigurationReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InitialConfigurationReq struct{}"
	}

	return strings.Join([]string{"InitialConfigurationReq", string(data)}, " ")
}
