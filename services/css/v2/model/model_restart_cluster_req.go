package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RestartClusterReq struct {
	// 操作角色

	Type string `json:"type"`
	// 节点类型

	Value string `json:"value"`
}

func (o RestartClusterReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestartClusterReq struct{}"
	}

	return strings.Join([]string{"RestartClusterReq", string(data)}, " ")
}
