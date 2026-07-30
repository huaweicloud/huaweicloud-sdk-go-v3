package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDatasetsRequest Request Object
type ListDatasetsRequest struct {

	// **参数解释:** LakeFormation实例ID。创建实例时自动生成。例如：2180518f-42b8-4947-b20b-adfc53981a25。 **约束限制:** 不涉及 **取值范围:** 不涉及 **默认取值:** 不涉及
	InstanceId string `json:"instance_id"`

	// **参数解释:** catalog名称。 **约束限制:** 只能包含字母、数字和下划线，且长度为1~256个字符。 **取值范围:** 长度为1~256个字符 **默认取值:** 不涉及
	CatalogName string `json:"catalog_name"`

	// **参数解释:** 数据库名称。 **约束限制:** 只能包含中文、字母、数字、下划线、中划线，且长度为1~128个字符。 **取值范围:** 长度为1~128个字符 **默认取值:** 不涉及
	DatabaseName string `json:"database_name"`

	// **参数解释:** 查询返回条数。 **约束限制:** 取值为0~1000 **取值范围:** 取值为0~1000 **默认取值:** 1000
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释:** 查询的起始记录ID。 **约束限制:** 长度为0~256个字符 **取值范围:** 长度为0~256个字符 **默认取值:** 不涉及
	Marker *string `json:"marker,omitempty"`

	// **参数解释:** 是否查询上一页。 **约束限制:** 不涉及 **取值范围:** 不涉及 **默认取值:** false
	ReversePage *bool `json:"reverse_page,omitempty"`

	// **参数解释:** 数据集名称通配符，用于模糊查询。 **约束限制:** 只能包含中文、字母、数字和_|*.-特殊字符，且长度为1~256个字符。 **取值范围:** 长度为1~256个字符
	NamePartern *string `json:"name_partern,omitempty"`

	// 数据格式 描述文件的组织方式：行存储/文本/图片/音频/视频/自定义
	Format *string `json:"format,omitempty"`
}

func (o ListDatasetsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDatasetsRequest struct{}"
	}

	return strings.Join([]string{"ListDatasetsRequest", string(data)}, " ")
}
