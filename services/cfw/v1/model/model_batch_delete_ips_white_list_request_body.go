package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteIpsWhiteListRequestBody struct {

	// **参数解释**：  要删除的白名单ID列表  **约束限制**：  不涉及  **取值范围**：  不涉及  **默认取值**：  不涉及
	ListIds []string `json:"list_ids"`
}

func (o BatchDeleteIpsWhiteListRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteIpsWhiteListRequestBody struct{}"
	}

	return strings.Join([]string{"BatchDeleteIpsWhiteListRequestBody", string(data)}, " ")
}
