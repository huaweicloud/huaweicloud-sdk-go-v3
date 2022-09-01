package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AssetOperateRequest struct {

	// 应用列表
	Apps []string `json:"apps" xml:"apps"`

	// 任务列表
	Tasks []AssetOperateRequestTasks `json:"tasks" xml:"tasks"`
}

func (o AssetOperateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssetOperateRequest struct{}"
	}

	return strings.Join([]string{"AssetOperateRequest", string(data)}, " ")
}
