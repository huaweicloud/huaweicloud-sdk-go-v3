package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCollectorChannelGroupResponse Response Object
type CreateCollectorChannelGroupResponse struct {

	// 创建成功
	Success        *bool `json:"success,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o CreateCollectorChannelGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCollectorChannelGroupResponse struct{}"
	}

	return strings.Join([]string{"CreateCollectorChannelGroupResponse", string(data)}, " ")
}
