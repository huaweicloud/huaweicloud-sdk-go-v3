package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ChargeInfo struct {
	// 计费模式。 取值范围： - prePaid：预付费，即包年/包月。暂不支持。 - postPaid：后付费，即按需计费。

	ChargeMode string `json:"charge_mode"`
}

func (o ChargeInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChargeInfo struct{}"
	}

	return strings.Join([]string{"ChargeInfo", string(data)}, " ")
}
