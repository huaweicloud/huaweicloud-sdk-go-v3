package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LoadSchemaMetadataReq struct {

	// 逻辑库信息。
	CompressedDatabasesInfo string `json:"compressed_databases_info"`

	// 关联的后端DN信息。
	DnInstance []DnInstance `json:"dn_instance"`
}

func (o LoadSchemaMetadataReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LoadSchemaMetadataReq struct{}"
	}

	return strings.Join([]string{"LoadSchemaMetadataReq", string(data)}, " ")
}
