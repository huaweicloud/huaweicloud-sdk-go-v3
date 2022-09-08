package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListProtectionGroupsResponse struct {

	// 保护组的信息列表。
	ServerGroups *[]ShowProtectionGroupParams `json:"server_groups,omitempty"`

	// 此参数为满足过滤条件的列表中包含的保护组个数。
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListProtectionGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProtectionGroupsResponse struct{}"
	}

	return strings.Join([]string{"ListProtectionGroupsResponse", string(data)}, " ")
}
