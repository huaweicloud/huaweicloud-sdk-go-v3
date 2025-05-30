package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateLineGroupResponse Response Object
type CreateLineGroupResponse struct {

	// 线路分组名称。
	Name *string `json:"name,omitempty"`

	// 线路分组包含的线路列表。 解析线路ID。
	Lines *[]string `json:"lines,omitempty"`

	// 资源状态。 取值范围：PENDING_CREATE，ACTIVE，PENDING_DELETE，PENDING_UPDATE，ERROR，FREEZE，DISABLE。
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

func (o CreateLineGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLineGroupResponse struct{}"
	}

	return strings.Join([]string{"CreateLineGroupResponse", string(data)}, " ")
}
