package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Clusters struct {
	CustomerConfig *CustomerConfig `json:"customerConfig,omitempty" xml:"customerConfig"`

	Datastore *Datastore `json:"datastore,omitempty" xml:"datastore"`

	// 集群的节点信息，请参见instances参数说明
	Instances *[]ClusterDetailInstance `json:"instances,omitempty" xml:"instances"`

	// az名称
	AzName *string `json:"azName,omitempty" xml:"azName"`

	// 数据库用户
	Dbuser *string `json:"dbuser,omitempty" xml:"dbuser"`

	// 规格名称
	FlavorName *string `json:"flavorName,omitempty" xml:"flavorName"`

	// 事件数
	RecentEvent *int32 `json:"recentEvent,omitempty" xml:"recentEvent"`

	// 自动关机
	IsAutoOff *bool `json:"isAutoOff,omitempty" xml:"isAutoOff"`

	// 选择是否启用定时开关机功能。定时开关机功能和自动关机功能不可同时开启
	IsScheduleBootOff *bool `json:"isScheduleBootOff,omitempty" xml:"isScheduleBootOff"`

	// 集群模式：sharding(分片集群)
	ClusterMode *string `json:"clusterMode,omitempty" xml:"clusterMode"`

	// 命名空间
	Namespace *string `json:"namespace,omitempty" xml:"namespace"`

	Task *ClusterTask `json:"task,omitempty" xml:"task"`

	// 集群绑定的EIP
	PublicEndpoint *string `json:"publicEndpoint,omitempty" xml:"publicEndpoint"`

	ActionProgress *ActionProgress `json:"actionProgress,omitempty" xml:"actionProgress"`

	// 集群创建时间，格式为ISO8601：YYYY-MM-DDThh:mm:ssZ
	Created string `json:"created" xml:"created"`

	// 开始时间
	BakExpectedStartTime *string `json:"bakExpectedStartTime,omitempty" xml:"bakExpectedStartTime"`

	// 保留时间
	BakKeepDay *int32 `json:"bakKeepDay,omitempty" xml:"bakKeepDay"`

	// 集群名称
	Name string `json:"name" xml:"name"`

	// 集群状态描述：Normal（正常）
	StatusDetail *string `json:"statusDetail,omitempty" xml:"statusDetail"`

	// 集群ID
	Id string `json:"id" xml:"id"`

	// 集群是否冻结：0：否 1：是
	IsFrozen string `json:"isFrozen" xml:"isFrozen"`

	// 集群更新时间，格式为ISO8601：YYYY-MM-DDThh:mm:ssZ
	Updated string `json:"updated" xml:"updated"`

	// 集群状态： - 100：创建中 - 200：正常 - 300：失败 - 303：创建失败 - 500：重启中 - 800：冻结 - 900：已关机 - 910：正在关机 - 920：正在开机
	Status string `json:"status" xml:"status"`

	FailedReasons *FailedReasons `json:"failedReasons,omitempty" xml:"failedReasons"`
}

func (o Clusters) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Clusters struct{}"
	}

	return strings.Join([]string{"Clusters", string(data)}, " ")
}
