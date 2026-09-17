package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPushChannelsRequest Request Object
type ListPushChannelsRequest struct {

	// 查询的起始位置，取值范围为非负整数，默认为0
	Offset *int32 `json:"offset,omitempty"`

	// 每页记录数，取值范围为非负整数，默认值为10
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListPushChannelsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPushChannelsRequest struct{}"
	}

	return strings.Join([]string{"ListPushChannelsRequest", string(data)}, " ")
}
