package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AccessConfigHostGroupIdList 日志接入主机组ID列表
type AccessConfigHostGroupIdList struct {

	// 主机组ID列表
	HostGroupIdList []string `json:"host_group_id_list"`
}

func (o AccessConfigHostGroupIdList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessConfigHostGroupIdList struct{}"
	}

	return strings.Join([]string{"AccessConfigHostGroupIdList", string(data)}, " ")
}
