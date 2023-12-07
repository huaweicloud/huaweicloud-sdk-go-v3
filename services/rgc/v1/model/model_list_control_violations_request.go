package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListControlViolationsRequest Request Object
type ListControlViolationsRequest struct {

	// 如果正在使用临时安全凭据，则此header是必需的，该值是临时安全凭据的安全令牌（会话令牌）。
	XSecurityToken *string `json:"X-Security-Token,omitempty"`

	// 账户Id。
	AccountId *string `json:"account_id,omitempty"`

	// OU ID。
	OrganizationUnitId *string `json:"organization_unit_id,omitempty"`
}

func (o ListControlViolationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListControlViolationsRequest struct{}"
	}

	return strings.Join([]string{"ListControlViolationsRequest", string(data)}, " ")
}
