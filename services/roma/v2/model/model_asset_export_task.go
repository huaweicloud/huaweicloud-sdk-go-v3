package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 任务ID
type AssetExportTask struct {
}

func (o AssetExportTask) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssetExportTask struct{}"
	}

	return strings.Join([]string{"AssetExportTask", string(data)}, " ")
}
