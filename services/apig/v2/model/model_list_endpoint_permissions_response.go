package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEndpointPermissionsResponse Response Object
type ListEndpointPermissionsResponse struct {

	// 本次返回的列表长度
	Size int32 `json:"size"`

	// 满足条件的记录数
	Total int64 `json:"total"`

	// 白名单记录列表
	Permissions *[]EndpointPermission `json:"permissions,omitempty"`

	XRequestId     *string `json:"x-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListEndpointPermissionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEndpointPermissionsResponse struct{}"
	}

	return strings.Join([]string{"ListEndpointPermissionsResponse", string(data)}, " ")
}
