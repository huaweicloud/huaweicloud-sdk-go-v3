package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListCustomAuthorizersV2Request struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 偏移量，表示从此偏移量开始查询，偏移量小于0时，自动转换为0
	Offset *int64 `json:"offset,omitempty"`

	// 每页显示的条目数量
	Limit *int32 `json:"limit,omitempty"`

	// 编号
	Id *string `json:"id,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	// 类型
	Type *string `json:"type,omitempty"`
}

func (o ListCustomAuthorizersV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCustomAuthorizersV2Request struct{}"
	}

	return strings.Join([]string{"ListCustomAuthorizersV2Request", string(data)}, " ")
}
