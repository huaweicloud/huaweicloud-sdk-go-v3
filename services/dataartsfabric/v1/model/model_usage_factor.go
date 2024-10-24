package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UsageFactor 使用量因子
type UsageFactor struct {
}

func (o UsageFactor) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UsageFactor struct{}"
	}

	return strings.Join([]string{"UsageFactor", string(data)}, " ")
}
