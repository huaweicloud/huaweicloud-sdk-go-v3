package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AppAsset struct {
	// 应用列表

	Apps *[]string `json:"apps,omitempty"`
	// 任务列表

	Tasks *[]AppAssetTasks `json:"tasks,omitempty"`
}

func (o AppAsset) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppAsset struct{}"
	}

	return strings.Join([]string{"AppAsset", string(data)}, " ")
}
