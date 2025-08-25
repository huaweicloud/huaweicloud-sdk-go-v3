package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CcspClusterInfo struct {

	// 任务ID
	TaskId string `json:"task_id"`

	// 项目ID
	ProjectId string `json:"project_id"`

	// 账号ID
	DomainId string `json:"domain_id"`

	// 任务详情
	TaskDetails []interface{} `json:"task_details"`

	// ccsp集群id
	CcspId string `json:"ccsp_id"`

	// 分布类型
	DistributedType string `json:"distributed_type"`

	// 集群ID
	ClusterId string `json:"cluster_id"`

	// 集群名称
	ClusterName string `json:"cluster_name"`

	// 集群所属的服务
	ServiceType string `json:"service_type"`

	// 集群的类型： - **SHARED** : 应用共享； - **EXCLUSIVE** : 应用独享
	Type string `json:"type"`

	// 集群中包含的服务实例数量
	InstanceNum int32 `json:"instance_num"`

	// 集群状态
	Status string `json:"status"`

	// 进度信息
	ProgressInfo *string `json:"progress_info,omitempty"`

	// 集群使用VSM集群的vsm实例数量
	VsmNum *int32 `json:"vsm_num,omitempty"`

	// 集群的创建时间，UNIX时间戳，单位毫秒
	CreateTime int64 `json:"create_time"`

	// 是否共享ccsp
	SharedCcsp bool `json:"shared_ccsp"`

	// 企业项目ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// az
	Az *string `json:"az,omitempty"`
}

func (o CcspClusterInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CcspClusterInfo struct{}"
	}

	return strings.Join([]string{"CcspClusterInfo", string(data)}, " ")
}
