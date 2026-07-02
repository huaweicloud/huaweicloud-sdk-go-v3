package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IntelligentKillSessionsReq struct {

	// **参数解释**：  节点ID，此参数是节点的唯一标识。  获取方法请参见[查询实例详情](https://support.huaweicloud.com/api-taurusdb/ShowGaussMySqlInstanceInfoUnifyStatus.html)。  **约束限制**：  不涉及。  **取值范围**：  只能由英文字母、数字组成，后缀为no07，长度为36个字符。  **默认取值**：  不涉及。
	NodeId string `json:"node_id"`
}

func (o IntelligentKillSessionsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IntelligentKillSessionsReq struct{}"
	}

	return strings.Join([]string{"IntelligentKillSessionsReq", string(data)}, " ")
}
