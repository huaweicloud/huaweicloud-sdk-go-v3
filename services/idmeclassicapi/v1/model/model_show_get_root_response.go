package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowGetRootResponse Response Object
type ShowGetRootResponse struct {

	// **参数解释：**  请求结果。  **取值范围：**  - SUCCESS：请求成功。 - FAIL：请求失败。
	Result *string `json:"result,omitempty"`

	// **参数解释：**  请求数据，返回目标数据实例所在树形结构的根节点信息列表。 数组按从直接父节点到根节点的顺序排列，包含各父节点的完整列表属性。  **取值范围：**  不涉及。
	Data *[]BasicObjectQueryViewDto `json:"data,omitempty"`

	// **参数解释：**  异常信息，当请求失败时返回具体的错误描述。  **取值范围：**  不涉及。
	Errors *[]string `json:"errors,omitempty"`

	PageInfo       *PageInfoViewDto `json:"pageInfo,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ShowGetRootResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGetRootResponse struct{}"
	}

	return strings.Join([]string{"ShowGetRootResponse", string(data)}, " ")
}
