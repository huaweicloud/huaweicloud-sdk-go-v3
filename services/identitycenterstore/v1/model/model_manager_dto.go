package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ManagerDto struct {

	// 经理
	Value *string `json:"value,omitempty"`
}

func (o ManagerDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ManagerDto struct{}"
	}

	return strings.Join([]string{"ManagerDto", string(data)}, " ")
}
