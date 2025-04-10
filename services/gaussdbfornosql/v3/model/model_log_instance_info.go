package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LogInstanceInfo struct {

	// 实例ID。
	Id *string `json:"id,omitempty"`

	// 实例名称。
	Name *string `json:"name,omitempty"`

	// 实例状态。 取值： - normal，表示实例正常。 - abnormal，表示实例异常。 - creating，表示实例创建中。 - frozen，表示实例被冻结。 - data_disk_full，表示实例磁盘已满。 - createfail，表示实例创建失败。 - enlargefail，表示实例扩容节点失败。
	Status *string `json:"status,omitempty"`

	// 实例类型。   -  取值为“Cluster”，表示GeminiDB Redis经典部署模式集群实例类型。   -  取值为“CloudNativeCluster”，表示GeminiDB Redis云原生部署模式集群实例类型。
	Mode *string `json:"mode,omitempty"`

	Datastore *InstancesDatastoreResult `json:"datastore,omitempty"`

	// 实例正在执行的动作。
	Actions *[]string `json:"actions,omitempty"`

	// 企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 日志类型。slow_log表示慢日志，audit_log表示审计日志。
	SupportedLogTypes *[]string `json:"supported_log_types,omitempty"`
}

func (o LogInstanceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LogInstanceInfo struct{}"
	}

	return strings.Join([]string{"LogInstanceInfo", string(data)}, " ")
}
