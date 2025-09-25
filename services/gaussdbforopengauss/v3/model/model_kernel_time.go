package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KernelTime 内核模块耗时
type KernelTime struct {

	// **参数解释**: 总耗时（单位：微秒）。 **取值范围**: 不涉及。
	AllTime int64 `json:"all_time"`

	KernelTimeDetails *KernelTimeDetails `json:"kernel_time_details"`
}

func (o KernelTime) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KernelTime struct{}"
	}

	return strings.Join([]string{"KernelTime", string(data)}, " ")
}
