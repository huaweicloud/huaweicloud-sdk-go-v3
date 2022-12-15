package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DataStoreList struct {

	// 数据库引擎。
	DatastoreName string `json:"datastore_name"`

	// 数据库引擎版本。
	Version string `json:"version"`
}

func (o DataStoreList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataStoreList struct{}"
	}

	return strings.Join([]string{"DataStoreList", string(data)}, " ")
}
