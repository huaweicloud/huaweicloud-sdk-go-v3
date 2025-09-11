package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SmnUrnResp **参数解释**： 通知组、通知对象对应的SMN主题URN。 **取值范围**： 只能包含字母、数字、“-”、“_” 、“:”。
type SmnUrnResp struct {
}

func (o SmnUrnResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SmnUrnResp struct{}"
	}

	return strings.Join([]string{"SmnUrnResp", string(data)}, " ")
}
