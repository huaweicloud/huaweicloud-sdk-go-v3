package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Limit 单页最大结果数。
type Limit struct {
}

func (o Limit) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Limit struct{}"
	}

	return strings.Join([]string{"Limit", string(data)}, " ")
}
