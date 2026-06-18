package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateGroupResourcePermissionsRequest Request Object
type UpdateGroupResourcePermissionsRequest struct {

	// **参数解释：** 代码组id，代码组首页，Group ID后的数字Id **默认取值：** 不涉及。
	GroupId int32 `json:"group_id"`

	// **参数解释：** 资源Id，通过获取代码组权限资源点列表获取的数据中的Id **默认取值：** 不涉及。
	ResourceId int32 `json:"resource_id"`

	Body *UpdatePermissionBodyDto `json:"body,omitempty"`
}

func (o UpdateGroupResourcePermissionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGroupResourcePermissionsRequest struct{}"
	}

	return strings.Join([]string{"UpdateGroupResourcePermissionsRequest", string(data)}, " ")
}
