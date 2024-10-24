package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SystemTagParam 系统标签查询参数。
type SystemTagParam struct {

	// 键，固定值。
	Key string `json:"key"`

	// 系统标签的值列表。
	Values []string `json:"values"`
}

func (o SystemTagParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SystemTagParam struct{}"
	}

	return strings.Join([]string{"SystemTagParam", string(data)}, " ")
}
