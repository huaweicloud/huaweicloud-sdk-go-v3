package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPermissionsResponse Response Object
type ListPermissionsResponse struct {

	// 请求ID。
	RequestId string `json:"request_id"`

	PageInfo *PageInfo `json:"page_info,omitempty"`

	// 权限实例列表。
	Permissions    []Permission `json:"permissions"`
	HttpStatusCode int          `json:"-"`
}

func (o ListPermissionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPermissionsResponse struct{}"
	}

	return strings.Join([]string{"ListPermissionsResponse", string(data)}, " ")
}
