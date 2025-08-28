package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExportRisksRequestBody struct {

	// **参数解释**: 导出集群环境安全数据的表头信息列表 **取值范围**: 最小值1，最大值10000
	ExportHeaders [][]string `json:"export_headers"`

	IacItems *IacRequestInfo `json:"iac_items,omitempty"`
}

func (o ExportRisksRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportRisksRequestBody struct{}"
	}

	return strings.Join([]string{"ExportRisksRequestBody", string(data)}, " ")
}
