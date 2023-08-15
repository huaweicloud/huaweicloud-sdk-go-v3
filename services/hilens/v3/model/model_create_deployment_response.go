package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDeploymentResponse Response Object
type CreateDeploymentResponse struct {
	Template *DeploymentTemplate `json:"template,omitempty"`

	// 部署成功失败的理由
	Reason *string `json:"reason,omitempty"`

	// 已经就绪的实例节点数
	ReadyReplicas *int32 `json:"ready_replicas,omitempty"`

	// 实例节点节点数
	Replicas *int32 `json:"replicas,omitempty"`

	// 应用部署描述，最大长度255，不允许^ ~ # $ % & * < > ( ) [ ] { } ' \" \\
	Description *string `json:"description,omitempty"`

	// 创建时间
	CreatedAt *string `json:"created_at,omitempty"`

	// 应用部署来源: HiLens市场(hlm) or aigallery(aig) or 自定义(userdefined)
	Source *string `json:"source,omitempty"`

	// 应用部署版本
	ApiVersion *string `json:"api_version,omitempty"`

	// 应用部署的指定节点，与clouster_id二选一
	NodeIds *[]string `json:"node_ids,omitempty"`

	// 应用部署的集群ID，与node_id二选一
	ClusterId *string `json:"cluster_id,omitempty"`

	// 更新时间
	UpdatedAt *string `json:"updated_at,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 部署名称
	Name *string `json:"name,omitempty"`

	// 部署ID
	Id *string `json:"id,omitempty"`

	// RUNNING：运行， FREEZE：冻结， UNFREEZE: 解冻， CREATING：创建中， CREATE_FAILED：创建失败， STARTING：启动中， START_FAILED：启动失败， STOPPING：停止中 STOP_FAILED：停止失败 DELETING：删除中 DELETE_FIALED：删除失败 HIBERNATED：休眠
	State *string `json:"state,omitempty"`

	// 部署的节点数，最小为1，集群部署也为1
	NodeNum *int32 `json:"node_num,omitempty"`

	// 每个节点的部署结果
	Result         *[]NodeResult `json:"result,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o CreateDeploymentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDeploymentResponse struct{}"
	}

	return strings.Join([]string{"CreateDeploymentResponse", string(data)}, " ")
}
