package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 集群行为进度，显示创建和扩容进度的百分比。
type ClusterDetailActionProgress struct {
	// 进度百分比。

	Creating *string `json:"CREATING,omitempty"`
}

func (o ClusterDetailActionProgress) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterDetailActionProgress struct{}"
	}

	return strings.Join([]string{"ClusterDetailActionProgress", string(data)}, " ")
}
