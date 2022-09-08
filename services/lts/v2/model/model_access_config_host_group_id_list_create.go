package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 日志接入主机组ID列表
type AccessConfigHostGroupIdListCreate struct {

	// 主机组ID列表
	HostGroupIdList []string `json:"host_group_id_list"`
}

func (o AccessConfigHostGroupIdListCreate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessConfigHostGroupIdListCreate struct{}"
	}

	return strings.Join([]string{"AccessConfigHostGroupIdListCreate", string(data)}, " ")
}
