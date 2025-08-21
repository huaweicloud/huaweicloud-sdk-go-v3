package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssetInfo 资产信息
type AssetInfo struct {

	// 资产类型；服务器设备/网络设备/配件/耗材/其他
	Type string `json:"type"`

	// 数量
	Amount int32 `json:"amount"`

	// 备注
	Remarks *string `json:"remarks,omitempty"`
}

func (o AssetInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssetInfo struct{}"
	}

	return strings.Join([]string{"AssetInfo", string(data)}, " ")
}
