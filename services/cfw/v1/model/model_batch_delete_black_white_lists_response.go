package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteBlackWhiteListsResponse Response Object
type BatchDeleteBlackWhiteListsResponse struct {

	// **参数解释**： 批量删除黑白名单列表id **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Data           *[]string `json:"data,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o BatchDeleteBlackWhiteListsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteBlackWhiteListsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteBlackWhiteListsResponse", string(data)}, " ")
}
