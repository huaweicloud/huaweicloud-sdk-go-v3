package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EsdbCheckRdsConnectionsRequestV3 请求体
type EsdbCheckRdsConnectionsRequestV3 struct {

	// **参数解释**：  实例ID，此参数是实例的唯一标识。  **约束限制**：  不涉及。  **取值范围**：  只能由英文字母、数字组成，后缀为in09，长度为36个字符。  **默认取值**：  不涉及。
	InstanceId *string `json:"instance_id,omitempty"`

	// **参数解释**：  rds连通性检查相关信息的集合。  **参数范围**：  不涉及。
	Infos *[]EsdbCheckRdsConnectionRequestV3 `json:"infos,omitempty"`
}

func (o EsdbCheckRdsConnectionsRequestV3) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EsdbCheckRdsConnectionsRequestV3 struct{}"
	}

	return strings.Join([]string{"EsdbCheckRdsConnectionsRequestV3", string(data)}, " ")
}
