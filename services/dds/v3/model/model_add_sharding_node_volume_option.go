package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddShardingNodeVolumeOption struct {
	// 指定新增的所有shard组的磁盘容量。取值范围：10GB~2000GB。

	Size string `json:"size"`
}

func (o AddShardingNodeVolumeOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddShardingNodeVolumeOption struct{}"
	}

	return strings.Join([]string{"AddShardingNodeVolumeOption", string(data)}, " ")
}
