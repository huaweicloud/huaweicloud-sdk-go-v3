package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRestoreTablesResponse Response Object
type ShowRestoreTablesResponse struct {

	// **参数解释**：  已恢复库表信息列表。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	DataList       *[]RestoreTablesInfo `json:"data_list,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ShowRestoreTablesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRestoreTablesResponse struct{}"
	}

	return strings.Join([]string{"ShowRestoreTablesResponse", string(data)}, " ")
}
