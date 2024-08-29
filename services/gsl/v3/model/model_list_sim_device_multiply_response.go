package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSimDeviceMultiplyResponse Response Object
type ListSimDeviceMultiplyResponse struct {

	// 每页记录数
	Limit *int64 `json:"limit,omitempty"`

	// 页码
	Offset *int64 `json:"offset,omitempty"`

	// 总数
	Count *int64 `json:"count,omitempty"`

	// 三网卡数据集合
	SimCards       *[]SimDeviceMultiplyVo `json:"sim_cards,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListSimDeviceMultiplyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSimDeviceMultiplyResponse struct{}"
	}

	return strings.Join([]string{"ListSimDeviceMultiplyResponse", string(data)}, " ")
}
