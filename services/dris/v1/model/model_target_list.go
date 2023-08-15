package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TargetList **参数说明**：目标设备类型对应id的列表。
type TargetList struct {

	// **参数说明**：通过LTE-PC5传输通道（by_lte_pc5参数值为true）下发的rsu id列表。如果rsu id列表为空，则匹配事件范围内的在线rsu进行下发。
	TargetRsuIds *[]string `json:"target_rsu_ids,omitempty"`

	// **参数说明**：通过LTE-Uu的传输通道（by_lte_uu参数值为true）下发的obu id列表。注意obu id列表不能为空。
	TargetObuIds *[]string `json:"target_obu_ids,omitempty"`
}

func (o TargetList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TargetList struct{}"
	}

	return strings.Join([]string{"TargetList", string(data)}, " ")
}
