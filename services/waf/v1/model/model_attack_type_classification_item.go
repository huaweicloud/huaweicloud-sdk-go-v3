package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AttackTypeClassificationItem struct {

	// AttackTypeItem的总数量
	Total *int32 `json:"total,omitempty" xml:"total"`

	// AttackTypeItem详细信息
	Items *[]AttackTypeItem `json:"items,omitempty" xml:"items"`
}

func (o AttackTypeClassificationItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttackTypeClassificationItem struct{}"
	}

	return strings.Join([]string{"AttackTypeClassificationItem", string(data)}, " ")
}
