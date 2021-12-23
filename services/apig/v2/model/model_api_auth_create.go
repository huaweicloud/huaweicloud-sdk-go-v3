package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApiAuthCreate struct {
	// 需要授权的环境编号

	EnvId string `json:"env_id"`
	// APP的编号列表

	AppIds []string `json:"app_ids"`
	// API的编号列表，可以选择租户自己的API，也可以选择从云市场上购买的API。

	ApiIds []string `json:"api_ids"`
}

func (o ApiAuthCreate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiAuthCreate struct{}"
	}

	return strings.Join([]string{"ApiAuthCreate", string(data)}, " ")
}
