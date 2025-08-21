package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListOperationRequest Request Object
type ListOperationRequest struct {

	// 纳管账号ID。
	AccountId *string `json:"account_id,omitempty"`

	// 注册OU ID。
	OrganizationalUnitId *string `json:"organizational_unit_id,omitempty"`
}

func (o ListOperationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOperationRequest struct{}"
	}

	return strings.Join([]string{"ListOperationRequest", string(data)}, " ")
}
