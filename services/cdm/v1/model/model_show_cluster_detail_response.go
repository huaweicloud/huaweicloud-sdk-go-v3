package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowClusterDetailResponse struct {

	// 集群绑定的EIP
	PublicEndpoint *string `json:"publicEndpoint,omitempty" xml:"publicEndpoint"`

	// 集群的节点信息，请参见instances参数说明
	Instances *[]ClusterDetailInstance `json:"instances,omitempty" xml:"instances"`

	// 安全组id
	SecurityGroupId *string `json:"security_group_id,omitempty" xml:"security_group_id"`

	// 子网id
	SubnetId *string `json:"subnet_id,omitempty" xml:"subnet_id"`

	// 虚拟私有云ID
	VpcId *string `json:"vpc_id,omitempty" xml:"vpc_id"`

	CustomerConfig *CustomerConfig `json:"customerConfig,omitempty" xml:"customerConfig"`

	Datastore *Datastore `json:"datastore,omitempty" xml:"datastore"`

	// 自动关机
	IsAutoOff *bool `json:"isAutoOff,omitempty" xml:"isAutoOff"`

	// 集群绑定的EIP域名
	PublicEndpointDomainName *string `json:"publicEndpointDomainName,omitempty" xml:"publicEndpointDomainName"`

	// 开始时间
	BakExpectedStartTime *string `json:"bakExpectedStartTime,omitempty" xml:"bakExpectedStartTime"`

	// 保留时间
	BakKeepDay *int32 `json:"bakKeepDay,omitempty" xml:"bakKeepDay"`

	MaintainWindow *CdmQueryClusterDetailsRepsonseMaintainWindow `json:"maintainWindow,omitempty" xml:"maintainWindow"`

	// 事件数
	RecentEvent *int32 `json:"recentEvent,omitempty" xml:"recentEvent"`

	// 规格名称
	FlavorName *string `json:"flavorName,omitempty" xml:"flavorName"`

	// az名称
	AzName *string `json:"azName,omitempty" xml:"azName"`

	// 对端域名
	EndpointDomainName *string `json:"endpointDomainName,omitempty" xml:"endpointDomainName"`

	PublicEndpointStatus *CdmQueryClusterDetailsRepsonsePublicEndpointStatus `json:"publicEndpointStatus,omitempty" xml:"publicEndpointStatus"`

	// 选择是否启用定时开关机功能。定时开关机功能和自动关机功能不可同时开启
	IsScheduleBootOff *bool `json:"isScheduleBootOff,omitempty" xml:"isScheduleBootOff"`

	// 命名空间
	Namespace *string `json:"namespace,omitempty" xml:"namespace"`

	// 弹性ip id
	EipId *string `json:"eipId,omitempty" xml:"eipId"`

	FailedReasons *FailedReasons `json:"failedReasons,omitempty" xml:"failedReasons"`

	// 数据库用户
	Dbuser *string `json:"dbuser,omitempty" xml:"dbuser"`

	Links *[]ClusterLinks `json:"links,omitempty" xml:"links"`

	// 集群模式：sharding(分片集群)
	ClusterMode *string `json:"clusterMode,omitempty" xml:"clusterMode"`

	Task *ClusterTask `json:"task,omitempty" xml:"task"`

	// 集群创建时间，格式为ISO8601：YYYY-MM-DDThh:mm:ssZ
	Created *string `json:"created,omitempty" xml:"created"`

	// 集群状态描述：Normal（正常）
	StatusDetail *string `json:"statusDetail,omitempty" xml:"statusDetail"`

	// 集群配置状态： - In-Sync：配置已同步。 - Applying：配置中。 - Sync-Failure：配置失败
	ConfigStatus *string `json:"config_status,omitempty" xml:"config_status"`

	ActionProgress *ActionProgress `json:"actionProgress,omitempty" xml:"actionProgress"`

	// 集群名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 集群ID
	Id *string `json:"id,omitempty" xml:"id"`

	// 集群是否冻结：0：否1：是
	IsFrozen *string `json:"isFrozen,omitempty" xml:"isFrozen"`

	// 集群配置状态：In-Sync：配置已同步。Applying：配置中。Sync-Failure：配置失败
	Actions *[]string `json:"actions,omitempty" xml:"actions"`

	// 集群更新时间，格式为 ISO8601：YYYY-MM-DDThh:mm:ssZ
	Updated *string `json:"updated,omitempty" xml:"updated"`

	// 集群状态： - 100：创建中 - 200：正常 - 300：失败 - 303：创建失败 - 800：冻结 - 900：已关机 - 910：正在关机 - 920：正在开机
	Status         *string `json:"status,omitempty" xml:"status"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowClusterDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowClusterDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowClusterDetailResponse", string(data)}, " ")
}
