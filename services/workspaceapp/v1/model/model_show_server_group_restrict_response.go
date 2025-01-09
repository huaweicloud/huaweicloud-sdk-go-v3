package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowServerGroupRestrictResponse Response Object
type ShowServerGroupRestrictResponse struct {

	// 单台服务器最大的链接会话数。
	MaxSession *int32 `json:"max_session,omitempty"`

	// 该租户可创建的最多服务器组数量。
	MaxGroupCount  *int32 `json:"max_group_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowServerGroupRestrictResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowServerGroupRestrictResponse struct{}"
	}

	return strings.Join([]string{"ShowServerGroupRestrictResponse", string(data)}, " ")
}
