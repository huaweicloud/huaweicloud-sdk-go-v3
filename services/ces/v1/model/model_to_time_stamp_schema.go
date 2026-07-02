package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ToTimeStampSchema **参数解释** 查询数据截止时间，UNIX时间戳，单位毫秒 **约束限制** from 必须小于to **取值范围** 毫秒级时间戳范围为[1111111111111,9999999999999] **默认取值** 不涉及
type ToTimeStampSchema struct {
}

func (o ToTimeStampSchema) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ToTimeStampSchema struct{}"
	}

	return strings.Join([]string{"ToTimeStampSchema", string(data)}, " ")
}
