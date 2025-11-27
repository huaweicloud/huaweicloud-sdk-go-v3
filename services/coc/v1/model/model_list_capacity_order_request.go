package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCapacityOrderRequest Request Object
type ListCapacityOrderRequest struct {

	// **参数解释：** 应用id。 **约束限制：** 不涉及。 **取值范围：** 需要查询容量数据排名的应用id。 **默认取值：** 不涉及
	ApplicationId *string `json:"application_id,omitempty"`

	// **参数解释：** 组件id。 **约束限制：** 不涉及。 **取值范围：** 需要查询容量数据排名的组件id。 **默认取值：** 不涉及
	ComponentId *string `json:"component_id,omitempty"`

	// **参数解释：** 分组id。 **约束限制：** 不涉及。 **取值范围：** 需要查询容量数据排名的分组id **默认取值：** 不涉及
	GroupId *string `json:"group_id,omitempty"`

	// **参数解释：** 云服务名称。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Provider string `json:"provider"`

	// **参数解释：** 资源类型名称。 **约束限制：** 不涉及。 **取值范围：** 资源类型较多，根据实际业务选择资源类型、常用资源类型如下： - cloudservers：弹性云服务器。 - servers：裸金属服务器。 - clusters：云容器引擎。 - instances：云数据库。 **默认取值：** 不涉及。
	Type string `json:"type"`
}

func (o ListCapacityOrderRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCapacityOrderRequest struct{}"
	}

	return strings.Join([]string{"ListCapacityOrderRequest", string(data)}, " ")
}
