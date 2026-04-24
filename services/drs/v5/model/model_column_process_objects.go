package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ColumnProcessObjects 列加工对象
type ColumnProcessObjects struct {

	// 选择的源库对象名
	ObjectSourceNames *[]string `json:"object_source_names,omitempty"`

	// 映射后的对象名
	ObjectAliasName *string `json:"object_alias_name,omitempty"`

	// 附加列是否已下发
	IsSent *bool `json:"is_sent,omitempty"`

	// 附加列信息
	ExtraColumnInfos *[]AddColumnInfo `json:"extra_column_infos,omitempty"`
}

func (o ColumnProcessObjects) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ColumnProcessObjects struct{}"
	}

	return strings.Join([]string{"ColumnProcessObjects", string(data)}, " ")
}
