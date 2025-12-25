package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceResponse Response Object
type ShowInstanceResponse struct {

	// **参数解释**： 配置状态。 **取值范围**： - In-Sync：应用完成。 - Applying：应用中。 - Pending-Reboot：待重启生效。 - Sync-Failure：修改失败。
	ConfigurationStatus *string `json:"configuration_status,omitempty"`

	// **参数解释**： 参数组ID。 **取值范围**： 不涉及。
	ParamsGroupId *string `json:"params_group_id,omitempty"`

	// **参数解释**： 类型。 **取值范围**： - dws-cn：cn节点。 - dws：dn节点。 - cms：cms节点。 - gtm：gtm节点。 - vw：vw节点。
	Type *string `json:"type,omitempty"`

	// **参数解释**： 子网ID。 **取值范围**： 不涉及。
	SubnetId *string `json:"subnet_id,omitempty"`

	// **参数解释**： 角色。 **取值范围**： 不涉及。
	Role *string `json:"role,omitempty"`

	// **参数解释**： 内部子网ID。 **取值范围**： 不涉及。
	InternalSubnetId *string `json:"internal_subnet_id,omitempty"`

	// **参数解释**： 分组信息。 **取值范围**： 不涉及。
	Group *string `json:"group,omitempty"`

	// **参数解释**： 安全组。 **取值范围**： 不涉及。
	SecureGroup *string `json:"secure_group,omitempty"`

	// **参数解释**： VPC ID。 **取值范围**： 不涉及。
	Vpc *string `json:"vpc,omitempty"`

	// **参数解释**： 可用区编码。 **取值范围**： 不涉及。
	Azcode *string `json:"azcode,omitempty"`

	// **参数解释**： 局点编码。 **取值范围**： 不涉及。
	Region *string `json:"region,omitempty"`

	// **参数解释**： 集群ID。 **取值范围**： 不涉及。
	ClusterId *string `json:"cluster_id,omitempty"`

	// **参数解释**： 创建时间。 **取值范围**： 不涉及。
	Created *string `json:"created,omitempty"`

	// **参数解释**： 更新时间。 **取值范围**： 不涉及。
	Updated *string `json:"updated,omitempty"`

	// **参数解释**： 状态。 **取值范围**： - BUILD：创建中 - ACTIVE：可用 - FAILED：不可用 - DELETED：已删除 - REBOOTING：重启中 - ERROR：创建失败 - SHUTDOWN：关机
	Status *string `json:"status,omitempty"`

	// **参数解释**： 节点名称。 **取值范围**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 链接信息。 **取值范围**： 不涉及。
	Links *[]LinkResp `json:"links,omitempty"`

	// **参数解释**： 节点ID。 **取值范围**： 不涉及。
	Id *string `json:"id,omitempty"`

	Flavor *ClusterFlavorResp `json:"flavor,omitempty"`

	Volume *CompatibleInstanceVolumeResp `json:"volume,omitempty"`

	Datastore *CompatibleDataStoreResp `json:"datastore,omitempty"`

	Fault *CompatibleFaultResp `json:"fault,omitempty"`

	Configuration *CompatibleConfigurationResp `json:"configuration,omitempty"`

	// **参数解释**： 废弃字段，无实际含义。 **取值范围**： 不涉及。
	Locality *string `json:"locality,omitempty"`

	// **参数解释**： 废弃字段，无实际含义。 **取值范围**： 不涉及。
	Replicas *[]CompatibleReplicasResp `json:"replicas,omitempty"`

	// **参数解释**： 数据库用户。 **取值范围**： 不涉及。
	DbUser *string `json:"db_user,omitempty"`

	// **参数解释**： 存储引擎。 **取值范围**： 不涉及。
	StorageEngine *string `json:"storage_engine,omitempty"`

	// **参数解释**： 付款方式。 **取值范围**： 不涉及。
	PayModel *int32 `json:"pay_model,omitempty"`

	// **参数解释**： 公网IP。 **取值范围**： 不涉及。
	PublicIp *string `json:"public_ip,omitempty"`

	// **参数解释**： 流量IP。 **取值范围**： 不涉及。
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
