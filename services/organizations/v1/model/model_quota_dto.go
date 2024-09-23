package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QuotaDto 组织的配额。
type QuotaDto struct {

	// 配额类型。account：账号；organizational_unit：组织单元；policy：策略。
	Type string `json:"type"`

	// 配额数量。
	Quota int32 `json:"quota"`

	// 最小配额。
	Min int32 `json:"min"`

	// 最大配额。
	Max int32 `json:"max"`

	// 已使用数量。
	Used int32 `json:"used"`
}

func (o QuotaDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotaDto struct{}"
	}

	return strings.Join([]string{"QuotaDto", string(data)}, " ")
}
