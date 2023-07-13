package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RestoreInstanceFromCollectionRequestBodyCollections struct {

	// 恢复前表名。
	OldName string `json:"old_name"`

	// 恢复后表名。
	NewName *string `json:"new_name,omitempty"`

	// 数据库集合恢复时间点。UNIX时间戳格式，单位是毫秒，时区是UTC。
	RestoreCollectionTime string `json:"restore_collection_time"`
}

func (o RestoreInstanceFromCollectionRequestBodyCollections) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreInstanceFromCollectionRequestBodyCollections struct{}"
	}

	return strings.Join([]string{"RestoreInstanceFromCollectionRequestBodyCollections", string(data)}, " ")
}
