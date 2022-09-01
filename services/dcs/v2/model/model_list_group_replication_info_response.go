package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListGroupReplicationInfoResponse struct {

	// 分片列表
	GroupList *[]InstanceGroupListInfo `json:"group_list,omitempty" xml:"group_list"`

	// 实例分片总数。
	GroupCount     *int32 `json:"group_count,omitempty" xml:"group_count"`
	HttpStatusCode int    `json:"-"`
}

func (o ListGroupReplicationInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGroupReplicationInfoResponse struct{}"
	}

	return strings.Join([]string{"ListGroupReplicationInfoResponse", string(data)}, " ")
}
