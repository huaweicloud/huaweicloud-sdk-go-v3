package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QueryTag struct {

	// 键。
	Key string `json:"key"`

	// 值列表。
	Values []string `json:"values"`
}

func (o QueryTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryTag struct{}"
	}

	return strings.Join([]string{"QueryTag", string(data)}, " ")
}
