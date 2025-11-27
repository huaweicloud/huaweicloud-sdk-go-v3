package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountResourcesRequest Request Object
type CountResourcesRequest struct {

	// **参数解释：** 云服务名称。 **约束限制：** 不涉及。 **取值范围：** 可选ecs,cce，rds等服务资源。 **默认取值：** 不涉及。
	Provider string `json:"provider"`

	// **参数解释：** 资源类型名称。 **约束限制：** 不涉及。 **取值范围：** 资源类型较多，根据实际业务选择资源类型、常用资源类型如下： - cloudservers：弹性云服务器 。 - servers：裸金属服务器。 - clusters：云容器引擎。 - instances：云数据库。 **默认取值：** 不涉及。
	Type string `json:"type"`

	// **参数解释：** 资源id列表。 **约束限制：** 不涉及。 **取值范围：** 自定义,资源id组成的集合。 **默认取值：** 不涉及。
	ResourceIdList *[]string `json:"resource_id_list,omitempty"`

	// **参数解释：** 云资源名称。 **约束限制：** 不涉及。 **取值范围：** 自定义，可参考：裸金属服务器 BMS。 **默认取值：** 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 区域id。 **约束限制：** 不涉及。 **取值范围：** 关联的区域region的id。 **默认取值：** 不涉及。
	RegionId *string `json:"region_id,omitempty"`

	// **参数解释：** 可用区id。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	AzId *string `json:"az_id,omitempty"`

	// **参数解释：** ip类型。 **约束限制：** 不涉及。 **取值范围：** fixed：内网IP，floating：弹性公网IP。 **默认取值：** 不涉及。
	IpType *string `json:"ip_type,omitempty"`

	// **参数解释：** 云资源IP。 **约束限制：** 不涉及。 **取值范围：** IPv4地址过滤结果，匹配规则为模糊匹配。 **默认取值：** 不涉及。
	Ip *string `json:"ip,omitempty"`

	// **参数解释：** 云资源状态。 **约束限制：** 不涉及。 **取值范围：** 请选择[[弹性云服务器 ECS](https://support.huaweicloud.com/api-ecs/ecs_08_0002.html)](tag:hws)[[弹性云服务器 ECS](https://support.huaweicloud.com/api-ecs/ecs_08_0002.html)](tag:hws_hk)中存在的云服务器状态。 **默认取值：** 不涉及。
	Status *string `json:"status,omitempty"`

	// **参数解释：** agent状态。 **约束限制：** 不涉及。 **取值范围：** - 运行中。 - 异常。 - 安装中。 - 安装失败。 - 已卸载。 - 未安装。 **默认取值：** 不涉及。
	AgentState *string `json:"agent_state,omitempty"`

	// **参数解释：** 镜像名称，模糊匹配。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	ImageName *string `json:"image_name,omitempty"`

	// **参数解释：** 系统类型。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	OsType *string `json:"os_type,omitempty"`

	// **参数解释：** 云资源的标签。 **约束限制：** 标签的格式为“key.value”。其中，key的长度不超过36个字符，value的长度不超过43个字符。 **取值范围：** 标签命名时，需满足如下要求：标签的key值只能包含大写字母（A~Z）、小写字母（a~z）、数字（0-9）、下划线（）、中划线（-）以及中文字符。 标签的value值只能包含大写字母（A~Z）、小写字母（a~z）、数字（0-9）、下划线（）、中划线（-）、小数点（.）以及中文字符。 **默认取值：** 不涉及。
	Tag *string `json:"tag,omitempty"`

	// **参数解释：** 云资源的标签key。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	TagKey *string `json:"tag_key,omitempty"`

	// **参数解释：** 云资源下的分组ID。 **约束限制：** 不涉及。 **取值范围：** 设置null值时获取未绑定分组的资源。 **默认取值：** 不涉及。
	GroupId *string `json:"group_id,omitempty"`

	// **参数解释：** 云资源下组件ID。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	ComponentId *string `json:"component_id,omitempty"`

	// **参数解释：** 云资源下应用ID。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	ApplicationId *string `json:"application_id,omitempty"`

	// **参数解释：** cce集群ID。 **约束限制：** 不涉及。 **取值范围：** 资源属于的cce的ID。 **默认取值：** 不涉及。
	CceClusterId *string `json:"cce_cluster_id,omitempty"`

	// **参数解释：** 企业项目id。 **约束限制：** 不涉及。 **取值范围：** 请选择[[企业管理](https://support.huaweicloud.com/usermanual-em/em_eps_qs_0400.html)](tag:hws)[[企业管理](https://support.huaweicloud.com/intl/zh-cn/usermanual-em/em_eps_qs_0400.html)](tag:hws_hk)中存在的项目ID。 **默认取值：** 不涉及。
	EpId *string `json:"ep_id,omitempty"`

	// **参数解释：** 是否已托管。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	IsDelegated *string `json:"is_delegated,omitempty"`

	// **参数解释：** 是否已收藏。 **约束限制：** 不涉及。 **取值范围：** - true：已收藏的企业项目。 - false：未收藏的企业项目。 **默认取值：** 不涉及。
	IsCollected *bool `json:"is_collected,omitempty"`

	// **参数解释：** 云资源规格名称。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	FlavorName *string `json:"flavor_name,omitempty"`

	// **参数解释：** 云服务器的计费类型。 **约束限制：** 不涉及。 **取值范围：** 计费模式： - 0：按需计费。 - 1：包年包月。 - 2：竞价计费。 **默认取值：** 不涉及。
	ChargingMode *string `json:"charging_mode,omitempty"`
}

func (o CountResourcesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountResourcesRequest struct{}"
	}

	return strings.Join([]string{"CountResourcesRequest", string(data)}, " ")
}
