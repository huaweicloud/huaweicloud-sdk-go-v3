package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateCollectorChannelGroupResponse Response Object
type UpdateCollectorChannelGroupResponse struct {

	// 创建成功
	Success        *bool `json:"success,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o UpdateCollectorChannelGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCollectorChannelGroupResponse struct{}"
	}

	return strings.Join([]string{"UpdateCollectorChannelGroupResponse", string(data)}, " ")
}
