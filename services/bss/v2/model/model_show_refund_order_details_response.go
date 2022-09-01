package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowRefundOrderDetailsResponse struct {

	// 查询总数。
	TotalCount *int32 `json:"total_count,omitempty" xml:"total_count"`

	// 资源信息列表。 具体请参见表2。
	RefundInfos    *[]OrderRefundInfoV2 `json:"refund_infos,omitempty" xml:"refund_infos"`
	HttpStatusCode int                  `json:"-"`
}

func (o ShowRefundOrderDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRefundOrderDetailsResponse struct{}"
	}

	return strings.Join([]string{"ShowRefundOrderDetailsResponse", string(data)}, " ")
}
