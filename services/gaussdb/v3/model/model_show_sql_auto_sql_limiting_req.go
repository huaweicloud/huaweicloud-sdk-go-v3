package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowSqlAutoSqlLimitingReq struct {

	// **参数解释**：  节点ID列表。  获取方法请参见[查询实例详情](https://support.huaweicloud.com/api-taurusdb/ShowGaussMySqlInstanceInfoUnifyStatus.html)。  **约束限制**：  节点角色必须为主节点。  **取值范围**：  列表元素为节点ID，只能由英文字母、数字组成，前面为UUID，后缀为no07，长度为36个字符。  **默认取值**：  不涉及。
	NodeIds *[]string `json:"node_ids,omitempty"`
}

func (o ShowSqlAutoSqlLimitingReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSqlAutoSqlLimitingReq struct{}"
	}

	return strings.Join([]string{"ShowSqlAutoSqlLimitingReq", string(data)}, " ")
}
