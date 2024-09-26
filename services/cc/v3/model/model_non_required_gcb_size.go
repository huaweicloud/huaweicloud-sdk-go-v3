package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NonRequiredGcbSize 全域互联带宽实例中的带宽值大小。
type NonRequiredGcbSize struct {

	// 功能说明：全域互联带宽实例中的带宽值大小，单位Mbit/s。 取值范围：2-300Mbit/s
	Size *int32 `json:"size,omitempty"`
}

func (o NonRequiredGcbSize) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NonRequiredGcbSize struct{}"
	}

	return strings.Join([]string{"NonRequiredGcbSize", string(data)}, " ")
}
