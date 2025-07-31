package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportSize 导出数据条数
type ExportSize struct {
}

func (o ExportSize) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportSize struct{}"
	}

	return strings.Join([]string{"ExportSize", string(data)}, " ")
}
