package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReclaimToPartnerAccountResponse Response Object
type ReclaimToPartnerAccountResponse struct {

	// 事务流水ID，只有成功响应才会返回。
	TransId        *string `json:"trans_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ReclaimToPartnerAccountResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReclaimToPartnerAccountResponse struct{}"
	}

	return strings.Join([]string{"ReclaimToPartnerAccountResponse", string(data)}, " ")
}
