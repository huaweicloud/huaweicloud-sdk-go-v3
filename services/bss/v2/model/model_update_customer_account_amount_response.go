package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateCustomerAccountAmountResponse Response Object
type UpdateCustomerAccountAmountResponse struct {

	// 事务流水ID，只有成功响应才会返回。
	TransferId     *string `json:"transfer_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateCustomerAccountAmountResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCustomerAccountAmountResponse struct{}"
	}

	return strings.Join([]string{"UpdateCustomerAccountAmountResponse", string(data)}, " ")
}
