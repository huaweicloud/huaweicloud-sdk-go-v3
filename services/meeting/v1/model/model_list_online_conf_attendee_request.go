package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListOnlineConfAttendeeRequest Request Object
type ListOnlineConfAttendeeRequest struct {

	// 会议ID
	ConfId string `json:"conf_id"`

	// 记录数偏移
	Offset *int32 `json:"offset,omitempty"`

	// 返回的与会者记录数，最大500条
	Limit *int32 `json:"limit,omitempty"`

	// 查询条件,支持third-account查询返回
	SearchKey *string `json:"search_key,omitempty"`
}

func (o ListOnlineConfAttendeeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOnlineConfAttendeeRequest struct{}"
	}

	return strings.Join([]string{"ListOnlineConfAttendeeRequest", string(data)}, " ")
}
