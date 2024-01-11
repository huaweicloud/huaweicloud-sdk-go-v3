package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SchemaInfo 集群模式空间信息
type SchemaInfo struct {

	// Schema名称。
	SchemaName string `json:"schema_name"`

	// 数据库名称。
	DatabaseName string `json:"database_name"`

	// 集群schema使用空间总值。
	TotalValue int32 `json:"total_value"`

	// Schema空间阈值。
	PermSpace int32 `json:"perm_space"`

	// 倾斜率。
	SkewPercent float64 `json:"skew_percent"`

	// 最小值。
	MinValue int32 `json:"min_value"`

	// 最大值。
	MaxValue int32 `json:"max_value"`

	// 最小dn节点。
	MinDn string `json:"min_dn"`

	// 最大cn节点。
	MaxDn string `json:"max_dn"`

	// dn节点数量。
	DnNum int32 `json:"dn_num"`
}

func (o SchemaInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SchemaInfo struct{}"
	}

	return strings.Join([]string{"SchemaInfo", string(data)}, " ")
}
