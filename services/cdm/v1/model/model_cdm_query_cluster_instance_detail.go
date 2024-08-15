package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CdmQueryClusterInstanceDetail struct {

	// 节点配置状态： - In-Sync：配置已同步。 - Applying：配置中。 - Sync-Failure：配置失败。
	ConfigurationStatus *string `json:"configurationStatus,omitempty"`

	// 配置ID
	ParamsGroupId *string `json:"paramsGroupId,omitempty"`

	// 配置服务类型，这里为cdm
	Type *string `json:"type,omitempty"`

	// 实例模式，这里为Standalone
	Role *string `json:"role,omitempty"`

	// 实例的子网ID
	Subnetid *string `json:"subnetid,omitempty"`

	// 安全组ID
	Securegroup *string `json:"securegroup,omitempty"`

	// 实例的VPC ID
	Vpc *string `json:"vpc,omitempty"`

	// 可用区名称
	Azcode *string `json:"azcode,omitempty"`

	// 局点名称
	Region *string `json:"region,omitempty"`

	// 实例创建时间，格式为ISO8601：YYYY-MM-DDThh:mm:ssZ
	Created *string `json:"created,omitempty"`

	// 实例更新时间，格式为ISO8601：YYYY-MM-DDThh:mm:ssZ
	Updated *string `json:"updated,omitempty"`

	// 实例名称
	Name *string `json:"name,omitempty"`

	// 实例ID
	Id *string `json:"id,omitempty"`

	Flavor *CdmQueryClusterInstanceDetailFlavor `json:"flavor,omitempty"`

	Datastore *Datastore `json:"datastore,omitempty"`

	// 数据库用户，这里为cdm。
	Dbuser *string `json:"dbuser,omitempty"`

	// 付费模式： - 0：按需 - 1：包周期
	PayModel *int32 `json:"payModel,omitempty"`

	// 集群绑定的公网地址
	PublicIp *string `json:"publicIp,omitempty"`

	// 集群的内网地址
	TrafficIp *string `json:"trafficIp,omitempty"`

	// 集群的内网IPv6地址
	TrafficIpv6 *string `json:"trafficIpv6,omitempty"`

	// 集群ID
	ClusterId *string `json:"cluster_id,omitempty"`
}

func (o CdmQueryClusterInstanceDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CdmQueryClusterInstanceDetail struct{}"
	}

	return strings.Join([]string{"CdmQueryClusterInstanceDetail", string(data)}, " ")
}
