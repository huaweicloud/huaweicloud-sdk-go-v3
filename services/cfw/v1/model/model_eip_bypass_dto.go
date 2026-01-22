package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EipBypassDto struct {

	// 防护操作类型，1表示一键关闭防护，0表示一键恢复防护
	BypassOperation int32 `json:"bypass_operation"`
}

func (o EipBypassDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EipBypassDto struct{}"
	}

	return strings.Join([]string{"EipBypassDto", string(data)}, " ")
}
