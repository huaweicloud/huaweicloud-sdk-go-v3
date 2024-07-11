package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListApplicationPermissionsResponse Response Object
type ListApplicationPermissionsResponse struct {

	// 角色应用权限
	Result *[]ApplicationPermissionVo `json:"result,omitempty"`

	// 请求成功失败状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListApplicationPermissionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApplicationPermissionsResponse struct{}"
	}

	return strings.Join([]string{"ListApplicationPermissionsResponse", string(data)}, " ")
}
