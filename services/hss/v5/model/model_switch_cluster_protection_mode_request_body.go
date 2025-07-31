package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SwitchClusterProtectionModeRequestBody struct {

	// 集群ID列表
	ClusterIds []string `json:"cluster_ids"`

	// - 1：开 - 0：关 - 2：关闭并卸载插件
	Opr int32 `json:"opr"`
}

func (o SwitchClusterProtectionModeRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchClusterProtectionModeRequestBody struct{}"
	}

	return strings.Join([]string{"SwitchClusterProtectionModeRequestBody", string(data)}, " ")
}
