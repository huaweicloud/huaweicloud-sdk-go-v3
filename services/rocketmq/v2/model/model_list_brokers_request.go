package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBrokersRequest Request Object
type ListBrokersRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 查询数量。
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量，表示从此偏移量开始查询，offset大于等于0。
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListBrokersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBrokersRequest struct{}"
	}

	return strings.Join([]string{"ListBrokersRequest", string(data)}, " ")
}
