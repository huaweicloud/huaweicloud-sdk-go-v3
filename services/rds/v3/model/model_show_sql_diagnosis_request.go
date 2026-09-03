package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSqlDiagnosisRequest Request Object
type ShowSqlDiagnosisRequest struct {

	// **参数解释**：  实例ID，此参数是实例的唯一标识。  **约束限制**：  不涉及。  **取值范围**：  只能由英文字母、数字组成，长度为36个字符。  **默认取值**：  不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**：  请求语言类型。  **约束限制**：  不涉及。  **取值范围**：  - en-us - zh-cn **默认取值**：  en-us。
	XLanguage *string `json:"X-Language,omitempty"`

	// **参数解释**：  实例诊断类型。  **约束限制**：  不涉及。  **取值范围**：  - disk (当前仅支持传入该项)  **默认取值**：  不涉及。
	Category string `json:"category"`

	// **参数解释**：  sql诊断类型。  **约束限制**：  不涉及。  **取值范围**：  - time (执行耗时长) - temp (临时表类) - sort (排序类) - ddl (DDL类)  **默认取值**：  不涉及。
	Subcategory string `json:"subcategory"`

	// **参数解释**：  索引位置，偏移量。从第一条数据偏移offset条数据后开始查询，默认为0（偏移0条数据，表示从第一条数据开始查询），必须为数字，不能为负数。  **约束限制**：  不涉及。  **取值范围**：  不涉及  **默认取值**：  0
	Offset *string `json:"offset,omitempty"`

	// **参数解释**：  查询记录数。默认为10，最小值为1，最大值为100。  **约束限制**：  不涉及。  **取值范围**：  不涉及  **默认取值**：  10
	Limit *string `json:"limit,omitempty"`
}

func (o ShowSqlDiagnosisRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSqlDiagnosisRequest struct{}"
	}

	return strings.Join([]string{"ShowSqlDiagnosisRequest", string(data)}, " ")
}
