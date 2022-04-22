package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 删除复制对请求数据结构
type DeleteReplicationRequestParams struct {

	// 保护组的ID。
	ServerGroupId *string `json:"server_group_id,omitempty"`

	// 是否删除容灾站点磁盘，默认值为false。
	DeleteTargetVolume *bool `json:"delete_target_volume,omitempty"`
}

func (o DeleteReplicationRequestParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteReplicationRequestParams struct{}"
	}

	return strings.Join([]string{"DeleteReplicationRequestParams", string(data)}, " ")
}
