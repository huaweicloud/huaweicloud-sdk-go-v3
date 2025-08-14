package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HostId **参数解释**： 主机ID **取值范围**： 字符长度1-64位
type HostId struct {
}

func (o HostId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HostId struct{}"
	}

	return strings.Join([]string{"HostId", string(data)}, " ")
}
