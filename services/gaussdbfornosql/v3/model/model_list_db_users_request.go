package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDbUsersRequest Request Object
type ListDbUsersRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 数据库账号名。若传此参数，则查询指定账号的信息，否则返回所有数据库账号信息。
	Name *string `json:"name,omitempty"`

	// 索引位置，偏移量。    - 从第一条数据偏移offset条数据后开始查询，默认为0。   - 取值必须为数字，且不能为负数。
	Offset *int32 `json:"offset,omitempty"`

	// 查询个数上限值。  - 取值范围：1~100。 - 不传该参数时，默认查询前100条信息。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListDbUsersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDbUsersRequest struct{}"
	}

	return strings.Join([]string{"ListDbUsersRequest", string(data)}, " ")
}
