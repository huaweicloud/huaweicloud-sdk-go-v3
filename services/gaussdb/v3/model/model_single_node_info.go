package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SingleNodeInfo struct {

	// 节点ID。
	NodeId string `json:"node_id"`

	// 节点名称。支持中文、数字、字母、连接符-和_，长度为4-128。
	Name string `json:"name"`
}

func (o SingleNodeInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SingleNodeInfo struct{}"
	}

	return strings.Join([]string{"SingleNodeInfo", string(data)}, " ")
}
