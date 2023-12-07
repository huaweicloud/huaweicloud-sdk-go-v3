package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowControlRequest Request Object
type ShowControlRequest struct {

	// 控制策略ID。
	ControlId string `json:"control_id"`

	// 如果正在使用临时安全凭据，则此header是必需的，该值是临时安全凭据的安全令牌（会话令牌）。
	XSecurityToken *string `json:"X-Security-Token,omitempty"`
}

func (o ShowControlRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowControlRequest struct{}"
	}

	return strings.Join([]string{"ShowControlRequest", string(data)}, " ")
}
