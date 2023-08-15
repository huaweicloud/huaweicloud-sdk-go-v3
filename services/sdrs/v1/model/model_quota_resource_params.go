package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QuotaResourceParams 配额资源数据结构
type QuotaResourceParams struct {

	// 资源类型server_groups：表示保护组资源类型。replications：表示复制对资源类型。
	Type string `json:"type"`

	// 已经使用的资源个数。
	Used int32 `json:"used"`

	// 资源配额。-1：表示无穷大。
	Quota int32 `json:"quota"`

	// 设置该资源配额允许的最小值。
	Min int32 `json:"min"`

	// 设置该资源配额允许的最大值。-1：表示无穷大。
	Max int32 `json:"max"`
}

func (o QuotaResourceParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotaResourceParams struct{}"
	}

	return strings.Join([]string{"QuotaResourceParams", string(data)}, " ")
}
