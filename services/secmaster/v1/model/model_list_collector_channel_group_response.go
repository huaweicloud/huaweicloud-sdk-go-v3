package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCollectorChannelGroupResponse Response Object
type ListCollectorChannelGroupResponse struct {

	// 响应结果
	Body           *[]Group `json:"body,omitempty"`
	HttpStatusCode int      `json:"-"`
}

func (o ListCollectorChannelGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCollectorChannelGroupResponse struct{}"
	}

	return strings.Join([]string{"ListCollectorChannelGroupResponse", string(data)}, " ")
}
