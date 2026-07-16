package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NodePoolMetadata 节点池的metadata信息。
type NodePoolMetadata struct {

	// **参数解释**： 节点池名称。用户可进行指定，若未指定将会使用默认名称。 **取值范围**： 不涉及。
	Name string `json:"name"`
}

func (o NodePoolMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodePoolMetadata struct{}"
	}

	return strings.Join([]string{"NodePoolMetadata", string(data)}, " ")
}
