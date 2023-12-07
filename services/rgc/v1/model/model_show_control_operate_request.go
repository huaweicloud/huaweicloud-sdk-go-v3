package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowControlOperateRequest Request Object
type ShowControlOperateRequest struct {

	// 如果正在使用临时安全凭据，则此header是必需的，该值是临时安全凭据的安全令牌（会话令牌）。
	XSecurityToken *string `json:"X-Security-Token,omitempty"`

	// 操作控制策略的请求ID。
	ControlOperateRequestId string `json:"control_operate_request_id"`
}

func (o ShowControlOperateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowControlOperateRequest struct{}"
	}

	return strings.Join([]string{"ShowControlOperateRequest", string(data)}, " ")
}
