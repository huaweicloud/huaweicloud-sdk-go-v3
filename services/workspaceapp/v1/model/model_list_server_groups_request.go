package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListServerGroupsRequest Request Object
type ListServerGroupsRequest struct {

	// 查询的偏移量
	Offset int32 `json:"offset"`

	// 查询的数量，值区间[1-100]
	Limit int32 `json:"limit"`

	// 服务器组名称
	ServerGroupName *string `json:"server_group_name,omitempty"`

	// 服务器组唯一标识
	ServerGroupId *string `json:"server_group_id,omitempty"`
}

func (o ListServerGroupsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServerGroupsRequest struct{}"
	}

	return strings.Join([]string{"ListServerGroupsRequest", string(data)}, " ")
}
