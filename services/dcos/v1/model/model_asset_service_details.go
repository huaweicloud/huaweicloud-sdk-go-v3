package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssetServiceDetails 资产服务信息
type AssetServiceDetails struct {

	// 资产类型；服务器设备/网络设备/配件/耗材/其他
	Type *string `json:"type,omitempty"`

	// 数量
	Amount *int32 `json:"amount,omitempty"`

	// 备注
	Remarks *string `json:"remarks,omitempty"`

	// 成功或异常
	TaskStatus *string `json:"task_status,omitempty"`

	// 异常原因
	Description *string `json:"description,omitempty"`
}

func (o AssetServiceDetails) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssetServiceDetails struct{}"
	}

	return strings.Join([]string{"AssetServiceDetails", string(data)}, " ")
}
