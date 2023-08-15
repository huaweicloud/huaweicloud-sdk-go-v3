package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFlavorInfo 规格信息。
type ListFlavorInfo struct {

	// cpu核数。
	Vcpu int32 `json:"vcpu"`

	// 内存大小。
	Mem int32 `json:"mem"`
}

func (o ListFlavorInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFlavorInfo struct{}"
	}

	return strings.Join([]string{"ListFlavorInfo", string(data)}, " ")
}
