package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StartOnlineDdlInfoItem struct {

	// **参数解释**：  无锁变更的具体执行SQL。  **约束限制**:  满足ALTER TABLE Statement语法形式，多条SQL需要以英文分号隔开。  **取值范围**：   不涉及。  **默认取值**：   不涉及。
	Sql string `json:"sql"`
}

func (o StartOnlineDdlInfoItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartOnlineDdlInfoItem struct{}"
	}

	return strings.Join([]string{"StartOnlineDdlInfoItem", string(data)}, " ")
}
