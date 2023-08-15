package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Deletable 是否允许删除该策略组
type Deletable struct {
}

func (o Deletable) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Deletable struct{}"
	}

	return strings.Join([]string{"Deletable", string(data)}, " ")
}
