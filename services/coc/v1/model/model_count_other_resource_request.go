package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountOtherResourceRequest Request Object
type CountOtherResourceRequest struct {

	// **参数解释：** 是否已托管。 **约束限制：** 不涉及。 **取值范围：** - true：已经托管。 - false：未托管。 **默认取值：** 不涉及。
	IsDelegated *bool `json:"is_delegated,omitempty"`

	// **参数解释：** 资源类型。 **约束限制：** 不涉及。 **取值范围：** - vm：虚拟机。 - PM：物理机。 - Middleware：中间件设备。 **默认取值：** 不涉及。
	Type *string `json:"type,omitempty"`

	// **参数解释：** 是否为sre。 **约束限制：** 不涉及。 **取值范围：** - true：是sre。 - false：非sre。 **默认取值：** 不涉及。
	IsDelegateDomain *bool `json:"is_delegate_domain,omitempty"`

	// **参数解释：** 资源名称，支持模糊查询。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	NameList *[]string `json:"name_list,omitempty"`

	// **参数解释：** 区域id。 **约束限制：** 不涉及。 **取值范围：** 关联的区域region的id。 **默认取值：** 不涉及。
	RegionId *string `json:"region_id,omitempty"`

	// **参数解释：** 创建时间中的开始日期，参考ISO8601标准格式。 **约束限制：** 开始日期和结束日期，至少有一个日期存在。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	CreateSince *string `json:"create_since,omitempty"`

	// **参数解释：** 创建时间中的结束日期，参考ISO8601标准格式。 **约束限制：** 开始日期和结束日期，至少有一个日期存在。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	CreateUntil *string `json:"create_until,omitempty"`

	// **参数解释：** 私有ip。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Ip *string `json:"ip,omitempty"`

	// **参数解释：** 操作系统。 **约束限制：** 不涉及。 **取值范围：** - windows：windows操作系统类型。 - linux：linux操作系统类型。 **默认取值：** 不涉及。
	OsType *string `json:"os_type,omitempty"`

	// **参数解释：** OS版本。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	OsVersionList *string `json:"os_version_list,omitempty"`

	// **参数解释：** agent状态。 **约束限制：** 不涉及。 **取值范围：** - ONLINE：运行中。 - OFFLINE：异常。 - INSTALLING：安装中。 - FAILED：安装失败。 - UNINSTALLED：已卸载。 - null：未安装。 **默认取值：** 不涉及。
	AgentState *string `json:"agent_state,omitempty"`
}

func (o CountOtherResourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountOtherResourceRequest struct{}"
	}

	return strings.Join([]string{"CountOtherResourceRequest", string(data)}, " ")
}
