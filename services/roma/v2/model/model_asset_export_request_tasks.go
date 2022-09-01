package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AssetExportRequestTasks struct {

	// 任务ID
	Id *string `json:"id,omitempty" xml:"id"`
}

func (o AssetExportRequestTasks) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssetExportRequestTasks struct{}"
	}

	return strings.Join([]string{"AssetExportRequestTasks", string(data)}, " ")
}
