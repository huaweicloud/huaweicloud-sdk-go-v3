package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DefaultReport **参数解释**: 是否是默认的，默认的不能删除 **取值范围**: - true ：是。 - false ：否。
type DefaultReport struct {
}

func (o DefaultReport) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DefaultReport struct{}"
	}

	return strings.Join([]string{"DefaultReport", string(data)}, " ")
}
