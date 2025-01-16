package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListStarRocksDataReplicationConfigByDataBaseRequest Request Object
type ListStarRocksDataReplicationConfigByDataBaseRequest struct {

	// **参数解释**：  StarRocks实例ID，严格匹配UUID规则。  **约束限制**：  不涉及。  **取值范围**：  只能由英文字母、数字组成，后缀为in17，且长度为36个字符。  **默认值**：  不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**：  请求语言类型。  **约束限制**：  不涉及。  **取值范围**：  - en-us - zh-cn  **默认值**：  en-us。
	XLanguage string `json:"X-Language"`

	// **参数解释**：  目标数据库名。  **约束限制**：  不涉及。  **取值范围**：  字符长度限制3~128位，仅支持英文大小写字母、数字以及下划线。  **默认值**：  不涉及。
	Database string `json:"database"`
}

func (o ListStarRocksDataReplicationConfigByDataBaseRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStarRocksDataReplicationConfigByDataBaseRequest struct{}"
	}

	return strings.Join([]string{"ListStarRocksDataReplicationConfigByDataBaseRequest", string(data)}, " ")
}
