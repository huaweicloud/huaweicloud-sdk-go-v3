package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LatestCreateTime **参数解释**： 最近生成时间，毫秒(如果返回值为null，代表暂未生成) **取值范围**: 不涉及
type LatestCreateTime struct {
}

func (o LatestCreateTime) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LatestCreateTime struct{}"
	}

	return strings.Join([]string{"LatestCreateTime", string(data)}, " ")
}
