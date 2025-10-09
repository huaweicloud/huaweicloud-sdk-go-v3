package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DatastoreItem 数据库信息。
type DatastoreItem struct {

	// 数据库引擎。
	Type string `json:"type"`

	// 数据库版本号。
	Version string `json:"version"`

	// 是否有补丁版本的数据库支持升级，返回true时可以通过升级补丁接口进行数据库升级，否则不允许升级补丁。
	PatchAvailable bool `json:"patch_available"`

	// 数据库的完整版本号。
	WholeVersion string `json:"whole_version"`
}

func (o DatastoreItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DatastoreItem struct{}"
	}

	return strings.Join([]string{"DatastoreItem", string(data)}, " ")
}
