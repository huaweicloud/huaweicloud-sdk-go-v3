package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteCollectorChannelGroupResponse Response Object
type DeleteCollectorChannelGroupResponse struct {

	// 创建成功
	Success        *bool `json:"success,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o DeleteCollectorChannelGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCollectorChannelGroupResponse struct{}"
	}

	return strings.Join([]string{"DeleteCollectorChannelGroupResponse", string(data)}, " ")
}
