package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RestoreHbaInfoRequestBody struct {

	// **参数解释**: 客户端接入认证修改历史记录ID。 **取值范围**: 传空时表示恢复默认。
	HbaHistoryId *string `json:"hba_history_id,omitempty"`
}

func (o RestoreHbaInfoRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreHbaInfoRequestBody struct{}"
	}

	return strings.Join([]string{"RestoreHbaInfoRequestBody", string(data)}, " ")
}
