package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBaseModelRequest Request Object
type ListBaseModelRequest struct {

	// 限制量，单次查询总量，必须由数字组成，默认为100，取值范围[1,1000]
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量，查询起始偏移，必须由数字组成，默认为0，取值范围[0,100000000]
	Offset *int32 `json:"offset,omitempty"`

	// 排序规则 目前默认时间升序
	SortDir *string `json:"sort_dir,omitempty"`

	// 排序规则 目前默认时间升序，支持根据create_time排序
	SortKey *string `json:"sort_key,omitempty"`
}

func (o ListBaseModelRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBaseModelRequest struct{}"
	}

	return strings.Join([]string{"ListBaseModelRequest", string(data)}, " ")
}
