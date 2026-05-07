package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HostStatusRes **参数解释**： 服务器状态 **取值范围**： - ACTIVE：运行中 - SHUTOFF：关机 - BUILDING：创建中 - ERROR：故障
type HostStatusRes struct {
}

func (o HostStatusRes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HostStatusRes struct{}"
	}

	return strings.Join([]string{"HostStatusRes", string(data)}, " ")
}
