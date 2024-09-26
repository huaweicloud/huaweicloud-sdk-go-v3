package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Uuid64Identifier struct {

	// 实例ID。
	Id string `json:"id"`
}

func (o Uuid64Identifier) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Uuid64Identifier struct{}"
	}

	return strings.Join([]string{"Uuid64Identifier", string(data)}, " ")
}
