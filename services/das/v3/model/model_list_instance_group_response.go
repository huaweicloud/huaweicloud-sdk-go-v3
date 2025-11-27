package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceGroupResponse Response Object
type ListInstanceGroupResponse struct {

	// 总数
	Total *int32 `json:"total,omitempty"`

	// 实例组列表
	InstanceGroupList *[]InstanceGroup `json:"instance_group_list,omitempty"`
	HttpStatusCode    int              `json:"-"`
}

func (o ListInstanceGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceGroupResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceGroupResponse", string(data)}, " ")
}
