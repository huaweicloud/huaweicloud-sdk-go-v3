package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HostStatus **参数解释**： 主机状态 **取值范围**: - ACTIVE：正在运行 - SHUTOFF：关机 - BUILDING：创建中 - ERROR：故障
type HostStatus struct {
}

func (o HostStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HostStatus struct{}"
	}

	return strings.Join([]string{"HostStatus", string(data)}, " ")
}
