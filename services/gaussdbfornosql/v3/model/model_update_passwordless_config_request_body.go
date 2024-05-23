package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdatePasswordlessConfigRequestBody struct {

	// 要设置的免密配置,ip与网段的列表,空列表表示清空免密配置。
	ConfigIps []string `json:"config_ips"`
}

func (o UpdatePasswordlessConfigRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePasswordlessConfigRequestBody struct{}"
	}

	return strings.Join([]string{"UpdatePasswordlessConfigRequestBody", string(data)}, " ")
}
