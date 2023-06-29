package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCustomerselfResourceRecordsResponse Response Object
type ListCustomerselfResourceRecordsResponse struct {

	// 资源费用记录数据。 具体请参见表3。
	FeeRecords *[]ResFeeRecordV2 `json:"fee_records,omitempty"`

	// 结果集数量，只有成功才返回这个参数。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 货币单位代码： CNY：人民币
	Currency       *string `json:"currency,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListCustomerselfResourceRecordsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCustomerselfResourceRecordsResponse struct{}"
	}

	return strings.Join([]string{"ListCustomerselfResourceRecordsResponse", string(data)}, " ")
}
