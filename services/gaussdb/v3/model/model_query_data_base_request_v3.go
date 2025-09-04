package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QueryDataBaseRequestV3 查询HTAP主实例数据库请求体。
type QueryDataBaseRequestV3 struct {

	// **参数解释**：  查询的数据库名称。  **约束限制**：  仅支持英文大小写字母、数字以及下划线，模糊搜索时列表元素数目必须为1。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	Databases *[]string `json:"databases,omitempty"`

	// **参数解释**：  需要查询数据库的源实例ID，严格匹配UUID规则。  **约束限制**：  只能由英文字母、数字组成，且长度为36个字符。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	SourceInstanceId *string `json:"source_instance_id,omitempty"`
}

func (o QueryDataBaseRequestV3) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryDataBaseRequestV3 struct{}"
	}

	return strings.Join([]string{"QueryDataBaseRequestV3", string(data)}, " ")
}
