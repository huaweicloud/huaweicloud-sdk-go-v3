package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建新集群信息
type CceCreateInfo struct {

	// 集群节点数
	NodeNum int64 `json:"node_num"`

	// 集群节点规格
	NodeFlavor string `json:"node_flavor"`

	// CCE集群规格
	CceFlavor string `json:"cce_flavor"`

	// 节点初始密码
	InitNodePwd string `json:"init_node_pwd"`

	// 可用区
	Az string `json:"az"`

	// 集群CPU架构类型：X86（VirtualMachine），ARM（ARM64）
	ClusterPlatformType string `json:"cluster_platform_type"`
}

func (o CceCreateInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CceCreateInfo struct{}"
	}

	return strings.Join([]string{"CceCreateInfo", string(data)}, " ")
}
