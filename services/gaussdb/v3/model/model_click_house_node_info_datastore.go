package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClickHouseNodeInfoDatastore 数据存储信息。
type ClickHouseNodeInfoDatastore struct {

	// 引擎ID。
	Id string `json:"id"`

	// 引擎类型，现在只支持click-house。
	Type string `json:"type"`

	// 引擎版本。
	Version string `json:"version"`
}

func (o ClickHouseNodeInfoDatastore) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClickHouseNodeInfoDatastore struct{}"
	}

	return strings.Join([]string{"ClickHouseNodeInfoDatastore", string(data)}, " ")
}
