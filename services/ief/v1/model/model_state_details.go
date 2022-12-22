package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IEC/IES节点状态详情
type StateDetails struct {

	// IEC/IES节点注册状态
	RegisterStat *string `json:"register_stat,omitempty"`

	// IEC/IES节点状态
	PurchaseStat *string `json:"purchase_stat,omitempty"`

	// IEC/IES节点错误信息
	PurchaseError *string `json:"purchase_error,omitempty"`
}

func (o StateDetails) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StateDetails struct{}"
	}

	return strings.Join([]string{"StateDetails", string(data)}, " ")
}
