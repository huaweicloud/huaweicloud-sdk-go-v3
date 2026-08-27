package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OnlineDdlTaskContentItem struct {

	// **参数解释**：  无锁变更的目标数据库。  **取值范围**： 不涉及。
	Schema *string `json:"schema,omitempty"`

	// **参数解释**：  无锁变更的DDL信息。
	DdlInfo *[]OnlineDdlInfoItem `json:"ddl_info,omitempty"`
}

func (o OnlineDdlTaskContentItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OnlineDdlTaskContentItem struct{}"
	}

	return strings.Join([]string{"OnlineDdlTaskContentItem", string(data)}, " ")
}
