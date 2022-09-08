package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListInvoicesResponse struct {

	// 记录数，只有成功的时候才返回这个字段
	Count *int32 `json:"count,omitempty"`

	// 发票信息列表，参见表2。
	Invoices       *[]InvoiceRequestInfoIntl `json:"invoices,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ListInvoicesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInvoicesResponse struct{}"
	}

	return strings.Join([]string{"ListInvoicesResponse", string(data)}, " ")
}
