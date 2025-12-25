package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExportPlaybook struct {

	// 导出剧本ID列表
	Id *[]string `json:"id,omitempty"`
}

func (o ExportPlaybook) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportPlaybook struct{}"
	}

	return strings.Join([]string{"ExportPlaybook", string(data)}, " ")
}
