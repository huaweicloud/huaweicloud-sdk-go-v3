package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PersistableModelSaveAsDto struct {

	// **参数解释：**  新实例的唯一标识。如不填写，系统将自动生成。  **约束限制：**  不涉及。  **取值范围：**  -9223372036854775808到9223372036854775807的整数。  **默认取值：**  不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释：**  修改者名称，用于记录执行另存操作的用户。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Modifier *string `json:"modifier,omitempty"`

	// **参数解释：**  最后更新时间。使用UTC+0时间格式，格式为yyyy-MM-ddTHH:mm:ss.SSSZ。 例如：2025-12-04T03:50:40.599+0000。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	LastUpdateTime *string `json:"lastUpdateTime,omitempty"`

	// **参数解释：**  创建者账号。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Creator *string `json:"creator,omitempty"`

	// **参数解释：**  创建时间。使用UTC+0时间格式，格式为yyyy-MM-ddTHH:mm:ss.SSSZ。 例如：2025-12-04T03:50:40.599+0000。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	CreateTime *string `json:"createTime,omitempty"`

	// **参数解释：**  扩展类型，用于指定数据模型的具体扩展类型。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	RdmExtensionType *string `json:"rdmExtensionType,omitempty"`

	Tenant *ObjectReferenceParamDto `json:"tenant,omitempty"`

	// **参数解释：**  源模型的编号，用于指定数据来源模型。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	SourceEntityNumber *string `json:"sourceEntityNumber,omitempty"`

	// **参数解释：**  源实例的唯一标识。对于单实例，使用实例的id；对于版本实例，使用versionId。 获取方法请参见[分页查询实例 - ShowFindUsingPost](https://support.huaweicloud.com/api-idme/ShowFindUsingPost.html)。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	SourceInstanceId string `json:"sourceInstanceId"`

	// **参数解释：**  需要从源实例中置空的自定义属性列表，包括基本属性、扩展属性和分类属性。这些属性在新实例中将设置为空值。  **约束限制：**  属性名称长度不能超过1000个字符。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	NeedSetNullAttrs *[]string `json:"needSetNullAttrs,omitempty"`

	// **参数解释：**  需要覆盖的属性值对象。可指定源实例中某些属性的新值，未指定的属性保持与源实例一致。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	EntityToSave *interface{} `json:"entityToSave,omitempty"`

	// **参数解释：**  需要返回的属性对象，用于控制响应中返回的字段信息。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	EntityToReturn *interface{} `json:"entityToReturn,omitempty"`

	// **参数解释：**  数据模型的唯一键属性，用于业务层面的唯一标识。若指定，新实例将使用此值作为唯一键。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	UniqueKey *string `json:"uniqueKey,omitempty"`
}

func (o PersistableModelSaveAsDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PersistableModelSaveAsDto struct{}"
	}

	return strings.Join([]string{"PersistableModelSaveAsDto", string(data)}, " ")
}
