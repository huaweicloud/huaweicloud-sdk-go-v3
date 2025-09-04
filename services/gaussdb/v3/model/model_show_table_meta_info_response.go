package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTableMetaInfoResponse Response Object
type ShowTableMetaInfoResponse struct {

	// 要版本升级的批量实例。
	TableMetaInfos *[]TableMetaInfo `json:"table_meta_infos,omitempty"`

	// 数据表名称
	TableNames *[]string `json:"table_names,omitempty"`

	// 数据库名称
	DatabaseNames  *[]string `json:"database_names,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowTableMetaInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTableMetaInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowTableMetaInfoResponse", string(data)}, " ")
}
