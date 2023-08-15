package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUserRequest Request Object
type ListUserRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 偏移量，表示从此偏移量开始查询， offset大于等于0。
	Offset *string `json:"offset,omitempty"`

	// 查询数量。
	Limit *string `json:"limit,omitempty"`
}

func (o ListUserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUserRequest struct{}"
	}

	return strings.Join([]string{"ListUserRequest", string(data)}, " ")
}
