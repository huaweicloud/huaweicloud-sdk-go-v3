package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// **参数说明**：分发通道，至少指定一个通道。
type Channel struct {

	// **参数说明**：LTE-PC5传输通道。若通过LTE-PC5传输通道下发事件，该字段为true。
	ByLtePc5 *bool `json:"by_lte_pc5,omitempty"`

	// **参数说明**：LTE-Uu的传输通道。若通过LTE-Uu的传输通道下发事件，该字段为true。
	ByLteUu *bool `json:"by_lte_uu,omitempty"`
}

func (o Channel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Channel struct{}"
	}

	return strings.Join([]string{"Channel", string(data)}, " ")
}
