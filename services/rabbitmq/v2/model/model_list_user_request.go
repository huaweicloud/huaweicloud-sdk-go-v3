package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUserRequest Request Object
type ListUserRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 分页查询偏移量，表示从此偏移量开始查询，offset大于等于0，默认从0开始查询。
	Offset *string `json:"offset,omitempty"`

	// 分页查询单页数量，取值范围0~50，默认查询10条。
	Limit *string `json:"limit,omitempty"`
}

func (o ListUserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUserRequest struct{}"
	}

	return strings.Join([]string{"ListUserRequest", string(data)}, " ")
}
