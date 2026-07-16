package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkflowPoolOrder pool order struct
type WorkflowPoolOrder struct {

	// 订阅ID。
	Id *string `json:"id,omitempty"`

	Sku *SkuInfo `json:"sku"`

	// 订阅计数。
	SkuCount string `json:"sku_count"`
}

func (o WorkflowPoolOrder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkflowPoolOrder struct{}"
	}

	return strings.Join([]string{"WorkflowPoolOrder", string(data)}, " ")
}
