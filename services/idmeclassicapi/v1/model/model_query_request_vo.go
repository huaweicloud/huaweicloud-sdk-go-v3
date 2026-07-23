package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QueryRequestVo 查询请求对象。
type QueryRequestVo struct {

	// **参数解释：**  编码类型，用于指定查询的字符编码。  **约束限制：**  不涉及。  **取值范围：**  - ASCII：ASCII - BINARY：BINARY - GB2312：GB2312 - GBK：GBK - UTF8：UTF8 - UTF16：UTF16 - UTF32：UTF32 - UTF8MB4：UTF8MB4  **默认取值：**  UTF8。
	CharacterSet *string `json:"characterSet,omitempty"`

	// **参数解释：**  查询条件组。  此参数已废弃，不建议继续使用，建议使用替代参数filter。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Conditions *[]QueryCondition `json:"conditions,omitempty"`

	// **参数解释：**  是否对查询结果进行加密。  **约束限制：**  不涉及。  **取值范围：**  - true：加密。 - false：不加密。  **默认取值：**  不涉及。
	Decrypt *bool `json:"decrypt,omitempty"`

	// **参数解释：**  实体类型，用于指定查询的目标实体类型。  **约束限制：**  不涉及。  **取值范围：**  - ENTITY：数据实体。 - RRELATION：关系实体。  **默认取值：**  不涉及。
	EntityType *string `json:"entityType,omitempty"`

	Filter *QueryCondition `json:"filter,omitempty"`

	// **参数解释：**  是否需要查询总记录数。  **约束限制：**  不涉及。  **取值范围：**  - true：需要。 - false：不需要。  **默认取值：**  false。
	IsNeedTotal *bool `json:"isNeedTotal,omitempty"`

	// **参数解释：**  是否需要展示所有参考对象信息。  **约束限制：**  不涉及。  **取值范围：**  - true：需要。 - false：不需要。  **默认取值：**  false。
	IsPresentAll *bool `json:"isPresentAll,omitempty"`

	// **参数解释：**  需要展示详细信息的参考对象。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	NeedPresentDetail *[]string `json:"needPresentDetail,omitempty"`

	// **参数解释：**  排序字段，按某个指定字段对查询结果进行排序。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	OrderBy *string `json:"orderBy,omitempty"`

	// **参数解释：**  排序字段的表别名。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	OrderByTableAlias *string `json:"orderByTableAlias,omitempty"`

	// **参数解释：**  多租查询参数，用于控制是否包含公共租户数据。  **约束限制：**  不涉及。  **取值范围：**  - EXCLUDE_PUBLIC_DATA：不包括公共数据。 - INCLUDE_PUBLIC_DATA：包括公共数据。 - ONLY_NEED_PUBLIC_DATA：只有公共数据。  **默认取值：**  不涉及。
	PublicData *string `json:"publicData,omitempty"`

	// **参数解释：**  排序方向。  **约束限制：**  不涉及。  **取值范围：**  - ASC：表示升序。 - DESC：表示降序。  **默认取值：**  ASC。
	Sort *string `json:"sort,omitempty"`

	// **参数解释：**  多字段排序配置。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Sorts *[]SortInfoVo `json:"sorts,omitempty"`
}

func (o QueryRequestVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryRequestVo struct{}"
	}

	return strings.Join([]string{"QueryRequestVo", string(data)}, " ")
}
