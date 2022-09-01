package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// quota详细信息
type QuotaInfo struct {

	// 查询的资源类型。
	Type string `json:"type" xml:"type"`

	// 已使用的资源配额。
	Used int32 `json:"used" xml:"used"`

	// 查询出的资源的总配额。
	Quota int32 `json:"quota" xml:"quota"`

	// 资源的最小配额。
	Min int32 `json:"min" xml:"min"`

	// 资源的最大配额。
	Max int32 `json:"max" xml:"max"`
}

func (o QuotaInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotaInfo struct{}"
	}

	return strings.Join([]string{"QuotaInfo", string(data)}, " ")
}
