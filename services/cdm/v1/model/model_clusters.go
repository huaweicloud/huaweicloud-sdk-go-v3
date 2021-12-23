package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Clusters struct {
	CustomerConfig *CustomerConfig `json:"customerConfig,omitempty"`

	Datastore *Datastore `json:"datastore,omitempty"`
	// 集群的节点信息，请参见instances参数说明

	Instances *[]ClusterDetailInstance `json:"instances,omitempty"`
	// az名称

	AzName *string `json:"azName,omitempty"`
	// 数据库用户

	Dbuser *string `json:"dbuser,omitempty"`
	// 规格名称

	FlavorName *string `json:"flavorName,omitempty"`
	// 事件数

	RecentEvent *int32 `json:"recentEvent,omitempty"`
	// 自动关机

	IsAutoOff *bool `json:"isAutoOff,omitempty"`
	// 选择是否启用定时开关机功能。定时开关机功能和自动关机功能不可同时开启

	IsScheduleBootOff *bool `json:"isScheduleBootOff,omitempty"`
	// 集群模式

	ClusterMode *string `json:"clusterMode,omitempty"`
	// 命名空间

	Namespace *string `json:"namespace,omitempty"`

	Task *ClusterTask `json:"task,omitempty"`
	// 集群绑定的EIP

	PublicEndpoint *string `json:"publicEndpoint,omitempty"`

	ActionProgress *ActionProgress `json:"actionProgress,omitempty"`
	// 集群创建时间，格式为ISO8601：YYYY-MM-DDThh:mm:ssZ

	Created string `json:"created"`
	// 集群名称

	Name string `json:"name"`
	// 集群状态描述

	StatusDetail *string `json:"statusDetail,omitempty"`
	// 集群ID

	Id string `json:"id"`
	// 集群是否冻结：0：否 1：是

	IsFrozen string `json:"isFrozen"`
	// 集群配置状态：In-Sync：配置已同步。Applying：配置中。Sync-Failure：配置失败

	ConfigStatus *string `json:"config_status,omitempty"`
	// 集群更新时间，格式为ISO8601：YYYY-MM-DDThh:mm:ssZ

	Updated string `json:"updated"`
	// 集群版本

	Version string `json:"version"`
	// 集群状态： - 100：创建中 - 200：正常 - 300：失败 - 303：创建失败 - 800：冻结 - 900：已关机 - 910：正在关机 - 920：正在开机

	Status string `json:"status"`

	FailedReasons *FailedReasons `json:"failedReasons,omitempty"`
}

func (o Clusters) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Clusters struct{}"
	}

	return strings.Join([]string{"Clusters", string(data)}, " ")
}
