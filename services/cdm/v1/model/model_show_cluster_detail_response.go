package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowClusterDetailResponse struct {
	// 集群绑定的EIP

	PublicEndpoint *string `json:"publicEndpoint,omitempty"`
	// 集群的节点信息，请参见instances参数说明

	Instances *[]ClusterDetailInstance `json:"instances,omitempty"`
	// 安全组id

	SecurityGroupId *string `json:"security_group_id,omitempty"`
	// 子网id

	SubnetId *string `json:"subnet_id,omitempty"`
	// 虚拟私有云ID

	VpcId *string `json:"vpc_id,omitempty"`

	CustomerConfig *CustomerConfig `json:"customerConfig,omitempty"`

	Datastore *Datastore `json:"datastore,omitempty"`
	// 自动关机

	IsAutoOff *bool `json:"isAutoOff,omitempty"`
	// 集群绑定的EIP域名

	PublicEndpointDomainName *string `json:"publicEndpointDomainName,omitempty"`
	// 开始时间

	BakExpectedStartTime *string `json:"bakExpectedStartTime,omitempty"`
	// 保留时间

	BakKeepDay *int32 `json:"bakKeepDay,omitempty"`

	MaintainWindow *CdmQueryClusterDetailsRepsonseMaintainWindow `json:"maintainWindow,omitempty"`
	// 事件数

	RecentEvent *int32 `json:"recentEvent,omitempty"`
	// 规格名称

	FlavorName *string `json:"flavorName,omitempty"`
	// az名称

	AzName *string `json:"azName,omitempty"`
	// 对端域名

	EndpointDomainName *string `json:"endpointDomainName,omitempty"`

	PublicEndpointStatus *CdmQueryClusterDetailsRepsonsePublicEndpointStatus `json:"publicEndpointStatus,omitempty"`
	// 选择是否启用定时开关机功能。定时开关机功能和自动关机功能不可同时开启

	IsScheduleBootOff *bool `json:"isScheduleBootOff,omitempty"`
	// 命名空间

	Namespace *string `json:"namespace,omitempty"`
	// 弹性ip id

	EipId *string `json:"eipId,omitempty"`

	FailedReasons *FailedReasons `json:"failedReasons,omitempty"`
	// 数据库用户

	Dbuser *string `json:"dbuser,omitempty"`

	Links *[]ClusterLinks `json:"links,omitempty"`
	// 集群模式

	ClusterMode *string `json:"clusterMode,omitempty"`

	Task *ClusterTask `json:"task,omitempty"`
	// 集群创建时间，格式为ISO8601：YYYY-MM-DDThh:mm:ssZ

	Created *string `json:"created,omitempty"`
	// 集群状态描述

	StatusDetail *string `json:"statusDetail,omitempty"`
	// 集群配置状态： - In-Sync：配置已同步。 - Applying：配置中。 - Sync-Failure：配置失败

	ConfigStatus *string `json:"config_status,omitempty"`

	ActionProgress *ActionProgress `json:"actionProgress,omitempty"`
	// 集群名称

	Name *string `json:"name,omitempty"`
	// 集群ID

	Id *string `json:"id,omitempty"`
	// 集群是否冻结：0：否1：是

	IsFrozen *string `json:"isFrozen,omitempty"`
	// 集群配置状态：In-Sync：配置已同步。Applying：配置中。Sync-Failure：配置失败

	Actions *[]string `json:"actions,omitempty"`
	// 集群更新时间，格式为 ISO8601：YYYY-MM-DDThh:mm:ssZ

	Updated *string `json:"updated,omitempty"`
	// 集群状态： - 100：创建中 - 200：正常 - 300：失败 - 303：创建失败 - 800：冻结 - 900：已关机 - 910：正在关机 - 920：正在开机

	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowClusterDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowClusterDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowClusterDetailResponse", string(data)}, " ")
}
