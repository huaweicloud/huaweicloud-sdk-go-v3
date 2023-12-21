package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SortInfoVo struct {
	CharacterSet *CharacterSetEnum `json:"characterSet,omitempty"`

	// 按某个字段进行排序。
	OrderBy *string `json:"orderBy,omitempty"`

	// 排序方向。 - ASC：表示升序。 - DESC：表示降序。
	Sort *string `json:"sort,omitempty"`

	// 排序信息。
	SortInfo *string `json:"sortInfo,omitempty"`

	// 排序信息字段。
	SortInfoOrderBy *string `json:"sortInfoOrderBy,omitempty"`
}

func (o SortInfoVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SortInfoVo struct{}"
	}

	return strings.Join([]string{"SortInfoVo", string(data)}, " ")
}
