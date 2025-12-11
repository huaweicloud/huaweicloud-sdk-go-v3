package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HostTaskStatus **参数解释**: 服务器扫描状态 **取值范围**: - scanning ：扫描中 - success ：扫描成功 - fail ：扫描失败 - cancel ：取消扫描
type HostTaskStatus struct {
}

func (o HostTaskStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HostTaskStatus struct{}"
	}

	return strings.Join([]string{"HostTaskStatus", string(data)}, " ")
}
