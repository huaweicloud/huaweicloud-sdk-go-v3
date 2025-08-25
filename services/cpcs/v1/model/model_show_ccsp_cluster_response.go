package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCcspClusterResponse Response Object
type ShowCcspClusterResponse struct {

	// 任务ID
	TaskId *string `json:"task_id,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 账号ID
	DomainId *string `json:"domain_id,omitempty"`

	// 任务详情
	TaskDetails *[]interface{} `json:"task_details,omitempty"`

	// ccsp集群id
	CcspId *string `json:"ccsp_id,omitempty"`

	// 分布类型
	DistributedType *string `json:"distributed_type,omitempty"`

	// 集群ID
	ClusterId *string `json:"cluster_id,omitempty"`

	// 集群名称
	ClusterName *string `json:"cluster_name,omitempty"`

	// 集群所属的服务
	ServiceType *string `json:"service_type,omitempty"`

	// 集群的类型： - **SHARED** : 应用共享； - **EXCLUSIVE** : 应用独享
	Type *string `json:"type,omitempty"`

	// 集群中包含的服务实例数量
	InstanceNum *int32 `json:"instance_num,omitempty"`

	// 集群状态
	Status *string `json:"status,omitempty"`

	// 进度信息
	ProgressInfo *string `json:"progress_info,omitempty"`

	// 集群使用VSM集群的vsm实例数量
	VsmNum *int32 `json:"vsm_num,omitempty"`

	// 集群的创建时间，UNIX时间戳，单位毫秒
	CreateTime *int64 `json:"create_time,omitempty"`

	// 是否共享ccsp
	SharedCcsp *bool `json:"shared_ccsp,omitempty"`

	// 企业项目ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// az
	Az             *string `json:"az,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowCcspClusterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCcspClusterResponse struct{}"
	}

	return strings.Join([]string{"ShowCcspClusterResponse", string(data)}, " ")
}
