package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Success **参数解释**: 是否成功 **取值范围**: - true  成功 - false 失败
type Success struct {
}

func (o Success) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Success struct{}"
	}

	return strings.Join([]string{"Success", string(data)}, " ")
}
