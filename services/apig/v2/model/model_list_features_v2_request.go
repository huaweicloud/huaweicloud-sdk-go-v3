package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFeaturesV2Request Request Object
type ListFeaturesV2Request struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
	InstanceId string `json:"instance_id"`

	// 偏移量，表示从此偏移量开始查询，偏移量小于0时，自动转换为0
	Offset *int64 `json:"offset,omitempty"`

	// 每页显示的条目数量，条目数量小于等于0时，自动转换为20，条目数量大于500时，自动转换为500
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListFeaturesV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFeaturesV2Request struct{}"
	}

	return strings.Join([]string{"ListFeaturesV2Request", string(data)}, " ")
}
