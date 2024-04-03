package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QueryRequestVo struct {
	CharacterSet *CharacterSetEnum `json:"characterSet,omitempty"`

	// 查询条件。
	Conditions *[]QueryCondition `json:"conditions,omitempty"`

	// 是否加密。 - true：加密。 - false：不加密。
	Decrypt *bool `json:"decrypt,omitempty"`

	// 实体类型。
	EntityType *string `json:"entityType,omitempty"`

	Filter *QueryCondition `json:"filter,omitempty"`

	// 是否需要查询总记录数。 - true：需要。 - false：不需要。
	IsNeedTotal *bool `json:"isNeedTotal,omitempty"`

	// 是否需要展示所有参考对象信息。 - true：需要。 - false：不需要。
	IsPresentAll *bool `json:"isPresentAll,omitempty"`

	// 需要展示详细信息的参考对象。
	NeedPresentDetail *[]string `json:"needPresentDetail,omitempty"`

	// 按某个字段进行排序。
	OrderBy *string `json:"orderBy,omitempty"`

	// 排序字段的表别名。
	OrderByTableAlias *string `json:"orderByTableAlias,omitempty"`

	// 多租查询参数。 - EXCLUDE_PUBLIC_DATA：不包括公共数据。 - INCLUDE_PUBLIC_DATA：包括公共数据。 - ONLY_NEED_PUBLIC_DATA：只有公共数据。
	PublicData *string `json:"publicData,omitempty"`

	// 排序方向。 - ASC：表示升序。 - DESC：表示降序。
	Sort *string `json:"sort,omitempty"`

	// 排序。
	Sorts *[]SortInfoVo `json:"sorts,omitempty"`
}

func (o QueryRequestVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryRequestVo struct{}"
	}

	return strings.Join([]string{"QueryRequestVo", string(data)}, " ")
}