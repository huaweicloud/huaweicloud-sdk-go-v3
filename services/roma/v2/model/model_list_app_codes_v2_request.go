package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListAppCodesV2Request struct {

	// 实例ID
	InstanceId string `json:"instance_id" xml:"instance_id"`

	// 应用编号
	AppId string `json:"app_id" xml:"app_id"`

	// 偏移量，表示从此偏移量开始查询，偏移量小于0时，自动转换为0
	Offset *int64 `json:"offset,omitempty" xml:"offset"`

	// 每页显示的条目数量
	Limit *int32 `json:"limit,omitempty" xml:"limit"`
}

func (o ListAppCodesV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAppCodesV2Request struct{}"
	}

	return strings.Join([]string{"ListAppCodesV2Request", string(data)}, " ")
}
