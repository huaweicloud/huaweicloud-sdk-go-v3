package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SpoidcConfig struct {

	// 重定向地址
	RedirectUrl string `json:"redirect_url"`
}

func (o SpoidcConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SpoidcConfig struct{}"
	}

	return strings.Join([]string{"SpoidcConfig", string(data)}, " ")
}
