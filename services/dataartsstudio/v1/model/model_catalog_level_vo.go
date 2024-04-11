package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CatalogLevelVo struct {

	// 编号。
	Id *int64 `json:"id,omitempty"`

	// 层级。
	Level *int32 `json:"level,omitempty"`

	// 中文名称。
	NameCh *string `json:"name_ch,omitempty"`

	// 英文名称。
	NameEn *string `json:"name_en,omitempty"`
}

func (o CatalogLevelVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CatalogLevelVo struct{}"
	}

	return strings.Join([]string{"CatalogLevelVo", string(data)}, " ")
}
