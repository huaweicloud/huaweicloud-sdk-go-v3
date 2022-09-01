package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RestoreInstanceFromCollectionRequestBody struct {

	// 数据库信息。
	RestoreCollections []RestoreInstanceFromCollectionRequestBodyRestoreCollections `json:"restore_collections" xml:"restore_collections"`
}

func (o RestoreInstanceFromCollectionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreInstanceFromCollectionRequestBody struct{}"
	}

	return strings.Join([]string{"RestoreInstanceFromCollectionRequestBody", string(data)}, " ")
}
