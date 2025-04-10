package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportAntiVirusResultRequestBody 导出数据的表头字段列表
type ExportAntiVirusResultRequestBody struct {

	// 导出表头集合
	ExportHeaders *[][]string `json:"export_headers,omitempty"`
}

func (o ExportAntiVirusResultRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportAntiVirusResultRequestBody struct{}"
	}

	return strings.Join([]string{"ExportAntiVirusResultRequestBody", string(data)}, " ")
}
