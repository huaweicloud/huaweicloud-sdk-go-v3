package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEnterpriseMultiAccountResponse Response Object
type ListEnterpriseMultiAccountResponse struct {

	// 记录条数。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 可回收余额信息，如果是余额账户，只会有一条记录。 具体请参见表2。
	AmountInfos    *[]RetrieveAmountInfoV2 `json:"amount_infos,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ListEnterpriseMultiAccountResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEnterpriseMultiAccountResponse struct{}"
	}

	return strings.Join([]string{"ListEnterpriseMultiAccountResponse", string(data)}, " ")
}
