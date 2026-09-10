package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateClusterReq 创建集群的请求参数。
type CreateClusterReq struct {

	// 新建的集群名称，名称只能包含数字、英文字母和下划线，但不能是纯数字，且不能以下划线开头。长度限制：1~128个字符。 说明：集群名称不区分大小写，系统会自动转换为小写。
	ClusterName string `json:"cluster_name"`

	// 集群的描述信息。
	Description *string `json:"description,omitempty"`

	// 集群规格。
	FlavorId string `json:"flavor_id"`

	// 集群的收费模式。只能设置为“1”，表示按照CU时收费。
	ChargeMode int32 `json:"charge_mode"`

	// 队列的虚拟私有云（VPC）的网段。建议使用网段：10.0.0.0/8~28，172.16.0.0/12~28，192.168.0.0/16~28。
	CidrInVpc *string `json:"cidr_in_vpc,omitempty"`

	// 集群需要绑定的工作空间ID。
	Workspaces *[]string `json:"workspaces,omitempty"`
}

func (o CreateClusterReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateClusterReq struct{}"
	}

	return strings.Join([]string{"CreateClusterReq", string(data)}, " ")
}
