package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateNodesInstallCmdV3RequestBody struct {

	// 集群节点名称列表
	NodeInfo *[]NodeConfig `json:"node_info,omitempty"`
}

func (o CreateNodesInstallCmdV3RequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNodesInstallCmdV3RequestBody struct{}"
	}

	return strings.Join([]string{"CreateNodesInstallCmdV3RequestBody", string(data)}, " ")
}
