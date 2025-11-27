package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceGroupRequest Request Object
type ListInstanceGroupRequest struct {

	// 数据库类型。支持MySQL、TaurusDB、GaussDB、MariaDB。
	DatastoreType string `json:"datastore_type"`

	// 实例组名称
	GroupName *string `json:"group_name,omitempty"`

	// 偏移量
	Offset *int32 `json:"offset,omitempty"`

	// 每页记录数
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListInstanceGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceGroupRequest struct{}"
	}

	return strings.Join([]string{"ListInstanceGroupRequest", string(data)}, " ")
}
