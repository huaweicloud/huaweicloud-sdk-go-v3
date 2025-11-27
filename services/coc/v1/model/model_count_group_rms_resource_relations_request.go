package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountGroupRmsResourceRelationsRequest Request Object
type CountGroupRmsResourceRelationsRequest struct {

	// **参数解释：** 应用id。 **约束限制：** 应用id，和分组id不能共存，且必填其中一个。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	ApplicationId *string `json:"application_id,omitempty"`

	// **参数解释：** 分组id。 **约束限制：** 分组id，和应用id不能共存，且必填其中一个。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	GroupId *string `json:"group_id,omitempty"`

	// **参数解释：** 云服务名称。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Provider string `json:"provider"`

	// **参数解释：** 云服务资源类型。 **约束限制：** 不涉及。 **取值范围：** 资源类型较多，根据实际业务选择资源类型、常用资源类型如下： - cloudservers：弹性云服务器。 - servers：裸金属服务器。 - clusters：云容器引擎。 - instances：云数据库。 **默认取值：** 不涉及。
	Type string `json:"type"`

	// **参数解释：** 厂商信息。 **约束限制：** 不涉及。 **取值范围：** - RMS：华为云厂商。 - ALI：阿里云厂商。 - OTHER：其他厂商。 **默认取值：** 不涉及
	Vendor string `json:"vendor"`

	// **参数解释：** 资源id列表。 **约束限制：** 不涉及。 **取值范围：** 用户选择的资源id组成的集合。 **默认取值：** 不涉及。
	ResourceIdList *[]string `json:"resource_id_list,omitempty"`

	// **参数解释：** 云资源名称。 **约束限制：** 不涉及。 **取值范围：** 自定义，可参考：裸金属服务器BMS。 **默认取值：** 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 区域id。 **约束限制：** 不涉及。 **取值范围：** 关联的区域region的id。 **默认取值：** 不涉及。
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

	// **参数解释：** 系统类型。 **约束限制：** 不涉及。 **取值范围：** - windows：windows操作系统类型。 - linux：linux操作系统类型。 **默认取值：** 不涉及。
	OsType *string `json:"os_type,omitempty"`

	// **参数解释：** 云资源的标签。 **约束限制：** 标签的格式为“key.value”。其中，key的长度不超过36个字符，value的长度不超过43个字符。 **取值范围：** 标签命名时，需满足如下要求：标签的key值只能包含大写字母（A~Z）、小写字母（a~z）、数字（0-9）、下划线（）、中划线（-）以及中文字符。 标签的value值只能包含大写字母（A~Z）、小写字母（a~z）、数字（0-9）、下划线（）、中划线（-）、小数点（.）以及中文字符。 **默认取值：** 不涉及。
	Tag *string `json:"tag,omitempty"`
}

func (o CountGroupRmsResourceRelationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountGroupRmsResourceRelationsRequest struct{}"
	}

	return strings.Join([]string{"CountGroupRmsResourceRelationsRequest", string(data)}, " ")
}
