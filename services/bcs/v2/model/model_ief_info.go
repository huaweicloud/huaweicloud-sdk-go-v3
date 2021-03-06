package model

import (
	"encoding/json"

	"strings"
)

type IefInfo struct {
	// BCS服务边缘节点部署模式，分为：随机模式（0），绑定模式（1）

	DeployMode *int32 `json:"deploy_mode,omitempty"`
}

func (o IefInfo) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "IefInfo struct{}"
	}

	return strings.Join([]string{"IefInfo", string(data)}, " ")
}
