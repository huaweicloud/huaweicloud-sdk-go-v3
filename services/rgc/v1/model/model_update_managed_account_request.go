package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateManagedAccountRequest Request Object
type UpdateManagedAccountRequest struct {

	// 纳管账号ID。
	ManagedAccountId string `json:"managed_account_id"`

	Body *UpdateManagedAccountRequestBody `json:"body,omitempty"`
}

func (o UpdateManagedAccountRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateManagedAccountRequest struct{}"
	}

	return strings.Join([]string{"UpdateManagedAccountRequest", string(data)}, " ")
}
