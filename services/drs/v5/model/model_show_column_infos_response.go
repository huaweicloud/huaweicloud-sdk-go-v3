package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowColumnInfosResponse Response Object
type ShowColumnInfosResponse struct {

	// 列映射信息
	ColumnMapInfos *[]ColumnMappingInfo `json:"column_map_infos,omitempty"`

	// 和列信息相关的对象
	ObjectWithColumnInfos *[]ObjectWithColumnInfo `json:"object_with_column_infos,omitempty"`
	HttpStatusCode        int                     `json:"-"`
}

func (o ShowColumnInfosResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowColumnInfosResponse struct{}"
	}

	return strings.Join([]string{"ShowColumnInfosResponse", string(data)}, " ")
}
