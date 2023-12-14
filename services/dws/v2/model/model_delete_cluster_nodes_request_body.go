package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteClusterNodesRequestBody 逻辑集群节点操作请求
type DeleteClusterNodesRequestBody struct {

	// 逻辑集群节点ID列表
	NodeList []string `json:"node_list"`

	// 操作类型，clear|delete
	OperateType string `json:"operate_type"`
}

func (o DeleteClusterNodesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteClusterNodesRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteClusterNodesRequestBody", string(data)}, " ")
}
