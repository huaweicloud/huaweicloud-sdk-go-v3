package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StarRocksInstanceInfoDatastore 数据库信息。
type StarRocksInstanceInfoDatastore struct {

	// 数据库ID。
	Id *string `json:"id,omitempty"`

	// 数据库类型。
	Type *string `json:"type,omitempty"`

	// 数据库版本。
	Version *string `json:"version,omitempty"`
}

func (o StarRocksInstanceInfoDatastore) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StarRocksInstanceInfoDatastore struct{}"
	}

	return strings.Join([]string{"StarRocksInstanceInfoDatastore", string(data)}, " ")
}
