package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 配额详细信息。
type QuotaDetail struct {

	// 已使用的数量。
	InUse int32 `json:"in_use" xml:"in_use"`

	// 最大的数量。
	Limit int32 `json:"limit" xml:"limit"`

	// 预留属性。
	Reserved int32 `json:"reserved" xml:"reserved"`
}

func (o QuotaDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotaDetail struct{}"
	}

	return strings.Join([]string{"QuotaDetail", string(data)}, " ")
}
