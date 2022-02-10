package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 集群行为进度，显示创建和扩容进度的百分比。CREATING表示创建的百分比。
type ClusterListActionProgress struct {
	// 进度百分比。

	Creating *string `json:"CREATING,omitempty"`
}

func (o ClusterListActionProgress) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterListActionProgress struct{}"
	}

	return strings.Join([]string{"ClusterListActionProgress", string(data)}, " ")
}
