package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPartnerAccountChangeRecordsResponse Response Object
type ListPartnerAccountChangeRecordsResponse struct {

	// 返回总条数。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 币种。 CNY：人民币。
	Currency *string `json:"currency,omitempty"`

	// 调账记录列表。 具体请参见表2。
	Records        *[]AccountChangeRecord `json:"records,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListPartnerAccountChangeRecordsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPartnerAccountChangeRecordsResponse struct{}"
	}

	return strings.Join([]string{"ListPartnerAccountChangeRecordsResponse", string(data)}, " ")
}
