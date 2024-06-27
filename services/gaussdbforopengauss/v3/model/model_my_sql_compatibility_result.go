package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MySqlCompatibilityResult struct {

	// MySQL兼容端口
	Port *string `json:"port,omitempty"`
}

func (o MySqlCompatibilityResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MySqlCompatibilityResult struct{}"
	}

	return strings.Join([]string{"MySqlCompatibilityResult", string(data)}, " ")
}
