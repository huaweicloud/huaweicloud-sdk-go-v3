package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 数据库信息。
type DatastoreItem struct {

	// 数据库引擎。
	Type string `json:"type" xml:"type"`

	// 数据库版本号。
	Version string `json:"version" xml:"version"`
}

func (o DatastoreItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DatastoreItem struct{}"
	}

	return strings.Join([]string{"DatastoreItem", string(data)}, " ")
}
