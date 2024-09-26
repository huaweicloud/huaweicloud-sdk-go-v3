package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TmsTagValues struct {

	// 键。
	Key string `json:"key"`

	// 值列表。
	Values []string `json:"values"`
}

func (o TmsTagValues) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TmsTagValues struct{}"
	}

	return strings.Join([]string{"TmsTagValues", string(data)}, " ")
}
