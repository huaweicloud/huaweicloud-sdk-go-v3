package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConnectionId **参数解释**：网络连接的ID。 **取值范围**：长度为1~64的英文、数字、中划线组合。
type ConnectionId struct {
}

func (o ConnectionId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConnectionId struct{}"
	}

	return strings.Join([]string{"ConnectionId", string(data)}, " ")
}
