package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VirusNum **参数解释** 新发现病毒数量 **取值范围** 非负整数，最小值0；单位：个
type VirusNum struct {
}

func (o VirusNum) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VirusNum struct{}"
	}

	return strings.Join([]string{"VirusNum", string(data)}, " ")
}
