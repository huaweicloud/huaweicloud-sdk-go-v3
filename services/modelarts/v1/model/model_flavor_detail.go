package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FlavorDetail 训练作业、算法的规格信息（该字段只有公共资源池存在）。
type FlavorDetail struct {

	// 资源规格的类型。可选值如下： - CPU - GPU - [Ascend](tag:hc,hk,fcs_super)
	FlavorType *string `json:"flavor_type,omitempty"`

	Billing *BillingInfo `json:"billing,omitempty"`

	FlavorInfo *FlavorInfo `json:"flavor_info,omitempty"`
}

func (o FlavorDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlavorDetail struct{}"
	}

	return strings.Join([]string{"FlavorDetail", string(data)}, " ")
}
