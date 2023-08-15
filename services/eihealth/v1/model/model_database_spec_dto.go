package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DatabaseSpecDto struct {

	// 规格编号
	Code string `json:"code"`

	// 规格名称
	Name string `json:"name"`
}

func (o DatabaseSpecDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DatabaseSpecDto struct{}"
	}

	return strings.Join([]string{"DatabaseSpecDto", string(data)}, " ")
}
