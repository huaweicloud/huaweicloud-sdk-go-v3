package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 企业项目配额响应
type EpQuotas struct {
	// 总配额

	Quota int32 `json:"quota"`
	// qutoa的资源类型

	Type string `json:"type"`
	// 配额使用量

	Used int32 `json:"used"`
}

func (o EpQuotas) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EpQuotas struct{}"
	}

	return strings.Join([]string{"EpQuotas", string(data)}, " ")
}
