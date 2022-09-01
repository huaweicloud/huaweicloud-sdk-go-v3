package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListProjectCofigsV2Request struct {

	// 实例ID
	InstanceId string `json:"instance_id" xml:"instance_id"`

	// 偏移量，表示从此偏移量开始查询，偏移量小于0时，自动转换为0
	Offset *int64 `json:"offset,omitempty" xml:"offset"`

	// 每页显示的条目数量
	Limit *int32 `json:"limit,omitempty" xml:"limit"`
}

func (o ListProjectCofigsV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectCofigsV2Request struct{}"
	}

	return strings.Join([]string{"ListProjectCofigsV2Request", string(data)}, " ")
}
