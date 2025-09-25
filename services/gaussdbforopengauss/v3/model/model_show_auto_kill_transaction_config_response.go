package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAutoKillTransactionConfigResponse Response Object
type ShowAutoKillTransactionConfigResponse struct {

	// **参数解释**: 配置类型。 **取值范围**: 不涉及。
	Type *string `json:"type,omitempty"`

	// **参数解释**: 筛选用户名。
	Usernames *[]string `json:"usernames,omitempty"`

	// **参数解释**: 阈值。 **取值范围**: 不涉及。
	Threshold *int64 `json:"threshold,omitempty"`

	// **参数解释**: 配置是否开启。 **取值范围**: 不涉及。
	AutoStop *bool `json:"auto_stop,omitempty"`

	// **参数解释**: 当前实例包含的所有数据库。
	DatabaseNames *[]string `json:"database_names,omitempty"`

	// **参数解释**: 筛选数据库名。
	DatabaseNamesSelect *[]string `json:"database_names_select,omitempty"`
	HttpStatusCode      int       `json:"-"`
}

func (o ShowAutoKillTransactionConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAutoKillTransactionConfigResponse struct{}"
	}

	return strings.Join([]string{"ShowAutoKillTransactionConfigResponse", string(data)}, " ")
}
