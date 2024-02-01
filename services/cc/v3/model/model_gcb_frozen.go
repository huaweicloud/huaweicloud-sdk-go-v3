package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GcbFrozen struct {

	// 功能说明: 全域互联带宽是否冻结。 取值范围：     true-冻结     false-非冻结
	Frozen *bool `json:"frozen,omitempty"`
}

func (o GcbFrozen) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GcbFrozen struct{}"
	}

	return strings.Join([]string{"GcbFrozen", string(data)}, " ")
}
