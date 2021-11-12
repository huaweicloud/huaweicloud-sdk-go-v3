package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AppAuthReq struct {
	// 需要授权的环境编号

	EnvId string `json:"env_id"`
	// APP的编号列表

	AppIds []string `json:"app_ids"`
	// API的编号列表，可以选择租户自己的API，也可以选择从云市场上购买的API。

	ApiIds []string `json:"api_ids"`
}

func (o AppAuthReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppAuthReq struct{}"
	}

	return strings.Join([]string{"AppAuthReq", string(data)}, " ")
}
