package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCategoryStatusResponse Response Object
type ShowCategoryStatusResponse struct {

	// **参数解释**： 查询结果总数。  **取值范围**： 不涉及。
	Total *int32 `json:"total,omitempty"`

	Result *StatusResponseResult `json:"result,omitempty"`

	// **参数解释**： 状态码。  **取值范围**： 不涉及。
	Status *string `json:"status,omitempty"`

	// **参数解释**： 响应信息。  **取值范围**： 不涉及。
	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowCategoryStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCategoryStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowCategoryStatusResponse", string(data)}, " ")
}
