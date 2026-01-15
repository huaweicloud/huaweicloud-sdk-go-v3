package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateLineGroupsResponse Response Object
type UpdateLineGroupsResponse struct {

	// 线路分组名称。
	Name *string `json:"name,omitempty"`

	// 线路分组包含的线路列表。 解析线路ID。
	Lines *[]string `json:"lines,omitempty"`

	// **参数解释：** 资源状态。 **取值范围：** - ACTIVE：正常 - PENDING_CREATE：创建中 - PENDING_UPDATE：更新中 - PENDING_DELETE：删除中 - PENDING_FREEZE：冻结中 - FREEZE：冻结 - ERROR：失败
	Status *string `json:"status,omitempty"`

	// 线路分组的描述信息
	Description *string `json:"description,omitempty"`

	// 线路分组的ID。
	LineId *string `json:"line_id,omitempty"`

	// 创建时间。 格式：yyyy-MM-dd'T'HH:mm:ss.SSS。
	CreatedAt *string `json:"created_at,omitempty"`

	// 更新时间。 格式：yyyy-MM-dd'T'HH:mm:ss.SSS。
	UpdatedAt      *string `json:"updated_at,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateLineGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLineGroupsResponse struct{}"
	}

	return strings.Join([]string{"UpdateLineGroupsResponse", string(data)}, " ")
}
