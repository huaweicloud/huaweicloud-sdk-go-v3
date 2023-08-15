package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RestoreInstanceFromCollectionRequestBodyRestoreCollections struct {

	// 数据库名称。
	Database string `json:"database"`

	// 数据库恢复时间点。如果是数据库级恢复，该参数必传，UNIX时间戳格式，单位是毫秒，时区是UTC。
	RestoreDatabaseTime *string `json:"restore_database_time,omitempty"`

	// 集合信息。
	Collections *[]RestoreInstanceFromCollectionRequestBodyCollections `json:"collections,omitempty"`
}

func (o RestoreInstanceFromCollectionRequestBodyRestoreCollections) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreInstanceFromCollectionRequestBodyRestoreCollections struct{}"
	}

	return strings.Join([]string{"RestoreInstanceFromCollectionRequestBodyRestoreCollections", string(data)}, " ")
}
