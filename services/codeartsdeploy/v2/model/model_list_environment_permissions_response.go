package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEnvironmentPermissionsResponse Response Object
type ListEnvironmentPermissionsResponse struct {

	// 权限数据，list类型数据
	Body           *[]DevUcEnvironmentPermission `json:"body,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o ListEnvironmentPermissionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEnvironmentPermissionsResponse struct{}"
	}

	return strings.Join([]string{"ListEnvironmentPermissionsResponse", string(data)}, " ")
}
