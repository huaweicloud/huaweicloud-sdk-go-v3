package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ObjectNumRes **参数解释**: 防护对象个数 **取值范围**: 取值0-100000
type ObjectNumRes struct {
}

func (o ObjectNumRes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ObjectNumRes struct{}"
	}

	return strings.Join([]string{"ObjectNumRes", string(data)}, " ")
}
