package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DetailIsUsedRes **参数解释**: 是否启用 **取值范围**: - false：否 - true：是
type DetailIsUsedRes struct {
}

func (o DetailIsUsedRes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetailIsUsedRes struct{}"
	}

	return strings.Join([]string{"DetailIsUsedRes", string(data)}, " ")
}
