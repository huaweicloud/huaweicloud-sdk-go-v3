package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoliciesPeripheralsParallelPortRedirection 并口重定向。
type PoliciesPeripheralsParallelPortRedirection struct {

	// 是否开启并口重定向。取值为： false：表示关闭。 true：表示开启。
	ParallelPortEnable *bool `json:"parallel_port_enable,omitempty"`
}

func (o PoliciesPeripheralsParallelPortRedirection) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoliciesPeripheralsParallelPortRedirection struct{}"
	}

	return strings.Join([]string{"PoliciesPeripheralsParallelPortRedirection", string(data)}, " ")
}
