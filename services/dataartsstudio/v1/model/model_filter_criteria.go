package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FilterCriteria struct {

	// 当前可选值：database
	Name string `json:"name"`

	// database的名称
	Value string `json:"value"`

	// 操作者
	Operator string `json:"operator"`
}

func (o FilterCriteria) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FilterCriteria struct{}"
	}

	return strings.Join([]string{"FilterCriteria", string(data)}, " ")
}
