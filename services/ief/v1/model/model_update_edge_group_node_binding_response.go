package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateEdgeGroupNodeBindingResponse struct {

	// 边缘节点组ID
	Id *string `json:"id,omitempty"`

	// 边缘节点组名称
	Name *string `json:"name,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 创建时间
	CreatedAt *string `json:"created_at,omitempty"`

	// 更新时间
	UpdatedAt *string `json:"updated_at,omitempty"`

	// 边缘节点组所属的项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 边缘节点组所属账号的IAM权限，没有铂金版权限的账号无法使用节点组功能
	IamRole *string `json:"iam_role,omitempty"`

	// 边缘节点组CPU总数，为边缘节点组所绑定的边缘节点的CPU数目之和
	Cpu *int32 `json:"cpu,omitempty"`

	// 边缘节点组内存总数，为边缘节点组所绑定的边缘节点的内存之和
	Memory *int32 `json:"memory,omitempty"`

	// 边缘节点组GPU总数，为边缘节点组所绑定的边缘节点的GPU数目之和
	GpuNum *int32 `json:"gpu_num,omitempty"`

	// 绑定的边缘节点详情
	Nodes *[]EdgeNode `json:"nodes,omitempty"`

	// 绑定的边缘应用详情
	Deployments *[]GroupDeployment `json:"deployments,omitempty"`

	// 属性
	Attributes *[]Attributes `json:"attributes,omitempty"`

	// 标签
	Tags *[]Attributes `json:"tags,omitempty"`

	// 绑定操作成功的节点ID列表
	SuccessNodeAdd *[]string `json:"success_node_add,omitempty"`

	// 解绑操作成功的节点ID列表
	SuccessNodeDel *[]string `json:"success_node_del,omitempty"`

	// 绑定操作失败的节点ID列表
	FailedNodeAdd *[]string `json:"failed_node_add,omitempty"`

	// 解绑操作失败的节点ID列表
	FailedNodeDel  *[]string `json:"failed_node_del,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o UpdateEdgeGroupNodeBindingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEdgeGroupNodeBindingResponse struct{}"
	}

	return strings.Join([]string{"UpdateEdgeGroupNodeBindingResponse", string(data)}, " ")
}
