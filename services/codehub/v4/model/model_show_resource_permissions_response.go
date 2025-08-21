package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowResourcePermissionsResponse Response Object
type ShowResourcePermissionsResponse struct {

	// **参数解释：** 资源权限列表。
	Body           *[]PermissionDto `json:"body,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ShowResourcePermissionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResourcePermissionsResponse struct{}"
	}

	return strings.Join([]string{"ShowResourcePermissionsResponse", string(data)}, " ")
}
