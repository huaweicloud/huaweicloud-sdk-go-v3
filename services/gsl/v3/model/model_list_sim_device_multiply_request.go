package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSimDeviceMultiplyRequest Request Object
type ListSimDeviceMultiplyRequest struct {

	// cid
	Cid *string `json:"cid,omitempty"`

	// 在线运营商标识
	OnlineCarrier *int32 `json:"online_carrier,omitempty"`

	// sim卡id
	SimCardId *int64 `json:"sim_card_id,omitempty"`

	// 批次号
	OrderId *int64 `json:"order_id,omitempty"`

	// 三网卡版本信息
	Version *int32 `json:"version,omitempty"`

	// 分页查询时每页显示的记录数，默认值为10，取值范围为10-500的整数
	Limit *int64 `json:"limit,omitempty"`

	// 分页查询时的页码数，默认值为1，取值范围为1-1000000的整数
	Offset *int64 `json:"offset,omitempty"`
}

func (o ListSimDeviceMultiplyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSimDeviceMultiplyRequest struct{}"
	}

	return strings.Join([]string{"ListSimDeviceMultiplyRequest", string(data)}, " ")
}
