package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExtensionStatistics struct {

	// 下载量
	Install *int32 `json:"install,omitempty"`

	// 评星
	Stars *float64 `json:"stars,omitempty"`
}

func (o ExtensionStatistics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExtensionStatistics struct{}"
	}

	return strings.Join([]string{"ExtensionStatistics", string(data)}, " ")
}
