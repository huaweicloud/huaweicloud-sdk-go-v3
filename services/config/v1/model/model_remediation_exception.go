package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RemediationException 合规规则修正例外。
type RemediationException struct {

	// 资源ID。
	ResourceId *string `json:"resource_id,omitempty"`

	// 加入合规规则修正例外的原因。
	Message *string `json:"message,omitempty"`

	// 加入合规规则修正例外的时间。
	JoinedAt *string `json:"joined_at,omitempty"`

	// 合规规则修正例外的创建者。
	CreatedBy *string `json:"created_by,omitempty"`
}

func (o RemediationException) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemediationException struct{}"
	}

	return strings.Join([]string{"RemediationException", string(data)}, " ")
}
