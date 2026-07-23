package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GenericLinkQueryDto struct {

	// **参数解释：**  查询角色，用于指定按哪种角色维度查询关系实例。  **约束限制：**  必须与objectId配合使用。  **取值范围：**  - TARGET：目标模型。 - SOURCE：源模型。 - TARGET_TYPE：目标模型对应的英文名称。 - SOURCE_TYPE：源模型对应的英文名称。  **默认取值：**  不涉及。
	Role *string `json:"role,omitempty"`

	// **参数解释：**  查询对象标识，与role参数配合使用。  **约束限制：**  - 当role为SOURCE或TARGET时，值为角色对应数据实例ID。 - 当role为SOURCE_TYPE或TARGET_TYPE时，值为角色对应的数据模型的英文名称。  **取值范围：**  - 当role为SOURCE或TARGET时，取值范围为：-9223372036854775808到9223372036854775807的整数。 - 当role为SOURCE_TYPE或TARGET_TYPE时，以大写字母开头，只能包含字母、数字、“_”，且长度为1-60个字符。  **默认取值：**  不涉及。
	ObjectId *string `json:"objectId,omitempty"`

	// **参数解释：**  是否仅返回关联的最新版本目标模型数据实例。此参数仅当源模型或目标模型为M-V模型实体时生效。  **约束限制：**  仅对M-V模型实体有效。  **取值范围：**  - true：仅返回关联的最新版本目标模型数据实例。 - false：返回关联的所有版本目标模型数据实例。  **默认取值：**  false。
	LatestOnly *bool `json:"latestOnly,omitempty"`

	// **参数解释：**  是否需要查询总记录数。开启后响应中的pageInfo将包含准确的totalRows和totalPages，但可能影响查询性能。  **约束限制：**  不涉及。  **取值范围：**  - true：需要。 - false：不需要。  **默认取值：**  false。
	IsNeedTotal *bool `json:"isNeedTotal,omitempty"`
}

func (o GenericLinkQueryDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GenericLinkQueryDto struct{}"
	}

	return strings.Join([]string{"GenericLinkQueryDto", string(data)}, " ")
}
