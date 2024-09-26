package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Uuid32Identifier struct {

	// 实例ID。
	Id string `json:"id"`
}

func (o Uuid32Identifier) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Uuid32Identifier struct{}"
	}

	return strings.Join([]string{"Uuid32Identifier", string(data)}, " ")
}
