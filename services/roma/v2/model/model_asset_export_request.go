package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AssetExportRequest struct {

	// 应用列表
	Apps []AssetExportRequestApps `json:"apps" xml:"apps"`

	// 任务列表
	Tasks *[]AssetExportRequestTasks `json:"tasks,omitempty" xml:"tasks"`
}

func (o AssetExportRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssetExportRequest struct{}"
	}

	return strings.Join([]string{"AssetExportRequest", string(data)}, " ")
}
