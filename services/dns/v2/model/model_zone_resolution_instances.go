package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ZoneResolutionInstances struct {

	// **参数解释：** 域名ID。 **取值范围：** 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释：** 域名。 **取值范围：** 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 域名状态。 **取值范围：** - ACTIVE：正常 - PENDING_CREATE：创建中 - PENDING_UPDATE：更新中 - PENDING_DELETE：删除中 - PENDING_FREEZE：冻结中 - FREEZE：冻结 - ILLEGAL：违规冻结 - POLICE：公安冻结 - PENDING_DISABLE：暂停中 - DISABLE：暂停 - ERROR：失败
	Status *string `json:"status,omitempty"`

	// **参数解释：** 域名所属的区域。 **取值范围：** 不涉及。
	Region *string `json:"region,omitempty"`

	// **参数解释：** 企业项目ID。 **取值范围：** 不涉及。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ZoneResolutionInstances) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ZoneResolutionInstances struct{}"
	}

	return strings.Join([]string{"ZoneResolutionInstances", string(data)}, " ")
}
