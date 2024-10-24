package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Stride 步长
type Stride struct {
}

func (o Stride) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Stride struct{}"
	}

	return strings.Join([]string{"Stride", string(data)}, " ")
}
