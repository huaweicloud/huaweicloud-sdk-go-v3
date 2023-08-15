package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListServicePermissionsDetailsResponse Response Object
type ListServicePermissionsDetailsResponse struct {

	// permission列表。
	Permissions *[]PermissionObject `json:"permissions,omitempty"`

	// 满足查询条件的终端节点服务的白名单总条数，不受分页（即limit、offset参数）影响。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListServicePermissionsDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServicePermissionsDetailsResponse struct{}"
	}

	return strings.Join([]string{"ListServicePermissionsDetailsResponse", string(data)}, " ")
}
