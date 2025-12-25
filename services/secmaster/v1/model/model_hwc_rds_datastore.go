package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HwcRdsDatastore 数据库信息。
type HwcRdsDatastore struct {

	// 数据库引擎，不区分大小写： MySQL PostgreSQL SQLServer
	Type string `json:"type"`

	// 数据库版本。
	Version string `json:"version"`

	// 数据库完整版本号。仅在数据库引擎是“PostgreSQL”时返回。
	CompleteVersion *string `json:"complete_version,omitempty"`
}

func (o HwcRdsDatastore) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HwcRdsDatastore struct{}"
	}

	return strings.Join([]string{"HwcRdsDatastore", string(data)}, " ")
}
