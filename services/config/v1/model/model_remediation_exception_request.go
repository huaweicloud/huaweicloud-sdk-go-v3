package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RemediationExceptionRequest 合规规则修正例外。
type RemediationExceptionRequest struct {

	// 资源ID。
	ResourceId string `json:"resource_id"`

	// 加入合规规则修正例外的原因。
	Message *string `json:"message,omitempty"`
}

func (o RemediationExceptionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemediationExceptionRequest struct{}"
	}

	return strings.Join([]string{"RemediationExceptionRequest", string(data)}, " ")
}
