package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AssetOperateRequest struct {

	// 应用列表
	Apps []AssetExportRequestApps `json:"apps"`

	// 任务列表
	Tasks []AssetExportRequestTasks `json:"tasks"`
}

func (o AssetOperateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssetOperateRequest struct{}"
	}

	return strings.Join([]string{"AssetOperateRequest", string(data)}, " ")
}
