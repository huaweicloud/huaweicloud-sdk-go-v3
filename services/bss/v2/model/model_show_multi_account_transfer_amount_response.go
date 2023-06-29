package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMultiAccountTransferAmountResponse Response Object
type ShowMultiAccountTransferAmountResponse struct {

	// 记录条数。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 可拨款余额信息，如果是余额账户，只会有一条记录。 具体请参见表2。
	AmountInfos    *[]TransferAmountInfoV2 `json:"amount_infos,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ShowMultiAccountTransferAmountResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMultiAccountTransferAmountResponse struct{}"
	}

	return strings.Join([]string{"ShowMultiAccountTransferAmountResponse", string(data)}, " ")
}
