package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResetActiveCodeResponse Response Object
type ResetActiveCodeResponse struct {

	// 激活码ID。
	ActiveCodeId *string `json:"active_code_id,omitempty"`

	// 激活码。
	ActiveCode *string `json:"active_code,omitempty"`

	// 应用ID。
	RobotId *string `json:"robot_id,omitempty"`

	// 智能交互对话ID。
	RoomId *string `json:"room_id,omitempty"`

	// 有效天数（0表示长期有效）。
	ValidPeriod *int32 `json:"valid_period,omitempty"`

	// 过期时间，格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"。
	ExpireTime *string `json:"expire_time,omitempty"`

	// 创建时间，格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"。
	CreateTime *string `json:"create_time,omitempty"`

	// 更新时间，格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"。
	UpdateTime *string `json:"update_time,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ResetActiveCodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetActiveCodeResponse struct{}"
	}

	return strings.Join([]string{"ResetActiveCodeResponse", string(data)}, " ")
}
