package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TmsMatch struct {

	// 键。第一期限定为resource_name,后续扩展。
	Key string `json:"key"`

	// 值。
	Value string `json:"value"`
}

func (o TmsMatch) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TmsMatch struct{}"
	}

	return strings.Join([]string{"TmsMatch", string(data)}, " ")
}
