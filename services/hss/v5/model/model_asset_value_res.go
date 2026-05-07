package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssetValueRes **参数解释**: 资产重要性 **约束限制**: 不涉及 **取值范围**： - important：重要资产 - common：一般资产 - test：测试资产  **默认取值**: 不涉及
type AssetValueRes struct {
}

func (o AssetValueRes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssetValueRes struct{}"
	}

	return strings.Join([]string{"AssetValueRes", string(data)}, " ")
}
