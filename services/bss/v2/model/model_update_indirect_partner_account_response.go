package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateIndirectPartnerAccountResponse Response Object
type UpdateIndirectPartnerAccountResponse struct {

	// 事务流水ID，只有成功响应才会返回。
	TransferId     *string `json:"transfer_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateIndirectPartnerAccountResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateIndirectPartnerAccountResponse struct{}"
	}

	return strings.Join([]string{"UpdateIndirectPartnerAccountResponse", string(data)}, " ")
}
