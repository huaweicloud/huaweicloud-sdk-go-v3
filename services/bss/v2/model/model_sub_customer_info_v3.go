package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SubCustomerInfoV3 struct {

	// 客户账号ID。
	CustomerId *string `json:"customer_id,omitempty"`
}

func (o SubCustomerInfoV3) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubCustomerInfoV3 struct{}"
	}

	return strings.Join([]string{"SubCustomerInfoV3", string(data)}, " ")
}
