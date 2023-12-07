package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowOperationRequest Request Object
type ShowOperationRequest struct {

	// 操作ID。
	OperationId string `json:"operation_id"`

	// 如果正在使用临时安全凭据，则此header是必需的，该值是临时安全凭据的安全令牌（会话令牌）。
	XSecurityToken *string `json:"X-Security-Token,omitempty"`
}

func (o ShowOperationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOperationRequest struct{}"
	}

	return strings.Join([]string{"ShowOperationRequest", string(data)}, " ")
}
