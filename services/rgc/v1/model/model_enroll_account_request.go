package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnrollAccountRequest Request Object
type EnrollAccountRequest struct {

	// 纳管账号ID。
	ManagedAccountId string `json:"managed_account_id"`

	Body *EnrollAccountRequestBody `json:"body,omitempty"`
}

func (o EnrollAccountRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnrollAccountRequest struct{}"
	}

	return strings.Join([]string{"EnrollAccountRequest", string(data)}, " ")
}
