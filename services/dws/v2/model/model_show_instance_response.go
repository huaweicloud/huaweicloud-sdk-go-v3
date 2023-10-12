package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceResponse Response Object
type ShowInstanceResponse struct {

	// 配置状态
	ConfigurationStatus *string `json:"configuration_status,omitempty"`

	// 参数组ID
	ParamsGroupId *string `json:"params_group_id,omitempty"`

	// 类型
	Type *string `json:"type,omitempty"`

	// 子网ID
	SubnetId *string `json:"subnet_id,omitempty"`

	// 角色
	Role *string `json:"role,omitempty"`

	// 内部子网ID
	InternalSubnetId *string `json:"internal_subnet_id,omitempty"`

	// 组
	Group *string `json:"group,omitempty"`

	// 安全组
	SecureGroup *string `json:"secure_group,omitempty"`

	// VPC
	Vpc *string `json:"vpc,omitempty"`

	// 编码
	Azcode *string `json:"azcode,omitempty"`

	// 区域
	Region *string `json:"region,omitempty"`

	// 集群ID
	ClusterId *string `json:"cluster_id,omitempty"`

	// 被创建的
	Created *string `json:"created,omitempty"`

	// 被更新的
	Updated *string `json:"updated,omitempty"`

	// 状态
	Status *string `json:"status,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	// 连接
	Links *[]LinkResp `json:"links,omitempty"`

	// ID
	Id *string `json:"id,omitempty"`

	Flavor *ClusterFlavorResp `json:"flavor,omitempty"`

	Volume *CompatibleInstanceVolumeResp `json:"volume,omitempty"`

	Datastore *CompatibleDataStoreResp `json:"datastore,omitempty"`

	Fault *CompatibleFaultResp `json:"fault,omitempty"`

	Configuration *CompatibleConfigurationResp `json:"configuration,omitempty"`

	// 地点
	Locality *string `json:"locality,omitempty"`

	// 备份
	Replicas *[]CompatibleReplicasResp `json:"replicas,omitempty"`

	// 数据库用户
	DbUser *string `json:"db_user,omitempty"`

	// 存储引擎
	StorageEngine *string `json:"storage_engine,omitempty"`

	// 付款方式
	PayModel *int32 `json:"pay_model,omitempty"`

	// 公网IP
	PublicIp *string `json:"public_ip,omitempty"`

	// 流量IP
	TrafficIp      *string `json:"traffic_ip,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceResponse", string(data)}, " ")
}
