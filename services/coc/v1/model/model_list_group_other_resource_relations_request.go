package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGroupOtherResourceRelationsRequest Request Object
type ListGroupOtherResourceRelationsRequest struct {

	// **参数解释：** 分组关联的应用id。 **约束限制：** 应用id和分组id，组件id不能共存,且必填其中一个。 **取值范围：** 分组关联的应用id。 **默认取值：** 不涉及。
	ApplicationId *string `json:"application_id,omitempty"`

	// **参数解释：** 组件id。 **约束限制：** 应用id和分组id，组件id不能共存,且必填其中一个。 **取值范围：** 分组关联的组件id。 **默认取值：** 不涉及。
	ComponentId *string `json:"component_id,omitempty"`

	// **参数解释：** 分组id。 **约束限制：** 应用id和分组id，组件id不能共存,且必填其中一个。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	GroupId *string `json:"group_id,omitempty"`

	// **参数解释：** 厂商信息。 **约束限制：** 不涉及。 **取值范围：** - RMS：华为云厂商。 - ALI：阿里云厂商。 - OTHER：其他厂商。 **默认取值：** 不涉及。
	Vendor string `json:"vendor"`

	// **参数解释：** 资源类型。 **约束限制：** 不涉及。 **取值范围：** 自定义，云资源类型。 **默认取值：** 不涉及。
	Type string `json:"type"`

	// **参数解释：** 用于分页查询，查询数量。 **约束限制：** 不涉及。 **取值范围：** 自定义，在1-500范围。 **默认取值：** 不涉及。
	Limit int32 `json:"limit"`

	// **参数解释：** 分页查询偏移量，表示从此偏移量开始查询。 **约束限制：** 不涉及。 **取值范围：** 0-2147483647。 **默认取值：** 0。
	Offset *string `json:"offset,omitempty"`

	// **参数解释：** 资源id列表。 **约束限制：** 不涉及。 **取值范围：** 资源id列表，最大值100。 **默认取值：** 不涉及。
	ResourceIdList *[]string `json:"resource_id_list,omitempty"`

	// **参数解释：** 云资源名称。 **约束限制：** 不涉及。 **取值范围：** 自定义，可参考：裸金属服务器BMS。 **默认取值：** 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 区域id列表。 **约束限制：** 不涉及。 **取值范围：** 区域id列表，最大值100。 **默认取值：** 不涉及。
	RegionId *string `json:"region_id,omitempty"`

	// **参数解释：** 可用区id。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	AzId *string `json:"az_id,omitempty"`

	// **参数解释：** ip类型。 **约束限制：** 不涉及。 **取值范围：** fixed：内网IP。 floating：弹性公网IP。 **默认取值：** 不涉及。
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

	// **参数解释：** 云服务器的计费类型。 **约束限制：** 不涉及。 **取值范围：** 计费模式： - 0：按需计费。 - 1：包年包月。 - 2：竞价计费。 **默认取值：** 不涉及。
	ChargingMode *string `json:"charging_mode,omitempty"`

	// **参数解释：** 规格名称。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	FlavorName *string `json:"flavor_name,omitempty"`

	// **参数解释：** ip列表。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	IpList *[]string `json:"ip_list,omitempty"`

	// **参数解释：** 是否已收藏。 **约束限制：** 不涉及。 **取值范围：** - true：查询已收藏分组管理的资源关联。 - false：查询未收藏分组管理的资源关联。 **默认取值：** 不涉及。
	IsCollected *bool `json:"is_collected,omitempty"`
}

func (o ListGroupOtherResourceRelationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGroupOtherResourceRelationsRequest struct{}"
	}

	return strings.Join([]string{"ListGroupOtherResourceRelationsRequest", string(data)}, " ")
}
