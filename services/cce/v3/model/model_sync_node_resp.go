package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SyncNodeResp **参数解释**： 固定值\"Sync node success\"，表示同步节点成功。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
type SyncNodeResp struct {
}

func (o SyncNodeResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SyncNodeResp struct{}"
	}

	return strings.Join([]string{"SyncNodeResp", string(data)}, " ")
}
