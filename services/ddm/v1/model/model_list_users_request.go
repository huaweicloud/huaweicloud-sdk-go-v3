package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListUsersRequest struct {

	// DDM实例ID。
	InstanceId string `json:"instance_id" xml:"instance_id"`

	// 分页参数：起始值 [大于等于0] 。默认值是0。
	Offset *int32 `json:"offset,omitempty" xml:"offset"`

	// 分页参数：每页多少条 [大于0且小于等于128]。默认值是128。
	Limit *int32 `json:"limit,omitempty" xml:"limit"`
}

func (o ListUsersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUsersRequest struct{}"
	}

	return strings.Join([]string{"ListUsersRequest", string(data)}, " ")
}
