package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AssetOperateRequest struct {
	// 应用列表

	Apps []string `json:"apps"`
	// 任务列表

	Tasks []AssetOperateRequestTasks `json:"tasks"`
}

func (o AssetOperateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssetOperateRequest struct{}"
	}

	return strings.Join([]string{"AssetOperateRequest", string(data)}, " ")
}
