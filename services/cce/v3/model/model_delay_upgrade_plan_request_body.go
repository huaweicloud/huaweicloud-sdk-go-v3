package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DelayUpgradePlanRequestBody struct {

	// 新的自动升级计划启动时间，必须早于notBeforeDeadline
	NotBefore string `json:"notBefore"`
}

func (o DelayUpgradePlanRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DelayUpgradePlanRequestBody struct{}"
	}

	return strings.Join([]string{"DelayUpgradePlanRequestBody", string(data)}, " ")
}
