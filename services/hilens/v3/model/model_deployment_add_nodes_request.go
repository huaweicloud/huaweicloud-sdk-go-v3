package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeploymentAddNodesRequest struct {

	// 应用部署到指定节点
	NodeIds *[]string `json:"node_ids,omitempty"`

	// 应用部署ID
	DeploymentId string `json:"deployment_id"`

	// 添加节点的标签
	NodeTags *[]DeploymentTag `json:"node_tags,omitempty"`

	// 添加的节点数量
	NodeNum int32 `json:"node_num"`
}

func (o DeploymentAddNodesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeploymentAddNodesRequest struct{}"
	}

	return strings.Join([]string{"DeploymentAddNodesRequest", string(data)}, " ")
}
