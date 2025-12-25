package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportAopworkflowInfo 导出流程的请求对象
type ExportAopworkflowInfo struct {

	// **参数解释**: 导出流程的ID列表 **约束限制**: 不涉及
	Id []string `json:"id"`
}

func (o ExportAopworkflowInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportAopworkflowInfo struct{}"
	}

	return strings.Join([]string{"ExportAopworkflowInfo", string(data)}, " ")
}
