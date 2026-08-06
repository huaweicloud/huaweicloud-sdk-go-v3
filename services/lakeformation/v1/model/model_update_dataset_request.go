package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDatasetRequest Request Object
type UpdateDatasetRequest struct {

	// **参数解释:** LakeFormation实例ID。创建实例时自动生成。例如：2180518f-42b8-4947-b20b-adfc53981a25。 **约束限制:** **取值范围:** 不涉及
	InstanceId string `json:"instance_id"`

	// **参数解释:** catalog名称。 **约束限制:** 只能包含字母、数字和下划线，且长度为1~256个字符。 **取值范围:** 长度为1~256个字符 **默认取值:** 不涉及
	CatalogName string `json:"catalog_name"`

	// **参数解释:** 数据集名称。 **约束限制:** 只能包含中文、字母、数字和_-特殊字符，且长度为1~256个字符。 **取值范围:** 长度为1~256个字符 **默认取值:** 不涉及
	DatasetName string `json:"dataset_name"`

	// **参数解释:** 数据库名称。 **约束限制:** 只能包含中文、字母、数字、下划线、中划线，且长度为1~128个字符。 **取值范围:** 长度为1~128个字符 **默认取值:** 不涉及
	DatabaseName string `json:"database_name"`

	Body *AlterDatasetInput `json:"body,omitempty"`
}

func (o UpdateDatasetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDatasetRequest struct{}"
	}

	return strings.Join([]string{"UpdateDatasetRequest", string(data)}, " ")
}
