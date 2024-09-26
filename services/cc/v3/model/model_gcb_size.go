package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GcbSize 全域互联带宽实例中的带宽值大小。
type GcbSize struct {

	// 功能说明：全域互联带宽实例中的带宽值大小，单位Mbit/s。 取值范围：2-300Mbit/s
	Size int32 `json:"size"`
}

func (o GcbSize) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GcbSize struct{}"
	}

	return strings.Join([]string{"GcbSize", string(data)}, " ")
}
