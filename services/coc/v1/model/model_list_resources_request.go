package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResourcesRequest Request Object
type ListResourcesRequest struct {

	// **参数解释：** 云服务名称。 **约束限制：** - 创建的云资源数量（count字段对应的值）大于1时，可以使用“自动排序”和“正则排序”设置有序的云资源名称。 - 创建的云资源数量（count字段对应的值）大于1时，为区分不同云资源，创建过程中系统会自动在名称后加“-0000”的类似标记。若用户在名称后已指定“-0000”的类似标记，系统将从该标记后继续顺序递增编号。故此时名称的长度为[1-59]个字符。 **取值范围：** - 只能由中文字符、英文字母、数字及“_”、“-”、“.”组成，且长度为[1-128]个英文字符或[1-64]个中文字符。 **默认取值：** 不涉及。
	Provider string `json:"provider"`

	// **参数解释：** 资源类型。 **约束限制：** 不涉及。 **取值范围：** 资源类型较多，根据实际业务选择资源类型、常用资源类型如下： - cloudservers：弹性云服务器。 - servers：裸金属服务器。 - clusters：云容器引擎。 - instances：云数据库。 **默认取值：** 不涉及。
	Type string `json:"type"`

	// **参数解释：** 分页查询每页显示的条目数量。 **约束限制：** 不涉及。 **取值范围：** 自定义，在1-500范围。 **默认取值：** 不涉及。
	Limit int32 `json:"limit"`

	// **参数解释：** 用于分页查询。 **约束限制：** 不涉及。 **取值范围：** 分页参数，通过上一个请求中返回的marker信息作为输入，获取当前页。 **默认取值：** 不涉及。
	Marker *string `json:"marker,omitempty"`

	// **参数解释：** 资源id列表。 **约束限制：** 不涉及。 **取值范围：** 资源id列表，最大值100。 **默认取值：** 不涉及。
	ResourceIdList *[]string `json:"resource_id_list,omitempty"`

	// **参数解释：** ip列表。 **约束限制：** 不涉及。 **取值范围：** 列表，最大值100。 **默认取值：** 不涉及。
	IpList *[]string `json:"ip_list,omitempty"`

	// **参数解释：** 云资源名称。 **约束限制：** 不涉及。 **取值范围：** 字符串，可参考：裸金属服务器BMS。 **默认取值：** 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 区域id。 **约束限制：** 不涉及。 **取值范围：** 关联的区域region的id。 **默认取值：** 不涉及。
	RegionId *string `json:"region_id,omitempty"`

	// **参数解释：** 可用区id。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	AzId *string `json:"az_id,omitempty"`

	// **参数解释：** ip类型。 **约束限制：** 不涉及。 **取值范围：** - fixed：内网IP。 - floating：弹性公网IP。 **默认取值：** 不涉及。
	IpType *string `json:"ip_type,omitempty"`

	// **参数解释：** 云资源IP。 **约束限制：** 不涉及。 **取值范围：** IPv4地址过滤结果，匹配规则为模糊匹配。 **默认取值：** 不涉及。
	Ip *string `json:"ip,omitempty"`

	// **参数解释：** 云资源状态。 **约束限制：** 不涉及。 **取值范围：** 请选择[[弹性云服务器 ECS](https://support.huaweicloud.com/api-ecs/ecs_08_0002.html)](tag:hws)[[弹性云服务器 ECS](https://support.huaweicloud.com/api-ecs/ecs_08_0002.html)](tag:hws_hk)中存在的云服务器状态。 **默认取值：** 不涉及。
	Status *string `json:"status,omitempty"`

	// **参数解释：** agent状态。 **约束限制：** 不涉及。 **取值范围：** - ONLINE：运行中。 - OFFLINE：异常。 - INSTALLING：安装中。 - FAILED：安装失败。 - UNINSTALLED：已卸载。 - null：未安装。 **默认取值：** 不涉及。
	AgentState *string `json:"agent_state,omitempty"`

	// **参数解释：** 镜像名称，模糊匹配。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	ImageName *string `json:"image_name,omitempty"`

	// **参数解释：** 系统类型。 **约束限制：** 不涉及。 **取值范围：** - windows：windows操作系统类型。 - linux：linux操作系统类型。 **默认取值：** 不涉及。
	OsType *string `json:"os_type,omitempty"`

	// **参数解释：** 云资源的标签。 **约束限制：** 标签的格式为“key.value”。其中，key的长度不超过36个字符，value的长度不超过43个字符。 **取值范围：** 标签命名时，需满足如下要求：标签的key值只能包含大写字母（A~Z）、小写字母（a~z）、数字（0-9）、下划线（）、中划线（-）以及中文字符。 标签的value值只能包含大写字母（A~Z）、小写字母（a~z）、数字（0-9）、下划线（）、中划线（-）、小数点（.）以及中文字符。 **默认取值：** 不涉及。
	Tag *string `json:"tag,omitempty"`

	// **参数解释：** 云资源的标签key。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	TagKey *string `json:"tag_key,omitempty"`

	// **参数解释：** 云资源下的分组ID。 **约束限制：** 传分组id，就查询分组下的资源数量。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	GroupId *string `json:"group_id,omitempty"`

	// **参数解释：** 云资源下组件ID。 **约束限制：** 传组件id，就查询组件下的资源数量。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	ComponentId *string `json:"component_id,omitempty"`

	// **参数解释：** 云资源下应用ID。 **约束限制：** 传应用id，就查询应用下的资源数量。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	ApplicationId *string `json:"application_id,omitempty"`

	// **参数解释：** cce集群ID。 **约束限制：** 不涉及。 **取值范围：** 资源属于的cce的ID。 **默认取值：** 不涉及。
	CceClusterId *string `json:"cce_cluster_id,omitempty"`

	// **参数解释：** 待创建云资源所属虚拟私有云（简称VPC），需要指定已创建VPC的ID，UUID格式。。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	VpcId *string `json:"vpc_id,omitempty"`

	// **参数解释：** 企业项目id。 **约束限制：** 不涉及。 **取值范围：** 请选择[[企业管理](https://support.huaweicloud.com/usermanual-em/em_eps_qs_0400.html)](tag:hws)[[企业管理](https://support.huaweicloud.com/intl/zh-cn/usermanual-em/em_eps_qs_0400.html)](tag:hws_hk)中存在的项目ID。 **默认取值：** 不涉及。
	EpId *string `json:"ep_id,omitempty"`

	// **参数解释：** 是否已托管。 **约束限制：** 不涉及。 **取值范围：** - true：已经托管。 - false：未托管。 **默认取值：** 不涉及。
	IsDelegated *bool `json:"is_delegated,omitempty"`

	// **参数解释：** 是否已收藏。 **约束限制：** 不涉及。 **取值范围：** - true：已收藏的企业项目。 - false：未收藏的企业项目。 **默认取值：** 不涉及。
	IsCollected *bool `json:"is_collected,omitempty"`

	// **参数解释：** 云资源规格名称。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	FlavorName *string `json:"flavor_name,omitempty"`

	// **参数解释：** 云服务器的计费类型。 **约束限制：** 不涉及。 **取值范围：** 计费模式： - 0：按需计费。 - 1：包年包月。 - 2：竞价计费。 **默认取值：** 不涉及。
	ChargingMode *string `json:"charging_mode,omitempty"`

	// **参数解释：** 分页查询偏移量，表示从此偏移量开始查询。 **约束限制：** 不涉及。 **取值范围：** 0-2147483647。 **默认取值：** 0。
	Offset *string `json:"offset,omitempty"`

	// **参数解释：** 企业项目id。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释：** 根据排序字段对资源进行排序显示。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	OrderField *string `json:"order_field,omitempty"`

	// **参数解释：** 排序。 **约束限制：** 不涉及。 **取值范围：** - DESC:倒序。 - ASC:正序。 **默认取值：** 不涉及。
	Direction *string `json:"direction,omitempty"`

	// **参数解释：** 显示关联应用。 **约束限制：** 不涉及。 **取值范围：** - true：显示关联应用信息。 - false：不显示关联应用信息。 **默认取值：** 不涉及。
	ShowAssociatedGroups *string `json:"show_associated_groups,omitempty"`

	// **参数解释：** 用户定义资源是否可运维实例。 **约束限制：** 不涉及。 **取值范围：** - ENABLE：可运维实例。 - DISABLE：不可运维实例operable字段不存在。 **默认取值：** 不涉及。
	Operable *string `json:"operable,omitempty"`

	// **参数解释：** 创建时间中的开始日期，参考ISO8601标准格式。 **约束限制：** 开始日期和结束日期，至少有一个日期存在。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	CreateSince *string `json:"create_since,omitempty"`

	// **参数解释：** 创建时间中的结束日期，参考ISO8601标准格式。 **约束限制：** 开始日期和结束日期，至少有一个日期存在。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	CreateUntil *string `json:"create_until,omitempty"`
}

func (o ListResourcesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourcesRequest struct{}"
	}

	return strings.Join([]string{"ListResourcesRequest", string(data)}, " ")
}
