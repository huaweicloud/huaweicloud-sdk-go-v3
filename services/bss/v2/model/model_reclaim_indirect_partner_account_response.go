package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReclaimIndirectPartnerAccountResponse Response Object
type ReclaimIndirectPartnerAccountResponse struct {

	// 事务流水ID，只有成功响应才会返回。
	TransId        *string `json:"trans_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ReclaimIndirectPartnerAccountResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReclaimIndirectPartnerAccountResponse struct{}"
	}

	return strings.Join([]string{"ReclaimIndirectPartnerAccountResponse", string(data)}, " ")
}
