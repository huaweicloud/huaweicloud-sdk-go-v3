package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCustomerAccountChangeRecordsResponse Response Object
type ListCustomerAccountChangeRecordsResponse struct {

	// |参数名称：总条数| |参数的约束及描述：总条数|
	TotalCount *int32 `json:"total_count,omitempty"`

	// |参数名称：币种| |参数约束及描述：币种|
	Currency *string `json:"currency,omitempty"`

	// |参数名称：收支明细记录列表| |参数约束以及描述：收支明细记录列表|
	Records        *[]CustomerAccountChangeRecord `json:"records,omitempty"`
	HttpStatusCode int                            `json:"-"`
}

func (o ListCustomerAccountChangeRecordsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCustomerAccountChangeRecordsResponse struct{}"
	}

	return strings.Join([]string{"ListCustomerAccountChangeRecordsResponse", string(data)}, " ")
}
