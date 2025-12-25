package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Invalidate 失效，包含如下:   - true ：是   - false ：否
type Invalidate struct {
}

func (o Invalidate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Invalidate struct{}"
	}

	return strings.Join([]string{"Invalidate", string(data)}, " ")
}
