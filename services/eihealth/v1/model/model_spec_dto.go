package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SpecDto struct {

	// 规格编号
	Code string `json:"code"`

	// 规格名称
	Name string `json:"name"`
}

func (o SpecDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SpecDto struct{}"
	}

	return strings.Join([]string{"SpecDto", string(data)}, " ")
}
