package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPermissionsResponse Response Object
type ListPermissionsResponse struct {

	// 权限的详细信息。
	Permissions *[]Permission `json:"permissions,omitempty"`

	PageInfo *PageInfo `json:"page_info,omitempty"`

	// 请求ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListPermissionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPermissionsResponse struct{}"
	}

	return strings.Join([]string{"ListPermissionsResponse", string(data)}, " ")
}
