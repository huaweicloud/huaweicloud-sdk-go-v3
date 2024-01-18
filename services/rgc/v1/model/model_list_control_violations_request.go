package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListControlViolationsRequest Request Object
type ListControlViolationsRequest struct {

	// 账户ID，用于过滤不合规资源。
	AccountId *string `json:"account_id,omitempty"`

	// 注册OU ID，用于过滤不合规资源。
	OrganizationUnitId *string `json:"organization_unit_id,omitempty"`
}

func (o ListControlViolationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListControlViolationsRequest struct{}"
	}

	return strings.Join([]string{"ListControlViolationsRequest", string(data)}, " ")
}
