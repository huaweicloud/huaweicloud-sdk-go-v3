package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCloudServersRequest Request Object
type ListCloudServersRequest struct {

	// 云服务器ID，格式为UUID，匹配规则为精确匹配。
	Id *string `json:"id,omitempty"`

	// 云服务器名称，匹配规则为模糊匹配。
	Name *string `json:"name,omitempty"`

	// 云服务器状态。  取值范围：  ACTIVE， BUILD，ERROR，HARD_REBOOT，MIGRATING，REBOOT，RESIZE，REVERT_RESIZE，SHELVED，SHELVED_OFFLOADED，SHUTOFF，UNKNOWN，VERIFY_RESIZE  弹性云服务器状态说明请参考[云服务器状态](https://support.huaweicloud.com/api-ecs/ecs_08_0002.html)
	Status *string `json:"status,omitempty"`

	// 云服务器是否处于回收站中
	InRecycleBin *bool `json:"in_recycle_bin,omitempty"`

	// 共池裸机按整机柜发放的同一批次的批创id。
	SpodId *string `json:"spod_id,omitempty"`

	// 云服务器规格名称。
	FlavorName *string `json:"flavor_name,omitempty"`

	// 镜像ID。
	ImageId *string `json:"image_id,omitempty"`

	// 元数据过滤，支持key=value过滤。
	Metadata *string `json:"metadata,omitempty"`

	// 元数据key过滤。
	MetadataKey *string `json:"metadata-key,omitempty"`

	// 查询tag字段中包含该值的云服务器。
	Tags *string `json:"tags,omitempty"`

	//  查询tag字段中不包含该值的云服务器
	NotTags *string `json:"not-tags,omitempty"`

	// 云服务器所在的AZ，匹配规则为模糊匹配。
	AvailabilityZone *string `json:"availability_zone,omitempty"`

	// 云服务器所在的AZ，匹配规则为精确匹配。
	AvailabilityZoneEq *string `json:"availability_zone_eq,omitempty"`

	// 云服务器的计费类型。
	ChargingMode *string `json:"charging_mode,omitempty"`

	// 云服务器使用的密钥对名称。
	KeyName *string `json:"key_name,omitempty"`

	// 过滤在launched_since时间之后启动的云服务器。格式为ISO8601时间格式，例如：2013-06-09T06:42:18Z。
	LaunchedSince *string `json:"launched_since,omitempty"`

	// 过滤绑定某个企业项目的云服务器。 若需要查询当前用户所有企业项目绑定的云服务，请传参all_granted_eps。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 控制查询输出的字段。在默认字段的基础上选择是否查询。   launched_at：云服务器启动时间。   key_name：云服务器使用的密钥对名称。   locked：云服务器是否为锁定状态。   root_device_name：云服务器系统盘的设备名称。   tenancy：在专属主机或共享池中创建云服务器。   dedicated_host_id：专属主机ID。   enterprise_project_id：查询绑定某个企业项目的云服务器。   tags：云服务器的标签列表。   metadata：云服务器元数据。   addresses：云服务器对应的网络地址信息。   security_groups：云服务器的安全组信息。   volumes_attached：云服务器挂载磁盘信息。   image：云服务器镜像信息。   power_state：云服务器电源状态。   cpu_options：自定义CPU选项。   market_info：云服务器计费信息，包含计费类型、到期时间等字段。
	ExpectFields *[]string `json:"expect-fields,omitempty"`

	// 查询返回VM数量限制。 limit 默认为10，最大为100。
	Limit *int32 `json:"limit,omitempty"`

	// 以单页最后一条server的ID作为分页标记。
	Marker *string `json:"marker,omitempty"`
}

func (o ListCloudServersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCloudServersRequest struct{}"
	}

	return strings.Join([]string{"ListCloudServersRequest", string(data)}, " ")
}
